package v1

import (
	"context"
	"github.com/kumahq/kuma/app/kuma-prometheus-sd/pkg/discovery/xds"
	"io"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"go.uber.org/multierr"

	"github.com/prometheus/prometheus/discovery/targetgroup"

	mads_v1_client "github.com/kumahq/kuma/pkg/mads/v1/client"
)

type streamDiscoverer struct {
	log     logr.Logger
	config  xds.DiscoveryConfig
	handler *Handler
}

type streamFactory struct {
	handler Handler
}

// CreateDiscoverer implements xds.DiscovererFactory
func (f *streamFactory) CreateDiscoverer(config xds.DiscoveryConfig, log logr.Logger) xds.DiscovererE {
	return &streamDiscoverer{
		log:     log,
		config:  config,
		handler: &f.handler,
	}
}

func NewFactory() xds.DiscovererFactory {
	return &streamFactory{}
}

func (s *streamDiscoverer) Run(ctx context.Context, ch chan<- []*targetgroup.Group) (errs error) {
	s.log.Info("creating a gRPC client for Monitoring Assignment Discovery Service (MADS) server ...")
	client, err := mads_v1_client.New(s.config.ServerURL)
	if err != nil {
		return errors.Wrap(err, "failed to connect to gRPC server")
	}
	defer func() {
		s.log.Info("closing a connection ...")
		if err := client.Close(); err != nil {
			errs = multierr.Append(errs, errors.Wrapf(err, "failed to close a connection"))
		}
	}()

	s.log.Info("starting an xDS stream ...")
	stream, err := client.StartStream()
	if err != nil {
		return errors.Wrap(err, "failed to start an xDS stream")
	}
	defer func() {
		s.log.Info("closing an xDS stream ...")
		if err := stream.Close(); err != nil {
			errs = multierr.Append(errs, errors.Wrapf(err, "failed to close an xDS stream"))
		}
	}()

	s.log.Info("sending first discovery request on a new xDS stream ...")
	err = stream.RequestAssignments(s.config.ClientName)
	if err != nil {
		return errors.Wrap(err, "failed to send a discovery request")
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		s.log.Info("waiting for a discovery response ...")
		assignments, err := stream.WaitForAssignments()
		if err != nil {
			return errors.Wrap(err, "failed to receive a discovery response")
		}
		s.log.Info("received monitoring assignments", "len", len(assignments))
		s.log.V(1).Info("received monitoring assignments", "assignments", assignments)

		s.handler.Handle(assignments, ch)

		if err := stream.ACK(); err != nil {
			if err == io.EOF {
				break
			}
			return errors.Wrap(err, "failed to ACK a discovery response")
		}
	}

	return nil
}
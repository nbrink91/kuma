(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d21ad62"],{bce0:function(t,n,a){"use strict";a.r(n);var e=function(){var t=this,n=t.$createElement,a=t._self._c||n;return a("div",{staticClass:"dataplanes-detail"},[a("YamlView",{attrs:{title:"Entity Overview","has-error":t.hasError,"is-loading":t.isLoading,"is-empty":t.isEmpty,content:t.content}})],1)},i=[],r=a("ff9d"),o={name:"TrafficTraceDetail",metaInfo:{title:"Traffic Trace Details"},components:{YamlView:r["a"]},data:function(){return{content:null,hasError:!1,isLoading:!0,isEmpty:!1}},watch:{$route:function(t,n){this.bootstrap()}},beforeMount:function(){this.bootstrap()},methods:{bootstrap:function(){var t=this,n=this.$route.params.mesh,a=this.$route.params.traffictrace;return this.$api.getTrafficTrace(n,a).then((function(n){n?t.content=n:t.$router.push("/404")})).catch((function(n){t.hasError=!0,console.error(n)})).finally((function(){setTimeout((function(){t.isLoading=!1}),"500")}))}}},s=o,c=a("2877"),u=Object(c["a"])(s,e,i,!1,null,null,null);n["default"]=u.exports}}]);
(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([[8],{vAWX:function(e,t,r){"use strict";var n=r("tAuX"),o=r("g09b");Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0,r("IzEo");var a=o(r("bx4M")),i=o(r("2Taf")),c=o(r("vZ4D")),u=o(r("l4Ni")),d=o(r("ujKo")),l=o(r("MhPg")),s=r("MuoO"),f=n(r("q1tI")),p=r("y1Nh"),v=function(e,t,r,n){var o,a=arguments.length,i=a<3?t:null===n?n=Object.getOwnPropertyDescriptor(t,r):n;if("object"===typeof Reflect&&"function"===typeof Reflect.decorate)i=Reflect.decorate(e,t,r,n);else for(var c=e.length-1;c>=0;c--)(o=e[c])&&(i=(a<3?o(i):a>3?o(t,r,i):o(t,r))||i);return a>3&&i&&Object.defineProperty(t,r,i),i},h=function(e){function t(){var e;return(0,i.default)(this,t),e=(0,u.default)(this,(0,d.default)(t).apply(this,arguments)),e.state={width:"100%"},e.resizeFooterToolbar=function(){requestAnimationFrame(function(){var t=document.querySelectorAll(".ant-layout-sider")[0];if(t){var r="calc(100% - ".concat(t.style.width,")"),n=e.state.width;n!==r&&e.setState({width:r})}})},e}return(0,l.default)(t,e),(0,c.default)(t,[{key:"componentDidMount",value:function(){window.addEventListener("resize",this.resizeFooterToolbar,{passive:!0});this.props.dispatch}},{key:"componentWillUnmount",value:function(){window.removeEventListener("resize",this.resizeFooterToolbar)}},{key:"render",value:function(){return f.default.createElement("div",null,f.default.createElement(p.PageHeaderWrapper,{content:"\u51fa\u73b0\u8fc7\u7684\u6240\u6709\u5546\u6237\u3002"},f.default.createElement(a.default,{title:"\u5546\u6237",bordered:!1})))}}]),t}(f.Component);h=v([(0,s.connect)(function(e){var t=e.record,r=e.creditcard,n=e.business,o=e.loading;return{record:t,creditcard:r,business:n,loading:o.effects["record/fetch"]}})],h);var w=h;t.default=w}}]);
(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([[9],{"8mMf":function(e,t,r){"use strict";var a=r("g09b");Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0;var n=a(r("p0pE")),u=a(r("d6i3"));r("miYZ");var s=a(r("tsqr")),c=r("AkU7"),d={namespace:"merchantList",state:{list:[]},effects:{add:u.default.mark(function e(t,r){var a,n,d,i;return u.default.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return a=t.payload,n=t.callback,d=r.call,r.put,e.next=4,d(c.addMerchant,a);case 4:if(i=e.sent,i){e.next=7;break}return e.abrupt("return");case 7:if(i.success){e.next=10;break}return s.default.error(i.error),e.abrupt("return");case 10:s.default.success("\u6e05\u52a0\u6210\u529f"),n&&n();case 12:case"end":return e.stop()}},e)})},reducers:{saveList:function(e,t){var r=t.payload;return(0,n.default)({},e,r)}}},i=d;t.default=i},AkU7:function(e,t,r){"use strict";var a=r("g09b");Object.defineProperty(t,"__esModule",{value:!0}),t.addMerchant=c;var n=a(r("d6i3")),u=a(r("1l/V")),s=a(r("sy1d"));function c(e){return d.apply(this,arguments)}function d(){return d=(0,u.default)(n.default.mark(function e(t){return n.default.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,s.default)("/merchant",{method:"POST",data:t}));case 1:case"end":return e.stop()}},e)})),d.apply(this,arguments)}}}]);
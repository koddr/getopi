(window.webpackJsonp=window.webpackJsonp||[]).push([[3],{MRB6:function(n,t,o){var e=o("QRet"),u=o("RJJO"),i="undefined"!=typeof window?e.useLayoutEffect:e.useEffect;n.exports=function(){var n=[].slice.call(arguments),t=e.useContext(u);var o=e.useState({});return i((function(){return t.on("@changed",(function(t,e){n.some((function(n){return n in e}))&&o[1]({})}))}),[]),e.useMemo((function(){var o=t.get(),e={};return n.forEach((function(n){e[n]=o[n]})),e.dispatch=t.dispatch,e}),[o[0]])}},QRet:function(n,t,o){"use strict";function e(n){E.options.__h&&E.options.__h(y);var t=y.__H||(y.__H={t:[],u:[]});return n>=t.t.length&&t.t.push({}),t.t[n]}function u(n){return i(g,n)}function i(n,t,o){var u=e(b++);return u.__c||(u.__c=y,u.i=[o?o(t):g(void 0,t),function(t){var o=n(u.i[0],t);u.i[0]!==o&&(u.i[0]=o,u.__c.setState({}))}]),u.i}function r(n,t){var o=e(b++);m(o.o,t)&&(o.i=n,o.o=t,y.__H.u.push(o))}function c(n,t){var o=e(b++);m(o.o,t)&&(o.i=n,o.o=t,y.__h.push(o))}function f(n){return s((function(){return{current:n}}),[])}function a(n,t,o){c((function(){"function"==typeof n?n(t()):n&&(n.current=t())}),null==o?o:o.concat(n))}function s(n,t){var o=e(b++);return m(o.o,t)?(o.o=t,o.v=n,o.i=n()):o.i}function _(n,t){return s((function(){return n}),t)}function l(n){var t=y.context[n.__c];if(!t)return n.__;var o=e(b++);return null==o.i&&(o.i=!0,t.sub(y)),t.props.value}function p(n,t){E.options.useDebugValue&&E.options.useDebugValue(t?t(n):n)}function h(){H.some((function(n){n.__P&&(n.__H.u.forEach(v),n.__H.u.forEach(d),n.__H.u=[])})),H=[]}function v(n){n.m&&n.m()}function d(n){var t=n.i();"function"==typeof t&&(n.m=t)}function m(n,t){return!n||t.some((function(t,o){return t!==n[o]}))}function g(n,t){return"function"==typeof t?t(n):t}o.r(t),o.d(t,"useState",(function(){return u})),o.d(t,"useReducer",(function(){return i})),o.d(t,"useEffect",(function(){return r})),o.d(t,"useLayoutEffect",(function(){return c})),o.d(t,"useRef",(function(){return f})),o.d(t,"useImperativeHandle",(function(){return a})),o.d(t,"useMemo",(function(){return s})),o.d(t,"useCallback",(function(){return _})),o.d(t,"useContext",(function(){return l})),o.d(t,"useDebugValue",(function(){return p}));var b,y,O,E=o("hosL"),H=[],j=E.options.__r,w=E.options.diffed,J=E.options.__c,x=E.options.unmount;E.options.__r=function(n){j&&j(n),b=0,(y=n.__c).__H&&(y.__H.u.forEach(v),y.__H.u.forEach(d),y.__H.u=[])},E.options.diffed=function(n){w&&w(n);var t=n.__c;if(t){var o=t.__H;o&&o.u.length&&(1!==H.push(t)&&O===E.options.requestAnimationFrame||((O=E.options.requestAnimationFrame)||function(n){var t,o=function(){clearTimeout(e),cancelAnimationFrame(t),setTimeout(n)},e=setTimeout(o,100);"undefined"!=typeof window&&(t=requestAnimationFrame(o))})(h))}},E.options.__c=function(n,t){t.some((function(n){n.__h.forEach(v),n.__h=n.__h.filter((function(n){return!n.i||d(n)}))})),J&&J(n,t)},E.options.unmount=function(n){x&&x(n);var t=n.__c;if(t){var o=t.__H;o&&o.t.forEach((function(n){return n.m&&n.m()}))}}},RJJO:function(n,t,o){var e=o("hosL");n.exports=e.createContext()},"Tn/a":function(){},a5dq:function(n){n.exports=function(n,t){n=n||[];var o=(t=t||{}).key||"storeon";return function(t){var e=!1;t.on("@init",(function(){e=!0;try{var n=localStorage.getItem(o);if(null!==n)return JSON.parse(n)}catch(n){}})),t.on("@dispatch",(function(t){if(e){var u={};0===n.length?u=t:n.forEach((function(n){u[n]=t[n]}));try{var i=JSON.stringify(u);localStorage.setItem(o,i)}catch(n){}}}))}}},j0dZ:function(n){function t(){return(t=Object.assign||function(n){for(var t=1;t<arguments.length;t++){var o=arguments[t];for(var e in o)Object.prototype.hasOwnProperty.call(o,e)&&(n[e]=o[e])}return n}).apply(this,arguments)}n.exports=function(n){var o={},e={},u=function n(u,i){if("@dispatch"!==u&&n("@dispatch",[u,i,o[u]]),o[u]){var r,c={};o[u].forEach((function(n){var o=n(e,i);o&&"function"!=typeof o.then&&(r=e=t({},e,o),t(c,o))})),r&&n("@changed",c)}},i={dispatch:u,get:function(){return e},on:function(n,t){return(o[n]||(o[n]=[])).push(t),function(){o[n]=o[n].filter((function(n){return n!==t}))}}};return n.forEach((function(n){n&&n(i)})),u("@init"),i}},m55C:function(){},n6wl:function(n,t,o){"use strict";o.r(t);var e=o("hosL"),u=o("RJJO"),i=o.n(u),r=o("j0dZ"),c=o.n(r),f=o("a5dq"),a=o.n(f),s=c()([n=>{n.on("@init",()=>({email:""})),n.on("login/pre-save/email",(n,t)=>({email:t}))},a()(["email"])]),_=o("MRB6"),l=o.n(_),p=o("Tn/a"),h=o.n(p),v=()=>{var{dispatch:n,email:t}=l()("email");return Object(e.h)("div",{class:h.a.form},Object(e.h)("label",null,"Email"),Object(e.h)("input",{type:"text",value:t,onInput:t=>{n("login/pre-save/email",t.target.value)}}),t)},d=o("m55C"),m=o.n(d);t.default=()=>Object(e.h)(i.a.Provider,{value:s},Object(e.h)("div",{class:m.a.login},Object(e.h)("h2",null,"Login"),Object(e.h)("p",null,"This is the Login component."),Object(e.h)(v,null)))}}]);
//# sourceMappingURL=route-login.chunk.0f3f3.esm.js.map
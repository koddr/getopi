(window.webpackJsonp=window.webpackJsonp||[]).push([[4],{E6Je:function(n,t,e){"use strict";e.r(t);var o=e("hosL"),u=e("RJJO"),r=e.n(u),i=e("j0dZ"),c=e.n(i),a=e("a5dq"),f=e.n(a),s=c()([n=>{n.on("@init",()=>({email:"",name:""})),n.on("register/pre-save/email",(n,t)=>({email:t})),n.on("register/pre-save/name",(n,t)=>({name:t}))},f()(["email","name"])]),l=e("MRB6"),p=e.n(l),_=e("dk4e"),h=e.n(_),v=()=>{var{dispatch:n,email:t,name:e}=p()("email","name");return Object(o.h)("div",{class:h.a.form},Object(o.h)("label",null,"Email"),Object(o.h)("input",{type:"text",value:t,onInput:t=>{n("register/pre-save/email",t.target.value)}}),t,Object(o.h)("br",null),Object(o.h)("label",null,"Name"),Object(o.h)("input",{type:"text",value:e,onInput:t=>{n("register/pre-save/name",t.target.value)}}),e)},m=e("zecI"),d=e.n(m);t.default=()=>Object(o.h)(r.a.Provider,{value:s},Object(o.h)("div",{class:d.a.home},Object(o.h)("h2",null,"Register"),Object(o.h)("p",null,"This is the Home component."),Object(o.h)(v,null)))},MRB6:function(n,t,e){var o=e("QRet"),u=e("RJJO"),r="undefined"!=typeof window?o.useLayoutEffect:o.useEffect;n.exports=function(){var n=[].slice.call(arguments),t=o.useContext(u);var e=o.useState({});return r((function(){return t.on("@changed",(function(t,o){n.some((function(n){return n in o}))&&e[1]({})}))}),[]),o.useMemo((function(){var e=t.get(),o={};return n.forEach((function(n){o[n]=e[n]})),o.dispatch=t.dispatch,o}),[e[0]])}},QRet:function(n,t,e){"use strict";function o(n){E.options.__h&&E.options.__h(O);var t=O.__H||(O.__H={t:[],u:[]});return n>=t.t.length&&t.t.push({}),t.t[n]}function u(n){return r(b,n)}function r(n,t,e){var u=o(g++);return u.__c||(u.__c=O,u.i=[e?e(t):b(void 0,t),function(t){var e=n(u.i[0],t);u.i[0]!==e&&(u.i[0]=e,u.__c.setState({}))}]),u.i}function i(n,t){var e=o(g++);d(e.o,t)&&(e.i=n,e.o=t,O.__H.u.push(e))}function c(n,t){var e=o(g++);d(e.o,t)&&(e.i=n,e.o=t,O.__h.push(e))}function a(n){return s((function(){return{current:n}}),[])}function f(n,t,e){c((function(){"function"==typeof n?n(t()):n&&(n.current=t())}),null==e?e:e.concat(n))}function s(n,t){var e=o(g++);return d(e.o,t)?(e.o=t,e.v=n,e.i=n()):e.i}function l(n,t){return s((function(){return n}),t)}function p(n){var t=O.context[n.__c];if(!t)return n.__;var e=o(g++);return null==e.i&&(e.i=!0,t.sub(O)),t.props.value}function _(n,t){E.options.useDebugValue&&E.options.useDebugValue(t?t(n):n)}function h(){j.some((function(n){n.__P&&(n.__H.u.forEach(v),n.__H.u.forEach(m),n.__H.u=[])})),j=[]}function v(n){n.m&&n.m()}function m(n){var t=n.i();"function"==typeof t&&(n.m=t)}function d(n,t){return!n||t.some((function(t,e){return t!==n[e]}))}function b(n,t){return"function"==typeof t?t(n):t}e.r(t),e.d(t,"useState",(function(){return u})),e.d(t,"useReducer",(function(){return r})),e.d(t,"useEffect",(function(){return i})),e.d(t,"useLayoutEffect",(function(){return c})),e.d(t,"useRef",(function(){return a})),e.d(t,"useImperativeHandle",(function(){return f})),e.d(t,"useMemo",(function(){return s})),e.d(t,"useCallback",(function(){return l})),e.d(t,"useContext",(function(){return p})),e.d(t,"useDebugValue",(function(){return _}));var g,O,y,E=e("hosL"),j=[],H=E.options.__r,w=E.options.diffed,x=E.options.__c,J=E.options.unmount;E.options.__r=function(n){H&&H(n),g=0,(O=n.__c).__H&&(O.__H.u.forEach(v),O.__H.u.forEach(m),O.__H.u=[])},E.options.diffed=function(n){w&&w(n);var t=n.__c;if(t){var e=t.__H;e&&e.u.length&&(1!==j.push(t)&&y===E.options.requestAnimationFrame||((y=E.options.requestAnimationFrame)||function(n){var t,e=function(){clearTimeout(o),cancelAnimationFrame(t),setTimeout(n)},o=setTimeout(e,100);"undefined"!=typeof window&&(t=requestAnimationFrame(e))})(h))}},E.options.__c=function(n,t){t.some((function(n){n.__h.forEach(v),n.__h=n.__h.filter((function(n){return!n.i||m(n)}))})),x&&x(n,t)},E.options.unmount=function(n){J&&J(n);var t=n.__c;if(t){var e=t.__H;e&&e.t.forEach((function(n){return n.m&&n.m()}))}}},RJJO:function(n,t,e){var o=e("hosL");n.exports=o.createContext()},a5dq:function(n){n.exports=function(n,t){n=n||[];var e=(t=t||{}).key||"storeon";return function(t){var o=!1;t.on("@init",(function(){o=!0;try{var n=localStorage.getItem(e);if(null!==n)return JSON.parse(n)}catch(n){}})),t.on("@dispatch",(function(t){if(o){var u={};0===n.length?u=t:n.forEach((function(n){u[n]=t[n]}));try{var r=JSON.stringify(u);localStorage.setItem(e,r)}catch(n){}}}))}}},dk4e:function(){},j0dZ:function(n){function t(){return(t=Object.assign||function(n){for(var t=1;t<arguments.length;t++){var e=arguments[t];for(var o in e)Object.prototype.hasOwnProperty.call(e,o)&&(n[o]=e[o])}return n}).apply(this,arguments)}n.exports=function(n){var e={},o={},u=function n(u,r){if("@dispatch"!==u&&n("@dispatch",[u,r,e[u]]),e[u]){var i,c={};e[u].forEach((function(n){var e=n(o,r);e&&"function"!=typeof e.then&&(i=o=t({},o,e),t(c,e))})),i&&n("@changed",c)}},r={dispatch:u,get:function(){return o},on:function(n,t){return(e[n]||(e[n]=[])).push(t),function(){e[n]=e[n].filter((function(n){return n!==t}))}}};return n.forEach((function(n){n&&n(r)})),u("@init"),r}},zecI:function(n){n.exports={profile:"profile__2vFXq"}}}]);
//# sourceMappingURL=route-register.chunk.0a547.esm.js.map
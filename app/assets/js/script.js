!function(e){var t={};function n(o){if(t[o])return t[o].exports;var r=t[o]={i:o,l:!1,exports:{}};return e[o].call(r.exports,r,r.exports,n),r.l=!0,r.exports}n.m=e,n.c=t,n.d=function(e,t,o){n.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:o})},n.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},n.t=function(e,t){if(1&t&&(e=n(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var o=Object.create(null);if(n.r(o),Object.defineProperty(o,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var r in e)n.d(o,r,function(t){return e[t]}.bind(null,r));return o},n.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return n.d(t,"a",t),t},n.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},n.p="",n(n.s=0)}([function(e,t,n){"use strict";n.r(t);n(1);var o=window;o.requestAnimationFrame=o.requestAnimationFrame||o.mozRequestAnimationFrame||o.webkitRequestAnimationFrame||o.msRequestAnimationFrame||function(e){setTimeout(e,17)};var r=function(e){return(e||o).innerWidth},s=function(e){return(e||o).innerHeight};r(),s(),r(),s(),document.documentElement&&document.documentElement.scrollTop||document.body.scrollTop},function(e,t){var n=function(e,t,n,o){return(e/=o/2)<1?n/2*e*e*e+t:n/2*((e-=2)*e*e+2)+t};!function(){var e=document.querySelectorAll('a[href^="#"]');if(e.length>0)for(var t=0;t<e.length;t++)e[t].addEventListener("click",o,!1);function o(e){e.preventDefault();var t=this.getAttribute("href");t=(t=t.split("#"))[1];var o=document.documentElement.scrollTop||document.body.scrollTop,r=document.getElementById(t).getBoundingClientRect().top+o,s=(document.body.getBoundingClientRect().bottom,window.innerHeight,o||0),c=30,i=0,l=s,u=r-s,a=500/c,d=!0;function m(){d=!1}function f(){window.removeEventListener("mousewheel",m,!1)}window.addEventListener("mousewheel",m,!1),function e(){if(!d)return f(),!1;s=n(i,l,u,a),window.scrollTo(0,s),i<a?(i++,window.setTimeout(e,1e3/c)):f()}()}}();var o=document.querySelector(".nav-bg"),r=document.querySelector(".nav__local"),s=document.querySelector(".nav__global"),c=document.querySelectorAll(".button-menuL"),i=document.querySelectorAll(".button-menuG");if(c.length>0)for(var l=0;l<c.length;l++)c[l].addEventListener("click",a,!1);if(i.length>0)for(l=0;l<i.length;l++)i[l].addEventListener("click",d,!1);var u=document.querySelectorAll(".nav-bg,.nav__close");if(u.length>0)for(l=0;l<u.length;l++)u[l].addEventListener("click",m,!1);function a(){c[0].classList.contains("opened")?m():(r.classList.add("block"),o.classList.add("block"),s.classList.remove("show"),c[0].classList.add("opened"),setTimeout((function(){r.classList.add("show"),o.classList.add("show")}),50),setTimeout((function(){s.classList.remove("block")}),300))}function d(){i[0].classList.contains("opened")?m():(s.classList.add("block"),o.classList.add("block"),r.classList.remove("show"),i[0].classList.add("opened"),setTimeout((function(){s.classList.add("show"),o.classList.add("show")}),50),setTimeout((function(){r.classList.remove("block")}),300))}function m(){r.classList.remove("show"),s.classList.remove("show"),o.classList.remove("show"),setTimeout((function(){r.classList.remove("block"),s.classList.remove("block"),o.classList.remove("block"),c[0].classList.remove("opened"),i[0].classList.remove("opened")}),300)}var f=document.querySelectorAll(".accordion__title");if(f.length>0)for(l=0;l<f.length;l++)f[l].addEventListener("click",(function(e){e.target.parentNode.classList.contains("open")?e.target.parentNode.classList.remove("open"):e.target.parentNode.classList.add("open")}),!1);var v=document.querySelector(".confirm"),L=(document.querySelector(".confirm.show2"),document.querySelectorAll(".modal-open"));if(L.length>0)for(l=0;l<L.length;l++)L[l].addEventListener("click",(function(e){var t=this.dataset.modal;document.getElementById(t).classList.add("show"),setTimeout((function(){document.getElementById(t).classList.add("show2")}),50)}),!1);var h=document.querySelectorAll(".confirm__bg,.confirm__close");if(h.length>0)for(l=0;l<h.length;l++)h[l].addEventListener("click",(function(e){v.classList.remove("show2"),setTimeout((function(){v.classList.remove("show")}),200)}),!1)}]);
//# sourceMappingURL=script.js.map
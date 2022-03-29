"use strict";(self.webpackChunkjuno_docs=self.webpackChunkjuno_docs||[]).push([[332],{3905:function(e,t,r){r.d(t,{Zo:function(){return p},kt:function(){return m}});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function a(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function c(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?a(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):a(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function l(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},a=Object.keys(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var i=n.createContext({}),u=function(e){var t=n.useContext(i),r=t;return e&&(r="function"==typeof e?e(t):c(c({},t),e)),r},p=function(e){var t=u(e.components);return n.createElement(i.Provider,{value:t},e.children)},s={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},f=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,a=e.originalType,i=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),f=u(r),m=o,g=f["".concat(i,".").concat(m)]||f[m]||s[m]||a;return r?n.createElement(g,c(c({ref:t},p),{},{components:r})):n.createElement(g,c({ref:t},p))}));function m(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=r.length,c=new Array(a);c[0]=f;var l={};for(var i in t)hasOwnProperty.call(t,i)&&(l[i]=t[i]);l.originalType=e,l.mdxType="string"==typeof e?e:o,c[1]=l;for(var u=2;u<a;u++)c[u]=r[u];return n.createElement.apply(null,c)}return n.createElement.apply(null,r)}f.displayName="MDXCreateElement"},8335:function(e,t,r){r.r(t),r.d(t,{frontMatter:function(){return l},contentTitle:function(){return i},metadata:function(){return u},toc:function(){return p},default:function(){return f}});var n=r(7462),o=r(3366),a=(r(7294),r(3905)),c=["components"],l={title:"Coverage"},i=void 0,u={unversionedId:"coverage",id:"coverage",title:"Coverage",description:"We use Codecov for coverage",source:"@site/docs/coverage.mdx",sourceDirName:".",slug:"/coverage",permalink:"/coverage",tags:[],version:"current",frontMatter:{title:"Coverage"},sidebar:"docs",previous:{title:"Useful Commands",permalink:"/useful-commands"},next:{title:"testing",permalink:"/testing"}},p=[],s={toc:p};function f(e){var t=e.components,r=(0,o.Z)(e,c);return(0,a.kt)("wrapper",(0,n.Z)({},s,r,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("p",null,"We use ",(0,a.kt)("a",{parentName:"p",href:"https://about.codecov.io/"},"Codecov")," for coverage"),(0,a.kt)("h1",{id:"installation"},"Installation"),(0,a.kt)("p",null,"Install the ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/DemerzelSolutions/courtney"},"courtney")," command-line tool to check configuration validity."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"make install-deps\n")),(0,a.kt)("p",null,"For coverage generation:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"make codecov-test\n")),(0,a.kt)("p",null,"To display locally coverage tests:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"go tool cover -html=coverage/coverage.out\n")))}f.isMDXComponent=!0}}]);
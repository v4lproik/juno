"use strict";(self.webpackChunkjuno_docs=self.webpackChunkjuno_docs||[]).push([[9671],{3905:function(e,n,t){t.d(n,{Zo:function(){return p},kt:function(){return m}});var r=t(7294);function o(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function i(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function l(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?i(Object(t),!0).forEach((function(n){o(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function a(e,n){if(null==e)return{};var t,r,o=function(e,n){if(null==e)return{};var t,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||(o[t]=e[t]);return o}(e,n);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}var c=r.createContext({}),u=function(e){var n=r.useContext(c),t=n;return e&&(t="function"==typeof e?e(n):l(l({},n),e)),t},p=function(e){var n=u(e.components);return r.createElement(c.Provider,{value:n},e.children)},s={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},d=r.forwardRef((function(e,n){var t=e.components,o=e.mdxType,i=e.originalType,c=e.parentName,p=a(e,["components","mdxType","originalType","parentName"]),d=u(t),m=o,f=d["".concat(c,".").concat(m)]||d[m]||s[m]||i;return t?r.createElement(f,l(l({ref:n},p),{},{components:t})):r.createElement(f,l({ref:n},p))}));function m(e,n){var t=arguments,o=n&&n.mdxType;if("string"==typeof e||o){var i=t.length,l=new Array(i);l[0]=d;var a={};for(var c in n)hasOwnProperty.call(n,c)&&(a[c]=n[c]);a.originalType=e,a.mdxType="string"==typeof e?e:o,l[1]=a;for(var u=2;u<i;u++)l[u]=t[u];return r.createElement.apply(null,l)}return r.createElement.apply(null,t)}d.displayName="MDXCreateElement"},9881:function(e,n,t){t.r(n),t.d(n,{frontMatter:function(){return a},contentTitle:function(){return c},metadata:function(){return u},toc:function(){return p},default:function(){return d}});var r=t(7462),o=t(3366),i=(t(7294),t(3905)),l=["components"],a={title:"Welcome",sidebar_position:1},c="Welcome to Juno",u={unversionedId:"intro",id:"intro",title:"Welcome",description:"Let's discover Juno in less than 5 minutes.",source:"@site/docs/intro.md",sourceDirName:".",slug:"/intro",permalink:"/docs/intro",editUrl:"https://github.com/NethermindEth/juno/tree/main/docs/docs/intro.md",tags:[],version:"current",sidebarPosition:1,frontMatter:{title:"Welcome",sidebar_position:1},sidebar:"tutorialSidebar",next:{title:"Getting Started",permalink:"/docs/category/getting-started"}},p=[{value:"What You&#39;ll Need",id:"what-youll-need",children:[{value:"Installing",id:"installing",children:[],level:3}],level:2},{value:"Running Juno",id:"running-juno",children:[{value:"Compiling Directly",id:"compiling-directly",children:[],level:3},{value:"Using Docker",id:"using-docker",children:[],level:3}],level:2}],s={toc:p};function d(e){var n=e.components,t=(0,o.Z)(e,l);return(0,i.kt)("wrapper",(0,r.Z)({},s,t,{components:n,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"welcome-to-juno"},"Welcome to Juno"),(0,i.kt)("p",null,"Let's discover ",(0,i.kt)("strong",{parentName:"p"},"Juno in less than 5 minutes"),"."),(0,i.kt)("h2",{id:"what-youll-need"},"What You'll Need"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://go.dev/doc/install"},"Golang")," version 1.18 for build and run the project."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://www.cairo-lang.org/docs/quickstart.html"},"Cairo-lang")," if you want to do ",(0,i.kt)("inlineCode",{parentName:"li"},"starknet_call")," command.")),(0,i.kt)("h3",{id:"installing"},"Installing"),(0,i.kt)("p",null,"You can get all the dependencies running the next command:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"$ go get -u github.com/NethermindEth/juno \n")),(0,i.kt)("h2",{id:"running-juno"},"Running Juno"),(0,i.kt)("h3",{id:"compiling-directly"},"Compiling Directly"),(0,i.kt)("p",null,"Compile Juno:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"$ make compile\n")),(0,i.kt)("p",null,"Run Juno:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"$ make run\n")),(0,i.kt)("h3",{id:"using-docker"},"Using Docker"),(0,i.kt)("p",null,"If you prefer to use docker, you can follow ",(0,i.kt)("a",{parentName:"p",href:"/docs/running/docker"},"this")," guide."))}d.isMDXComponent=!0}}]);
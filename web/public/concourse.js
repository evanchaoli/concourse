!function(e){"use strict";"function"==typeof define&&define.amd?define(e):"undefined"!=typeof module&&"undefined"!=typeof module.exports?module.exports=e():"undefined"!=typeof Package?Sortable=e():window.Sortable=e()}(function(){"use strict";function e(e,t){this.el=e,this.options=t=v({},t),e[U]=this;var i={group:Math.random(),sort:!0,disabled:!1,store:null,handle:null,scroll:!0,scrollSensitivity:30,scrollSpeed:10,draggable:/[uo]l/i.test(e.nodeName)?"li":">*",ghostClass:"sortable-ghost",ignore:"a, img",filter:null,animation:0,setData:function(e,t){e.setData("Text",t.textContent)},dropBubble:!1,dragoverBubble:!1,dataIdAttr:"data-id",delay:0};for(var o in i)!(o in t)&&(t[o]=i[o]);var r=t.group;r&&"object"==typeof r||(r=t.group={name:r}),["pull","put"].forEach(function(e){e in r||(r[e]=!0)}),t.groups=" "+r.name+(r.put.join?" "+r.put.join(" "):"")+" ";for(var s in this)"_"===s.charAt(0)&&(this[s]=n(this,this[s]));a(e,"mousedown",this._onTapStart),a(e,"touchstart",this._onTapStart),a(e,"dragover",this),a(e,"dragenter",this),X.push(this._onDragOver),t.store&&this.sort(t.store.get(this))}function t(e){_&&_.state!==e&&(l(_,"display",e?"none":""),!e&&_.state&&y.insertBefore(_,m),_.state=e)}function n(e,t){var n=Y.call(arguments,2);return t.bind?t.bind.apply(t,[e].concat(n)):function(){return t.apply(e,n.concat(Y.call(arguments)))}}function i(e,t,n){if(e){n=n||I,t=t.split(".");var i=t.shift().toUpperCase(),o=new RegExp("\\s("+t.join("|")+")\\s","g");do if(">*"===i&&e.parentNode===n||(""===i||e.nodeName.toUpperCase()==i)&&(!t.length||((" "+e.className+" ").match(o)||[]).length==t.length))return e;while(e!==n&&(e=e.parentNode))}return null}function o(e){e.dataTransfer.dropEffect="move",e.preventDefault()}function a(e,t,n){e.addEventListener(t,n,!1)}function r(e,t,n){e.removeEventListener(t,n,!1)}function s(e,t,n){if(e)if(e.classList)e.classList[n?"add":"remove"](t);else{var i=(" "+e.className+" ").replace(P," ").replace(" "+t+" "," ");e.className=(i+(n?" "+t:"")).replace(P," ")}}function l(e,t,n){var i=e&&e.style;if(i){if(void 0===n)return I.defaultView&&I.defaultView.getComputedStyle?n=I.defaultView.getComputedStyle(e,""):e.currentStyle&&(n=e.currentStyle),void 0===t?n:n[t];t in i||(t="-webkit-"+t),i[t]=n+("string"==typeof n?"":"px")}}function u(e,t,n){if(e){var i=e.getElementsByTagName(t),o=0,a=i.length;if(n)for(;a>o;o++)n(i[o],o);return i}return[]}function d(e){e.draggable=!1}function p(){G=!1}function c(e,t){var n=e.lastElementChild,i=n.getBoundingClientRect();return t.clientY-(i.top+i.height)>5&&n}function f(e){for(var t=e.tagName+e.className+e.src+e.href+e.textContent,n=t.length,i=0;n--;)i+=t.charCodeAt(n);return i.toString(36)}function h(e){for(var t=0;e&&(e=e.previousElementSibling);)"TEMPLATE"!==e.nodeName.toUpperCase()&&t++;return t}function g(e,t){var n,i;return function(){void 0===n&&(n=arguments,i=this,setTimeout(function(){1===n.length?e.call(i,n[0]):e.apply(i,n),n=void 0},t))}}function v(e,t){if(e&&t)for(var n in t)t.hasOwnProperty(n)&&(e[n]=t[n]);return e}var m,b,_,y,S,C,D,$,w,x,E,N,T,j,B={},P=/\s+/g,U="Sortable"+(new Date).getTime(),A=window,I=A.document,k=A.parseInt,O=!!("draggable"in I.createElement("div")),G=!1,L=function(e,t,n,i,o,a,r){var s=I.createEvent("Event"),l=(e||t[U]).options,u="on"+n.charAt(0).toUpperCase()+n.substr(1);s.initEvent(n,!0,!0),s.item=i||t,s.from=o||t,s.clone=_,s.oldIndex=a,s.newIndex=r,l[u]&&l[u].call(e,s),t.dispatchEvent(s)},R=Math.abs,Y=[].slice,X=[],M=g(function(e,t,n){if(n&&t.scroll){var i,o,a,r,s=t.scrollSensitivity,l=t.scrollSpeed,u=e.clientX,d=e.clientY,p=window.innerWidth,c=window.innerHeight;if(D!==n&&(C=t.scroll,D=n,C===!0)){C=n;do if(C.offsetWidth<C.scrollWidth||C.offsetHeight<C.scrollHeight)break;while(C=C.parentNode)}C&&(i=C,o=C.getBoundingClientRect(),a=(R(o.right-u)<=s)-(R(o.left-u)<=s),r=(R(o.bottom-d)<=s)-(R(o.top-d)<=s)),a||r||(a=(s>=p-u)-(s>=u),r=(s>=c-d)-(s>=d),(a||r)&&(i=A)),(B.vx!==a||B.vy!==r||B.el!==i)&&(B.el=i,B.vx=a,B.vy=r,clearInterval(B.pid),i&&(B.pid=setInterval(function(){i===A?A.scrollTo(A.pageXOffset+a*l,A.pageYOffset+r*l):(r&&(i.scrollTop+=r*l),a&&(i.scrollLeft+=a*l))},24)))}},30);return e.prototype={constructor:e,_onTapStart:function(e){var t=this,n=this.el,o=this.options,a=e.type,r=e.touches&&e.touches[0],s=(r||e).target,l=s,u=o.filter;if(!("mousedown"===a&&0!==e.button||o.disabled)&&(s=i(s,o.draggable,n))){if(x=h(s),"function"==typeof u){if(u.call(this,e,s,this))return L(t,l,"filter",s,n,x),void e.preventDefault()}else if(u&&(u=u.split(",").some(function(e){return e=i(l,e.trim(),n),e?(L(t,e,"filter",s,n,x),!0):void 0})))return void e.preventDefault();(!o.handle||i(l,o.handle,n))&&this._prepareDragStart(e,r,s)}},_prepareDragStart:function(e,t,n){var i,o=this,r=o.el,s=o.options,l=r.ownerDocument;n&&!m&&n.parentNode===r&&(T=e,y=r,m=n,S=m.nextSibling,N=s.group,i=function(){o._disableDelayedDrag(),m.draggable=!0,s.ignore.split(",").forEach(function(e){u(m,e.trim(),d)}),o._triggerDragStart(t)},a(l,"mouseup",o._onDrop),a(l,"touchend",o._onDrop),a(l,"touchcancel",o._onDrop),s.delay?(a(l,"mousemove",o._disableDelayedDrag),a(l,"touchmove",o._disableDelayedDrag),o._dragStartTimer=setTimeout(i,s.delay)):i())},_disableDelayedDrag:function(){var e=this.el.ownerDocument;clearTimeout(this._dragStartTimer),r(e,"mousemove",this._disableDelayedDrag),r(e,"touchmove",this._disableDelayedDrag)},_triggerDragStart:function(e){e?(T={target:m,clientX:e.clientX,clientY:e.clientY},this._onDragStart(T,"touch")):O?(a(m,"dragend",this),a(y,"dragstart",this._onDragStart)):this._onDragStart(T,!0);try{I.selection?I.selection.empty():window.getSelection().removeAllRanges()}catch(t){}},_dragStarted:function(){y&&m&&(s(m,this.options.ghostClass,!0),e.active=this,L(this,y,"start",m,y,x))},_emulateDragOver:function(){if(j){l(b,"display","none");var e=I.elementFromPoint(j.clientX,j.clientY),t=e,n=" "+this.options.group.name,i=X.length;if(t)do{if(t[U]&&t[U].options.groups.indexOf(n)>-1){for(;i--;)X[i]({clientX:j.clientX,clientY:j.clientY,target:e,rootEl:t});break}e=t}while(t=t.parentNode);l(b,"display","")}},_onTouchMove:function(e){if(T){var t=e.touches?e.touches[0]:e,n=t.clientX-T.clientX,i=t.clientY-T.clientY,o=e.touches?"translate3d("+n+"px,"+i+"px,0)":"translate("+n+"px,"+i+"px)";j=t,l(b,"webkitTransform",o),l(b,"mozTransform",o),l(b,"msTransform",o),l(b,"transform",o),e.preventDefault()}},_onDragStart:function(e,t){var n=e.dataTransfer,i=this.options;if(this._offUpEvents(),"clone"==N.pull&&(_=m.cloneNode(!0),l(_,"display","none"),y.insertBefore(_,m)),t){var o,r=m.getBoundingClientRect(),s=l(m);b=m.cloneNode(!0),l(b,"top",r.top-k(s.marginTop,10)),l(b,"left",r.left-k(s.marginLeft,10)),l(b,"width",r.width),l(b,"height",r.height),l(b,"opacity","0.8"),l(b,"position","fixed"),l(b,"zIndex","100000"),y.appendChild(b),o=b.getBoundingClientRect(),l(b,"width",2*r.width-o.width),l(b,"height",2*r.height-o.height),"touch"===t?(a(I,"touchmove",this._onTouchMove),a(I,"touchend",this._onDrop),a(I,"touchcancel",this._onDrop)):(a(I,"mousemove",this._onTouchMove),a(I,"mouseup",this._onDrop)),this._loopId=setInterval(this._emulateDragOver,150)}else n&&(n.effectAllowed="move",i.setData&&i.setData.call(this,n,m)),a(I,"drop",this);setTimeout(this._dragStarted,0)},_onDragOver:function(e){var n,o,a,r=this.el,s=this.options,u=s.group,d=u.put,f=N===u,h=s.sort;if(void 0!==e.preventDefault&&(e.preventDefault(),!s.dragoverBubble&&e.stopPropagation()),N&&!s.disabled&&(f?h||(a=!y.contains(m)):N.pull&&d&&(N.name===u.name||d.indexOf&&~d.indexOf(N.name)))&&(void 0===e.rootEl||e.rootEl===this.el)){if(M(e,s,this.el),G)return;if(n=i(e.target,s.draggable,r),o=m.getBoundingClientRect(),a)return t(!0),void(_||S?y.insertBefore(m,_||S):h||y.appendChild(m));if(0===r.children.length||r.children[0]===b||r===e.target&&(n=c(r,e))){if(n){if(n.animated)return;v=n.getBoundingClientRect()}t(f),r.appendChild(m),this._animate(o,m),n&&this._animate(v,n)}else if(n&&!n.animated&&n!==m&&void 0!==n.parentNode[U]){$!==n&&($=n,w=l(n));var g,v=n.getBoundingClientRect(),C=v.right-v.left,D=v.bottom-v.top,x=/left|right|inline/.test(w.cssFloat+w.display),E=n.offsetWidth>m.offsetWidth,T=n.offsetHeight>m.offsetHeight,j=(x?(e.clientX-v.left)/C:(e.clientY-v.top)/D)>.5,B=n.nextElementSibling;G=!0,setTimeout(p,30),t(f),g=x?n.previousElementSibling===m&&!E||j&&E:B!==m&&!T||j&&T,g&&!B?r.appendChild(m):n.parentNode.insertBefore(m,g?B:n),this._animate(o,m),this._animate(v,n)}}},_animate:function(e,t){var n=this.options.animation;if(n){var i=t.getBoundingClientRect();l(t,"transition","none"),l(t,"transform","translate3d("+(e.left-i.left)+"px,"+(e.top-i.top)+"px,0)"),t.offsetWidth,l(t,"transition","all "+n+"ms"),l(t,"transform","translate3d(0,0,0)"),clearTimeout(t.animated),t.animated=setTimeout(function(){l(t,"transition",""),l(t,"transform",""),t.animated=!1},n)}},_offUpEvents:function(){var e=this.el.ownerDocument;r(I,"touchmove",this._onTouchMove),r(e,"mouseup",this._onDrop),r(e,"touchend",this._onDrop),r(e,"touchcancel",this._onDrop)},_onDrop:function(t){var n=this.el,i=this.options;clearInterval(this._loopId),clearInterval(B.pid),clearTimeout(this.dragStartTimer),r(I,"drop",this),r(I,"mousemove",this._onTouchMove),r(n,"dragstart",this._onDragStart),this._offUpEvents(),t&&(t.preventDefault(),!i.dropBubble&&t.stopPropagation(),b&&b.parentNode.removeChild(b),m&&(r(m,"dragend",this),d(m),s(m,this.options.ghostClass,!1),y!==m.parentNode?(E=h(m),L(null,m.parentNode,"sort",m,y,x,E),L(this,y,"sort",m,y,x,E),L(null,m.parentNode,"add",m,y,x,E),L(this,y,"remove",m,y,x,E)):(_&&_.parentNode.removeChild(_),m.nextSibling!==S&&(E=h(m),L(this,y,"update",m,y,x,E),L(this,y,"sort",m,y,x,E))),e.active&&L(this,y,"end",m,y,x,E)),y=m=b=S=_=C=D=T=j=$=w=N=e.active=null,this.save())},handleEvent:function(e){var t=e.type;"dragover"===t||"dragenter"===t?m&&(this._onDragOver(e),o(e)):("drop"===t||"dragend"===t)&&this._onDrop(e)},toArray:function(){for(var e,t=[],n=this.el.children,o=0,a=n.length,r=this.options;a>o;o++)e=n[o],i(e,r.draggable,this.el)&&t.push(e.getAttribute(r.dataIdAttr)||f(e));return t},sort:function(e){var t={},n=this.el;this.toArray().forEach(function(e,o){var a=n.children[o];i(a,this.options.draggable,n)&&(t[e]=a)},this),e.forEach(function(e){t[e]&&(n.removeChild(t[e]),n.appendChild(t[e]))})},save:function(){var e=this.options.store;e&&e.set(this)},closest:function(e,t){return i(e,t||this.options.draggable,this.el)},option:function(e,t){var n=this.options;return void 0===t?n[e]:void(n[e]=t)},destroy:function(){var e=this.el;e[U]=null,r(e,"mousedown",this._onTapStart),r(e,"touchstart",this._onTapStart),r(e,"dragover",this),r(e,"dragenter",this),Array.prototype.forEach.call(e.querySelectorAll("[draggable]"),function(e){e.removeAttribute("draggable")}),X.splice(X.indexOf(this._onDragOver),1),this._onDrop(),this.el=e=null}},e.utils={on:a,off:r,css:l,find:u,bind:n,is:function(e,t){return!!i(e,t,e)},extend:v,throttle:g,closest:i,toggleClass:s,index:h},e.version="1.2.0",e.create=function(t,n){return new e(t,n)},e});var concourse={};$(".js-expandable").on("click",function(){$(this).parent().hasClass("expanded")?$(this).parent().removeClass("expanded"):$(this).parent().addClass("expanded")}),concourse.Build=function(e){this.$el=e,this.$abortBtn=this.$el.find(".js-abortBuild"),this.buildID=this.$el.data("build-id"),this.abortEndpoint="/api/v1/builds/"+this.buildID+"/abort"},concourse.Build.prototype.bindEvents=function(){var e=this;this.$abortBtn.on("click",function(t){e.abort()})},concourse.Build.prototype.abort=function(){var e=this;$.ajax({method:"POST",url:e.abortEndpoint}).done(function(t,n){e.$abortBtn.remove()}).error(function(t){e.$abortBtn.addClass("errored")})},$(function(){if($(".js-build").length){var e=new concourse.Build($(".js-build"));e.bindEvents()}}),function(e){e.fn.pausePlayBtn=function(){var t=e(this);return{loading:function(){t.removeClass("disabled enabled").addClass("loading"),t.find("i").removeClass("fa-pause").addClass("fa-circle-o-notch fa-spin")},enable:function(){t.removeClass("loading").addClass("enabled"),t.find("i").removeClass("fa-circle-o-notch fa-spin").addClass("fa-play")},error:function(){t.removeClass("loading").addClass("errored"),t.find("i").removeClass("fa-circle-o-notch fa-spin").addClass("fa-pause")},disable:function(){t.removeClass("loading").addClass("disabled"),t.find("i").removeClass("fa-circle-o-notch fa-spin").addClass("fa-pause")}}}}(jQuery),$(function(){if($(".js-job").length){var e=new concourse.PauseUnpause($(".js-job"));e.bindEvents(),$(".js-build").each(function(e,t){var n,i,o=$(t),a=o.data("status"),r=o.find(".js-build-times"),s=r.data("start-time"),l=r.data("end-time"),u=$("<time>"),d=$("<time>");if(void 0===window.moment)return void console.log("moment library not included, cannot parse durations");if(s>0&&(n=moment.unix(s),u.text(n.fromNow()),u.attr("datetime",n.format()),u.attr("title",n.format("lll Z")),$("<div/>").text("started: ").append(u).appendTo(r)),i=moment.unix(l),d.text(i.fromNow()),d.attr("datetime",i.format()),d.attr("title",i.format("lll Z")),$("<div/>").text(a+": ").append(d).appendTo(r),l>0&&s>0){var p=moment.duration(i.diff(n)),c=$("<span>");c.addClass("duration"),c.text(p.format("h[h]m[m]s[s]")),$("<div/>").text("duration: ").append(c).appendTo(r)}})}}),concourse.PauseUnpause=function(e,t,n){this.$el=e,this.pauseCallback=void 0===t?function(){}:t,this.unpauseCallback=void 0===n?function(){}:n,this.pauseBtn=this.$el.find(".js-pauseUnpause").pausePlayBtn(),this.pauseEndpoint="/api/v1/"+this.$el.data("endpoint")+"/pause",this.unPauseEndpoint="/api/v1/"+this.$el.data("endpoint")+"/unpause"},concourse.PauseUnpause.prototype.bindEvents=function(){var e=this;e.$el.delegate(".js-pauseUnpause.disabled","click",function(t){e.pause()}),e.$el.delegate(".js-pauseUnpause.enabled","click",function(t){e.unpause()})},concourse.PauseUnpause.prototype.pause=function(e){var t=this;t.pauseBtn.loading(),$.ajax({method:"PUT",url:t.pauseEndpoint}).done(function(e,n){t.pauseBtn.enable(),t.pauseCallback()}).error(function(e){t.pauseBtn.error(),401==e.status&&t.redirect("/login")})},concourse.PauseUnpause.prototype.unpause=function(e){var t=this;t.pauseBtn.loading(),$.ajax({method:"PUT",url:this.unPauseEndpoint}).done(function(e){t.pauseBtn.disable(),t.unpauseCallback()}).error(function(e){t.pauseBtn.error(),401==e.status&&t.redirect("/login")})},concourse.PauseUnpause.prototype.redirect=function(e){window.location=e},function(e){concourse.PipelinesNav=function(e){this.$el=$(e),this.$toggle=e.find($(".js-pipelinesNav-toggle")),this.$list=e.find($(".js-pipelinesNav-list")),this.pipelinesEndpoint="/api/v1/pipelines"},concourse.PipelinesNav.prototype.bindEvents=function(){var t=this;t.$toggle.on("click",function(){t.toggle()}),e.create(t.$list[0],{onUpdate:function(){t.onSort()}}),t.loadPipelines()},concourse.PipelinesNav.prototype.onSort=function(){var e=this,t=e.$list.find("a").toArray().map(function(e){return e.innerHTML});$.ajax({method:"PUT",url:e.pipelinesEndpoint+"/ordering",contentType:"application/json",data:JSON.stringify(t)})},concourse.PipelinesNav.prototype.toggle=function(){$("body").toggleClass("pipelinesNav-visible")},concourse.PipelinesNav.prototype.loadPipelines=function(){var e=this;$.ajax({method:"GET",url:e.pipelinesEndpoint}).done(function(t,n){$(t).each(function(t,n){var i=$("<li>"),o=n.paused?"enabled":"disabled",a=n.paused?"play":"pause";i.html('<span class="btn-pause fl '+o+' js-pauseUnpause"><i class="fa fa-fw fa-'+a+'"></i></span><a href="'+n.url+'">'+n.name+"</a>"),i.data("endpoint","pipelines/"+n.name),i.data("pipelineName",n.name),i.addClass("clearfix"),e.$list.append(i),e.newPauseUnpause(i),concourse.pipelineName===n.name&&n.paused&&e.$el.find(".js-groups").addClass("paused")})})},concourse.PipelinesNav.prototype.newPauseUnpause=function(e){var t=this,n=new concourse.PauseUnpause(e,function(){e.data("pipelineName")===concourse.pipelineName&&t.$el.find(".js-groups").addClass("paused")},function(){e.data("pipelineName")===concourse.pipelineName&&t.$el.find(".js-groups").removeClass("paused")});n.bindEvents()}}(Sortable),$(function(){if($(".js-pipelinesNav").length){var e=new concourse.PipelinesNav($(".js-pipelinesNav"));e.bindEvents()}}),$(function(){if($(".js-resource").length){var e=new concourse.PauseUnpause($(".js-resource"),function(){},function(){});e.bindEvents()}}),function(){concourse.StepData=function(e){return this.data=void 0===e?{}:e,this.idCounter=1,this.parallelGroupStore={},this};var e={updateIn:function(e,t){var n,i=jQuery.extend(!0,{},this.data);n=Array.isArray(e)?e.join("."):e.id;var o=i[n];i[n]=t(i[n]);var a=i[n];return o===a?this:new concourse.StepData(i)},getIn:function(e){return Array.isArray(e)?this.data[e.join(".")]:this.data[e.id]},setIn:function(e,t){var n=jQuery.extend(!0,{},this.data);return Array.isArray(e)?n[e.join(".")]=t:n[e.id]=t,new concourse.StepData(n)},forEach:function(e){for(var t in this.data)e(this.data[t])},getSorted:function(){var e=[];for(var t in this.data)e.push([t,this.data[t]]);return e=e.sort(function(e,t){for(var n=e[0].split("."),i=t[0].split("."),o=0;o<n.length;o++){var a=parseInt(n[o]),r=parseInt(i[o]);if(a>r)return 1}return-1}),e=e.map(function(e){return e[1]})},translateLocation:function(e,t){if(!Array.isArray(e))return e;var n,i=0,o=0;if(e.length>1){var a=e.slice(0,e.length-1).join(".");if(void 0===this.parallelGroupStore[a]&&(this.parallelGroupStore[a]=this.idCounter,this.idCounter++),i=this.parallelGroupStore[a],e.length>2){var r=e.slice(0,e.length-2).join(".");o=void 0===this.parallelGroupStore[r]?0:this.parallelGroupStore[r]}}return n=this.idCounter,this.idCounter++,t&&(o=n-1,i=0),{id:n,parallel_group:i,parent_id:o}},getRenderableData:function(){for(var e=this,t=[],n=[],i=e.getSorted(),o=function(e,t,i,o){void 0===n[e]?n[e]=o:n[e].hold&&(o.groupSteps=n[e].groupSteps,o.children=n[e].children,n[e]=o),e>t&&(n[e].groupSteps[s.id]=n[s.id]),0!==t&&e>t&&(n[t].groupSteps[e]=n[e]),0!==i&&(r.isHook()?n[i].children[e]=n[e]:n[i].groupSteps[e]=n[e])},a=0;a<i.length;a++){var r=i[a],s=e.translateLocation(r.origin().location,r.origin().substep),l=r.logs(),u=l.lines,d={key:s.id,step:r,location:s,logLines:u,children:[]};n[s.id]=d,0!==s.parent_id&&void 0===n[s.parent_id]&&(n[s.parent_id]={hold:!0,groupSteps:[],children:[]}),0!==s.parallel_group&&void 0===n[s.parallel_group]&&(n[s.parallel_group]={hold:!0,groupSteps:[],children:[]}),s.serial_group=s.serial_group?s.serial_group:0,0!==s.serial_group&&(renderSerialGroup={serial:!0,step:r,location:s,key:s.serial_group,groupSteps:[],children:[]},o(s.serial_group,s.parallel_group,s.parent_id,renderSerialGroup)),0!==s.parallel_group&&(renderParallelGroup={aggregate:!0,step:r,location:s,key:s.parallel_group,groupSteps:[],children:[]},o(s.parallel_group,s.serial_group,s.parent_id,renderParallelGroup)),0!==s.parallel_group&&(0===s.serial_group||s.serial_group>s.parallel_group)?t[s.parallel_group]=n[s.parallel_group]:0!==s.serial_group?t[s.serial_group]=n[s.serial_group]:(t[s.id]=n[s.id],0!==s.parent_id&&(n[s.parent_id].children[s.id]=n[s.id]))}return t}};concourse.StepData.prototype=e}();
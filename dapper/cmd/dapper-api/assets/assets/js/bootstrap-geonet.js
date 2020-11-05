!function(a){a(function(){a.support.transition=(function(){var b=(function(){var e=document.createElement("bootstrap"),d={WebkitTransition:"webkitTransitionEnd",MozTransition:"transitionend",OTransition:"oTransitionEnd",msTransition:"MSTransitionEnd",transition:"transitionend"},c;for(c in d){if(e.style[c]!==undefined){return d[c]}}}());return b&&{end:b}})()})}(window.jQuery);!function(c){var b='[data-dismiss="alert"]',a=function(d){c(d).on("click",b,this.close)};a.prototype.close=function(i){var h=c(this),f=h.attr("data-target"),g;if(!f){f=h.attr("href");f=f&&f.replace(/.*(?=#[^\s]*$)/,"")}g=c(f);i&&i.preventDefault();g.length||(g=h.hasClass("alert")?h:h.parent());g.trigger(i=c.Event("close"));if(i.isDefaultPrevented()){return}g.removeClass("in");function d(){g.trigger("closed").remove()}c.support.transition&&g.hasClass("fade")?g.on(c.support.transition.end,d):d()};c.fn.alert=function(d){return this.each(function(){var f=c(this),e=f.data("alert");if(!e){f.data("alert",(e=new a(this)))}if(typeof d=="string"){e[d].call(f)}})};c.fn.alert.Constructor=a;c(function(){c("body").on("click.alert.data-api",b,a.prototype.close)})}(window.jQuery);!function(b){var a=function(d,c){this.$element=b(d);this.options=b.extend({},b.fn.button.defaults,c)};a.prototype.setState=function(f){var h="disabled",c=this.$element,e=c.data(),g=c.is("input")?"val":"html";f=f+"Text";e.resetText||c.data("resetText",c[g]());c[g](e[f]||this.options[f]);setTimeout(function(){f=="loadingText"?c.addClass(h).attr(h,h):c.removeClass(h).removeAttr(h)},0)};a.prototype.toggle=function(){var c=this.$element.parent('[data-toggle="buttons-radio"]');c&&c.find(".active").removeClass("active");this.$element.toggleClass("active")};b.fn.button=function(c){return this.each(function(){var f=b(this),e=f.data("button"),d=typeof c=="object"&&c;if(!e){f.data("button",(e=new a(this,d)))}if(c=="toggle"){e.toggle()}else{if(c){e.setState(c)}}})};b.fn.button.defaults={loadingText:"loading..."};b.fn.button.Constructor=a;b(function(){b("body").on("click.button.data-api","[data-toggle^=button]",function(d){var c=b(d.target);if(!c.hasClass("btn")){c=c.closest(".btn")}c.button("toggle")})})}(window.jQuery);!function(a){var b=function(d,c){this.$element=a(d);this.options=c;this.options.slide&&this.slide(this.options.slide);this.options.pause=="hover"&&this.$element.on("mouseenter",a.proxy(this.pause,this)).on("mouseleave",a.proxy(this.cycle,this))};b.prototype={cycle:function(c){if(!c){this.paused=false}this.options.interval&&!this.paused&&(this.interval=setInterval(a.proxy(this.next,this),this.options.interval));return this},to:function(g){var c=this.$element.find(".active"),d=c.parent().children(),e=d.index(c),f=this;if(g>(d.length-1)||g<0){return}if(this.sliding){return this.$element.one("slid",function(){f.to(g)})}if(e==g){return this.pause().cycle()}return this.slide(g>e?"next":"prev",a(d[g]))},pause:function(c){if(!c){this.paused=true}clearInterval(this.interval);this.interval=null;return this},next:function(){if(this.sliding){return}return this.slide("next")},prev:function(){if(this.sliding){return}return this.slide("prev")},slide:function(j,d){var l=this.$element.find(".active"),c=d||l[j](),i=this.interval,k=j=="next"?"left":"right",f=j=="next"?"first":"last",g=this,h=a.Event("slide");this.sliding=true;i&&this.pause();c=c.length?c:this.$element.find(".item")[f]();if(c.hasClass("active")){return}if(a.support.transition&&this.$element.hasClass("slide")){this.$element.trigger(h);if(h.isDefaultPrevented()){return}c.addClass(j);c[0].offsetWidth;l.addClass(k);c.addClass(k);this.$element.one(a.support.transition.end,function(){c.removeClass([j,k].join(" ")).addClass("active");l.removeClass(["active",k].join(" "));g.sliding=false;setTimeout(function(){g.$element.trigger("slid")},0)})}else{this.$element.trigger(h);if(h.isDefaultPrevented()){return}l.removeClass("active");c.addClass("active");this.sliding=false;this.$element.trigger("slid")}i&&this.cycle();return this}};a.fn.carousel=function(c){return this.each(function(){var f=a(this),e=f.data("carousel"),d=a.extend({},a.fn.carousel.defaults,typeof c=="object"&&c);if(!e){f.data("carousel",(e=new b(this,d)))}if(typeof c=="number"){e.to(c)}else{if(typeof c=="string"||(c=d.slide)){e[c]()}else{if(d.interval){e.cycle()}}}})};a.fn.carousel.defaults={interval:5000,pause:"hover"};a.fn.carousel.Constructor=b;a(function(){a("body").on("click.carousel.data-api","[data-slide]",function(h){var g=a(this),d,c=a(g.attr("data-target")||(d=g.attr("href"))&&d.replace(/.*(?=#[^\s]+$)/,"")),f=!c.data("modal")&&a.extend({},c.data(),g.data());c.carousel(f);h.preventDefault()})})}(window.jQuery);!function(a){var b=function(d,c){this.$element=a(d);this.options=a.extend({},a.fn.collapse.defaults,c);if(this.options.parent){this.$parent=a(this.options.parent)}this.options.toggle&&this.toggle()};b.prototype={constructor:b,dimension:function(){var c=this.$element.hasClass("width");return c?"width":"height"},show:function(){var f,c,e,d;if(this.transitioning){return}f=this.dimension();c=a.camelCase(["scroll",f].join("-"));e=this.$parent&&this.$parent.find("> .accordion-group > .in");if(e&&e.length){d=e.data("collapse");if(d&&d.transitioning){return}e.collapse("hide");d||e.data("collapse",null)}this.$element[f](0);this.transition("addClass",a.Event("show"),"shown");this.$element[f](this.$element[0][c])},hide:function(){var c;if(this.transitioning){return}c=this.dimension();this.reset(this.$element[c]());this.transition("removeClass",a.Event("hide"),"hidden");this.$element[c](0)},reset:function(c){var d=this.dimension();this.$element.removeClass("collapse")[d](c||"auto")[0].offsetWidth;this.$element[c!==null?"addClass":"removeClass"]("collapse");return this},transition:function(g,d,e){var f=this,c=function(){if(d.type=="show"){f.reset()}f.transitioning=0;f.$element.trigger(e)};this.$element.trigger(d);if(d.isDefaultPrevented()){return}this.transitioning=1;this.$element[g]("in");a.support.transition&&this.$element.hasClass("collapse")?this.$element.one(a.support.transition.end,c):c()},toggle:function(){this[this.$element.hasClass("in")?"hide":"show"]()}};a.fn.collapse=function(c){return this.each(function(){var f=a(this),e=f.data("collapse"),d=typeof c=="object"&&c;if(!e){f.data("collapse",(e=new b(this,d)))}if(typeof c=="string"){e[c]()}})};a.fn.collapse.defaults={toggle:true};a.fn.collapse.Constructor=b;a(function(){a("body").on("click.collapse.data-api","[data-toggle=collapse]",function(h){var g=a(this),c,f=g.attr("data-target")||h.preventDefault()||(c=g.attr("href"))&&c.replace(/.*(?=#[^\s]+$)/,""),d=a(f).data("collapse")?"toggle":g.data();a(f).collapse(d)})})}(window.jQuery);!function(d){var b='[data-toggle="dropdown"]',a=function(f){var e=d(f).on("click.dropdown.data-api",this.toggle);d("html").on("click.dropdown.data-api",function(){e.parent().removeClass("open")})};a.prototype={constructor:a,toggle:function(j){var i=d(this),h,f,g;if(i.is(".disabled, :disabled")){return}f=i.attr("data-target");if(!f){f=i.attr("href");f=f&&f.replace(/.*(?=#[^\s]*$)/,"")}h=d(f);h.length||(h=i.parent());g=h.hasClass("open");c();if(!g){h.toggleClass("open")}return false}};function c(){d(b).parent().removeClass("open")}d.fn.dropdown=function(e){return this.each(function(){var g=d(this),f=g.data("dropdown");if(!f){g.data("dropdown",(f=new a(this)))}if(typeof e=="string"){f[e].call(g)}})};d.fn.dropdown.Constructor=a;d(function(){d("html").on("click.dropdown.data-api",c);d("body").on("click.dropdown",".dropdown form",function(f){f.stopPropagation()}).on("click.dropdown.data-api",b,a.prototype.toggle)})}(window.jQuery);!function(e){var a=function(i,h){this.options=h;this.$element=e(i).delegate('[data-dismiss="modal"]',"click.dismiss.modal",e.proxy(this.hide,this))};a.prototype={constructor:a,toggle:function(){return this[!this.isShown?"show":"hide"]()},show:function(){var h=this,i=e.Event("show");this.$element.trigger(i);if(this.isShown||i.isDefaultPrevented()){return}e("body").addClass("modal-open");this.isShown=true;d.call(this);c.call(this,function(){var j=e.support.transition&&h.$element.hasClass("fade");h.$element.appendTo(document.body).removeClass("invisible");if(j){h.$element[0].offsetWidth}h.$element.addClass("in");j?h.$element.one(e.support.transition.end,function(){h.$element.trigger("shown")}):h.$element.trigger("shown")})},hide:function(i){i&&i.preventDefault();var h=this;i=e.Event("hide");this.$element.trigger(i);if(!this.isShown||i.isDefaultPrevented()){return}this.isShown=false;e("body").removeClass("modal-open");d.call(this);this.$element.removeClass("in");e.support.transition&&this.$element.hasClass("fade")?g.call(this):f.call(this)}};function g(){var h=this,i=setTimeout(function(){h.$element.off(e.support.transition.end);f.call(h)},500);this.$element.one(e.support.transition.end,function(){clearTimeout(i);f.call(h)})}function f(h){this.$element.addClass("invisible").trigger("hidden");c.call(this)}function c(k){var j=this,i=this.$element.hasClass("fade")?"fade":"";if(this.isShown&&this.options.backdrop){var h=e.support.transition&&i;this.$backdrop=e('<div class="modal-backdrop '+i+'" />').appendTo(document.body);if(this.options.backdrop!="static"){this.$backdrop.click(e.proxy(this.hide,this))}if(h){this.$backdrop[0].offsetWidth}this.$backdrop.addClass("in");h?this.$backdrop.one(e.support.transition.end,k):k()}else{if(!this.isShown&&this.$backdrop){this.$backdrop.removeClass("in");e.support.transition&&this.$element.hasClass("fade")?this.$backdrop.one(e.support.transition.end,e.proxy(b,this)):b.call(this)}else{if(k){k()}}}}function b(){this.$backdrop.remove();this.$backdrop=null}function d(){var h=this;if(this.isShown&&this.options.keyboard){e(document).on("keyup.dismiss.modal",function(i){i.which==27&&h.hide()})}else{if(!this.isShown){e(document).off("keyup.dismiss.modal")}}}e.fn.modal=function(h){return this.each(function(){var k=e(this),j=k.data("modal"),i=e.extend({},e.fn.modal.defaults,k.data(),typeof h=="object"&&h);if(!j){k.data("modal",(j=new a(this,i)))}if(typeof h=="string"){j[h]()}else{if(i.show){j.show()}}})};e.fn.modal.defaults={backdrop:true,keyboard:true,show:true};e.fn.modal.Constructor=a;e(function(){e("body").on("click.modal.data-api",'[data-toggle="modal"]',function(l){var k=e(this),i,h=e(k.attr("data-target")||(i=k.attr("href"))&&i.replace(/.*(?=#[^\s]+$)/,"")),j=h.data("modal")?"toggle":e.extend({},h.data(),k.data());l.preventDefault();h.modal(j)})})}(window.jQuery);!function(b){var a=function(d,c){this.init("tooltip",d,c)};a.prototype={constructor:a,init:function(f,e,d){var g,c;this.type=f;this.$element=b(e);this.options=this.getOptions(d);this.enabled=true;if(this.options.trigger!="manual"){g=this.options.trigger=="hover"?"mouseenter":"focus";c=this.options.trigger=="hover"?"mouseleave":"blur";this.$element.on(g,this.options.selector,b.proxy(this.enter,this));this.$element.on(c,this.options.selector,b.proxy(this.leave,this))}this.options.selector?(this._options=b.extend({},this.options,{trigger:"manual",selector:""})):this.fixTitle()},getOptions:function(c){c=b.extend({},b.fn[this.type].defaults,c,this.$element.data());if(c.delay&&typeof c.delay=="number"){c.delay={show:c.delay,hide:c.delay}}return c},enter:function(d){var c=b(d.currentTarget)[this.type](this._options).data(this.type);if(!c.options.delay||!c.options.delay.show){return c.show()}clearTimeout(this.timeout);c.hoverState="in";this.timeout=setTimeout(function(){if(c.hoverState=="in"){c.show()}},c.options.delay.show)},leave:function(d){var c=b(d.currentTarget)[this.type](this._options).data(this.type);if(this.timeout){clearTimeout(this.timeout)}if(!c.options.delay||!c.options.delay.hide){return c.hide()}c.hoverState="out";this.timeout=setTimeout(function(){if(c.hoverState=="out"){c.hide()}},c.options.delay.hide)},show:function(){var g,c,i,e,h,d,f;if(this.hasContent()&&this.enabled){g=this.tip();this.setContent();if(this.options.animation){g.addClass("fade")}d=typeof this.options.placement=="function"?this.options.placement.call(this,g[0],this.$element[0]):this.options.placement;c=/in/.test(d);g.remove().css({top:0,left:0,display:"block"}).appendTo(c?this.$element:document.body);i=this.getPosition(c);e=g[0].offsetWidth;h=g[0].offsetHeight;switch(c?d.split(" ")[1]:d){case"bottom":f={top:i.top+i.height,left:i.left+i.width/2-e/2};break;case"top":f={top:i.top-h,left:i.left+i.width/2-e/2};break;case"left":f={top:i.top+i.height/2-h/2,left:i.left-e};break;case"right":f={top:i.top+i.height/2-h/2,left:i.left+i.width};break}g.css(f).addClass(d).addClass("in")}},isHTML:function(c){return typeof c!="string"||(c.charAt(0)==="<"&&c.charAt(c.length-1)===">"&&c.length>=3)||/^(?:[^<]*<[\w\W]+>[^>]*$)/.exec(c)},setContent:function(){var d=this.tip(),c=this.getTitle();d.find(".tooltip-inner")[this.isHTML(c)?"html":"text"](c);d.removeClass("fade in top bottom left right")},hide:function(){var c=this,d=this.tip();d.removeClass("in");function e(){var f=setTimeout(function(){d.off(b.support.transition.end).remove()},500);d.one(b.support.transition.end,function(){clearTimeout(f);d.remove()})}b.support.transition&&this.$tip.hasClass("fade")?e():d.remove()},fixTitle:function(){var c=this.$element;if(c.attr("title")||typeof(c.attr("data-original-title"))!="string"){c.attr("data-original-title",c.attr("title")||"").removeAttr("title")}},hasContent:function(){return this.getTitle()},getPosition:function(c){return b.extend({},(c?{top:0,left:0}:this.$element.offset()),{width:this.$element[0].offsetWidth,height:this.$element[0].offsetHeight})},getTitle:function(){var e,c=this.$element,d=this.options;e=c.attr("data-original-title")||(typeof d.title=="function"?d.title.call(c[0]):d.title);return e},tip:function(){return this.$tip=this.$tip||b(this.options.template)},validate:function(){if(!this.$element[0].parentNode){this.hide();this.$element=null;this.options=null}},enable:function(){this.enabled=true},disable:function(){this.enabled=false},toggleEnabled:function(){this.enabled=!this.enabled},toggle:function(){this[this.tip().hasClass("in")?"hide":"show"]()}};b.fn.tooltip=function(c){return this.each(function(){var f=b(this),e=f.data("tooltip"),d=typeof c=="object"&&c;if(!e){f.data("tooltip",(e=new a(this,d)))}if(typeof c=="string"){e[c]()}})};b.fn.tooltip.Constructor=a;b.fn.tooltip.defaults={animation:true,placement:"top",selector:false,template:'<div class="tooltip"><div class="tooltip-arrow"></div><div class="tooltip-inner"></div></div>',trigger:"hover",title:"",delay:0}}(window.jQuery);!function(b){var a=function(d,c){this.init("popover",d,c)};a.prototype=b.extend({},b.fn.tooltip.Constructor.prototype,{constructor:a,setContent:function(){var e=this.tip(),d=this.getTitle(),c=this.getContent();e.find(".popover-title")[this.isHTML(d)?"html":"text"](d);e.find(".popover-content > *")[this.isHTML(c)?"html":"text"](c);e.removeClass("fade top bottom left right in")},hasContent:function(){return this.getTitle()||this.getContent()},getContent:function(){var d,c=this.$element,e=this.options;d=c.attr("data-content")||(typeof e.content=="function"?e.content.call(c[0]):e.content);return d},tip:function(){if(!this.$tip){this.$tip=b(this.options.template)}return this.$tip}});b.fn.popover=function(c){return this.each(function(){var f=b(this),e=f.data("popover"),d=typeof c=="object"&&c;if(!e){f.data("popover",(e=new a(this,d)))}if(typeof c=="string"){e[c]()}})};b.fn.popover.Constructor=a;b.fn.popover.defaults=b.extend({},b.fn.tooltip.defaults,{placement:"right",content:"",template:'<div class="popover"><div class="arrow"></div><div class="popover-inner"><h3 class="popover-title"></h3><div class="popover-content"><p></p></div></div></div>'})}(window.jQuery);!function(b){function a(f,e){var g=b.proxy(this.process,this),c=b(f).is("body")?b(window):b(f),d;this.options=b.extend({},b.fn.scrollspy.defaults,e);this.$scrollElement=c.on("scroll.scroll.data-api",g);this.selector=(this.options.target||((d=b(f).attr("href"))&&d.replace(/.*(?=#[^\s]+$)/,""))||"")+" .nav li > a";this.$body=b("body");this.refresh();this.process()}a.prototype={constructor:a,refresh:function(){var c=this,d;this.offsets=b([]);this.targets=b([]);d=this.$body.find(this.selector).map(function(){var f=b(this),e=f.data("target")||f.attr("href"),g=/^#\w/.test(e)&&b(e);return(g&&e.length&&[[g.position().top,e]])||null}).sort(function(f,e){return f[0]-e[0]}).each(function(){c.offsets.push(this[0]);c.targets.push(this[1])})},process:function(){var h=this.$scrollElement.scrollTop()+this.options.offset,e=this.$scrollElement[0].scrollHeight||this.$body[0].scrollHeight,g=e-this.$scrollElement.height(),f=this.offsets,c=this.targets,j=this.activeTarget,d;if(h>=g){return j!=(d=c.last()[0])&&this.activate(d)}for(d=f.length;d--;){j!=c[d]&&h>=f[d]&&(!f[d+1]||h<=f[d+1])&&this.activate(c[d])}},activate:function(e){var d,c;this.activeTarget=e;b(this.selector).parent(".active").removeClass("active");c=this.selector+'[data-target="'+e+'"],'+this.selector+'[href="'+e+'"]';d=b(c).parent("li").addClass("active");if(d.parent(".dropdown-menu")){d=d.closest("li.dropdown").addClass("active")}d.trigger("activate")}};b.fn.scrollspy=function(c){return this.each(function(){var f=b(this),e=f.data("scrollspy"),d=typeof c=="object"&&c;if(!e){f.data("scrollspy",(e=new a(this,d)))}if(typeof c=="string"){e[c]()}})};b.fn.scrollspy.Constructor=a;b.fn.scrollspy.defaults={offset:10};b(function(){b('[data-spy="scroll"]').each(function(){var c=b(this);c.scrollspy(c.data())})})}(window.jQuery);!function(b){var a=function(c){this.element=b(c)};a.prototype={constructor:a,show:function(){var i=this.element,f=i.closest("ul:not(.dropdown-menu)"),d=i.attr("data-target"),g,c,h;if(!d){d=i.attr("href");d=d&&d.replace(/.*(?=#[^\s]*$)/,"")}if(i.parent("li").hasClass("active")){return}g=f.find(".active a").last()[0];h=b.Event("show",{relatedTarget:g});i.trigger(h);if(h.isDefaultPrevented()){return}c=b(d);this.activate(i.parent("li"),f);this.activate(c,c.parent(),function(){i.trigger({type:"shown",relatedTarget:g})})},activate:function(e,d,h){var c=d.find("> .active"),g=h&&b.support.transition&&c.hasClass("fade");function f(){c.removeClass("active").find("> .dropdown-menu > .active").removeClass("active");e.addClass("active");if(g){e[0].offsetWidth;e.addClass("in")}else{e.removeClass("fade")}if(e.parent(".dropdown-menu")){e.closest("li.dropdown").addClass("active")}h&&h()}g?c.one(b.support.transition.end,f):f();c.removeClass("in")}};b.fn.tab=function(c){return this.each(function(){var e=b(this),d=e.data("tab");if(!d){e.data("tab",(d=new a(this)))}if(typeof c=="string"){d[c]()}})};b.fn.tab.Constructor=a;b(function(){b("body").on("click.tab.data-api",'[data-toggle="tab"], [data-toggle="pill"]',function(c){c.preventDefault();b(this).tab("show")})})}(window.jQuery);!function(a){var b=function(d,c){this.$element=a(d);this.options=a.extend({},a.fn.typeahead.defaults,c);this.matcher=this.options.matcher||this.matcher;this.sorter=this.options.sorter||this.sorter;this.highlighter=this.options.highlighter||this.highlighter;this.updater=this.options.updater||this.updater;this.$menu=a(this.options.menu).appendTo("body");this.source=this.options.source;this.shown=false;this.listen()};b.prototype={constructor:b,select:function(){var c=this.$menu.find(".active").attr("data-value");this.$element.val(this.updater(c)).change();return this.hide()},updater:function(c){return c},show:function(){var c=a.extend({},this.$element.offset(),{height:this.$element[0].offsetHeight});this.$menu.css({top:c.top+c.height,left:c.left});this.$menu.show();this.shown=true;return this},hide:function(){this.$menu.hide();this.shown=false;return this},lookup:function(e){var d=this,c,f;this.query=this.$element.val();if(!this.query){return this.shown?this.hide():this}c=a.grep(this.source,function(g){return d.matcher(g)});c=this.sorter(c);if(!c.length){return this.shown?this.hide():this}return this.render(c.slice(0,this.options.items)).show()},matcher:function(c){return ~c.toLowerCase().indexOf(this.query.toLowerCase())},sorter:function(e){var f=[],d=[],c=[],g;while(g=e.shift()){if(!g.toLowerCase().indexOf(this.query.toLowerCase())){f.push(g)}else{if(~g.indexOf(this.query)){d.push(g)}else{c.push(g)}}}return f.concat(d,c)},highlighter:function(c){var d=this.query.replace(/[\-\[\]{}()*+?.,\\\^$|#\s]/g,"\\$&");return c.replace(new RegExp("("+d+")","ig"),function(e,f){return"<strong>"+f+"</strong>"})},render:function(c){var d=this;c=a(c).map(function(e,f){e=a(d.options.item).attr("data-value",f);e.find("a").html(d.highlighter(f));return e[0]});c.first().addClass("active");this.$menu.html(c);return this},next:function(d){var e=this.$menu.find(".active").removeClass("active"),c=e.next();if(!c.length){c=a(this.$menu.find("li")[0])}c.addClass("active")},prev:function(d){var e=this.$menu.find(".active").removeClass("active"),c=e.prev();if(!c.length){c=this.$menu.find("li").last()}c.addClass("active")},listen:function(){this.$element.on("blur",a.proxy(this.blur,this)).on("keypress",a.proxy(this.keypress,this)).on("keyup",a.proxy(this.keyup,this));if(a.browser.webkit||a.browser.msie){this.$element.on("keydown",a.proxy(this.keypress,this))}this.$menu.on("click",a.proxy(this.click,this)).on("mouseenter","li",a.proxy(this.mouseenter,this))},keyup:function(c){switch(c.keyCode){case 40:case 38:break;case 9:case 13:if(!this.shown){return}this.select();break;case 27:if(!this.shown){return}this.hide();break;default:this.lookup()}c.stopPropagation();c.preventDefault()},keypress:function(c){if(!this.shown){return}switch(c.keyCode){case 9:case 13:case 27:c.preventDefault();break;case 38:if(c.type!="keydown"){break}c.preventDefault();this.prev();break;case 40:if(c.type!="keydown"){break}c.preventDefault();this.next();break}c.stopPropagation()},blur:function(d){var c=this;setTimeout(function(){c.hide()},150)},click:function(c){c.stopPropagation();c.preventDefault();this.select()},mouseenter:function(c){this.$menu.find(".active").removeClass("active");a(c.currentTarget).addClass("active")}};a.fn.typeahead=function(c){return this.each(function(){var f=a(this),e=f.data("typeahead"),d=typeof c=="object"&&c;if(!e){f.data("typeahead",(e=new b(this,d)))}if(typeof c=="string"){e[c]()}})};a.fn.typeahead.defaults={source:[],items:8,menu:'<ul class="typeahead dropdown-menu"></ul>',item:'<li><a href="#"></a></li>'};a.fn.typeahead.Constructor=b;a(function(){a("body").on("focus.typeahead.data-api",'[data-provide="typeahead"]',function(d){var c=a(this);if(c.data("typeahead")){return}d.preventDefault();c.typeahead(c.data())})})}(window.jQuery);jQuery.noConflict();jQuery("document").ready(function(){if(screen.width<=699){jQuery("html, body").animate({scrollTop:jQuery("#heading-anchor").offset().top},1000);jQuery("tr").click(function(){var a=jQuery(this).find("a").attr("href");if(a){window.location=a}});jQuery("div.quake-row").click(function(){var a=jQuery(this).find("a").attr("href");if(a){window.location=a}})}});
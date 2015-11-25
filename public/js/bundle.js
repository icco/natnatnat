/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId])
/******/ 			return installedModules[moduleId].exports;
/******/
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			exports: {},
/******/ 			id: moduleId,
/******/ 			loaded: false
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.loaded = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(0);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	__webpack_require__(1);
	__webpack_require__(2);


/***/ },
/* 1 */
/***/ function(module, exports) {

	// https://github.com/embedly/embedly-jquery
	// https://wordpress.stackexchange.com/questions/15445/is-there-a-built-in-function-to-see-if-a-urlis-oembed-compatible
	$(document).ready(function() {
	  // Default
	  var defaultUrlRe = /(http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?/;
	
	  // Youtube and Vimeo only
	  var videoUrlRe = /https?:\/\/(www\.)?(youtu|vimeo|soundcloud).+/i;
	  $.embedly.defaults = {
	    key: 'be853af8968a408eb7ec368d2019614a',
	    query: {
	      maxwidth: Math.min($(window).width(), 720),
	      words: 20,
	    },
	    secure: true,
	    method: $.noop,
	    urlRe: videoUrlRe
	  };
	
	  $('.markdown a').each(function(i, el) {
	    var url = $(el).attr('href');
	    $.embedly.oembed([url]).done(function(results) {
	      // console.log(results[0]);
	      var addon = $('<div/>', {class: "embed pure-u-1", data: url});
	      if (results[0].error) {
	        return;
	      }
	
	      if (results[0].html != undefined) {
	        addon.append(results[0].html);
	        addon.removeClass('embed');
	      } else {
	        title = $('<h1/>').append(results[0].title);
	        desc = $('<p/>').append(results[0].description);
	        addon.append(title);
	        addon.append(desc);
	      }
	      var prent = $(el).parents('.post').children('.addons');
	      prent.append(addon);
	    });
	  });
	
	  $('li.link').each(function(i, el) {
	    $(el).children('a.adder').click(function() {
	      link = $(el).children('a.actual')[0];
	      title = $(link).text();
	      tags = $(this).data('tags');
	      url =  $(link).attr('href');
	      mkd = "\n[" + title + "](" + url + ")\n\n" + tags;
	      ta = $('textarea[name="text"]');
	      ta.val(ta.val() + mkd);
	    });
	  });
	
	  // Markdown Preview
	  var md_text_name = "textarea[name=text]";
	  if ($(md_text_name).length) {
	    $(md_text_name).bind('input propertychange', function() {
	      jQuery.post('/md', {'text': $(this).val()}, function (data) {
	        $('#rendered').html(data);
	      });
	    });
	  }
	  var title_name = "input[name=title]";
	  if ($(title_name).length) {
	    $(title_name).bind('input propertychange', function() {
	      $('#rendered_title').text($(title_name).val());
	    });
	  }
	
	  $(".slidingDiv").hide();
	  $('.show_hide').click(function() {
	    $("#rendered").slideToggle();
	  });
	
	  // Stats Graph
	  if ($("#statsgraph").length && false) {
	    // Chart dimensions
	    var m = [20, 80, 20, 80]; // Margins
	    var w = 960 - m[1] - m[3];
	    var h = 150 - m[0] - m[2];
	
	    // Scales. Nice functions which auto resize things.
	    // Also defines the ranges for the graph (top and bottom numbers)
	    var x = d3.time.scale().range([0, w]);
	    var y = d3.scale.linear().range([h, 0]);
	
	    var xAxis = d3.svg.axis().scale(x).tickSize(-h).tickSubdivide(true);
	    var yAxis = d3.svg.axis().scale(y).ticks(4).orient("right");
	
	    // An area generator, for the light fill.
	    var area = d3.svg.area()
	      .interpolate("basis")
	      .x(function(d) { return x(d.x); })
	      .y0(h)
	      .y1(function(d) { return y(d.y); });
	
	    var line = d3.svg.line()
	      .interpolate("basis")
	      .x(function(d) { return x(d.x); })
	      .y(function(d) { return y(d.y); });
	
	    var color = "black";
	
	    d3.json("/posts.json", function(error, data) {
	      var values = data.map(function(d) {
	        return { x: new Date(d.date), y: +d.id };
	      });
	
	      // Compute the minimum and maximum date, and the maximum y value.
	      x.domain([new Date(data[0].date), new Date(data[data.length - 1].date)]);
	      y.domain([0, d3.max(values, function(d) { return d.y; })]).nice();
	
	      // Add an SVG element with the desired dimensions and margin.
	      var svg = d3.select("#statsgraph").append("svg:svg")
	        .attr("width", w + m[1] + m[3])
	        .attr("height", h + m[0] + m[2])
	        .append("svg:g")
	        .attr("transform", "translate(" + m[3] + "," + m[0] + ")");
	
	      // TODO(icco): get this to work.
	      var barPadding = 1;
	      svg.selectAll("rect")
	        .data(values)
	        .enter()
	        .append("rect")
	        .attr("fill", color)
	        .attr("x", function(d) { return x(d.x); })
	        .attr("y", function(d) { return y(d.y); })
	        .attr("width", 1)
	        .attr("height", function(d) { return h - y(d.y); });
	
	      // Add the clip path.
	      svg.append("svg:clipPath")
	        .attr("id", "clip")
	        .append("svg:rect")
	        .attr("width", w)
	        .attr("height", h);
	
	      // Add the x-axis.
	      svg.append("svg:g")
	        .attr("class", "x axis")
	        .attr("transform", "translate(0," + h + ")")
	        .call(xAxis);
	
	      // Add the y-axis.
	      svg.append("svg:g")
	        .attr("class", "y axis")
	        .attr("transform", "translate(" + w + ",0)")
	        .call(yAxis);
	
	      // Add a small label for the name.
	      svg.append("svg:text")
	        .attr("x", w - 6)
	        .attr("y", m[0] - 12)
	        .attr("text-anchor", "end")
	        .text("commits/day");
	    });
	  }
	
	  // Change the time elements to be relative times.
	  // Requires moment.js
	  $('time').each(function(i, el) {
	    var t = moment($(el).attr('datetime'), "YYYY-MM-DDTHH:mm:ss.SSSSSSZ");
	    $(el).text(t.fromNow());
	    $(el).attr("title", t.format());
	  });
	
	  $('#links h2').each(function(i, el) {
	    var t = moment($(el).text(), "YYYY-MM-DD HH:mm:ssZ");
	    $(el).text(t.utc().format("MMMM Do YYYY"));
	  });
	
	  // https://highlightjs.org/usage/
	  hljs.initHighlightingOnLoad();
	});


/***/ },
/* 2 */
/***/ function(module, exports, __webpack_require__) {

	// style-loader: Adds some css to the DOM by adding a <style> tag
	
	// load the styles
	var content = __webpack_require__(3);
	if(typeof content === 'string') content = [[module.id, content, '']];
	// add the styles to the DOM
	var update = __webpack_require__(5)(content, {});
	if(content.locals) module.exports = content.locals;
	// Hot Module Replacement
	if(false) {
		// When the styles change, update the <style> tags
		if(!content.locals) {
			module.hot.accept("!!./../../node_modules/css-loader/index.js?sourceMap!./../../node_modules/sass-loader/index.js?sourceMap!./style.scss", function() {
				var newContent = require("!!./../../node_modules/css-loader/index.js?sourceMap!./../../node_modules/sass-loader/index.js?sourceMap!./style.scss");
				if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
				update(newContent);
			});
		}
		// When the module is disposed, remove the <style> tags
		module.hot.dispose(function() { update(); });
	}

/***/ },
/* 3 */
/***/ function(module, exports, __webpack_require__) {

	exports = module.exports = __webpack_require__(4)();
	// imports
	
	
	// module
	exports.push([module.id, "/* This is broken in Firefox.\np {\n  -ms-word-break: break-all;\n  word-break: break-all;\n\n  // Non standard for webkit.\n  word-break: break-word;\n\n  -webkit-hyphens: auto;\n  -moz-hyphens: auto;\n  -ms-hyphens: auto;\n  hyphens: auto;\n}\n*/\na, a:link, a:visited {\n  transition: color .4s;\n  color: #265C83; }\n\nh1 a {\n  text-decoration: none; }\n\nh1 {\n  line-height: 1.3; }\n\na:hover {\n  color: #7FDBFF; }\n\na:active {\n  transition: color .3s;\n  color: #007BE6; }\n\n.show_hide {\n  cursor: pointer; }\n\n.pager {\n  padding-left: 0;\n  margin: 20px 0;\n  list-style: none;\n  text-align: center; }\n  .pager li {\n    display: inline; }\n    .pager li > a,\n    .pager li > span {\n      display: inline-block;\n      padding: 5px 14px; }\n  .pager .next > a,\n  .pager .next > span {\n    float: right; }\n  .pager .previous > a,\n  .pager .previous > span {\n    float: left; }\n  .pager .disabled > a,\n  .pager .disabled > a:hover,\n  .pager .disabled > a:focus,\n  .pager .disabled > span {\n    cursor: not-allowed; }\n\n.post .addons h1 {\n  margin: 0; }\n\n.post .addons .embed {\n  padding: 10px; }\n\n.pure-form textarea {\n  resize: vertical; }\n\n.pure-form label[for=\"option-one\"] {\n  float: right; }\n\n.markdown img {\n  max-width: 100%; }\n\nli.link a {\n  cursor: pointer; }\n\n.tag-header {\n  margin-top: 30px; }\n\nul.inline {\n  display: inline;\n  list-style-type: none;\n  margin: 0;\n  padding: 0; }\n  ul.inline li {\n    text-align: left;\n    display: inline;\n    list-style: none;\n    display: inline;\n    margin: 0; }\n  ul.inline li:after {\n    content: \", \"; }\n  ul.inline li:last-child:after {\n    content: \"\"; }\n  ul.inline li:nth-last-child(2):after {\n    content: \" & \"; }\n", "", {"version":3,"sources":["/./src/scss/src/scss/style.scss"],"names":[],"mappings":"AACA;;;;;;;;;;;;;EAaE;AAGF;EACE,sBAAsB;EACtB,eAAe,EAChB;;AAED;EACE,sBAAsB,EACvB;;AAED;EACE,iBAAiB,EAClB;;AAED;EACE,eAAe,EAChB;;AAED;EACE,sBAAsB;EACtB,eAAe,EAChB;;AAED;EACE,gBAAgB,EACjB;;AAED;EACE,gBAAgB;EAChB,eAAe;EACf,iBAAiB;EACjB,mBAAmB,EAqCpB;EAzCD;IAOI,gBAAgB,EAUjB;IAjBH;;MAUM,sBAAsB;MACtB,kBAAkB,EACnB;EAZL;;IAsBM,aAAa,EACd;EAvBL;;IA6BM,YAAY,EACb;EA9BL;;;;IAsCM,oBAAoB,EACrB;;AAIL;EAGM,UAAU,EACX;;AAJL;EAOM,cAAc,EACf;;AAIL;EAEI,iBAAiB,EAClB;;AAHH;EAMI,aAAa,EACd;;AAGH;EAEI,gBAAgB,EACjB;;AAGH;EAEI,gBAAgB,EACjB;;AAGH;EACE,iBAAiB,EAClB;;AAED;EAqBE,gBAAgB;EAChB,sBAAsB;EACtB,UAAU;EACV,WAAW,EACZ;EAzBD;IAEI,iBAAiB;IACjB,gBAAgB;IAChB,iBAAiB;IACjB,gBAAgB;IAChB,UAAU,EACX;EAPH;IAUI,cAAc,EACf;EAXH;IAcI,YAAY,EACb;EAfH;IAkBI,eAAe,EAChB","file":"style.scss","sourcesContent":["// Enable word-breaking for weird text.\n/* This is broken in Firefox.\np {\n  -ms-word-break: break-all;\n  word-break: break-all;\n\n  // Non standard for webkit.\n  word-break: break-word;\n\n  -webkit-hyphens: auto;\n  -moz-hyphens: auto;\n  -ms-hyphens: auto;\n  hyphens: auto;\n}\n*/\n\n// Based on http://mrmrs.io/links/\na, a:link, a:visited {\n  transition: color .4s;\n  color: #265C83;\n}\n\nh1 a {\n  text-decoration: none;\n}\n\nh1 {\n  line-height: 1.3;\n}\n\na:hover {\n  color: #7FDBFF;\n}\n\na:active {\n  transition: color .3s;\n  color: #007BE6;\n}\n\n.show_hide {\n  cursor: pointer;\n}\n\n.pager {\n  padding-left: 0;\n  margin: 20px 0;\n  list-style: none;\n  text-align: center;\n\n  li {\n    display: inline;\n    > a,\n    > span {\n      display: inline-block;\n      padding: 5px 14px;\n    }\n\n    > a:hover,\n    > a:focus {\n    }\n  }\n\n  .next {\n    > a,\n    > span {\n      float: right;\n    }\n  }\n\n  .previous {\n    > a,\n    > span {\n      float: left;\n    }\n  }\n\n  .disabled {\n    > a,\n    > a:hover,\n    > a:focus,\n    > span {\n      cursor: not-allowed;\n    }\n  }\n}\n\n.post {\n  .addons {\n    h1 {\n      margin: 0;\n    }\n\n    .embed {\n      padding: 10px;\n    }\n  }\n}\n\n.pure-form {\n  textarea {\n    resize: vertical;\n  }\n\n  label[for=\"option-one\"] {\n    float: right;\n  }\n}\n\n.markdown {\n  img {\n    max-width: 100%;\n  }\n}\n\nli.link {\n  a {\n    cursor: pointer;\n  }\n}\n\n.tag-header {\n  margin-top: 30px;\n}\n\nul.inline {\n  li  {\n    text-align: left;\n    display: inline;\n    list-style: none;\n    display: inline;\n    margin: 0;\n  }\n\n  li:after {\n    content: \", \";\n  }\n\n  li:last-child:after {\n    content: \"\";\n  }\n\n  li:nth-last-child(2):after {\n    content: \" & \";\n  }\n\n  display: inline;\n  list-style-type: none;\n  margin: 0;\n  padding: 0;\n}\n"],"sourceRoot":"webpack://"}]);
	
	// exports


/***/ },
/* 4 */
/***/ function(module, exports) {

	/*
		MIT License http://www.opensource.org/licenses/mit-license.php
		Author Tobias Koppers @sokra
	*/
	// css base code, injected by the css-loader
	module.exports = function() {
		var list = [];
	
		// return the list of modules as css string
		list.toString = function toString() {
			var result = [];
			for(var i = 0; i < this.length; i++) {
				var item = this[i];
				if(item[2]) {
					result.push("@media " + item[2] + "{" + item[1] + "}");
				} else {
					result.push(item[1]);
				}
			}
			return result.join("");
		};
	
		// import a list of modules into the list
		list.i = function(modules, mediaQuery) {
			if(typeof modules === "string")
				modules = [[null, modules, ""]];
			var alreadyImportedModules = {};
			for(var i = 0; i < this.length; i++) {
				var id = this[i][0];
				if(typeof id === "number")
					alreadyImportedModules[id] = true;
			}
			for(i = 0; i < modules.length; i++) {
				var item = modules[i];
				// skip already imported module
				// this implementation is not 100% perfect for weird media query combinations
				//  when a module is imported multiple times with different media queries.
				//  I hope this will never occur (Hey this way we have smaller bundles)
				if(typeof item[0] !== "number" || !alreadyImportedModules[item[0]]) {
					if(mediaQuery && !item[2]) {
						item[2] = mediaQuery;
					} else if(mediaQuery) {
						item[2] = "(" + item[2] + ") and (" + mediaQuery + ")";
					}
					list.push(item);
				}
			}
		};
		return list;
	};


/***/ },
/* 5 */
/***/ function(module, exports, __webpack_require__) {

	/*
		MIT License http://www.opensource.org/licenses/mit-license.php
		Author Tobias Koppers @sokra
	*/
	var stylesInDom = {},
		memoize = function(fn) {
			var memo;
			return function () {
				if (typeof memo === "undefined") memo = fn.apply(this, arguments);
				return memo;
			};
		},
		isOldIE = memoize(function() {
			return /msie [6-9]\b/.test(window.navigator.userAgent.toLowerCase());
		}),
		getHeadElement = memoize(function () {
			return document.head || document.getElementsByTagName("head")[0];
		}),
		singletonElement = null,
		singletonCounter = 0,
		styleElementsInsertedAtTop = [];
	
	module.exports = function(list, options) {
		if(false) {
			if(typeof document !== "object") throw new Error("The style-loader cannot be used in a non-browser environment");
		}
	
		options = options || {};
		// Force single-tag solution on IE6-9, which has a hard limit on the # of <style>
		// tags it will allow on a page
		if (typeof options.singleton === "undefined") options.singleton = isOldIE();
	
		// By default, add <style> tags to the bottom of <head>.
		if (typeof options.insertAt === "undefined") options.insertAt = "bottom";
	
		var styles = listToStyles(list);
		addStylesToDom(styles, options);
	
		return function update(newList) {
			var mayRemove = [];
			for(var i = 0; i < styles.length; i++) {
				var item = styles[i];
				var domStyle = stylesInDom[item.id];
				domStyle.refs--;
				mayRemove.push(domStyle);
			}
			if(newList) {
				var newStyles = listToStyles(newList);
				addStylesToDom(newStyles, options);
			}
			for(var i = 0; i < mayRemove.length; i++) {
				var domStyle = mayRemove[i];
				if(domStyle.refs === 0) {
					for(var j = 0; j < domStyle.parts.length; j++)
						domStyle.parts[j]();
					delete stylesInDom[domStyle.id];
				}
			}
		};
	}
	
	function addStylesToDom(styles, options) {
		for(var i = 0; i < styles.length; i++) {
			var item = styles[i];
			var domStyle = stylesInDom[item.id];
			if(domStyle) {
				domStyle.refs++;
				for(var j = 0; j < domStyle.parts.length; j++) {
					domStyle.parts[j](item.parts[j]);
				}
				for(; j < item.parts.length; j++) {
					domStyle.parts.push(addStyle(item.parts[j], options));
				}
			} else {
				var parts = [];
				for(var j = 0; j < item.parts.length; j++) {
					parts.push(addStyle(item.parts[j], options));
				}
				stylesInDom[item.id] = {id: item.id, refs: 1, parts: parts};
			}
		}
	}
	
	function listToStyles(list) {
		var styles = [];
		var newStyles = {};
		for(var i = 0; i < list.length; i++) {
			var item = list[i];
			var id = item[0];
			var css = item[1];
			var media = item[2];
			var sourceMap = item[3];
			var part = {css: css, media: media, sourceMap: sourceMap};
			if(!newStyles[id])
				styles.push(newStyles[id] = {id: id, parts: [part]});
			else
				newStyles[id].parts.push(part);
		}
		return styles;
	}
	
	function insertStyleElement(options, styleElement) {
		var head = getHeadElement();
		var lastStyleElementInsertedAtTop = styleElementsInsertedAtTop[styleElementsInsertedAtTop.length - 1];
		if (options.insertAt === "top") {
			if(!lastStyleElementInsertedAtTop) {
				head.insertBefore(styleElement, head.firstChild);
			} else if(lastStyleElementInsertedAtTop.nextSibling) {
				head.insertBefore(styleElement, lastStyleElementInsertedAtTop.nextSibling);
			} else {
				head.appendChild(styleElement);
			}
			styleElementsInsertedAtTop.push(styleElement);
		} else if (options.insertAt === "bottom") {
			head.appendChild(styleElement);
		} else {
			throw new Error("Invalid value for parameter 'insertAt'. Must be 'top' or 'bottom'.");
		}
	}
	
	function removeStyleElement(styleElement) {
		styleElement.parentNode.removeChild(styleElement);
		var idx = styleElementsInsertedAtTop.indexOf(styleElement);
		if(idx >= 0) {
			styleElementsInsertedAtTop.splice(idx, 1);
		}
	}
	
	function createStyleElement(options) {
		var styleElement = document.createElement("style");
		styleElement.type = "text/css";
		insertStyleElement(options, styleElement);
		return styleElement;
	}
	
	function createLinkElement(options) {
		var linkElement = document.createElement("link");
		linkElement.rel = "stylesheet";
		insertStyleElement(options, linkElement);
		return linkElement;
	}
	
	function addStyle(obj, options) {
		var styleElement, update, remove;
	
		if (options.singleton) {
			var styleIndex = singletonCounter++;
			styleElement = singletonElement || (singletonElement = createStyleElement(options));
			update = applyToSingletonTag.bind(null, styleElement, styleIndex, false);
			remove = applyToSingletonTag.bind(null, styleElement, styleIndex, true);
		} else if(obj.sourceMap &&
			typeof URL === "function" &&
			typeof URL.createObjectURL === "function" &&
			typeof URL.revokeObjectURL === "function" &&
			typeof Blob === "function" &&
			typeof btoa === "function") {
			styleElement = createLinkElement(options);
			update = updateLink.bind(null, styleElement);
			remove = function() {
				removeStyleElement(styleElement);
				if(styleElement.href)
					URL.revokeObjectURL(styleElement.href);
			};
		} else {
			styleElement = createStyleElement(options);
			update = applyToTag.bind(null, styleElement);
			remove = function() {
				removeStyleElement(styleElement);
			};
		}
	
		update(obj);
	
		return function updateStyle(newObj) {
			if(newObj) {
				if(newObj.css === obj.css && newObj.media === obj.media && newObj.sourceMap === obj.sourceMap)
					return;
				update(obj = newObj);
			} else {
				remove();
			}
		};
	}
	
	var replaceText = (function () {
		var textStore = [];
	
		return function (index, replacement) {
			textStore[index] = replacement;
			return textStore.filter(Boolean).join('\n');
		};
	})();
	
	function applyToSingletonTag(styleElement, index, remove, obj) {
		var css = remove ? "" : obj.css;
	
		if (styleElement.styleSheet) {
			styleElement.styleSheet.cssText = replaceText(index, css);
		} else {
			var cssNode = document.createTextNode(css);
			var childNodes = styleElement.childNodes;
			if (childNodes[index]) styleElement.removeChild(childNodes[index]);
			if (childNodes.length) {
				styleElement.insertBefore(cssNode, childNodes[index]);
			} else {
				styleElement.appendChild(cssNode);
			}
		}
	}
	
	function applyToTag(styleElement, obj) {
		var css = obj.css;
		var media = obj.media;
		var sourceMap = obj.sourceMap;
	
		if(media) {
			styleElement.setAttribute("media", media)
		}
	
		if(styleElement.styleSheet) {
			styleElement.styleSheet.cssText = css;
		} else {
			while(styleElement.firstChild) {
				styleElement.removeChild(styleElement.firstChild);
			}
			styleElement.appendChild(document.createTextNode(css));
		}
	}
	
	function updateLink(linkElement, obj) {
		var css = obj.css;
		var media = obj.media;
		var sourceMap = obj.sourceMap;
	
		if(sourceMap) {
			// http://stackoverflow.com/a/26603875
			css += "\n/*# sourceMappingURL=data:application/json;base64," + btoa(unescape(encodeURIComponent(JSON.stringify(sourceMap)))) + " */";
		}
	
		var blob = new Blob([css], { type: "text/css" });
	
		var oldSrc = linkElement.href;
	
		linkElement.href = URL.createObjectURL(blob);
	
		if(oldSrc)
			URL.revokeObjectURL(oldSrc);
	}


/***/ }
/******/ ]);
//# sourceMappingURL=bundle.js.map
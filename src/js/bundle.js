/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};

/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {

/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId])
/******/ 			return installedModules[moduleId].exports;

/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			exports: {},
/******/ 			id: moduleId,
/******/ 			loaded: false
/******/ 		};

/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);

/******/ 		// Flag the module as loaded
/******/ 		module.loaded = true;

/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}


/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;

/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;

/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";

/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(0);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
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


/***/ }
/******/ ]);
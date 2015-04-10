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


  // Stats Graph
  if ($("#statsgraph").length) {
    d3.json("/posts.json", function(error, data) {
      var w = 500;
      var h = 500;
      var canvas = d3.select("#statsgraph")
        .append("svg")
        .attr("width", w)
        .attr("height", h)
        .attr("border", "black");

      var group = canvas.append("g")
        .attr("transform", "translate(100,10)");

      var x = d3.time.scale().range([0, w]);
      var y = d3.scale.linear().range([0, h]);

      var line = d3.svg.line()
        .interpolate("basis")
        .x(function(d) { return console.log(d); x(new Date(d.date)); })
        .y(function(d) { return y(d.id); }); 

      group.selectAll("path")
        .data(data).enter()
        .append("path")
        .attr("d", function(d) { console.log(y(d.id), y(new Date(d.date)), line(d)); return line(d); })
        .attr("fill", "none")
        .attr("stroke", "green")
        .attr("stroke-width", 3);
    });
  }
});

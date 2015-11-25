// Parses links and turns them into embeds.
//
// https://github.com/embedly/embedly-jquery
// https://wordpress.stackexchange.com/questions/15445/is-there-a-built-in-function-to-see-if-a-urlis-oembed-compatible
var Embedly = function() {
  var $ = require('jquery');

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
};
module.exports = Embedly;

// Takes a list of links and makes it so when you click them the appear as
// markdown in a textbox.
var LinkAdder = function() {
  var $ = require('jquery');
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
};
module.exports = LinkAdder;

// Markdown Preview generator.
var MarkdownPreview = function() {
  var $ = require('jquery');
  var md_text_name = "textarea[name=text]";
  if ($(md_text_name).length) {
    $(md_text_name).bind('input propertychange', function() {
      $.post('/md', {'text': $(this).val()}, function (data) {
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
};
module.exports = MarkdownPreview;

// Change the time elements to be relative times.
// Requires moment.js
var RelativeTimes = function() {
  var $ = require('jquery');
  var moment = require('moment');
  $('time').each(function(i, el) {
    var t = moment($(el).attr('datetime'), "YYYY-MM-DDTHH:mm:ss.SSSSSSZ");
    $(el).text(t.fromNow());
    $(el).attr("title", t.format());
  });

  $('#links h2').each(function(i, el) {
    var t = moment($(el).text(), "YYYY-MM-DD HH:mm:ssZ");
    $(el).text(t.utc().format("MMMM Do YYYY"));
  });
};
module.exports = RelativeTimes;

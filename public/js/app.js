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
});

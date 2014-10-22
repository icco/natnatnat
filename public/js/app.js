$.embedly.defaults = {
  key: 'be853af8968a408eb7ec368d2019614a',
  query: {
    maxwidth: 480,
    words: 20,
  },
  secure: true,
  method: $.noop
};

$('.markdown a').each(function(i, el) {
  var url = $(el).attr('href');
  $.embedly.oembed([url]).done(function(results) {
    console.log(results[0]);
    var addon = $('<div/>', {class: "embed pure-u-md-1-2 pure-u-1", data: url});
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

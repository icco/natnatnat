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

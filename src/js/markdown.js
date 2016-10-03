// Markdown Preview generator.
$(document).ready(function() {
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
});

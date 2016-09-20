// Takes a list of links and makes it so when you click them the appear as
// markdown in a textbox.
$(document).ready(function() {
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

// Doesn't let you close forms that have been edited.
var formHasChanged = false;
var submitted = false;

$('input, select, textarea').bind('input propertychange', function() {
  // console.log("form change");
  formHasChanged = true;
});

window.onbeforeunload = function (e) {
  if (formHasChanged && !submitted) {
    var message = "You have not saved your changes.", e = e || window.event;
    if (e) {
      e.returnValue = message;
    }
    return message;
  }
}

$("form").submit(function() {
  submitted = true;
});

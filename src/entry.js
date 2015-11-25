window.onload = function() {
  var time = require('./js/time.js');
  time();

  var embeds = require('./js/embeds.js');
  embeds();

  var links = require('./js/link_adder.js');
  links();

  var md = require('./js/markdown.js');
  md();
};

// https://highlightjs.org/usage/
var hljs = require('highlight.js');
// highlight css from https://github.com/isagalaev/highlight.js/blob/658226e69491f027b90ad55a73c9c8c4d6c4765b/src/styles/github.css
require('./scss/github.css');
hljs.initHighlightingOnLoad();

require('./scss/tachyons.css');
require('./scss/style.scss');
require('./img/natwelchlogo.png');

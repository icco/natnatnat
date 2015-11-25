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
hljs.initHighlightingOnLoad();

require('./scss/tachyons.css');
require('./scss/style.scss');
require('./img/natwelchlogo.png');

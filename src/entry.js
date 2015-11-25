var app = require('./js/app.js');
require('./scss/tachyons.css');
require('./scss/style.scss');
require('./img/natwelchlogo.png');

app();

// https://highlightjs.org/usage/
var hljs = require('highlight.js');
hljs.initHighlightingOnLoad();

<!DOCTYPE html>
<html>
  <head>
    <title>Nat? Nat. Nat! {{.Title}}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="/css/pure-min.0.5.0.css" />
    <link rel="stylesheet" href="/css/style.css" />

    <!-- same shit, old and new -->
    <link rel="webmention" href="/mention" />
    <link rel="http://webmention.org/" href="/mention" />
  </head>
  <body>
    <div class="wrap">
      <div class="pure-g">
        <div class="pure-u-1-12">
          <a href="/"><img src="/img/natwelchlogo2.png" class="pure-img" alt="Nat? Nat. Nat!"></a>
        </div>
        <div class="pure-u-11-12">
          <div class="menu">
            <ul>
              <li><a href="/about">About</a></li>
              <li><a href="/archives">Archives</a></li>
              <li><a href="/stats">Stats</a></li>

              {{if .IsAdmin}}
              <li><a href="/post/new">New Post</a></li>
              <li><a href="/settings">Settings</a></li>
              {{end}}
            </ul>
          </div>
        </div>
      </div>

      <div class="pure-g">
        <div class="pure-u-1">

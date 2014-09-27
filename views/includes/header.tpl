<!DOCTYPE html>
<html>
  <head>
    <title>Nat? Nat. Nat! {{.Title}}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="/css/pure-min.0.5.0.css" />
    <link rel="stylesheet" href="/css/style.css" />
  </head>
  <body>
    <div class="wrap">
      <div class="pure-g">
        <div class="pure-u-1-5">
          <div class="leftnav">
            <a href="/"><img src="/img/natwelchlogo2.png" class="logo"></a>
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

        <div class="pure-u-4-5">

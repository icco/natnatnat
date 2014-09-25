<!DOCTYPE html>
<html>
  <head>
    <title>Nat? Nat. Nat! {{.Title}}</title>
    <link rel="stylesheet" href="/css/pure-min.0.5.0.css" />
    <link rel="stylesheet" href="/css/style.css" />
  </head>
  <body>
    <div class="wrap">
      <div class="grid">
        <div class="grid__col grid_col--1-of-5 leftnav">
          <a href="/"><img src="/img/natwelchlogo2.png" class="logo"></a>
          <ul>
            <li><a href="/about">About</a></li>
            <li><a href="/archives">Archives</a></li>
            <li><a href="/stats">Stats</a></li>

            {{if .IsAdmin}}
              <div><a href="/post/new">New Post</a></div>
            {{end}}
          </ul>
        </div>

        <div class="grid__col grid__col--4-of-5">

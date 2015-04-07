{{ template "includes/header" printf "#%d" .Entry.Id }}

<div class="post">
  <div class="meta">
    <div class="id">
      #{{.Entry.Id}}
    </div>
    <div class="time">
      <a href="/post/{{.Entry.Id}}"><time datetime="{{.Entry.Datetime|jsontime}}">{{.Entry.Datetime|fmttime}}</time></a>
    </div>
  </div>

  <div class="post-content">
    {{ if not .Entry.Public }}
      <h2 style="color: red; font-weight: 800;">POST IS NOT PUBLIC.</h2>
    {{ end }}

    {{ if .Entry.Title }}
      <h1>{{.Entry.Title}}</h1>
    {{ end }}

    <div class="markdown">
      {{.Entry.Content|mrkdwn}}
    </div>
  </div>

  <div class="addons">
  </div>
</div>

<div class="post-nav">
  <ul class="pager">
    <li class="{{if not .Prev}}disabled{{end}}"><a class="prev" href="{{.Prev}}">&#171;</a></li>
    <li class="{{if not .Next}}disabled{{end}}"><a class="next" href="{{.Next}}">&#187;</a></li>
  </ul>
</div>

{{ template "includes/footer" }}

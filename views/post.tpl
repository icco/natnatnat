{{ template "includes/header" }}

<div class="post">
  <div class="time">
    # <a href="/post/{{.Entry.Id}}">{{.Entry.Datetime|fmttime}}</a>
  </div>

  <div class="post-content">
    {{ if .Entry.Title }}
      <h1>{{.Entry.Title}}</h1>
    {{ end }}

    <div class="markdown">
      {{.Entry.Content|mrkdwn}}
    </div>
  </div>
</div>

<div class="post-nav">
  <ul class="pager">
    <li><a class="prev" href="#">&#171;</a></li>
    <li><a class="next" href="#">&#187;</a></li>
  </ul>
</div>

{{ template "includes/footer" }}

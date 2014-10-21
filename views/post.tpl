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

{{ template "includes/footer" }}

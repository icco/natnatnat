{{ template "includes/header" }}

<div class="post">
  <h1>{{.Entry.Title}}</h1>
  <div class="markdown">{{.Entry.Content|mrkdwn}}</div>

  <div class="meta">
    <a href="/post/{{.Entry.Id}}">{{.Entry.Datetime|fmttime}}</a>
  </div>
</div>

{{ template "includes/footer" }}

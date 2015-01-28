{{ template "includes/header" printf "#%s" .Tag }}

<div class="tag-header">
    All posts that contain the tag <strong>{{.Tag}}</strong>.
</div>

{{ range $entry := .Posts }}
  <div class="front-page post">
    <div class="time">
      # <a href="/post/{{$entry.Id}}">{{$entry.Datetime|fmttime}}</a>
    </div>

    <div class="post-content">
      {{ if $entry.Title }}
        <h2><a href="/post/{{$entry.Id}}">{{$entry.Title}}</a></h2>
      {{ end }}

      <div class="markdown">
        {{$entry.Content|mrkdwn}}
      </div>
    </div>
  </div>
{{ end }}

{{ template "includes/footer" }}

{{ template "includes/header" }}

{{ range $entry := .Posts }}
  <div class="front-page post">
    <div class="time">
    </div>
    <div class="meta">
      <div class="id">
        #{{$entry.Id}}
      </div>
      <div class="time">
        <a href="/post/{{$entry.Id}}">{{$entry.Datetime|fmttime}}</a>
      </div>
    </div>

    <div class="post-content">
      {{ if $entry.Title }}
        <h1><a href="/post/{{$entry.Id}}">{{$entry.Title}}</a></h1>
      {{ end }}

      <div class="markdown">
        {{$entry.Content|mrkdwn}}
      </div>
    </div>
  </div>
{{ end }}

{{ template "includes/footer" }}

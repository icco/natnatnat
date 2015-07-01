{{ if .Page }}
  {{ template "includes/header" printf "Page %d" .Page }}
{{ else }}
  {{ template "includes/header" }}
{{ end }}

{{ range $entry := .Posts }}
  <div class="post">
    <div class="cf">
      <div class="fl dib tl">
        #{{$entry.Id}}
      </div>
      <div class="fr dib tr">
        <a href="/post/{{$entry.Id}}"><time datetime="{{$entry.Datetime|jsontime}}">{{$entry.Datetime|fmttime}}</time></a>
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

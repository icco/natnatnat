{{ template "includes/header" printf "%v" .Date}}

<div class="mvm">
  All posts from {{.Date}}
</div>

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

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
        {{ $entry.Content|summary }}

        <p><a href="/post/{{$entry.Id}}">Continue Reading...</a></p>
      </div>
    </div>
  </div>
{{ end }}

<div class="post-nav f2">
  <ul class="pager">
    {{if ge .Next 0}}
      <li class=""><a class="next" href="/page/{{.Next}}">&#171;</a></li>
    {{end}}
    {{if ge .Prev 0}}
      <li class=""><a class="prev" href="/page/{{.Prev}}">&#187;</a></li>
    {{end}}
  </ul>
</div>

{{ template "includes/footer" }}

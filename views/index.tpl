{{ if .Page }}
  {{ template "includes/header" printf "Page %d" .Page }}
{{ else }}
  {{ template "includes/header" }}
{{ end }}

{{ range $entry := .Posts }}
  <div class="post">
    {{ template "includes/post_meta" $entry }}

    <div class="post-content">
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

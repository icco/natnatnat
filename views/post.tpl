{{ if .Entry.Title }}
  {{ template "includes/header" printf "#%d - \"%s\"" .Entry.Id .Entry.Title }}
{{ else }}
  {{ template "includes/header" printf "#%d" .Entry.Id }}
{{ end }}

{{ if .Entry.Draft }}
<h2 class="mvm" style="color: red; font-weight: 800;">POST IS A DRAFT.</h2>
{{ end }}

<div class="post">
  <div class="cf">
    <div class="fl dib tl">
      <h1 class="f1 mvn"><a href="/post/{{.Entry.Id}}">{{.Entry.Title}}</a></h1>
      #{{ .Entry.Id }} / <time datetime="{{ .Entry.Datetime|jsontime }}">{{ .Entry.Datetime|fmttime }}</time>
    </div>
  </div>

  <div class="post-content">

    <div class="markdown">
      {{.Entry.Content|mrkdwn}}
    </div>
  </div>

  <div class="addons">
  </div>
</div>

<div class="post-nav f2">
  <ul class="pager">
    <li class="{{if not .Prev}}disabled{{end}}"><a class="prev" href="{{.Prev}}">&#171;</a></li>
    <li class="{{if not .Next}}disabled{{end}}"><a class="next" href="{{.Next}}">&#187;</a></li>
  </ul>
</div>

{{ template "includes/footer" }}

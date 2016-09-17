{{ if .Entry.Title }}
  {{ template "includes/header" printf "#%d - \"%s\"" .Entry.Id .Entry.Title }}
{{ else }}
  {{ template "includes/header" printf "#%d" .Entry.Id }}
{{ end }}

{{ if .Entry.Draft }}
<h2 class="mvm" style="color: red; font-weight: 800;">POST IS A DRAFT.</h2>
{{ end }}

<article class="pv0 ph3 pa4-m pa5-l oh pos-rel mt0-ns mt4">
  <p class="mb0 mt0">#{{ .Entry.Id }} / <time datetime="{{ .Entry.Datetime|jsontime }}">{{ .Entry.Datetime|fmttime }}</time></p>

  <h1 class="f1 f-headline-ns mt0 mb3"><a href="/post/{{.Entry.Id}}">{{ .Entry.Title }}</a></h1>

  <p class="gray f6 mb4 ttu tracked">By Nat Welch</p>

  <div class="lh-copy mw7">
    <div class="markdown">
      {{.Entry.Content|mrkdwn}}
    </div>
  </div>

  <div class="addons">
  </div>
</article>

<div class="post-nav f4">
  <ul class="pager">
    <li class="{{if not .Prev}}disabled{{end}}"><a class="prev" href="{{.Prev}}">&#171;</a></li>
    <li class="{{if not .Next}}disabled{{end}}"><a class="next" href="{{.Next}}">&#187;</a></li>
  </ul>
</div>

{{ template "includes/footer" }}

<div class="cf">
  <div class="fl dib tl">
    <h1 class="f1 mb0"><a href="/post/{{.Id}}">{{ .Title }}</a></h1>
    #{{ .Id }} / <time datetime="{{ .Datetime|jsontime }}">{{ .Datetime|fmttime }}</time> / <a href="http://natwelch.com">Nat Welch</a>
  </div>
</div>

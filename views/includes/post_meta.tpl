<div class="cf">
  <div class="fl dib tl">
    <h1 class="f1 mvn"><a href="/post/{{.Entry.Id}}">{{ .Title }}</a></h1>
    #{{ .Id }} / <time datetime="{{ .Entry.Datetime|jsontime }}">{{ .Datetime|fmttime }}</time> / <a href="http://natwelch.com">Nat Welch</a>
  </div>
</div>

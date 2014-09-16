{{ template "includes/header" }}

{{range $entry := .Posts }} 
  <div class="front-page post">
    {{ if $entry.Title }}
      <h2><a href="/post/{{$entry.Id}}">{{ $entry.Title }}</a></h2>
    {{ end }}
    {{ $entry.Content }} <a href="/post/{{$entry.Id}}">{{$entry.Datetime|fmttime}}</a>
  </div>
{{ end }}
{{ template "includes/footer" }}

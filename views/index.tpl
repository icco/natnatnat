{{ template "includes/header" }}

{{range $entry := .Posts }} 
  <div class="front-page post">
    {{ $entry.Content }} <a href="/post/{{$entry.Id}}">{{$entry.Datetime}}</a>
  </div>
{{ end }}
{{ template "includes/footer" }}

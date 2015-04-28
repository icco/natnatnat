{{ template "includes/header" "Links" }}

<h1 style="text-align: center;">Links</h1>

<div id="links">
  {{ range $day, $links := .Links}}
    <h2>{{$day}}</h2>
    <ul>
      {{ range $link := $links}}
        <li><a href="{{$link.Url}}">{{$link.Title}}</a></li>
      {{ end }}
    </ul>
  {{ end }}
</div>

{{ template "includes/footer" }}

{{ template "includes/header" "Links" }}

<h1 style="text-align: center;">Links</h1>

<p>This is a list of all links that I have read grouped by day. You can also see this data in a more comprehensible form on <a href="https://pinboard.in/u:icco">Pinboard.in</a>.

<div id="links">
  {{ with $root := . }}
    {{ range $day := $root.Days}}
      <h2>{{$day}}</h2>
      <ul>
        {{ range $l := (index $root.Links $day)}}
          <li><a href="{{$l.Url}}">{{$l.Title}}</a></li>
        {{ end }}
      </ul>
    {{ end }}
  {{ end }}
</div>

{{ template "includes/footer" }}

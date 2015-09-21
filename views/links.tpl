{{ template "includes/header" "Links" }}

<h1 style="text-align: center;">Links</h1>

<p>This is a list of all links that I have read grouped by day. You can also see this data in a more comprehensible form on <a href="https://pinboard.in/u:icco">Pinboard.in</a>.

<div id="links">
  {{ range $pair := .LinkDays}}
  <h2>{{$pair.Day}}</h2>
  <ul>
    {{ range $l := (index $pair.Links)}}
      <li><a href="{{$l.Url}}">{{$l.Title}}</a></li>
    {{ end }}
  </ul>
  {{ end }}
</div>

{{ template "includes/footer" }}

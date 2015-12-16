{{ template "includes/header" "#Admin" }}

<p>
Welcome, {{.User}}! (<a href="{{.LogoutUrl}}">sign out</a>)
</p>

<h2>Admin Navigation</h2>
<ul>
  <li><a href="/post/new">New Post?</a></li>
  <li><a href="/aliases">Edit Tag Aliases?</a></li>
  <li><a href="/settings">Edit Settings?</a></li>
</ul>

<h2>Longform Drafts</h2>
<ul>
  {{ range $entry := .Drafts }}
  <li>
  #{{$entry.Id}}: "{{$entry.Title}}" <a href="/post/{{$entry.Id}}"><time datetime="{{$entry.Datetime|jsontime}}">{{$entry.Datetime|fmttime}}</time></a>, <a href="/edit/{{$entry.Id}}">EDIT</a>
  </li>
  {{ end }}
</ul>

<h2>Longform Published</h2>
<ul>
  {{ range $entry := .Longform}}
    {{ if ! $entry.Draft }}
      <li>
      #{{$entry.Id}}: "{{$entry.Title}}" <a href="/post/{{$entry.Id}}"><time datetime="{{$entry.Datetime|jsontime}}">{{$entry.Datetime|fmttime}}</time></a>, <a href="/edit/{{$entry.Id}}">EDIT</a>
      </li>
    {{ end }}
  {{ end }}
</ul>

{{ template "includes/footer" }}

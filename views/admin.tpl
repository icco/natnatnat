{{ template "includes/header" "#Admin" }}

<article class="pv0 ph3 pa4-m pa5-l oh pos-rel mt0-ns mt4">
  <h1 class="f1 mt0 mb3">Admin</h1>

  <div class="lh-copy">
    <div class="markdown">
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
          {{ if $entry.Draft }}
          {{ else }}
            <li>
            #{{$entry.Id}}: "{{$entry.Title}}" <a href="/post/{{$entry.Id}}"><time datetime="{{$entry.Datetime|jsontime}}">{{$entry.Datetime|fmttime}}</time></a>, <a href="/edit/{{$entry.Id}}">EDIT</a>
            </li>
          {{ end }}
        {{ end }}
      </ul>
    </div>
  </div>
</article>

{{ template "includes/footer" }}

{{ template "includes/header" "#Admin" }}

<p>
Welcome, {{.User}}! (<a href="{{.LogoutUrl}}">sign out</a>)
</p>

<ul>
  <li><a href="/post/new">New Post?</a></li>
  <li><a href="/aliases">Edit Tag Aliases?</a></li>
  <li><a href="/settings">Edit Settings?</a></li>
</ul>

{{ template "includes/footer" }}

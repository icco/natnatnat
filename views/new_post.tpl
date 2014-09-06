{{ template "includes/header" }}

Welcome, {{.User}}! (<a href="{{.LogoutUrl}}">sign out</a>)

<form method="post" action="/post/new">
  <input type="text" name="title" placeholder="Title" />
  <textarea name="text"></textarea>
  <input type="submit" />
</form>

{{ template "includes/footer" }}

{{ template "includes/header" }}
Welcome, {{.User}}! (<a href="{{.LogoutUrl}}">sign out</a>)

<ul>
  <li>Session: {{.Session}}</li>
  <li>Twitter: {{.Twitter}}</li>
  <li>Version: {{.Version}}</li>
</ul>

{{ template "includes/footer" }}

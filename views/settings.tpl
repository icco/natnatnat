{{ template "includes/header" }}
Welcome, {{.User}}! (<a href="{{.LogoutUrl}}">sign out</a>)

<ul>
<li>Session: {{.Session}}</li>
<li>Twitter: {{.Twitter}}</li>
</ul>

{{ template "includes/footer" }}

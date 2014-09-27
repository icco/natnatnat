{{ template "includes/header" }}
Welcome, {{.User}}! (<a href="{{.LogoutUrl}}">sign out</a>)

	Session
	TwitterAccessToken
	TwitterAccessTokenSecret
	TwitterKey
	TwitterSecret
	User
	Version
	Xsrf


<form class="pure-form pure-form-aligned">
  <fieldset>
    <input type="hidden" value="{{.Xsrf}}" name="xsrf" />

    <div class="pure-control-group">
      <label for="name">Session</label>
      <input id="name" type="text" placeholder="Session" value="{{.Session}}">
    </div>

    <div class="pure-control-group">
      <label for="name">TwitterAccessToken</label>
      <input id="name" type="text" placeholder="TwitterAccessToken" value="{{.TwitterAccessToken}}">
    </div>

    <div class="pure-control-group">
      <label for="name">TwitterAccessTokenSecret</label>
      <input id="name" type="text" placeholder="TwitterAccessTokenSecret" value="{{.TwitterAccessTokenSecret}}">
    </div>

    <div class="pure-control-group">
      <label for="name">TwitterKey</label>
      <input id="name" type="text" placeholder="TwitterKey" value="{{.TwitterKey}}">
    </div>

    <div class="pure-control-group">
      <label for="name">TwitterSecret</label>
      <input id="name" type="text" placeholder="TwitterSecret" value="{{.TwitterSecret}}">
    </div>

    <div class="pure-control-group">
      <label for="name">Version</label>
      <input id="name" type="text" placeholder="Version" value="{{.Version}}" disabled>
    </div>

    <div class="pure-controls">
      <button type="submit" class="pure-button pure-button-primary">Submit</button>
    </div>
  </fieldset>
</form>

{{ template "includes/footer" }}

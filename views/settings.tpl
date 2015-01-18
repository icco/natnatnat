{{ template "includes/header" }}
<p>Welcome, {{.User}}! (<a href="{{.LogoutUrl}}">sign out</a>)</p>


<form class="pure-form pure-form-aligned" method="post" action="/settings">

  <fieldset>
    <input type="hidden" value="{{.Xsrf}}" name="xsrf" />

    <div class="pure-control-group">
      <label for="session">Session</label>
      <input id="session" name="session_key" type="text" placeholder="Session" value="{{.Session}}">
    </div>

    <div class="pure-control-group">
      <label for="twitter_key">Twitter Key</label>
      <input id="twitter_key" name="twitter_key" type="text" placeholder="Twitter Key" value="{{.TwitterKey}}">
    </div>

    <div class="pure-control-group">
      <label for="twitter_sec">Twitter Secret</label>
      <input id="twitter_sec" name="twitter_sec" type="text" placeholder="Twitter Secret" value="{{.TwitterSecret}}">
    </div>

    <div class="pure-control-group">
      <label for="twitter_atok">Twitter Access Token</label>
      <input id="twitter_atok" name="twitter_atok" type="text" placeholder="Twitter Access Token" value="{{.TwitterAccessToken}}">
    </div>

    <div class="pure-control-group">
      <label for="twitter_asec">Twitter Access Token Secret</label>
      <input id="twitter_asec" name="twitter_asec" type="text" placeholder="Twitter Access Token Secret" value="{{.TwitterAccessTokenSecret}}">
    </div>

    <div class="pure-control-group">
      <label for="pb_usr">Pinboard User</label>
      <input id="pb_usr" name="pb_usr" type="text" placeholder="Pinboard User" value="{{.PinboardUser}}">
    </div>

    <div class="pure-control-group">
      <label for="pb_tok">Pinboard Token</label>
      <input id="pb_tok" name="pb_tok" type="text" placeholder="Pinboard Token" value="{{.PinboardToken}}">
    </div>

    <div class="pure-control-group">
      <label for="version">Version</label>
      <input id="version" name="version" type="text" placeholder="Version" value="{{.Version}}" disabled>
    </div>

    <div class="pure-controls">
      <button type="submit" class="pure-button pure-button-primary">Submit</button>
    </div>
  </fieldset>

  <p>A random string for you: <br/><textarea class="pure-input-1" style="min-height: 50px;">{{.Random}}</textarea></p>
</form>

{{ template "includes/footer" }}

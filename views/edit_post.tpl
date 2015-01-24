{{ template "includes/header" printf "#editpost #%d" .Entry.Id }}

<p>
Welcome, {{.User}}! (<a href="{{.LogoutUrl}}">sign out</a>)
</p>

<form method="post" action="{{.EditUrl}}" class="pure-form pure-form-stacked">
  <input type="text" name="title" placeholder="Title"  class="pure-input-1" />

  <textarea name="text" class="pure-input-1" style="min-height: 200px;">{{.Entry.Content}}</textarea>

  <input type="hidden" value="{{.Xsrf}}" name="xsrf" />

  <div class="pure-g">
    <div class="pure-u-1-1 form-padding">
      <input type="submit" class="pure-button pure-button-primary" />
      <label for="option-one" class="pure-checkbox">
        <input id="option-one" type="checkbox" name="private" {{if not .Entry.Public}}checked{{end}}>
        Private?
      </label>
    </div>
  </div>
</form>

{{ template "includes/footer" }}

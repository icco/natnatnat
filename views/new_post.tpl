{{ template "includes/header" }}

<p>
Welcome, {{.User}}! (<a href="{{.LogoutUrl}}">sign out</a>)
</p>

<form method="post" action="/post/new" class="pure-form pure-form-stacked">
  <input type="text" name="title" placeholder="Title"  class="pure-input-1" />

  <textarea name="text" class="pure-input-1" style="min-height: 200px;"></textarea>

  <input type="hidden" value="{{.Xsrf}}" name="xsrf" />

  <div class="pure-g">
    <div class="pure-u-1-2">
      <input type="text" class="pure-input-1" name="tags" placeholder="tags" />
    </div>

    <div class="pure-u-1-2 form-padding">
      <label for="option-one" class="pure-checkbox">
        <input id="option-one" type="checkbox" value="" name="private">
        Private?
      </label>
    </div>

    <div class="pure-g">
      <div class="pure-u-1-4">
        <input type="submit" class="pure-button pure-button-primary" />
      </div>
    </div>
  </div>
</form>

{{ template "includes/footer" }}

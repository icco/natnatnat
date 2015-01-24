{{ template "includes/header" }}

<p>
Welcome, {{.User}}! (<a href="{{.LogoutUrl}}">sign out</a>)
</p>

<form method="post" action="/post/new" class="pure-form pure-form-stacked">
  <input type="text" name="title" placeholder="Title"  class="pure-input-1" />

  <textarea name="text" class="pure-input-1" style="min-height: 200px;"></textarea>

  <input type="hidden" value="{{.Xsrf}}" name="xsrf" />

  <div class="pure-g">
    <div class="pure-u-1-1 form-padding">
      <input type="submit" class="pure-button pure-button-primary" />
      <label for="option-one" class="pure-checkbox">
        <input id="option-one" type="checkbox" name="private">
        Private?
      </label>
    </div>
  </div>
</form>

<ul>
  {{ range $link := .Links }}
    {{ with $link }}
      <li class="link"><a class="adder" data-tags="{{.TagString}}">&plus;</a> &ndash; <a class="actual" href="{{.Url}}">{{.Title}}</a></li>
    {{end}}
  {{ end }}
</ul>

{{ template "includes/footer" }}

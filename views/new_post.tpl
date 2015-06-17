{{ template "includes/header" "#NewPost" }}

<p>
Welcome, {{.User}}! (<a href="{{.LogoutUrl}}">sign out</a>)
</p>

<form method="post" action="/post/new" class="">
  <input type="text" name="title" placeholder="Title" class="db w-100 pas mvm input-text" />

  <textarea name="text" class="db w-100 pas mvs input-text" style="min-height: 200px;"></textarea>

  <input type="hidden" value="{{.Xsrf}}" name="xsrf" />

  <div class="cf">
    <input type="submit" class="btn pas mrm" />
    <label for="option-one" class="tr pas fr">
      <input id="option-one" type="checkbox" name="private">
      Private?
    </label>
  </div>
</form>

<div class="preview">
  <div class="mvs">
    <small><a class="show_hide">Preview...</a></small>
  </div>
  <h1 id="rendered_title"></h1>
  <div id="rendered"></div>
</div>

<div class="links">
  <ul>
    {{ range $link := .Links }}
      {{ with $link }}
        <li class="link"><a class="adder" data-tags="{{.TagString}}">&plus;</a> &ndash; <a class="actual" href="{{.Url}}">{{.Title}}</a></li>
      {{end}}
    {{ end }}
  </ul>
</div>

{{ template "includes/footer" }}

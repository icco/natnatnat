{{ template "includes/header" printf "#editpost #%d" .Entry.Id }}

<p>
Welcome, {{.User}}! (<a href="{{.LogoutUrl}}">sign out</a>)
</p>

<form method="post" action="{{.EditUrl}}" class="">
  <input type="text" name="title" placeholder="Title"  class="db w-100 pas mvm input-text" value="{{.Entry.Title}}" />

  <textarea name="text" class="db w-100 pas mvs input-text" style="min-height: 17rem; resize: vertical;">{{.Entry.Content}}</textarea>

  <input type="hidden" value="{{.Xsrf}}" name="xsrf" />

  <div class="cf">
    <input type="submit" class="btn pas mrm btn--blue" />
    <label for="option-one" class="tr pas fr">
      <input id="option-one" type="checkbox" name="draft" {{if .Entry.Draft}}checked{{end}}>
      Draft?
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

{{ template "includes/footer" }}

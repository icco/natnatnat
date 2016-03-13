{{ template "includes/header" "#NewPost" }}

<p>
Welcome, {{.User}}! (<a href="{{.LogoutUrl}}">sign out</a>)
</p>

<form method="post" action="/post/new" class="">
  <input type="text" name="title" placeholder="Title" class="db w-100 pas mvm input-text" />

  <textarea name="text" class="db w-100 pas mvs input-text" style="min-height: 17rem; resize: vertical;"></textarea>

  <input type="hidden" value="{{.Xsrf}}" name="xsrf" />

  <div class="cf">
    <input type="submit" class="btn pas mrm btn--blue" />
    <label for="option-one" class="tr pas fr">
      <input id="option-one" type="checkbox" name="draft">
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

<div id="links" class="links">
  {{ range $pair := .Links}}
    <h2>{{$pair.Day}}</h2>
    <ul>
      {{ range $l := (index $pair.Links)}}
        <li class="link"><a class="adder" data-tags="{{$l.TagString}}">&plus;</a> &ndash; <a class="actual" href="{{$l.Url}}">{{$l.Title}}</a></li>
      {{ end }}
    </ul>
  {{ end }}
</div>

<script language="JavaScript">
  window.onbeforeunload = confirmExit;
  function confirmExit() {
    return "You are in the process of creating a post.";
  }
</script>

{{ template "includes/footer" }}

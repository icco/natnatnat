{{ template "includes/header" "#NewPost" }}

<article class="mh3">
  <p>
  Welcome, {{.User}}! (<a href="{{.LogoutUrl}}">sign out</a>)
  </p>

  <form method="post" action="/post/new" class="">
    <input type="text" name="title" placeholder="Title" class="db w-100 pa1 mv3 input-text" />

    <textarea name="text" class="db w-100 pa1 mv1 input-text" style="min-height: 17rem; resize: vertical;"></textarea>

    <input type="hidden" value="{{.Xsrf}}" name="xsrf" />

    <div class="cf">
      <input type="submit" class="f6 link dim br2 ph3 pv2 mb2 dib white bg-dark-blue pointer" />
      <label for="option-one" class="tr pa1 fr">
        <input id="option-one" type="checkbox" name="draft">
        Draft?
      </label>
    </div>
  </form>
</article>

<div class="preview mh3">
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

{{ template "includes/footer" }}

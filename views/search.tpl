{{ template "includes/header" "Search" }}

<article class="pv0 ph3 ph4-m ph5-l oh mt0-ns mt4 black-80">
  <form method="get" accept-charset="utf-8">
    <fieldset id="sign_up" class="ba b--transparent ph0 mh0">
      <legend class="ph0 mh0 fw6 clip">Search</legend>
      <div class="mt3">
        <input class="pa2 input-reset ba bg-transparent w-100 measure" type="text" name="s" id="s" value="{{ .Query }}">
      </div>
    </fieldset>
    <div class="mt3"><input class="b ph3 pv2 input-reset ba b--black bg-transparent grow pointer f6" type="submit" value="Search"></div>
  </form>

  <p>
  {{.Count}} Results
  </p>
</article>

{{ range $entry := .Results }}
  <article class="pv0 ph3 ph4-m ph5-l oh mt0-ns mt4">
    {{ template "includes/post_meta" $entry }}

    <div class="post-content">
      <div class="markdown">
        {{ $entry.Content|summary }}

        <p><a href="/post/{{$entry.Id}}">Continue Reading...</a></p>
      </div>
    </div>
  </article>
{{ end }}

{{ template "includes/footer" }}

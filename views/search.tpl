{{ template "includes/header" "Search" }}

<article class="pv0 ph3 ph4-m ph5-l oh mt0-ns mt4 black-80">
  <form class="mw7 pa2 br2-ns " method="get" accept-charset="utf-8">
    <fieldset class="cf bn ma0 pa0">
      <div class="cf">
        <input class="fl f6 f5-l input-reset ba b--black-70 black-80 bg-white fl pa3 lh-solid w-100 w-75-m w-80-l br2-ns br--left-ns" placeholder="Search query" type="text" name="s" value="{{ .Query }}" id="s">
        <input class="fl f6 f5-l button-reset pv3 tc ba b--black-70 bg-animate bg-black-70 hover-bg-black white pointer w-100 w-25-m w-20-l br2-ns br--right-ns" type="submit" value="Search">
      </div>
    </fieldset>
  </form>

  <p class="pa2">
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

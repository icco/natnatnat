{{ template "includes/header" printf "%v" .Date}}

<article class="pv0 ph3 ph4-m ph5-l oh pos-rel mt0-ns mt4">
  <h1 class="f1 mt0 mb3">All posts from {{.Date}}</h1>
</article>

{{ range $entry := .Posts }}
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

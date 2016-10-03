{{ template "includes/header" "Tags" }}

<article class="pv0 ph3 pa4-m pa5-l oh pos-rel mt0-ns mt4">
  <h1 class="f1 mt0 mb3">Tags</h1>

  <div class="lh-copy">
    <div class="markdown">
      <p>
      All the tags in posts, in no valuable order!
      </p>

      <ul>
        {{ range $tag, $count := .Tags }}
        <li><a href="/tags/{{$tag}}">{{$tag}}</a></li>
        {{ end }}
      </ul>
    </div>
  </div>
</article>

{{ template "includes/footer" }}

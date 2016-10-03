{{ template "includes/header" printf "#%s" .Tag }}

<article class="pv0 ph3 ph4-m ph5-l oh pos-rel mt0-ns mt4">
  <div class="lh-copy">
    <div class="markdown">
      <p>
      All posts that contain the tag <strong>{{.Tag}}</strong>, and its aliases.
      </p>
    </div>
  </div>
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

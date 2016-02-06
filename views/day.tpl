{{ template "includes/header" printf "%v" .Date}}

<div class="mvm">
  All posts from {{.Date}}
</div>

{{ range $entry := .Posts }}
  <div class="post">
    {{ template "includes/post_meta" $entry }}

    <div class="post-content">
      <div class="markdown">
        {{ $entry.Content|summary }}

        <p><a href="/post/{{$entry.Id}}">Continue Reading...</a></p>
      </div>
    </div>
  </div>
{{ end }}

{{ template "includes/footer" }}

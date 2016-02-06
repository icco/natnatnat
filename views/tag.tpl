{{ template "includes/header" printf "#%s" .Tag }}

<div class="mvm">
    All posts that contain the tag <strong>{{.Tag}}</strong>, and its aliases (<ul class="inline">
      {{ range $a := .Aliases }}
      <li>{{ $a }}</li>
      {{ end }}
    </ul>).
</div>

{{ range $entry := .Posts }}
  <div class="post">
    {{ template "includes/post_meta" $entry }}

    <div class="post-content">
      <div class="markdown">
        {{$entry.Content|mrkdwn}}
      </div>
    </div>
  </div>
{{ end }}

{{ template "includes/footer" }}

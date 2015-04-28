{{ template "includes/header" "Tags" }}

<h1 style="text-align: center;">Tags</h1>

<div class="tag-header">
    All the tags in posts, in no valuable order!
</div>

<ul>
  {{ range $tag, $count := .Tags }}
    <li><a href="/tags/{{$tag}}">{{$tag}}</a></li>
  {{ end }}
</ul>

{{ template "includes/footer" }}

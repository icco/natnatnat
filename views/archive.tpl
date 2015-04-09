{{ template "includes/header" "Archives" }}

<h1 style="text-align: center;">Archives</h1>

<div id="archives">
  <ul>
    {{ range $year, $months := .Years }}
      <div class="year">{{ $year }}</div>

      {{ range $month, $days := $months }}
        <div class="month">{{ $month }}</div>
        {{ range $day, $posts := $days }}
          <div class="day">{{ $day }}</div>
          <div class="days">
            {{ range $post := $posts }}
            <li><a href="/post/{{ $post.Id }}">{{ if $post.Title }}{{ $post.Title }}{{ else }}#{{ $post.Id }}{{ end }}</a></li>
            {{ end }}
          </div>
        {{ end }}
      {{ end }}
    {{ end }}
  </ul>
</div>

{{ template "includes/footer" }}

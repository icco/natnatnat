{{ template "includes/header" }}

<div id="archives">
  <ul>
    {{ range $year, $months := .Years }}
      <div class="year">{{ $year }}</div>

      {{ range $month, $days := $months }}
        <div class="month">{{ $month }}</div>
        {{ range $day, $posts := $days }}
          <div class="day">{{ $day }}</div>
          {{ range $post := $posts }}
            <li><a href="/post/{{ $post.Id }}">#{{ $post.Id }}</a></li>
          {{ end }}
        {{ end }}
      {{ end }}
    {{ end }}
  </ul>
</div>

{{ template "includes/footer" }}

{{ template "includes/header" "Archives" }}

<h1 class="tc lh-title">Archives</h1>

<h2 class="tc lh-title">Streaks</h2>

<div>
  {{ range $year, $months := .Years }}
    <div class="f3 lh-title">{{ $year }}</div>

    {{ .Months }}
    {{ range $null, $month := .Months }}
      {{with $days := index $months $month }}
        <div class="mvs">
          <div class="f4 lh-title dib mrm">{{ $month }}</div>
          {{ range $day, $posts := $days }}
            {{ if $posts }}
            <a href="/day/{{$year}}/{{printf "%d" $month}}/{{$day}}" class="none">
              <div class="w1 h1 dib v-mid bg-light-green ba b--lightest-green" title="{{$year}}/{{printf "%02d" $month}}/{{$day}} - {{len $posts}} posts"></div>
            </a>
            {{ else }}
            <div class="w1 h1 dib v-mid bg-blue ba b--lightest-blue" title="{{$year}}/{{printf "%02d" $month}}/{{$day}} No Posts"></div>
            {{ end }}
          {{ end }}
        </div>
      {{ end }}
    {{ end }}
  {{ end }}
</div>

<h2 class="tc lh-title">All Posts</h2>
<div class="lh-copy">
  <ul>
    {{ range $post := .Posts }}
      <li><a href="/post/{{ $post.Id }}">{{ if $post.Title }}{{ $post.Title }}{{ else }}#{{ $post.Id }}{{ end }}</a></li>
    {{ end }}
  </ul>
</div>

{{ template "includes/footer" }}

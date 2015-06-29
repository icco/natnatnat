{{ template "includes/header" "Archives" }}

<h1 class="tc lh-title">Archives</h1>

<h2 class="tc lh-title">Streaks</h2>

<div>
  {{ range $year, $months := .Years }}
    <div class="f3 lh-title">{{ $year }}</div>

    {{ range $month, $days := $months }}
      <div class="mvs">
        <div class="f4 lh-title dib mrm">{{ $month }}</div>
        {{ range $day, $posts := $days }}
          {{ if $posts }}
          <a href="/day/{{$year}}/{{printf "%d" $month}}/{{$day}}">
            <div class="w1 h1 dib v-mid bg-light-green ba b--lightest-green" title="{{$year}}/{{printf "%02d" $month}}/{{$day}} - {{len $posts}} posts"></div>
          </a>
          {{ else }}
          <div class="w1 h1 dib v-mid bg-red ba b--lightest-red" title="{{$year}}/{{printf "%02d" $month}}/{{$day}} No Posts"></div>
          {{ end }}
        {{ end }}
      </div>
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

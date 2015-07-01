{{ template "includes/header" "Archives" }}

<h1 class="tc lh-title">Archives</h1>

<div>
  {{ with $data := . }}
    {{ if $data.Years }}
      <h2 class="tc lh-title">Streaks</h2>
      {{ range $year, $months := $data.Years }}
        <div class="f3 lh-title">{{ $year }}</div>
        {{ range $null, $month := $data.Months }}
          {{ with $days := index $months $month }}
            <div class="mvs">
              <div class="f4 lh-title dib mrm mw4 tr">{{ $month }}</div>
              {{ range $day, $posts := $days }}
                {{ if gt $day 0 }}
                  {{ if $posts }}
                    <a href="/day/{{$year}}/{{m2i $month}}/{{$day}}" class="none">
                      <div class="w1 h1 dib v-mid bg-light-green ba b--lightest-green" title="{{$year}}/{{m2i $month}}/{{$day}} - {{$posts}} posts"></div>
                    </a>
                  {{ else }}
                    <div class="w1 h1 dib v-mid ba b--light-gray" title="{{$year}}/{{m2i $month}}/{{$day}} No Posts"></div>
                  {{ end }}
                {{ end }}
              {{ end }}
            </div>
          {{ end }}
        {{ end }}
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

{{ template "includes/header" "Archives" }}

<article class="pv0 ph3 pa4-m pa5-l oh pos-rel mt0-ns mt4">
  <h1 class="f1 mt0 mb3">Archives</h1>

  <div>
    {{ with $data := . }}
      {{ if $data.Years }}
        <h2 class="lh-title">Streaks</h2>
        {{ range $year, $months := $data.Years }}
          <div class="f3 lh-title">{{ $year }}</div>
          {{ range $null, $month := $data.Months }}
            {{ with $days := index $months $month }}
              <div class="mvs">
                <div class="f4 lh-title dib mh3 tr w4">{{ $month }}</div>
                {{ range $day, $posts := $days }}
                  {{ if gt $day 0 }}
                    {{ if $posts }}
                      <a href="/day/{{$year}}/{{m2i $month}}/{{$day}}" class="link">
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

  <h2 class="lh-title">All Posts</h2>
  <div class="lh-copy">
    <ul>
      {{ range $post := .Posts }}
        <li><a href="/post/{{ $post.Id }}">{{ if $post.Title }}{{ $post.Title }}{{ else }}#{{ $post.Id }}{{ end }}</a></li>
      {{ end }}
    </ul>
  </div>
</article>

{{ template "includes/footer" }}

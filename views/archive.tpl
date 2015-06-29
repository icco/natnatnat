{{ template "includes/header" "Archives" }}

<h1 style="text-align: center;">Archives</h1>

<div id="archives">
  {{ range $year, $months := .Years }}
    <div class="year">{{ $year }}</div>

    {{ range $month, $days := $months }}
      <div class="month">{{ $month }}</div>
      {{ range $day, $posts := $days }}
        {{ if $posts }}
        <a href="/day/{{$year}}/{{$month}}/{{$day}}">
          <div class="w1 h1 dib bg-light-green ba b--lightest-green" title="{{$year}}/{{$month}}/{{$day}} - {{len $posts}} posts"></div>
        </a>
        {{ else }}
        <div class="w1 h1 dib bg-red ba b--lightest-red" title="{{$year}}/{{$month}}/{{$day}} No Posts"></div>
        {{ end }}
      {{ end }}
    {{ end }}
  {{ end }}
</div>

<div id="archives">
  <ul>
    {{ range $post := .Posts }}
      <li><a href="/post/{{ $post.Id }}">{{ if $post.Title }}{{ $post.Title }}{{ else }}#{{ $post.Id }}{{ end }}</a></li>
    {{ end }}
</div>

{{ template "includes/footer" }}

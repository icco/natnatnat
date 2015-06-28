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
          <div class="w1 h1 dib bg-green" title="{{ $day }}"></div>
        </a>
        {{ else }}
          <div class="w1 h1 dib bg-black" title="{{ $day }}"></div>
        {{ end }}
      {{ end }}
    {{ end }}
  {{ end }}
</div>

{{ template "includes/footer" }}

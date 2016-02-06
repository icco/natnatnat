{{ template "includes/header" printf "%v" .Date}}

<div class="mvm">
  All posts from {{.Date}}
</div>

{{ range $entry := .Posts }}
  <div class="post">
    <div class="cf">
      <div class="fl dib tl">
        <h1 class="f1 mvn"><a href="/post/{{$entry.Id}}">{{$entry.Title}}</a></h1>
        #{{ $entry.Id }} / <time datetime="{{ $entry.Datetime|jsontime }}">{{ $entry.Datetime|fmttime }}</time>
      </div>
    </div>

    <div class="post-content">
      <div class="markdown">
        {{ $entry.Content|summary }}

        <p><a href="/post/{{$entry.Id}}">Continue Reading...</a></p>
      </div>
    </div>
  </div>
{{ end }}

{{ template "includes/footer" }}

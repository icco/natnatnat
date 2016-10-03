{{ template "includes/header" "Links" }}

<article class="pv0 ph3 pa4-m pa5-l oh pos-rel mt0-ns mt4">
  <h1 class="f1 mt0 mb3">Links</h1>

  <div class="lh-copy">
    <div class="markdown">
      <p>This is a list of all links that I have read grouped by day. You can also see this data in a more comprehensible form on <a href="https://pinboard.in/u:icco">Pinboard.in</a>.

      <div id="links">
        {{ range $pair := .LinkDays}}
        <h2>{{$pair.Day}}</h2>
        <ul>
          {{ range $l := (index $pair.Links)}}
          <li><a href="{{$l.Url}}">{{$l.Title}}</a></li>
          {{ end }}
        </ul>
        {{ end }}
      </div>
    </div>
  </div>
</article>

{{ template "includes/footer" }}

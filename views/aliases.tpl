{{ template "includes/header" "Tag Aliases" }}

<h1 style="text-align: center;">Tag Aliases</h1>

<form method="post" action="/post/new" class="pure-form pure-form-stacked">
  <input type="text" name="name" placeholder="From"  class="pure-input-1" />
  <input type="text" name="tag" placeholder="To"  class="pure-input-1" />

  <input type="hidden" value="{{.Xsrf}}" name="xsrf" />

  <div class="pure-g">
    <div class="pure-u-1-1 form-padding">
      <input type="submit" class="pure-button pure-button-primary" />
    </div>
  </div>
</form>

<ul>
  {{ range $from, $to:= .Tags }}
  <li>{{$from}} &rarr; <a href="/tags/{{$to}}">{{$to}}</a></li>
  {{ end }}
</ul>

{{ template "includes/footer" }}

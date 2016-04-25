{{ template "includes/header" "Search" }}

<form>
  <input type="text" name="s" class="f3 semibold input-text plm pvm db ba black-80 b--black-20 w-100 bg-black-05 bg-black-10-focus" value="{{ .Query }}" />
  <input type="submit" class="input-text bg-black-05 pam brs semibold ba b--black-20 ttu tracked-mega bg-black-10-focus mvs" value="Search" />
</form>

<p>
{{.Count}} Results
</p>

{{ range $entry := .Results }}
  <div class="post">
    {{ template "includes/post_meta" $entry }}

    <div class="post-content">
      <div class="markdown">
        {{ $entry.Content|summary }}

        <p><a href="/post/{{$entry.Id}}">Continue Reading...</a></p>
      </div>
    </div>
  </div>
{{ end }}

{{ template "includes/footer" }}

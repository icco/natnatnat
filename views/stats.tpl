{{ template "includes/header" "Stats" }}

<article class="pa3 pa5-ns" data-name="slab-stat-large">
  <h3 class="f6 ttu tracked">All Time Stats</h3>
  <div class="cf">
    <dl class="db dib-l w-auto-l lh-title mr6-l">
      <dd class="f6 fw4 ml0">Avg. posts made per day</dd>
      <dd class="f2 f-subheadline-l fw6 ml0">{{printf "%.2f" .PostsPerDay}}</dd>
    </dl>

    <dl class="db dib-l w-auto-l lh-title mr6-l">
      <dd class="f6 fw4 ml0">Avg. links read per day</dd>
      <dd class="f2 f-subheadline-l fw6 ml0">{{printf "%.2f" .LinksPerDay}}</dd>
    </dl>

    <dl class="db dib-l w-auto-l lh-title mr6-l">
      <dd class="f6 fw4 ml0">Days since first post</dd>
      <dd class="f2 f-subheadline-l fw6 ml0">{{printf "%.2f" .DaysSince}}</dd>
    </dl>

    <dl class="db dib-l w-auto-l lh-title mr6-l">
      <dd class="f6 fw4 ml0">Total number of posts</dd>
      <dd class="f2 f-subheadline-l fw6 ml0">{{.Posts}}</dd>
    </dl>

    <dl class="db dib-l w-auto-l lh-title mr6-l">
      <dd class="f6 fw4 ml0">Avg. words per post</dd>
      <dd class="f2 f-subheadline-l fw6 ml0">{{printf "%.2f" .WordsPerPost}}</dd>
    </dl>

    <dl class="db dib-l w-auto-l lh-title mr6-l">
      <dd class="f6 fw4 ml0">Avg. words per day</dd>
      <dd class="f2 f-subheadline-l fw6 ml0">{{printf "%.2f" .WordsPerDay}}</dd>
    </dl>
  </div>
</article>

<!-- years -->
{{ range $y := .Years }}
<article class="pa3 pa5-ns" data-name="slab-stat-large">
  <h3 class="f6 ttu tracked">{{ $y }} Stats</h3>
  <div class="cf">
    <dl class="db dib-l w-auto-l lh-title mr6-l">
      <dd class="f6 fw4 ml0">Posts this year</dd>
      <dd class="f2 f-subheadline-l fw6 ml0">{{ printf "%.0f" (index $.YearData $y 0) }}</dd>
    </dl>

    <dl class="db dib-l w-auto-l lh-title mr6-l">
      <dd class="f6 fw4 ml0">Avg. posts per week</dd>
      <dd class="f2 f-subheadline-l fw6 ml0">{{ printf "%.2f" (index $.YearData $y 1) }}</dd>
    </dl>

    <dl class="db dib-l w-auto-l lh-title mr6-l">
      <dd class="f6 fw4 ml0">Links saved</dd>
      <dd class="f2 f-subheadline-l fw6 ml0">{{ printf "%.0f" (index $.YearData $y 2) }}</dd>
    </dl>
  </div>
</article>
{{ end }}

{{ template "includes/footer" }}

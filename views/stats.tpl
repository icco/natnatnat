{{ template "includes/header" "Stats" }}

<div class="tc center">
  <h1>Stats</h1>
</div>

<div id="stats" class="lh-title tc dt">
  <div class="dtr-ns db">
    <div class="mw5-ns center tc pal dtc-ns">
      <div class="f1">{{printf "%.2f" .PostsPerDay}}</div>
      <div class="book">Avg. posts made per day</div>
    </div>

    <div class="mw5-ns center tc pal dtc-ns">
      <div class="f1">{{printf "%.2f" .LinksPerDay}}</div>
      <div class="book">Avg. links read per day</div>
    </div>

    <!-- TODO
    <div class="">
      <div class="f1">{{printf "%.2f" .LinksPerPost}}</div>
      <div class="book">Avg. links per post</div>
    </div>
    -->

    <div class="mw5-ns center tc pal dtc-ns">
      <div class="f1">{{printf "%.2f" .DaysSince}}</div>
      <div class="book">Days since first post</div>
    </div>
  </div>

  <div class="dtr-ns db">
    <div class="mw5-ns center tc pal dtc-ns">
      <div class="f1">{{printf "%.2f" .WordsPerDay}}</div>
      <div class="book">Avg. words per day</div>
    </div>

    <div class="mw5-ns center tc pal dtc-ns">
      <div class="f1">{{printf "%.2f" .WordsPerPost}}</div>
      <div class="book">Avg. words per post</div>
    </div>

    <div class="mw5-ns center tc pal dtc-ns">
      <div class="f1">{{.Posts}}</div>
      <div class="book">Total number of posts</div>
    </div>
  </div>
</div>

<div id="statsgraph">
  <!-- TODO -->
</div>

{{ template "includes/footer" }}

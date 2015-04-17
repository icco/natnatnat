{{ template "includes/header" "Stats" }}

<h1 style="text-align: center;">Stats</h1>

<div id="stats" class="pure-g">
  <div class="pure-u-1 pure-u-md-1-3 stat">
    <div class="big">{{printf "%.2f" .PostsPerDay}}</div>
    <div class="small">Avg. posts made per day</div>
  </div>

  <!-- TODO
  <div class="pure-u-1 pure-u-md-1-3 stat">
    <div class="big"><%= @avgs[:articles] %></div>
    <div class="small">Avg. articles read per day</div>
  </div>
  -->

  <!-- TODO
  <div class="pure-u-1 pure-u-md-1-3 stat">
    <div class="big"><%= @avgs[:links] %></div>
    <div class="small">Avg. links per post</div>
  </div>
  -->

  <div class="pure-u-1 pure-u-md-1-3 stat">
    <div class="big">{{printf "%.2f" .DaysSince}}</div>
    <div class="small">Days since first post</div>
  </div>

  <div class="pure-u-1 pure-u-md-1-3 stat">
    <div class="big">{{printf "%.2f" .WordsPerDay}}</div>
    <div class="small">Avg. words per day</div>
  </div>

  <div class="pure-u-1 pure-u-md-1-3 stat">
    <div class="big">{{printf "%.2f" .WordsPerPost}}</div>
    <div class="small">Avg. words per post</div>
  </div>

  <div class="pure-u-1 pure-u-md-1-3 stat">
    <div class="big">{{.Posts}}</div>
    <div class="small">Total number of posts</div>
  </div>
</div>

<div id="statsgraph">
</div>

{{ template "includes/footer" }}

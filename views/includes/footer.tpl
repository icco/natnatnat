        </div> <!-- /.unit -->
      </div> <!-- /.grid -->
      <div class="pure-g footer">
        <div class="pure-u-1">
          <div class="pure-menu pure-menu-open pure-menu-horizontal">
            <ul>
              <li><a href="/about">About</a></li>
              <li><a href="/archives">Archives</a></li>
              <li><a href="/stats">Stats</a></li>

              {{if .IsAdmin}}
              <li><a href="/post/new">New Post</a></li>
              <li><a href="/settings">Settings</a></li>
              {{end}}
            </ul>
          </div>
        </div>
      </div>
    </div> <!-- /.wrap -->

    <!-- Testing out embeds -->
    <script src="//ajax.googleapis.com/ajax/libs/jquery/1.8.3/jquery.min.js" type="text/javascript"></script>
    <script src="//cdn.embed.ly/jquery.embedly-3.1.1.min.js" type="text/javascript"></script>
    <script src="/js/app.js?v=20141102" type="text/javascript"></script>
  </body>
</html>

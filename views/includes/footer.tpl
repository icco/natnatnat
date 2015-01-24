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
    <script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js" type="text/javascript"></script>
    <script src="//cdn.embed.ly/jquery.embedly-3.1.1.min.js" type="text/javascript"></script>
    <script src="/js/app.js?v=20150124" type="text/javascript"></script>

    <script>
      (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
          (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
        m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
      })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
      ga('create', 'UA-333449-17', 'auto');
      ga('send', 'pageview');
    </script>
  </body>
</html>

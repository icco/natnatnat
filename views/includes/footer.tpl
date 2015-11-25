        </div> <!-- /.l-box -->
      </div> <!-- /.grid -->
      <div class="footer">
        <div class="">
          <ul>
            {{if .IsAdmin}}
            <li><a href="/post/new">New Post</a></li>
            <li><a href="/settings">Settings</a></li>
            {{end}}
          </ul>
        </div>
      </div>
    </div> <!-- /.wrap -->

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

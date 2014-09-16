{{ template "includes/header" }}
<p>
  Id: {{ .Entry.Id }}
</p>
<p>
  Title: {{ .Entry.Title }}
</p>
<p>
  Content: {{ .Entry.Content }}
</p>
<p>
  Datetime: {{ .Entry.Datetime|fmttime}}
</p>
{{ template "includes/footer" }}

<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd" xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://writing.natwelch.com/</loc>
    <lastmod>{{.Newest|jsontime}}</lastmod>
    <changefreq>daily</changefreq>
    <priority>1.0</priority>
  </url>
  <url>
    <loc>https://writing.natwelch.com/archives/</loc>
    <lastmod>{{.Newest|jsontime}}</lastmod>
    <changefreq>daily</changefreq>
    <priority>1.0</priority>
  </url>
  <url>
    <loc>https://writing.natwelch.com/tags/</loc>
    <lastmod>{{.Newest|jsontime}}</lastmod>
    <changefreq>daily</changefreq>
    <priority>0.5</priority>
  </url>
  <url>
    <loc>https://writing.natwelch.com/stats/</loc>
    <lastmod>{{.Newest|jsontime}}</lastmod>
    <changefreq>daily</changefreq>
    <priority>0.5</priority>
  </url>
  <url>
    <loc>https://writing.natwelch.com/links/</loc>
    <lastmod>{{.Newest|jsontime}}</lastmod>
    <changefreq>daily</changefreq>
    <priority>0.5</priority>
  </url>

  {{ range $entry := .Posts }}
    <url>
      <loc>https://writing.natwelch.com/post/{{$entry.Id}}</loc>
      <lastmod>{{$entry.Datetime|jsontime}}</lastmod>
      <changefreq>monthly</changefreq>
      <priority>0.7</priority>
    </url>
  {{ end }}
</urlset> 

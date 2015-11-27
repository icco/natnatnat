#! /usr/bin/env ruby

# This takes a folder of files and adds a line of "origin_url:
# http://pseudoweb.net/2012/11/17/what-i-learned-at-burning-man-2012/"

directory = "./posts/"

Dir.entries(directory).select {|f| f.start_with? "2" }.each do |f|
   parts = File.basename(f, ".md").split("-")
   year = parts.delete_at(0)
   month = parts.delete_at(0)
   day = parts.delete_at(0)
   stub = parts.join("-")

   p [year, month, day, stub]
   origin = "http://pseudoweb.net/#{year}/#{month}/#{day}/#{stub}/"

   lines = File.readlines(directory + f)
   lines.insert(2, "origin: #{origin}\n")
   File.write(directory + f, lines.join)
end

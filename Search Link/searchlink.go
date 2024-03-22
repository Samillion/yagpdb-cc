{{ $engines := sdict 
	"google" (sdict "name" "Google" "url" "https://www.google.com/search?hl=en&q=")
	"duck" (sdict "name" "DuckDuckGo" "url" "https://duckduckgo.com/?q=")
	"brave" (sdict "name" "Brave Search" "url" "https://search.brave.com/search?q=")
	"wiki" (sdict "name" "Wikipedia" "url" "https://www.wikipedia.org/wiki/")
}}

{{ $in := trimSpace .StrippedMsg }}
{{ $cmd := slice .Cmd 1 (len .Cmd) | lower }}

{{ $set := or ($engines.Get $cmd) $engines.google }}
{{ $name := $set.name }}
{{ $url := $set.url }}

{{ if $in }}
	{{ $keys := urlquery $in }}
	{{ if eq $cmd "wiki" }}
		{{ $keys = urlquery (joinStr "_" .CmdArgs) }}
	{{ end }}
	{{ $link := print $url $keys }}
	{{ sendMessage nil $link }}
{{ else }}
	{{ $desc := print "A simple link generator to search using **" $name "**." }}
	{{ $usage := print "**Usage:**" "```" (lower .Cmd) " <keywords>```" }}
	{{ $embed := cembed 
		"title" "Search"
		"description" (print $desc "\n\n" $usage )
		"footer" (sdict "text" "Supported: google, duck, brave, wiki")
	}}
	{{ sendMessage nil $embed }}
{{ end }}

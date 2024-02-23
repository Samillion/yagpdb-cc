{{ $engines := cslice
	(sdict "trigger" "google" "name" "Google" "url" "https://www.google.com/search?hl=en&q=")
	(sdict "trigger" "duck" "name" "DuckDuckGo" "url" "https://duckduckgo.com/?q=")
	(sdict "trigger" "brave" "name" "Brave Search" "url" "https://search.brave.com/search?q=")
}}

{{ $c := reReplace .ServerPrefix .Cmd "" | lower }}
{{ $name := "" }}
{{ $url := "" }}

{{ range $engines }}
	{{- if eq (lower .trigger) $c -}}
		{{- $name = .name -}}
		{{- $url = .url -}}
		{{ break }}
	{{- end -}}
{{- end }}

{{ $msg := trimSpace .StrippedMsg }}

{{ if $msg }}
	{{ if not $url }}
		{{ sendMessage nil "**Error:** Search engine __$url__ is not set." }}
		{{ return }}
	{{ end }}
	{{ $keys := (urlquery ($msg)) }}
	{{ $link := print $url $keys }}
	{{ sendMessage nil $link }}
{{ else }}
	{{ if not $name }}
		{{ sendMessage nil "**Error:** Search engine __$name__ is not set." }}
		{{ return }}
	{{ end }}
	{{ $desc := print "A simple link generator to search using **" $name "** because you're too lazy to open the browser and do it there." }}
	{{ $usage := print "**Usage:**" "```" (lower .Cmd) " <keywords>```" }}
	{{ $embed := cembed 
		"title" "Search"
		"description" (print $desc "\n\n" $usage )
		"footer" (sdict "text" "Supported: google, duck, brave")
	}}
	{{ sendMessage nil $embed }}
{{ end }}

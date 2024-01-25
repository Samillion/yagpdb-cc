{{ $engines := cslice
	(sdict "trigger" "google" "name" "Google" "url" "https://www.google.com/search?hl=en&q=")
	(sdict "trigger" "duck" "name" "DuckDuckGo" "url" "https://duckduckgo.com/?q=")
	(sdict "trigger" "steam" "name" "Steam" "url" "https://store.steampowered.com/search/?term=")
}}

{{ $c := reReplace .ServerPrefix .Cmd "" | lower }}
{{ $name := "" }}
{{ $url := "" }}

{{ range $engines }}
	{{- if eq .trigger $c -}}
		{{- $name = .name -}}
		{{- $url = .url -}}
	{{- end -}}
{{- end }}

{{ $msg := trimSpace .StrippedMsg }}

{{ if $msg }}
	{{ $keys := (urlquery ($msg)) }}
	{{ $link := print $url $keys }}
	{{ sendMessage nil $link }}
{{ else }}
	{{ $desc := print "A simple link generator to search using **" $name "** because you're too lazy to open the browser and do it there." }}
	{{ $usage := print "**Usage:**" "```" (lower .Cmd) " <keywords>```" }}
	{{ $embed := cembed 
		"title" "Search"
		"description" (print $desc "\n\n" $usage )
		"footer" (sdict "text" "Supported: google, duck, steam")
	}}
	{{ sendMessage nil $embed }}
{{ end }}

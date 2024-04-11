{{ $default := 15 }}
{{ $max := 60 }}

{{ $lower := "abcdefghijklmnopqrstuvwxyz" }}
{{ $upper := upper $lower }}
{{ $alpha := print $lower $upper }}
{{ $num := "0123456789" }}
{{ $sym := "~!@#$%&*" }}

{{ $chars := sdict 
	"alpha" $alpha
	"lower" $lower
	"upper" $upper
	"num" $num
	"sym" $sym
	"mix" (print $alpha $num $sym)
}}

{{ $amount := $default }}
{{ $list := $chars.mix }}
{{ $args := .CmdArgs }}
{{ $check := cslice "mix" }}

{{ if $args }}
	{{ $list = "" }}
	{{ $check = cslice }}
	{{ range $args }}
		{{- if reFind `^\d{1,}$` . -}}
			{{- $amount = . | toInt -}}
		{{- else -}}
			{{- $type := lower . -}}
			{{/* mix has all types, no need for other args */}}
			{{- if eq $type "mix" -}}
				{{- $list = $chars.mix -}}
				{{- $check = cslice "mix" -}}
				{{- break -}}
			{{- end -}}
			{{- if $chars.HasKey $type -}}
				{{- if not (in $check $type) -}}
					{{- $list = print $list ($chars.Get $type) -}}
					{{- $check = $check.Append $type -}}
				{{- end -}}
			{{- end -}}
		{{- end -}}
	{{- end }}
	
	{{/* alpha already has lower and upper */}}
	{{ if and (in $check "alpha") (or (in $check "upper") (in $check "lower")) }}
		{{ $list = "" }}
		{{ $new := cslice }}
		{{ range $check }}
			{{- if not (or (eq . "upper" "lower") (in $new .)) -}}
				{{- $list = print $list ($chars.Get .) -}}
				{{- $new = $new.Append . -}}
			{{- end -}}
		{{- end }}
		{{ $check = $new }}
	{{ end }}
	
	{{ if not $list }}
		{{ $list = $chars.mix }}
		{{ $check = cslice "mix" }}
	{{ end }}
{{ end }}

{{ $list = split $list "" }}
{{ $result := "" }}

{{ if or (lt $amount 1) (gt $amount $max) }}
	{{ $amount = $default }}
{{ end }}

{{ range (seq 0 $amount) }}
    {{- $result = print $result (index (shuffle $list) 0) -}}
{{- end }}

{{ $message := print 
	"```" $result "```" "\n"
	"**Generated output:**" "\n"
	"- Count: " $amount " characters" "\n" 
	"- Type: " (joinStr ", " $check.StringSlice) "\n"
}}

{{ $embed := cembed 
	"title" "Random String"
	"description" $message
	"footer" (sdict "text" "Options: alpha, lower, upper, num, sym, mix")
	"color" 10530378
}}

{{ sendMessage nil $embed }}

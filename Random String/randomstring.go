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
{{ $check := cslice }}

{{ if $args }}
	{{ $list = "" }}
	{{ range $args }}
		{{- if reFind `^\d{1,}$` . -}}
			{{- $amount = . | toInt -}}
		{{- else -}}
			{{- $type := lower . -}}
			{{/* mix has all types, break loop */}}
			{{- if eq $type "mix" -}}
				{{- $list = $chars.mix -}}
				{{- $check = cslice -}}
				{{- break -}}
			{{- end -}}
			{{- if $chars.HasKey $type -}}
				{{- if not (in $check $type) -}}
					{{- $list = print $list ($chars.Get $type) -}}
				{{- end -}}
				{{- $check = $check.Append $type -}}
			{{- end -}}
		{{- end -}}
	{{- end }}
	
	{{/* alpha already has lower and upper, so need to repeat it */}}
	{{ if and (in $check "alpha") (or (in $check "upper") (in $check "lower")) }}
		{{ $list = "" }}
		{{ range $check }}
			{{- if not (eq . "upper" "lower") -}}
				{{- $list = print $list ($chars.Get .) -}}
			{{- end -}}
		{{- end }}
	{{ end }}
	
	{{ if not $list }}
		{{ $list = $chars.mix }}
	{{ end }}
{{ end }}

{{ $list = split $list "" }}

{{ if or (lt $amount 1) (gt $amount $max) }}
	{{ $amount = $default }}
{{ end }}

{{ $result := "" }}

{{ range (seq 0 $amount) }}
    {{- $result = print $result (index (shuffle $list) 0) -}}
{{- end }}

{{ $result = print "```" $result "```" }}

{{ $embed := cembed 
	"title" "Random String"
	"description" (print "Generated an output with " $amount " characters" $result)
	"footer" (sdict "text" "Options: alpha, lower, upper, num, sym, mix")
	"color" 10530378
}}

{{ sendMessage nil $embed }}

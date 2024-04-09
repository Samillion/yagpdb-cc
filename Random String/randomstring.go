{{ $lower := "abcdefghijklmnopqrstuvwxyz" }}
{{ $upper := upper $lower }}
{{ $alpha := print $lower $upper }}
{{ $num := "0123456789" }}
{{ $sym := "~!@#$%&*" }}
{{ $default := 15 }}
{{ $max := 60 }}

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

{{ if $args }}
	{{ $first := index $args 0 }}
	{{ $second := "" }}
	
	{{ if gt (len $args) 1 }}
		{{ $second = index $args 1 }}
	{{ end }}
	
	{{ if reFind `^\d{1,2}$` $first }}
		{{ $amount = $first | toInt }}
	{{ else if reFind `^\d{1,2}$` $second }}
		{{ $amount = $second | toInt }}
	{{ end }}
	
	{{ $list = or ($chars.Get (lower $first)) ($chars.Get (lower $second)) $chars.mix }}

	{{ if eq $first $second }}
		{{ $amount = $default }}
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

{{ $explain := print 
	"**Options:**" "\n"
	"- alpha, lower, upper, num, sym, mix" "\n\n"
	"**Character Limits:**" "\n"
	"- Minimum 1, maximum " $max "\n"
}}

{{ $embed := cembed 
	"title" "Random String"
	"description" (print $amount " characters" $result "\n" $explain)
	"footer" (sdict "text" "Usage: -random <amount:int> <option:str>")
}}

{{ sendMessage nil $embed }}

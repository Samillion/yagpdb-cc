{{ $lower := "abcdefghijklmnopqrstuvwxyz" }}
{{ $upper := upper $lower }}
{{ $alpha := print $lower $upper }}
{{ $num := "0123456789" }}
{{ $sym := "~!@#$%&*" }}
{{ $max := 60 }}

{{ $chars := sdict 
	"alpha" $alpha
	"num" $num
	"sym" $sym
	"upper" $upper
	"lower" $lower
	"mix" (print $alpha $num $sym)
}}

{{ $amount := 15 }}
{{ $list := $chars.mix }}
{{ $args := .CmdArgs }}

{{ if $args }}
	{{ $amount = index $args 0 | toInt }}
	{{ if gt (len $args) 1 }}
		{{ $type := index $args 1 | lower }}
		{{ if $chars.HasKey $type }}
			{{ $list = $chars.Get $type }}
		{{ end }}
	{{ end }}
	{{ if $chars.HasKey (index $args 0 | lower) }}
		{{ $amount = 15 }}
		{{ $list = $chars.Get (index $args 0 | lower) }}
	{{ end }}
{{ end }}

{{ $list = split $list "" }}

{{ if or (lt $amount 1) (gt $amount $max) }}
	{{ $amount = 15 }}
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

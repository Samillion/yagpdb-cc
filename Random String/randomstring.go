{{ $list := split "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()_-+={[}]<,>.?" "" }}
{{ $text := "" }}
{{ $amount := 15 }}
{{ $note := "" }}

{{ if gt (len .Args) 1 }}
	{{ $input := toInt (index .CmdArgs 0) }}
	
	{{ if and (gt $input 0 ) (le $input 60) }}
		{{ $amount = $input }}
	{{ else if or (lt $input 1 ) (gt $input 60) }}
		{{ $note = (print 
			"**Note:**" "\n" 
			"- Character count min is 1 and max is 60." "\n" 
			"- Your input is: **" $input "**" "\n" 
			"- Defaulted back to 15 chars output." 
		) }}
	{{ end }}
{{ end }}

{{ range (seq 0 $amount) }}
    {{- $text = print $text (index (shuffle $list) 0) -}}
{{- end }}

{{ $text = print "**Random String:** (" $amount " chars)" "\n" "```" $text "```" }}
{{ $text = print $text "\n" $note }}
{{ sendMessage nil $text }}

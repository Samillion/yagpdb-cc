{{ $ch := 1198221854121402468 }}
{{ $cmd := .Cmd }}
{{ $msg := trimSpace .StrippedMsg }}
{{ $usage := print "**Usage:**" "\n" "```" $cmd " your confession```" }}

{{ $explain := print 
	"This commnd will relay your message **__anonymously__** into the <#" $ch "> channel." "\n\n"
	"**Explanation:**" "\n"
	"- Once you use `-confess` it will post your message in the relative channel." "\n"
	"- Your message will be deleted instantly, leaving only the anonymous one for others to read." "\n"
	"- If the confession is less than 50 characters, nothing will be posted." "\n"
	"- For your comfort, feel free to review the code of this command [here](https://github.com/Samillion/yagpdb-cc/tree/main/Confession)."
}}

{{ $main := cembed 
	"title" "Confession"
	"description" (print $explain "\n\n" $usage)
}}

{{ if not (getChannel $ch) }}
	{{ sendMessage nil (print "**Error:** Please edit the command code and adjust `$ch` to a valid confession channel ID.") }}
	{{ return }}
{{ end }}

{{ if $msg }}
	{{ if gt (len $msg) 50 }}
		{{ $note := print "Use " $cmd " to post your own, anonymously." }}
		{{ $embed := cembed 
			"title" "Anonymous Confession"
			"description" $msg 
			"color" 8421888
			"footer" (sdict "text" $note)
		}}
		{{ sendMessage $ch $embed }}
	{{ end }}
	{{ deleteTrigger 0 }}
{{ else }}
	{{ sendMessage nil $main }}
{{ end }}

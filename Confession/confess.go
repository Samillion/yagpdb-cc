{{ $ch := 1200964297518497903 }}
{{ $cd := 600 }}

{{ $cmd := .Cmd }}
{{ $msg := trimSpace .StrippedMsg }}
{{ $usage := print "**Usage:**" "\n" "```" $cmd " your confession```" }}

{{ $explain := print 
	"This commnd will post your message **__anonymously__** to <#" $ch ">." "\n\n"
	"**Explanation:**" "\n"
	"- Once you use `-confess` it will post your message in the relative channel." "\n"
	"- Your message will be deleted instantly, leaving only the anonymous one for others to read." "\n"
	"- If the confession is less than 50 characters, nothing will be posted." "\n"
	"- A limit of one confession every 10 minutes, globally." "\n"
	"- Your initial post, even if it fails (ie: on cooldown), will instantly be deleted." "\n"
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

{{ $db := "confessions" }}
{{ $ccid := .CCID }}

{{ if $msg }}
	{{ if gt (len $msg) 50 }}
		{{ if not (dbGet $ccid $db) }}
			{{ $note := print "Use " $cmd " to post your own, anonymously." }}
			{{ $embed := cembed 
				"title" "Anonymous Confession"
				"description" $msg 
				"color" 8421888
				"footer" (sdict "text" $note)
			}}
			{{ sendMessage $ch $embed }}
			{{ dbSetExpire $ccid $db "cooldown" $cd }}
		{{ else }}
			{{ $x := sendMessageRetID nil "There is a global 10 minutes cooldown per confession." }}
			{{ deleteMessage nil $x 10 }}
		{{ end }}
	{{ else }}
		{{ $x := sendMessageRetID nil "Your confession must contain at least 50 characters." }}
		{{ deleteMessage nil $x 10 }}
	{{ end }}
	{{ deleteTrigger 0 }}
{{ else }}
	{{ sendMessage nil $main }}
{{ end }}

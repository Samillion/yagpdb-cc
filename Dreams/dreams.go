{{ $ch := 1200964297518497903 }}
{{ $minimum := 20 }}
{{ $cd := 600 }}

{{ $cmd := .Cmd }}
{{ $msg := trimSpace .StrippedMsg }}
{{ $usage := print "**Usage:**" "\n" "```" $cmd " your dream```" }}

{{ $explain := print 
	"Had an interesting dream last night? Share it with us by using this command." "\n\n"
	"This command will post your message **__anonymously__** to <#" $ch ">." "\n\n"
	"**Explanation:**" "\n"
	"- Once you use `" $cmd "` it will post your message in the relative channel." "\n"
	"- Your message will be deleted instantly, leaving only the anonymous one for others to read." "\n"
	"- If the dream is less than " $minimum " characters, nothing will be posted." "\n"
	"- A limit of one dream every 10 minutes, globally." "\n"
	"- Your initial post, even if it fails (ie: on cooldown), will instantly be deleted." "\n"
	"- There are **__zero logs__**, it is 100% anonymous. (even Discord Audit Logs)" "\n"
	"- Feel free to review the code of this command [here](https://github.com/Samillion/yagpdb-cc/tree/main/Dreams)."
}}

{{ $main := cembed 
	"title" "Anonymous Dream"
	"description" (print $explain "\n\n" $usage)
}}

{{ if not (getChannel $ch) }}
	{{ sendMessage nil (print "**Error:** Please edit the command code and adjust `$ch` to a valid channel ID.") }}
	{{ deleteTrigger 0 }}
	{{ return }}
{{ end }}

{{ $db := "dreams" }}
{{ $ccid := .CCID }}

{{ if $msg }}
	{{ if gt (len $msg) $minimum }}
		{{ if not (dbGet $ccid $db) }}
			{{ $note := print "Use " $cmd " to post your own, anonymously." }}
			{{ $embed := cembed 
				"title" "Anonymous Dream"
				"description" $msg 
				"color" 8421888
				"footer" (sdict "text" $note)
			}}
			{{ sendMessage $ch $embed }}
			{{ dbSetExpire $ccid $db "cooldown" $cd }}
			
			{{ $cdtext := print 
				"A global cooldown is active for 10 minutes, meaning no one can post a dream during that time." "\n\n"
				"__This messaage will be deleted automatically when the cooldown is over__." "\n\n"
				"**Why a global cooldown?**" "\n"
				"If cooldown was per user, it would be a way to trace who posted the dream. A global cooldown guarantees anonymity."
			}}
			{{ $cdtime := cembed 
				"title" "Cooldown"
				"description" $cdtext
				"color" 11993101
			}}
			{{ $x := sendMessageRetID $ch $cdtime }}
			{{ deleteMessage $ch $x $cd }}
		{{ else }}
			{{ $x := sendMessageRetID nil "There is a global 10 minutes cooldown per confession." }}
			{{ deleteMessage nil $x 10 }}
		{{ end }}
	{{ else }}
		{{ $x := sendMessageRetID nil (print "Your dream post must contain at least " $minimum " characters.") }}
		{{ deleteMessage nil $x 10 }}
	{{ end }}
	{{ deleteTrigger 0 }}
{{ else }}
	{{ sendMessage nil $main }}
{{ end }}

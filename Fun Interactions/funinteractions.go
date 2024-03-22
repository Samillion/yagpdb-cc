{{ $cooldown := 60 }}

{{ $type := sdict 
	"hug" (sdict 
		"gifs" (cslice 
			"https://c.tenor.com/1fIciLLNlSQAAAAC/tenor.gif"
			"https://c.tenor.com/-7RUZhALnFcAAAAC/tenor.gif"
			"https://c.tenor.com/VQURjAnQwRkAAAAC/tenor.gif"
			"https://c.tenor.com/N_palirBY7EAAAAC/tenor.gif"
			"https://c.tenor.com/Ta7V9zu43h8AAAAC/tenor.gif"
			"https://c.tenor.com/sUdlhZ9Vx8AAAAAd/tenor.gif"
			"https://c.tenor.com/dzs3WPLgqoYAAAAC/tenor.gif"
			"https://c.tenor.com/5UBsAKy5NHIAAAAC/tenor.gif"
			"https://c.tenor.com/4yWA6WK5ugAAAAAC/tenor.gif"
		)
		"self" "https://c.tenor.com/mNZS9vIzMZYAAAAC/tenor.gif"
		"phrase" "hugs"
		"selfPhrase" "hugs themselves"
	)
	"fart" (sdict 
		"gifs" (cslice 
			"https://c.tenor.com/4nUZ8ZECPGoAAAAC/tenor.gif"
			"https://c.tenor.com/7j_zisuASDoAAAAC/tenor.gif"
			"https://c.tenor.com/5SUNJKDXrUcAAAAC/tenor.gif"
			"https://c.tenor.com/iv5qRp14eAcAAAAC/tenor.gif"
			"https://c.tenor.com/r821mUXnUkQAAAAC/tenor.gif"
			"https://c.tenor.com/bTMASnERNzcAAAAC/tenor.gif"
			"https://c.tenor.com/Ua7tk_pmkcIAAAAC/tenor.gif"
		)
		"self" "https://c.tenor.com/mU3s40s1j2AAAAAd/tenor.gif"
		"phrase" "farts on"
		"selfPhrase" "smells their own fart"
	)
	"slap" (sdict 
		"gifs" (cslice 
			"https://media.tenor.com/tKF3HiguDmEAAAAC/wrrruutchxxxxiii-slapt.gif"
			"https://media.tenor.com/__oycZBexeAAAAAC/slap.gif"
			"https://media.tenor.com/W0K0vteByOoAAAAC/slap-in-the-face-slap.gif"
			"https://media.tenor.com/GBShVmDnx9kAAAAC/anime-slap.gif"
			"https://media.tenor.com/bHqFiKBJGmoAAAAC/korone-dog-form.gif"
			"https://media.tenor.com/rvXlkAdJXIgAAAAC/paper-fan-slap.gif"
			"https://media.tenor.com/ra17G61QRQQAAAAC/tapa-slap.gif"
			"https://media.tenor.com/eU5H6GbVjrcAAAAC/slap-jjk.gif"
		)
		"self" "https://media.tenor.com/rwipJN-E7okAAAAC/slap-slapping-self.gif"
		"phrase" "slaps"
		"selfPhrase" "why are you slapping yourself?"
	)
	"spank" (sdict 
		"gifs" (cslice 
			"https://media.tenor.com/JMq5b6MWcPkAAAAC/spanking-spank.gif"
			"https://media.tenor.com/WDMDojVJY5EAAAAC/bunny-spanking.gif"
			"https://media.tenor.com/NcdgzWMhZXkAAAAC/spank-playful.gif"
			"https://media.tenor.com/V03y9g2bJ_kAAAAC/gto-retro-anime.gif"
			"https://media.tenor.com/Sp7yE5UzqFMAAAAC/spank-slap.gif"
			"https://media.tenor.com/CAesvxP0KyEAAAAd/shinobu-kocho-giyuu-tomioka.gif"
			"https://media.tenor.com/lNyexavajUEAAAAd/anime-spanking.gif"
			"https://media.tenor.com/LMKZH2bDKHsAAAAC/spank-anime.gif"
		)
		"self" "https://c.tenor.com/nn7A6kVJyzwAAAAC/tenor.gif"
		"phrase" "spanks"
		"selfPhrase" "spanks own butt"
	)
}}

{{/* Do not edit beyond this point unless you know what you're doing. */}}
{{ $usage := print "```Usage: " .Cmd " @user```" }}
{{ $cmd := slice .Cmd 1 (len .Cmd) | lower }}
{{ $set := or ($type.Get $cmd) $type.hug }}

{{ $on := "" }}
{{ if eq $cmd "fart" }}{{ $on = " on " }}{{ end }}
{{ $embed := sdict 
	"title" (print (title $cmd) " Command")
	"description" (print "Do you want to " $cmd $on " someone? Then this is the perfect command for you." "\n\n" $usage)
}}

{{ if .CmdArgs }}
	{{ $user := .User.ID }}
	{{ $target := 0 }}
	
	{{ if $x := .Message.Mentions }}
		{{ $target = (index $x 0).ID }}
		{{ if and (gt (len $x) 1) (.Message.ReferencedMessage) }}
			{{ $target = (index $x 1).ID }}
		{{ end }}
	{{ end }}
	
	{{ $target = userArg $target }}

	{{ if $target }}
		{{ $target = $target.ID }}
		{{ $db := print "fun_" $cmd }}
		{{ if not (dbGet $user $db) }}
			{{ if gt (toInt $cooldown) 0 }}
				{{ dbSetExpire $user $db "cooldown" $cooldown }}
			{{ end }}
			{{ $image := index $set.gifs (randInt (len $set.gifs)) }}
			{{ $desc := (print "<@" $user "> " $set.phrase " <@" $target ">.") }}
			{{ if eq $target $user }}
				{{ $desc = (print "<@" $user "> " $set.selfPhrase ".") }}
				{{ $image = $set.self }}
			{{ end }}
			{{ $col := 16777215 }}
			{{ $p := 0 }}
			{{ $r := .Member.Roles }}
			{{ range .Guild.Roles }}
				{{- if and (in $r .ID) (.Color) (lt $p .Position) -}}
					{{- $p = .Position -}}
					{{- $col = .Color -}}
				{{- end -}}
			{{- end }}
			{{ $embed = sdict 
				"description" $desc
				"image" (sdict "url" $image)
				"color" $col
			}}
		{{ end }}
	{{ else }}
		{{ $embed = sdict 
			"title" "Who dis?"
			"description" (print "Defined user is not a member in this server." "\n\n" $usage)
		}}
	{{ end }}
{{ end }}

{{ sendMessage nil (cembed $embed) }}

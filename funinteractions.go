{{/*
	Basic interactions for meaningless fun purposes.
	
	See <https://github.com/Samillion/yagpdb-fun-interactions> for more information.

	Author: Samillion <https://github.com/Samillion>
*/}}

{{ $cooldown := 30 }}

{{ $hugs := cslice
	"https://c.tenor.com/1fIciLLNlSQAAAAC/tenor.gif"
	"https://c.tenor.com/-7RUZhALnFcAAAAC/tenor.gif"
	"https://c.tenor.com/VQURjAnQwRkAAAAC/tenor.gif"
	"https://c.tenor.com/N_palirBY7EAAAAC/tenor.gif"
	"https://c.tenor.com/Ta7V9zu43h8AAAAC/tenor.gif"
	"https://c.tenor.com/sUdlhZ9Vx8AAAAAd/tenor.gif"
	"https://c.tenor.com/dzs3WPLgqoYAAAAC/tenor.gif"
	"https://c.tenor.com/5UBsAKy5NHIAAAAC/tenor.gif"
	"https://c.tenor.com/4yWA6WK5ugAAAAAC/tenor.gif"
}}
{{ $selfHug := "https://c.tenor.com/mNZS9vIzMZYAAAAC/tenor.gif" }}

{{ $farts := cslice
	"https://c.tenor.com/4nUZ8ZECPGoAAAAC/tenor.gif"
	"https://c.tenor.com/7j_zisuASDoAAAAC/tenor.gif"
	"https://c.tenor.com/5SUNJKDXrUcAAAAC/tenor.gif"
	"https://c.tenor.com/iv5qRp14eAcAAAAC/tenor.gif"
	"https://c.tenor.com/r821mUXnUkQAAAAC/tenor.gif"
	"https://c.tenor.com/bTMASnERNzcAAAAC/tenor.gif"
	"https://c.tenor.com/Ua7tk_pmkcIAAAAC/tenor.gif"
}}
{{ $selfFart := "https://c.tenor.com/mU3s40s1j2AAAAAd/tenor.gif" }}

{{ $slaps := cslice
	"https://media.tenor.com/tKF3HiguDmEAAAAC/wrrruutchxxxxiii-slapt.gif"
	"https://media.tenor.com/__oycZBexeAAAAAC/slap.gif"
	"https://media.tenor.com/W0K0vteByOoAAAAC/slap-in-the-face-slap.gif"
	"https://media.tenor.com/GBShVmDnx9kAAAAC/anime-slap.gif"
	"https://media.tenor.com/bHqFiKBJGmoAAAAC/korone-dog-form.gif"
	"https://media.tenor.com/rvXlkAdJXIgAAAAC/paper-fan-slap.gif"
	"https://media.tenor.com/ra17G61QRQQAAAAC/tapa-slap.gif"
	"https://media.tenor.com/eU5H6GbVjrcAAAAC/slap-jjk.gif"
}}
{{ $selfSlap := "https://media.tenor.com/rwipJN-E7okAAAAC/slap-slapping-self.gif" }}

{{ $spanks := cslice
	"https://media.tenor.com/JMq5b6MWcPkAAAAC/spanking-spank.gif"
	"https://media.tenor.com/WDMDojVJY5EAAAAC/bunny-spanking.gif"
	"https://media.tenor.com/NcdgzWMhZXkAAAAC/spank-playful.gif"
	"https://media.tenor.com/V03y9g2bJ_kAAAAC/gto-retro-anime.gif"
	"https://media.tenor.com/Sp7yE5UzqFMAAAAC/spank-slap.gif"
	"https://media.tenor.com/CAesvxP0KyEAAAAd/shinobu-kocho-giyuu-tomioka.gif"
	"https://media.tenor.com/lNyexavajUEAAAAd/anime-spanking.gif"
	"https://media.tenor.com/LMKZH2bDKHsAAAAC/spank-anime.gif"
}}
{{ $selfSpank := "https://c.tenor.com/nn7A6kVJyzwAAAAC/tenor.gif" }}

{{/* Do not edit beyond this point unless you know what you're doing. */}}
{{ $usage := (print "```Usage: " .Cmd " @user```") }}
{{ $cTitle := reReplace .ServerPrefix .Cmd "" }}
{{ $images := cslice }}
{{ $phrase := "" }}
{{ $selfPhrase := "" }}
{{ $self := "" }}

{{ if eq $cTitle "hug" }}
	{{ $images = $hugs }}
	{{ $self = $selfHug }}
	{{ $phrase = "hugs" }}
	{{ $selfPhrase = "self hugs" }}
{{ else if eq $cTitle "fart" }}
	{{ $images = $farts }}
	{{ $self = $selfFart }}
	{{ $phrase = "farts on" }}
	{{ $selfPhrase = "smells their own fart" }}
{{ else if eq $cTitle "slap" }}
	{{ $images = $slaps }}
	{{ $self = $selfSlap }}
	{{ $phrase = "slaps" }}
	{{ $selfPhrase = "why are you slapping yourself?" }}
{{ else if eq $cTitle "spank" }}
	{{ $images = $spanks }}
	{{ $self = $selfSpank }}
	{{ $phrase = "spanks" }}
	{{ $selfPhrase = "spanks own butt" }}
{{ end }}

{{ if gt (len .Args) 1 }}
	{{ $user := .User.ID }}
	{{ $target := 0 }}
	
	{{ if $x := .Message.Mentions }}
		{{ $target = (index $x 0).ID }}
	{{ end }}
	
	{{ $target = userArg ($target) }}

	{{ if $target }}
		{{ $target = $target.ID }}
		{{ $db := (str $cTitle) }}
		{{ if not (dbGet $user $db) }}
			{{ if gt $cooldown 0 }}
				{{ dbSetExpire $user $db "cooldown" $cooldown }}
			{{ end }}
			{{ $image := index $images (randInt (len $images)) }}
			{{ $desc := (print "<@" $user "> " $phrase " <@" $target ">.") }}
			{{ if eq $target $user }}
				{{ $desc = (print "<@" $user "> " $selfPhrase ".") }}
				{{ $image = $self }}
			{{ end }}
			{{ $col := 16777215 }}
			{{ $p := 0 }}
			{{ $r := .Member.Roles }}
			{{ range .Guild.Roles }}
				{{- if and (in $r .ID) (.Color) (lt $p .Position) }}
					{{- $p = .Position }}
					{{- $col = .Color }}
				{{- end -}}
			{{- end }}
			{{ $embed := cembed 
				"description" $desc
				"image" (sdict "url" $image)
				"color" $col
			}}
			{{ sendMessage nil $embed }}
		{{ end }}
	{{ else }}
		{{ $embed := cembed 
			"title" "Who dis?"
			"description" (print "Defined user is not a member in this server." "\n\n" $usage)
		}}
		{{ sendMessage nil $embed }}
	{{ end }}
{{ else }}
	{{ $on := "" }}
	{{ if eq $cTitle "fart" }}{{ $on = " on " }}{{ end }}
	{{ $embed := cembed 
		"title" (print (title $cTitle) " Command")
		"description" (print "Do you want to " $cTitle $on " someone? Then this is the perfect command for you." "\n\n" $usage)
	}}
	{{ sendMessage nil $embed }}
{{ end }}

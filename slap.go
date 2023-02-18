{{/*
	A basic slap command for meaningless fun purposes.
	
	See <https://github.com/Samillion/yagpdb-slap> for more information.

	Author: Samillion <https://github.com/Samillion>
*/}}

{{/* Adjust cooldown and GIFs (direct links) */}}
{{ $cooldown := 180 }}
{{ $images := cslice
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
{{ $botSlap := "https://media.tenor.com/J9flx980Ot4AAAAC/reggie-deflecting-criticism.gif" }}

{{/* Do not edit beyond this point unless you know what you're doing. */}}
{{ $usage := (print "```Usage: " .ServerPrefix "slap @user```") }}
{{ $args := parseArgs 1 $usage (carg "userid" "User ID or mention") }}
{{ $user := .User.ID }}
{{ $target := ($args.Get 0) }}
{{ if getMember $target }}
	{{ if not (dbGet $user "slap") }}
		{{ if gt $cooldown 0 }}
			{{ dbSetExpire $user "slap" "cooldown" $cooldown }}
		{{ end }}
		{{ $image := index $images (randInt (len $images)) }}
		{{ $desc := (print "<@" $user "> slapped <@" $target ">.") }}
		{{ if eq $target $user }}
			{{ $desc = (print "<@" $user ">, why are you slapping yourself?") }}
			{{ $image = $selfSlap }}
		{{ else if eq $target 204255221017214977 }}
			{{ $desc = (print "<@" $user ">, you think you can slap me? :rofl:") }}
			{{ $image = $botSlap }}
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
	{{ sendMessage nil (print "Defined user is not a member in this server." "\n" $usage) }}
{{ end }}

{{/*
	A basic slap command for meaningless fun purposes.
	Usage: -slap @user
	
	Default cooldown (per user) is 3 mins (180), for no limit set it to 0.
	
	See https://github.com/Samillion/yagpdb-slap for more information.
*/}}
{{ $usage := "```Usage: -slap @user```" }}
{{ $args := parseArgs 1 $usage (carg "userid" "User ID or mention") }}
{{ $cooldown := 180 }}
{{ $user := .User.ID }}
{{ $target := $args.Get 0 }}
{{ if getMember $target }}
	{{ $userNick := (getMember $user).Nick }}
	{{ $targetNick := (getMember $target).Nick }}
	{{ if not (dbGet $user "slap") }}
		{{ if gt $cooldown 0 }}
			{{ dbSetExpire $user "slap" "cooldown" $cooldown }}
		{{ end }}
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
		{{ $image := index $images (randInt (len $images)) }}
		{{ $desc := (print "**" $userNick "** slapped **" $targetNick "**.") }}
		{{ if eq $target $user }}
			{{ $desc = (print "**" $userNick "**, why are you slapping yourself?") }}
			{{ $image = "https://media.tenor.com/rwipJN-E7okAAAAC/slap-slapping-self.gif" }}
		{{ else if eq $target 204255221017214977 }}
			{{ $desc = (print "**" $userNick "**, you think you can slap me? :rofl:") }}
			{{ $image = "https://media.tenor.com/J9flx980Ot4AAAAC/reggie-deflecting-criticism.gif" }}
		{{ end }}
		{{/* Get user color: <https://yagpdb-cc.github.io/code-snippets/get-username-color> - Start */}}
		{{ $col := 16777215 }}
		{{ $p := 0 }}
		{{ $r := .Member.Roles }}
		{{ range .Guild.Roles }}
			{{- if and (in $r .ID) (.Color) (lt $p .Position) }}
				{{- $p = .Position }}
				{{- $col = .Color }}
			{{- end -}}
		{{- end }}
		{{/* Get user color - End */}}
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

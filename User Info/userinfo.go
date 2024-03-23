{{ $member := .Member }}
{{ $user := .User }}
{{ $args := .CmdArgs }}

{{ if $args }}
	{{ $member = getMember (index $args 0) }}
	
	{{ if not $member }}
		{{ sendMessage nil "User must be a member in this server." }}
		{{ return }}
	{{ end }}
	
	{{ $user = $member.User }}
{{ end }}

{{ $bot := "No" }}
{{ if $user.Bot }}{{ $bot = "Yes" }}{{ end }}

{{ $created := snowflakeToTime $user.ID }}
{{ $age := humanizeTimeSinceDays ($created) }}

{{ $joined := $member.JoinedAt.Parse }}
{{ $joinedAge := humanizeTimeSinceDays ($joined) }}

{{ $status := (index (exec "whois" $user.ID).Fields 6).Value }}

{{ $userAvatar := $member.User.AvatarURL "4096" }}
{{ $serverAvatar := "Not set" }}

{{ if $member.Avatar }}
	{{ $serverAvatar = (print "[Link](" ($member.AvatarURL "4096") ")") }}
{{ end }}

{{ $globalName := or $member.User.Globalname "Not set" }}
{{ $nickname := or $member.Nick "Not set" }}

{{ $embed := cembed 
	"fields" (cslice
		(sdict "name" "Username" "value" $user.String "inline" true)
		(sdict "name" "Display Name" "value" $globalName "inline" true)
		(sdict "name" "Server Name" "value" $nickname "inline" true)
		
		(sdict "name" "Tag" "value" $user.Mention "inline" true)
		(sdict "name" "User Avatar" "value" (print "[Link](" $userAvatar ")") "inline" true)
		(sdict "name" "Server Avatar" "value" $serverAvatar "inline" true)
		
		(sdict "name" "Account Created" "value" ($created.Format "02 Jan, 2006") "inline" true)
		(sdict "name" "Account Age" "value" $age "inline" true)
		(sdict "name" "Bot" "value" $bot "inline" true)
		
		(sdict "name" "Joined Server" "value" ($joined.Format "02 Jan, 2006") "inline" true)
		(sdict "name" "Joined Age" "value" $joinedAge "inline" true)
		(sdict "name" "User ID" "value" (printf "%d" $user.ID) "inline" true)
		
		(sdict "name" "Status" "value" $status)
	)
	"thumbnail" (sdict "url" $userAvatar)
}}

{{ sendMessage nil $embed }}

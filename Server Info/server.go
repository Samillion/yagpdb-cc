{{ $levels := cslice
	"None: Unrestricted"
	"Low: Must have a verified email on their Discord account."
	"Medium: Must also be registered on Discord for longer than 5 minutes."
	"(╯°□°）╯︵ ┻━┻: Must also be a member of this server for longer than 10 minutes."
	"┻━┻ ﾐヽ(ಠ益ಠ)ノ彡┻━┻: Must have a verified phone on their Discord account."
}}

{{ $afk := "Not set" }}
{{ if .Guild.AfkChannelID }}
	{{ $afk = printf "- Channel: <#%d>\n- Timeout: %s"
		.Guild.AfkChannelID
		(humanizeDurationSeconds (toDuration (mult .Guild.AfkTimeout .TimeSecond)))
	}}
{{ end }}

{{ $icon := .Guild.IconURL "256" }}
{{ $owner := userArg .Guild.OwnerID }}
{{ $createdAt := snowflakeToTime .Guild.ID }}

{{ $embed := cembed
	"thumbnail" (sdict "url" $icon)
	"fields" (cslice
		(sdict "name" "Server:" "value" .Guild.Name)
		(sdict "name" "Server ID:" "value" (str .Guild.ID))
		(sdict "name" "Verification Level" "value" (index $levels .Guild.VerificationLevel))
		(sdict "name" "Members" "value" (printf "- Total: %d Members\n- Online: %d Members" .Guild.MemberCount onlineCount))
		(sdict "name" "DB Limit" "value" (print dbCount "/" (mult .Guild.MemberCount 50 1)))
		(sdict "name" "Owner" "value" $owner.Mention)
		(sdict "name" "AFK" "value" $afk)
	)
	"footer" (sdict "text" "Created at")
	"timestamp" $createdAt
}}

{{ if .CmdArgs }}
	{{ if eq (index .CmdArgs 0) "icon" }}
		{{ $embed = cembed "title" "Server Icon" "image" (sdict "url" $icon) }}
	{{ end }}
{{ end }}

{{ sendMessage nil $embed }}

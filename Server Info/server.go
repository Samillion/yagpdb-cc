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

{{ $icon := "" }}
{{ if .Guild.Icon }}
	{{ $ext := "webp" }}
	{{ if eq (slice .Guild.Icon 0 2) "a_" }}{{ $ext = "gif" }}{{ end }}
	{{ $icon = printf "https://cdn.discordapp.com/icons/%d/%s.%s" .Guild.ID .Guild.Icon $ext }}
{{ end }}

{{ $owner := userArg .Guild.OwnerID }}
{{ $createdAt := div .Guild.ID 4194304 | add 1420070400000 | mult 1000000 | toDuration | (newDate 1970 1 1 0 0 0).Add }}

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

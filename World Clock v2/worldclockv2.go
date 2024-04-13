{{ $clocks := sdict
	"africa (central)" "Africa/Harare"
	"africa (east)" "Africa/Addis_Ababa"
	"africa (west)" "Africa/Casablanca"
	"asia (central)" "Asia/Bangkok"
	"asia (east)" "Asia/Tokyo"
	"asia (west)" "Asia/Kabul"
	"europe (eastern)" "EET"
	"europe (central)" "CET"
	"uk" "Europe/London"
	"us (eastern)" "US/Eastern"
	"us (pacific)" "US/Pacific"
	"utc" "Etc/UTC"
}}
{{ $cList := sdict
	"albania" "Europe/Tirane"
	"afghanistan" "Asia/Kabul"
	"angola" "Africa/Luanda"
	"argentina" "America/Argentina/Buenos_Aires"
	"armenia" "Asia/Yerevan"
	"austria" "Europe/Vienna"
	"australia" "Australia/Sydney"
	"bahrain" "Asia/Bahrain"
	"bangladesh" "Asia/Dhaka"
	"belgium" "Europe/Brussels"
	"bhutan" "Asia/Thimphu"
	"brazil" "America/Santarem"
	"bulgaria" "Europe/Sofia"
	"cambodia" "Asia/Phnom_Penh"
	"canada" "Canada/Eastern"
	"chad" "Africa/Ndjamena"
	"chile" "America/Santiago"
	"china" "Asia/Hong_Kong"
	"colombia" "America/Bogota"
	"croatia" "Europe/Zagreb"
	"cyprus" "Europe/Nicosia"
	"denmark" "Europe/Copenhagen"
	"ecuador" "America/Guayaquil"
	"egypt" "Africa/Cairo"
	"ethiopia" "Africa/Addis_Ababa"
	"eritrea" "Africa/Asmara"
	"estonia" "Europe/Tallinn"
	"fiji" "Pacific/Fiji"
	"finland" "Europe/Helsinki"
	"france" "Europe/Paris"
	"germany" "Europe/Berlin"
	"ghana" "Africa/Accra"
	"greece" "Europe/Athens"
	"guatemala" "America/Guatemala"
	"guinea" "Africa/Conakry"
	"hungary" "Europe/Budapest"
	"india" "IST"
	"indonesia" "Asia/Jakarta"
	"iran" "Asia/Tehran"
	"israel" "Asia/Jerusalem"
	"italy" "Europe/Rome"
	"japan" "Asia/Tokyo"
	"jordan" "Asia/Amman"
	"kenya" "Africa/Nairobi"
	"korea" "Asia/Seoul"
	"kosovo" "CET"
	"kuwait" "Asia/Kuwait"
	"laos" "Asia/Vientiane"
	"latvia" "Europe/Riga"
	"lebanon" "Asia/Beirut"
	"liberia" "Africa/Monrovia"
	"libya" "Africa/Tripoli"
	"luxembourg" "Europe/Luxembourg"
	"lithuania" "Europe/Vilnius"
	"macao" "Asia/Macao"
	"macedonia" "Europe/Skopje"
	"madagascar" "Indian/Antananarivo"
	"malawi" "Africa/Blantyre"
	"malaysia" "Asia/Kuala_Lumpur"
	"mali" "Africa/Bamako"
	"mauritania" "Africa/Nouakchott"
	"mauritius" "Indian/Mauritius"
	"mexico" "America/Mexico_City"
	"moldova" "Europe/Chisinau"
	"mongolia" "Asia/Ulaanbaatar"
	"morocco" "Africa/Casablanca"
	"mozambique" "Africa/Maputo"
	"namibia" "Africa/Windhoek"
	"netherlands" "Europe/Amsterdam"
	"nepal" "Asia/Kathmandu"
	"niger" "Africa/Niamey"
	"nigeria" "Africa/Lagos"
	"norway" "Europe/Oslo"
	"philippines" "Asia/Manila"
	"poland" "Europe/Warsaw"
	"portugal" "Europe/Lisbon"
	"qatar" "Asia/Qatar"
	"romania" "Europe/Bucharest"
	"rwanda" "Africa/Kigali"
	"saudi" "Asia/Riyadh"
	"senegal" "Africa/Dakar"
	"serbia" "Europe/Belgrade"
	"singapore" "Asia/Singapore"
	"slovakia" "Europe/Bratislava"
	"slovenia" "Europe/Ljubljana"
	"somalia" "Africa/Mogadishu"
	"spain" "Europe/Madrid"
	"sweden" "Europe/Stockholm"
	"switzerland" "Europe/Zurich"
	"syria" "Asia/Damascus"
	"taiwan" "Asia/Taipei"
	"tajikistan" "Asia/Dushanbe"
	"tanzania" "Africa/Dar_es_Salaam"
	"tunisia" "Africa/Tunis"
	"turkey" "Asia/Istanbul"
	"thailand" "Asia/Bangkok"
	"uae" "Asia/Dubai"
	"uganda" "Africa/Kampala"
	"ukraine" "Europe/Kiev"
	"usa" "US/Eastern"
	"uzbekistan" "Asia/Tashkent"
	"venezuela" "America/Caracas"
	"vietnam" "Asia/Ho_Chi_Minh"
	"yemen" "Asia/Aden"
	"zambia" "Africa/Lusaka"
	"zimbabwe" "Africa/Harare"
}}

{{ $cmd := .Cmd }}
{{ $args := .CmdArgs }}
{{ $embed := sdict 
	"title" ":alarm_clock:  World Clock" 
	"footer" (sdict "text" (print "Usage: " $cmd " help"))
}}

{{ define "reCountry" }}
	{{- if . -}}
		{{- $in := . -}}
		{{- if eq (len $in) 2 3 -}}
			{{- $in = upper $in -}}
		{{- else if $re := reFindAllSubmatches `(us)(\s+)\((eastern|pacific)\)` $in -}}
			{{- $re = index $re 0 -}}
			{{- $fi := index $re 1 -}}
			{{- $se := index $re 3 -}}
			{{- $in = (print (upper $fi) " (" (title $se) ")") -}}
		{{- end -}}
		{{- return (title $in) -}}
	{{- end -}}
{{ end }}

{{ if $args }}
	{{ $in := index $args 0 | lower }}
	{{ if eq $in "help" }}
		{{ $desc := print 
			"**Usage:**" "\n"
			"```" $cmd " ```"
			"```" $cmd " <country>```"
			"```" $cmd " <hour> <country>```"
			"```" $cmd " <country> <country>```" "\n"
			"Where `<hour>` is in 24h format and `<country>` is from a predefined list." "\n\n"
			"You can view the `<country>` list by using `" $cmd " list`" "\n\n"
			"**Examples:**" "\n"
			"```" $cmd " " "\n" "Lists current time for predefined timezones.```"
			"```" $cmd " France" "\n" "Lists current time in France.```"
			"```" $cmd " 14 Germany" "\n" "Lists the equivilant of 2 PM in Germany for different timezones.```"
			"```" $cmd " Italy Japan" "\n" "Compares time difference between Italy and Japan.```"
			"```" $cmd " 18 \"US (Eastern)\"" "\n" "For countries with space in their label, wrap it in quotes.```"
		}}
		{{ $embed.Set "description" $desc }}
	{{ else if eq $in "list" }}
		{{ $desc := "" }}
		{{ $count := 0 }}

		{{ range $k, $v := $cList }}
			{{- if eq (len $k) 2 3 -}}
				{{- $k = upper $k -}}
			{{- end -}}
		
			{{- $desc = joinStr ", " $desc (title $k) -}}
			{{- $count = add $count 1 -}}
		{{- end }}

		{{ $desc = print $desc "\n\n" "**Total:** " $count }}
		{{ $embed.Set "title" ":earth_africa:  Supported Countries" }}
		{{ $embed.Set "description" $desc }}
	{{ else if eq (len $args) 2 }}
		{{ $first := index .CmdArgs 0 | lower }}
		{{ $second := index .CmdArgs 1 | lower }}

		{{ if and (reFind `[\D]` $first) (reFind `[\D]` $second) }}
			{{ if or (and (not ($clocks.HasKey $first)) (not ($cList.HasKey $first))) (and (not ($clocks.HasKey $second)) (not ($cList.HasKey $second))) }}
				{{ sendMessage nil "Both inputs need to be from the supported countries list." }}
				{{ return }}
			{{ end }}

			{{ $fiOut := execTemplate "reCountry" $first }}
			{{ $seOut := execTemplate "reCountry" $second }}

			{{ $first = or ($clocks.Get $first) ($cList.Get $first) }}
			{{ $second = or ($clocks.Get $second) ($cList.Get $second) }}

			{{ $first = currentTime.In (loadLocation $first) }}
			{{ $second = currentTime.In (loadLocation $second) }}

			{{ $fiComp := newDate $first.Year $first.Month $first.Day $first.Hour (sub $first.Minute 1) $first.Second }}
			{{ $seComp := newDate $second.Year $second.Month $second.Day $second.Hour (sub $second.Minute 1) $second.Second }}

			{{ $diff := $fiComp.Sub $seComp }}

			{{ if lt $diff 0 }}
				{{ $diff = $seComp.Sub $fiComp }}
			{{ end }}

			{{ if eq $diff 0 }}
				{{ $diff = "no" }}
			{{ else }}
				{{ $diff = humanizeDurationMinutes ($diff) }}
			{{ end }}

			{{ $desc := print 
				"Current time in **" $fiOut "** is **" ($first.Format "Monday 3:04 PM") "**\n" 
				"Current time in **" $seOut "** is **" ($second.Format "Monday 3:04 PM") "**\n\n" 
				"There is **" $diff "** difference between **" $fiOut "** and **" $seOut "**."
			}}
			{{ $embed.Set "description" $desc }}
		{{ else if and (reFind `[\d]` $first) (reFind `[\D]` $second) }}
			{{ if and (not ($clocks.HasKey $second)) (not ($cList.HasKey $second)) }}
				{{ sendMessage nil "Country needs to be from the supported countries list." }}
				{{ return }}
			{{ end }}

			{{ $first = toInt $first }}
			{{ if or (lt $first 0 ) (gt $first 24) }}
				{{ $first = 0 }}
			{{ end }}

			{{ $cOut := execTemplate "reCountry" $second }}
			{{ $desc := print "Listing time results to match **" $cOut "** at **" $first ":00** (24h) for different timezones." }}
			{{ $second = or ($clocks.Get $second) ($cList.Get $second) }}

			{{ $embed.Set "fields" cslice }}
			{{ $embed.Set "description" $desc }}

			{{ range $country, $timezone := $clocks }}
				{{- $time := currentTime.In (loadLocation $timezone) -}}
				{{- $stamp := newDate $time.Year $time.Month $time.Day $first 0 0 $second -}}
				{{- $time = $stamp.In (loadLocation $timezone) -}}
				{{- $format := print ($time.Format "Monday 3:04 PM") -}}
				
				{{- $country = execTemplate "reCountry" $country -}}
				
				{{- $embed.fields.Append (sdict "name" $country "value" $format "inline" true) | $embed.Set "fields" -}}
			{{- end }}
		{{ end }}
	{{ else if eq (len $args) 1 }}
		{{ $input := index .CmdArgs 0 | lower }}

		{{ if and (not ($clocks.HasKey $input)) (not ($cList.HasKey $input)) }}
			{{ sendMessage nil "Country needs to be from the supported countries list." }}
			{{ return }}
		{{ end }}

		{{ $zone := or ($clocks.Get $input) ($cList.Get $input) }}
		{{ $curr := currentTime.In (loadLocation $zone) }}
		{{ $curr = print ($curr.Format "Monday 3:04 PM") }}

		{{ $country := execTemplate "reCountry" $input }}
		{{ $desc := print "Current time in **" $country "** is **" $curr "**" }}
		{{ $embed.Set "description" $desc }}
	{{ end }}
{{ else }}
	{{ $embed.Set "fields" cslice }}

	{{ range $country, $timezone := $clocks }}
		{{- $time := currentTime.In (loadLocation $timezone) -}}
		{{- $format := print ($time.Format "Monday 3:04 PM") -}}
		{{- $country = execTemplate "reCountry" $country -}}
		{{- $embed.fields.Append (sdict
			"name" $country
			"value" $format
			"inline" true
		) | $embed.Set "fields" -}}
	{{- end }}
{{ end }}

{{ sendMessage nil (cembed $embed) }}

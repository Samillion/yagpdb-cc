{{ $alphanum := sdict
    "0" "𝟢" "1" "𝟣" "2" "𝟤" "3" "𝟥" "4" "𝟦" "5" "𝟧" "6" "𝟨" "7" "𝟩" "8" "𝟪" "9" "𝟫"

    "a" "𝒶" "b" "𝒷" "c" "𝒸" "d" "𝒹" "e" "𝑒" "f" "𝒻" "g" "𝑔" "h" "𝒽" "i" "𝒾" "j" "𝒿"
    "k" "𝓀" "l" "𝓁" "m" "𝓂" "n" "𝓃" "o" "𝑜" "p" "𝓅" "q" "𝓆" "r" "𝓇" "s" "𝓈" "t" "𝓉"
    "u" "𝓊" "v" "𝓋" "w" "𝓌" "x" "𝓍" "y" "𝓎" "z" "𝓏"
    
    "A" "𝒜" "B" "𝐵" "C" "𝒞" "D" "𝒟" "E" "𝐸" "F" "𝐹" "G" "𝒢" "H" "𝐻" "I" "𝐼" "J" "𝒥"
    "K" "𝒦" "L" "𝐿" "M" "𝑀" "N" "𝒩" "O" "𝒪" "P" "𝒫" "Q" "𝒬" "R" "𝑅" "S" "𝒮" "T" "𝒯"
    "U" "𝒰" "V" "𝒱" "W" "𝒲" "X" "𝒳" "Y" "𝒴" "Z" "𝒵"
}}

{{ $in := trimSpace .StrippedMsg }}
{{ $usage := print "**Usage:**" "\n" "```" .Cmd " any text you'd like```" }}

{{ $message := cembed 
	"title" "Fancy Text"
	"description" (print "Convert any text into **𝐹𝒶𝓃𝒸𝓎 𝓉𝑒𝓍𝓉** easily." "\n\n" $usage)
}}

{{ if $in }}
	{{ if gt (len $in) 500 }}
		{{ $explain := print 
			"A maximum of 500 characters is allowed. Due to Discrd's handling of Unicode, there is no escaping this limit, unfortunately." "\n\n"
			"**Explanation:**" "\n"
			"- One fancy char = 4 regular chars. So `𝒽` is the same as `hhhh`." "\n"
			"- Discord's message limit is 2000 characters total. `2000 / 4 = 500`."
		}}
		
		{{ $message = cembed 
			"title" "Fancy Text"
			"description" (print $explain "\n\n" $usage)
		}}
	{{ else }}
		{{ range $old, $new := $alphanum }}
			{{- $in = joinStr $new (split $in $old) -}}
		{{- end }}

		{{ $message = $in }}
	{{ end }}
{{ end }}

{{ sendMessage nil $message }}

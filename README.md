# Slap command
As basic as it gets, and as seen on many bots. `-slap @user` will generate an embed with a humorous slap GIF from a list.

It will generate a different response for:
- If user mentions self. (ie `Sam: -slap @Sam`)
- if user mentions YAGPDB bot. (ie `Sam: -slap @YAGPDB.xyz`)
- If user mentions a non existent member on the server.

Embed will match the user's role color, thanks to [buthed010203](https://yagpdb-cc.github.io/code-snippets/get-username-color)'s code snippet.

## Setup
- Login to YAGPDB dashboard. (https://yagpdb.xyz/manage)
- Navigate: Core -> Custom Commands -> Create a new custom command.
- Trigger: `slap`
- Copy and paste [slap](https://github.com/Samillion/yagpdb-slap/blob/main/slap.go) code in the response field.
- Save.

## Modify
Cooldown: (`$cooldown`)
- By default it is set to 30 seconds, per user.
- If set to 0, there is no cooldown limit

GIFs: (`$images`)
- A predefined list (slice) of GIFs with a humorous theme.
- Direct links should be used there.
- Feel free to add or remove GIFs, make sure each URL is wraped with quotes. `"URL"`

The predefined GIFs:
```
https://media.tenor.com/tKF3HiguDmEAAAAC/wrrruutchxxxxiii-slapt.gif
https://media.tenor.com/__oycZBexeAAAAAC/slap.gif
https://media.tenor.com/W0K0vteByOoAAAAC/slap-in-the-face-slap.gif
https://media.tenor.com/GBShVmDnx9kAAAAC/anime-slap.gif
https://media.tenor.com/bHqFiKBJGmoAAAAC/korone-dog-form.gif
https://media.tenor.com/rvXlkAdJXIgAAAAC/paper-fan-slap.gif
https://media.tenor.com/ra17G61QRQQAAAAC/tapa-slap.gif
https://media.tenor.com/eU5H6GbVjrcAAAAC/slap-jjk.gif
```

`$selfSlap` is the GIF shown when a user slaps self. Predefined self slap GIF:
```
https://media.tenor.com/rwipJN-E7okAAAAC/slap-slapping-self.gif
```

`$botSlap` is the GIF shown when a user slaps YAGPDB bot. Predefined bot slap GIF:
```
https://media.tenor.com/J9flx980Ot4AAAAC/reggie-deflecting-criticism.gif
```

Preview:

![image](https://github.com/Samillion/yagpdb-slap/assets/17427046/9a9b35d6-cfab-413b-afea-921c9e97d664)


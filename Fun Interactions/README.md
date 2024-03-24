# Fun Interactions
As basic as it gets, and as seen on many bots. `-hug @user`, `-slap @user`, `-spank @user`, `-fart @user` will generate an embed with a humorous GIF from a list.

It will generate a different response for:
- If user mentions self. (ie: `Sam: -slap @Sam`)
- if user mentions another user. (ie: `-spank @YAGPDB.xyz`)
- It is possible to add a message as long as there is a valid user mention. (ie: `-fart on @Sam`)
- If user mentions a non existent member on the server. (Outputs invalid user error)
- If no mention or message is added. (ie: `-hug`) (Will output how to use command)

Embed will match the user's role color, thanks to [buthed010203](https://yagpdb-cc.github.io/code-snippets/get-username-color)'s code snippet.

## Modify
Cooldown: (`$cooldown`)
- By default it is set to 0 seconds, meaning no cooldown limit
- If set to 60, cooldown will be 60 seconds

GIFs:
- Predefined lists (slice) of GIFs with a humorous theme.
- Different lists for each interaction in `$type` (sdict) -> (hug|fart|slap|spank) -> `"gifs" (cslice)`.
- Direct links should be used in `"gifs" (cslice)`.
- Feel free to add or remove GIFs, make sure each URL is wraped with quotes. `"URL"`

## Setup
- Login to YAGPDB dashboard. (https://yagpdb.xyz/manage)
- Navigate: Custom Commands -> Commands -> Create a new custom command.
- Trigger type: Regex
- Trigger: `^-\b(hug|spank|slap|fart)\b(\s+|\z)`
- Copy and paste [code](https://raw.githubusercontent.com/Samillion/yagpdb-cc/main/Fun%20Interactions/funinteractions.go) in the response field.
- Save.

## Preview

![image](https://github.com/Samillion/yagpdb-slap/assets/17427046/9a9b35d6-cfab-413b-afea-921c9e97d664)


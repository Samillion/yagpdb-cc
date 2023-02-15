# Slap command
As basic as it gets, and as seen on many bots. `-slap @user` will generate an embed with a humorous slap GIF from a list.

Cooldown by default is set to 3 minuutes (180 secs), per user. To set it with no limit, change to 0.

It will generate a different response for:
- If user tags themselves.
- if user tags YAGPDB bot.
- If user tags a non existent member on the server.

Embed will match the user's role color, thanks to [buthed010203](https://yagpdb-cc.github.io/code-snippets/get-username-color)'s code snippet.

## Setup:
- Login to YAGPDB dashboard. (https://yagpdb.xyz/manage)
- Navigate: Core -> Custom Commands -> Create a new custom command.
- Trigger: `slap`
- Copy and paste [slap.go](https://github.com/Samillion/yagpdb-cc/blob/main/slap.go) code in the response field.
- Save.

Preview:

![image](https://user-images.githubusercontent.com/17427046/218953630-ff236c9c-fcfd-4c86-be50-e9038886389b.png)

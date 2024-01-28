# Anonymous Confessions
A command to post messages anonymously in a dedicated confessions channel.

There is a global cooldown limit per confession. The reasons it is global and not per user, is to not identify users with confession timestamps and to prevent spam.

Details:
- `$ch` should be the ID for the confessions channel.
- `$cd` is the amount of seconds for cooldown per confession. Default is 10 minutes. (600 seconds)

## Explanation and Preview

![image](https://github.com/Samillion/yagpdb-cc/assets/17427046/e09a35f3-d269-42e4-a970-8f417fb0a256)

![image](https://github.com/Samillion/yagpdb-cc/assets/17427046/4a2f9356-da2a-4a8b-8565-197186695577)


## Setup
- Login to YAGPDB dashboard. (https://yagpdb.xyz/manage)
- Navigate: Custom Commands -> Commands -> Create a new custom command.
- Trigger type: Command
- Trigger: `confess`
- Copy and paste [code](https://raw.githubusercontent.com/Samillion/yagpdb-cc/main/Confession/confess.go) in the response field.
- Save.

## Disclaimer
This command (code), as is, guarantees anonymity. Currently even Discord's Audit logs will not show the deleted log of the `-confess` message by the user.

However, it is possible for server admins to be monitoring actions. Such as message delete, edit and send. If so, then anonymity is compromised in this scenario.

Server admins using this command should respect the privacy of their members. This can be a funny command to mess around with or a genuine outlet to unload a memory or two. Think of the bigger picture.

I have no control over Discord, only what happens on my own Discord server. If you're a member in a server that has this exact command or a similar feature, then you should ask the admins directly about their privacy policy.

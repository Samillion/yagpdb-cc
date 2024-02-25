# Anonymous Dreams
A command to post messages anonymously in a dedicated dreams channel. The intended use is to post dreams you have when sleeping, though it can be used for dreams as in "hopes and wishes" as well.

There is a global cooldown limit per dream. The reason it is global and not per user, is to not identify users with dream timestamps and to prevent spam.

**Details:**
- `$ch` should be the ID for the dreams channel.
- `$cd` is the amount of seconds for cooldown per dream. Default is 600 seconds. (10 minutes)

On my Discord server, I have the `#dreams` channel in read-only mode. Meaning that "Send Message" permission is disabled for everyone, to keep the channel free from clutter.

## Explanation and Preview
![image](https://github.com/Samillion/yagpdb-cc/assets/17427046/f3eaadd5-dd9e-4bd9-a331-61996ed57552)

![image](https://github.com/Samillion/yagpdb-cc/assets/17427046/bf340cd1-f9ab-46b1-b45f-ff989a98c3e4)

## Setup
- Login to YAGPDB dashboard. (https://yagpdb.xyz/manage)
- Navigate: Custom Commands -> Commands -> Create a new custom command.
- Trigger type: Command
- Trigger: `dream`
- Copy and paste [code](https://raw.githubusercontent.com/Samillion/yagpdb-cc/main/Dreams/dreams.go) in the response field.
- Save.

## Disclaimer
This command (code), as is, guarantees anonymity. Currently even Discord's Audit logs will not show the deleted log of the `-dream` message by the user.

However, it is possible for server admins to be monitoring actions. Such as message delete, edit and send. If so, then anonymity is compromised in this scenario.

Server admins using this command should respect the privacy of their members. This can be a funny command to mess around with or a genuine outlet to unload a memory or two. Think of the bigger picture.

I have no control over Discord, only what happens on my own Discord server. If you're a member in a server that has this exact command or a similar feature, then you should ask the admins directly about their privacy policy then decide based on the answer whether you'd like to use the confession feature on their server.

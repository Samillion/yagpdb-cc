# Confession
A command to post messages anonymously in a dedicated confessions channel.

There is a global cooldown limit per confession. The reasons it is global and not per user, is to not identify users' timestamps with the cooldown limit and to prevent spam.

Details:
- `$ch` should be the ID for the confessions channel.
- `$cd` is the amount of seconds for cooldown per confession. Default is 10 minutes.

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


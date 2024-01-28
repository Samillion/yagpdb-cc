# Confession
A command to post messages anonymously in a dedicated confessions channel.

There is a global cooldown limit per confession set for 10 minutes. The reasons it is global and not per user, is to not identify users' timestamps with the cooldown limit and to prevent spam.

## Explanation and Preview

![image](https://github.com/Samillion/yagpdb-cc/assets/17427046/bcafad69-803a-4192-8aae-5c69239af2c7)

![image](https://github.com/Samillion/yagpdb-cc/assets/17427046/4a2f9356-da2a-4a8b-8565-197186695577)


## Setup
- Login to YAGPDB dashboard. (https://yagpdb.xyz/manage)
- Navigate: Custom Commands -> Commands -> Create a new custom command.
- Trigger type: Command
- Trigger: `confess`
- Copy and paste [code](https://raw.githubusercontent.com/Samillion/yagpdb-cc/main/Confession/confess.go) in the response field.
- Save.


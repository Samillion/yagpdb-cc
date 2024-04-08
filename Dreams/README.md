# Dreams
A command to post messages in a dedicated dreams channel. The intended use is to post dreams you have when sleeping, though it can be used for dreams as in "hopes" as well.

**Details:**
- `$ch` should be the ID for the dreams channel.
- `$cd` is the amount of seconds for cooldown per dream, default is 600 seconds. (10 minutes)
- `$minimum` is the minimum characters count required for a dream to be posted, default is 20.

On my Discord server, I have the `#dreams` channel in read-only mode. Meaning that `Send Message` permission is disabled for everyone, to keep the channel free from clutter.

<table>
 <tr>
    <td colspan="2">Explanation and Preview</td>
 </tr>
 <tr>
    <td><img src="https://github.com/Samillion/yagpdb-cc/assets/17427046/1d32d32d-c6d8-47c7-865a-937718ffa59e" alt=""></td>
    <td><img src="https://github.com/Samillion/yagpdb-cc/assets/17427046/42a4e0c3-e629-4400-b665-8de5659d4dcb" alt=""></td>
 </tr>
</table>

## Setup
- Login to YAGPDB dashboard. (https://yagpdb.xyz/manage)
- Navigate: Custom Commands -> Commands -> Create a new custom command.
- Trigger type: Command
- Trigger: `dream`
- Copy and paste [code](https://raw.githubusercontent.com/Samillion/yagpdb-cc/main/Dreams/dreams.go) in the response field.
- Save.

## Disclaimer
The basic idea is to post dreams anonymously, the code as is does not keep logs of any kind nor track user use, even Discord Audit log will not show the posted and deleted actions of this command.

However:
- Pop up notifications could reveal who used the command, if someone is online at the same time. (Thanks to Arksi pointing this out)
- Server admins could have third party loggers for message post, edit and delete.

Server admins using this command should respect the privacy of their members. This can be a funny command to mess around with or a genuine outlet to unload a memory or two. Think of the bigger picture.

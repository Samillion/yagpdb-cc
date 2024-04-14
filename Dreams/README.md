# Dreams & Confessions
A command to post messages in a dedicated dreams and confessions channel, the purpose is to confess something or to share an interesting dream you had.

**Usage:**
- `-dream details of your dream`
- `-confess your confession`

**Details:**
- `$ch` should be the ID for the dreams channel.
- `$cd` is the amount of seconds for cooldown per dream, default is 600 seconds. (10 minutes)
- `$minimum` is the minimum characters count required for a dream or a confession to be posted, default is 20.

On my Discord server, I have the `#dreams-confessions` channel in read-only mode. Meaning that `Send Message` permission is disabled for everyone, to keep the channel free from clutter.

<table>
 <tr>
    <td colspan="2">Explanation and Preview</td>
 </tr>
 <tr>
    <td><img src="https://github.com/Samillion/yagpdb-cc/assets/17427046/ecd4a040-5b6d-4409-8ac3-1a3b312424b7" alt=""></td>
    <td><img src="https://github.com/Samillion/yagpdb-cc/assets/17427046/481b49fc-3cb2-49e6-9101-9b21ae2a4cca" alt=""></td>
 </tr>
</table>

## Setup
- Login to YAGPDB dashboard. (https://yagpdb.xyz/manage)
- Navigate: Custom Commands -> Commands -> Create a new custom command.
- Trigger type: Regex
- Trigger: `^-\b(dream|confess)\b(\s+|\z)`
- Copy and paste [code](https://raw.githubusercontent.com/Samillion/yagpdb-cc/main/Dreams/dreams.go) in the response field.
- Save.

## Disclaimer
The basic idea is to post dreams and confessions anonymously, the code as is does not keep logs of any kind nor track use, even Discord Audit log will not show the post and delete actions of this command, because Discord ignores single actions by bots in the logs.

However:
- Pop up notifications could reveal who used the command, if someone is online at the same time. (Thanks to Arksi pointing this out)
- Server admins could have third party loggers for message post, edit and delete.

Server admins using this command should respect the privacy of their members. This can be a funny command to mess around with or a genuine outlet to unload a memory or two. Think of the bigger picture.

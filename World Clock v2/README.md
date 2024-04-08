# World Clock v2
This is an "advanced" version of the [World Clock](https://yagpdb-cc.github.io/utilities/world-clock) command made by [jo3-l](https://github.com/jo3-l). It is honestly a complete overkill, but it services my lazy side on that rare time I want this information and I don't have to open my browser for it.

Features:
- By default it will output time for regions that cover most of the globe
- You can display current time for a specific country
- You can display time in different timezones based on a specific time of a country
- You can compare time difference between two countries
- 12 default global timezones and 106 country specific timezones

## Setup
- Login to YAGPDB dashboard. (https://yagpdb.xyz/manage)
- Navigate: Custom Commands -> Commands -> Create a new custom command.
- Trigger type: Command
- Trigger: `time`
- Copy and paste [code](https://raw.githubusercontent.com/Samillion/yagpdb-cc/main/World%20Clock%20v2/worldclockv2.go) in the response field.
- Save.
- Run `-time help` on your server to discover the features.

## Modify
Global Timezones: (`$clocks`)
- An sdict type list containing timezones to cover most of the globe
- You can add, adjust or remove items in the list
- Make sure the name is in lowercase, preferably a single name without space

Country Timezones: (`$cList`)
- An sdict type list containing timezones to cover most of the globe
- You can add, adjust or remove items in the list
- Make sure the name is in lowercase, preferably a single name without space

<table>
 <tr>
    <td colspan="2">Preview</td>
 </tr>
 <tr>
    <td><img src="https://github.com/Samillion/yagpdb-cc/assets/17427046/56d9e38f-a492-47c8-9407-c9aa70abb5cf" alt=""></td>
    <td><img src="https://github.com/Samillion/yagpdb-cc/assets/17427046/55474c0b-30c4-4b73-bb1d-34611f2d2b0d" alt=""></td>
 </tr>
 <tr>
    <td><img src="https://github.com/Samillion/yagpdb-cc/assets/17427046/6a812335-b3cf-4aad-a283-c00c2d499a43" alt=""></td>
    <td><img src="https://github.com/Samillion/yagpdb-cc/assets/17427046/37dfae46-e4e6-4fb4-b90d-6e5f191dbb61" alt=""></td>
 </tr>
 <tr>
    <td><img src="https://github.com/Samillion/yagpdb-cc/assets/17427046/c87b2234-99e4-4170-8e28-116476d82c69" alt=""></td>
    <td><img src="https://github.com/Samillion/yagpdb-cc/assets/17427046/b752f8a8-9fa2-49c1-84b4-cc0a31f14f9b" alt=""></td>
 </tr>
</table>

# Random String
A simple and basic random string generator with options to control the type and amount of characters in output.

#### Output
- By default it will output 15 characters
- You can increase that by changing `{{ $default := 15 }}`
- You can output a maximum of 60 characters
- You can increase the maximum 60 limit by changing `{{ $max := 60 }}`

#### Options
- `alpha`: Alphabets, lower and upper case `[a-zA-Z]`
- `lower`: Alphabets, lower case `[a-z]`
- `upper`: Alphabets, upper case `[A-Z]`
- `num`: Numerical `[0-9]`
- `sym`: Symbols, by default `~!@#$%&*`
- `mix`: All the previous types combined (Default)

#### Arguments
The order of arguments is fluid, meaning all the following scenarios will work
```
-random
-random 30
-random num
-random 40 upper
-random lower 40
```

## Setup
- Login to YAGPDB dashboard. (https://yagpdb.xyz/manage)
- Navigate: Custom Commands -> Commands -> Create a new custom command.
- Trigger type: Command
- Trigger: `random`
- Copy and paste [code](https://raw.githubusercontent.com/Samillion/yagpdb-cc/main/Random%20String/randomstring.go) in the response field.
- Save.

<table>
 <tr>
    <td colspan="2">Preview</td>
 </tr>
 <tr>
    <td><img src="https://github.com/Samillion/yagpdb-cc/assets/17427046/27dd265d-7d44-4033-9c93-4638d2a9094b" alt=""></td>
    <td><img src="https://github.com/Samillion/yagpdb-cc/assets/17427046/1c46add3-8283-4ec7-84d6-38916ef6204f" alt=""></td>
 </tr>
</table>

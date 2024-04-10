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
The order of arguments is fluid and multiple character types can be used. Meaning, all the following scenarios will work:
```
-random
-random 30
-random num
-random 40 upper
-random lower 40
-random num lower
-random sym num 30
-random 25 upper sym
```
The arguments are processed, for example:
- If `alpha` and `upper` are used, only `alpha` is processed since it has `upper` already.
- If `mix` is used with other types (ie: `mix num`), only `mix` will be processed since it has all types already.

This is to ensure the generated output does not favor a specific type based on repeated strings.

## Setup
- Login to YAGPDB dashboard. (https://yagpdb.xyz/manage)
- Navigate: Custom Commands -> Commands -> Create a new custom command.
- Trigger type: Command
- Trigger: `random`
- Copy and paste [code](https://raw.githubusercontent.com/Samillion/yagpdb-cc/main/Random%20String/randomstring.go) in the response field.
- Save.

## Preview

![image](https://github.com/Samillion/yagpdb-cc/assets/17427046/05e0f1c5-795c-459f-b90e-af1169439d13)

# Search Link
A search link generator for Google, Duckduckgo, Brave and Wikipedia.

It simply generates a link based on the search engine you specify and the keywords you use. I made this to serve my lazy side to generate the link without opening a browser.

## Setup
- Login to YAGPDB dashboard. (https://yagpdb.xyz/manage)
- Navigate: Custom Commands -> Commands -> Create a new custom command.
- Trigger type: regex
- Trigger: `^-\b(google|duck|brave|wiki)\b`
- Copy and paste [code](https://raw.githubusercontent.com/Samillion/yagpdb-cc/main/Search%20Link/searchlink.go) in the response field.
- Save.

## Preview

![image](https://github.com/Samillion/yagpdb-cc/assets/17427046/67ef0773-e8fa-4a91-8858-686483079852)

# arkfbp-cli

Standard Tooling for ArkFBP Development

## Installation

```
go get -u github.com/arkfbp/arkfbp-cli
```

use ```arkfbp-cli``` to confirm your installation, for instance:

```
arkfbp-cli version
```

## Quickstart

Create a new ArkFBP project, run on the server side, written in Node.JS

```arkfbp-cli create --type server --language javascript --name helloworld```


Create a new ArkFBP project, run on the server side, written in Golang

```arkfbp-cli create --type server --language go --name helloworld --package github.com/rockl2e/helloworld```


Create a new ArkFBP web project written in Typescript

```arkfbp-cli create --type web --language typescript --name helloworld```

## ArkFBP Project Category References

### ArkFBP Project Type & Language Support(current status)

|   Type |            Description              |
| ------ | ----------------------------------- |
| server | supposed to run on the server side  |
|  web   | run on the browser side             |
| script | desktop & server side as the script |


#### Server

| Type   | Language   | Internal WEB Framework(optional) |
| ----   | --------   | -------------------------------- |
| server | javascript |              express             |
| server | typescript |              express             |
| server | go         |                                  |
| server | python     |         django / flask           |

#### WEB

| Type | Language   | Framework |
| ---- | --------   | -------   |
| web  | javascript | vue       |
| web  | typescript | vue       |

#### Script

| Type    | Language   |
| ----    | --------   |
| script  | javascript |
| script  | typescript |
| script  | python     |
| script  | go         |


## arkfbp-cli Command References

1. ```create```: create a new ArkFBP project
2. ```init```: reinit the meta information of the ArkFBP project, basically the .arkfbp folder
3. ```createflow```: create flow in the arkfbp project
4. ```createnode```: create node file of one flow
5. ```help```: show the help message
6. ```inspect```: inspect the info of the arkfbp project
7. ```version```: show the vresion of the arkfbp-cli
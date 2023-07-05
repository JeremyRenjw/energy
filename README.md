<p align="center">
    <img src="https://assets.yanghy.cn/energy-doc/energy-icon.png">
</p>

# Energy is the framework for Go to build desktop applications based on CEF

[中文](README.zh_CN.md) |
English

---
[![github](https://img.shields.io/github/last-commit/energye/energy/main.svg?logo=github&logoColor=green&label=commit)](https://github.com/energye/energy)
[![release](https://img.shields.io/github/v/release/energye/energy?logo=git&logoColor=green)](https://github.com/energye/energy/releases)
[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=red)](http://www.apache.org/licenses/LICENSE-2.0)
![repo](https://img.shields.io/github/repo-size/energye/energy.svg?logo=github&logoColor=green&label=repo-size)

![go-version](https://img.shields.io/github/go-mod/go-version/energye/energy?logo=git&logoColor=green)
---

### [Project Introduction](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/6350f94ca749ba0318943f25)
Energy is a framework developed by Golang based on CEF(Chromium Embedded Framework), embedded with [CEF](https://bitbucket.org/chromiumembedded/cef) binary
> [energy](https://github.com/energye/energy) is a framework developed by Golang based on CEF(Chromium Embedded Framework), embedded with [CEF](https://bitbucket.org/chromiumembedded/cef) binary
>
> Use Go and Web technology (HTML+CSS+JavaScript) to build cross-platform desktop applications that support Windows, Linux and MacOS
>
> Knowledge of the front-end technology stack and some knowledge of the Go language is required

### Characteristic

> - development environment is simple and the compilation speed is fast. Only the Go development environment and the CEF binary framework that Energy depends on are needed
> - cross-platform: A set of code can be packaged into Windows, domestic UOS, Deepin, Kylin, MacOS, Linux
> - Language responsibilities
>> - Go: Go is responsible for window creation, CEF configuration and function implementation, creation of various UI components, low-level system calls, and functions that JS cannot handle, such as file stream, security encryption, high-performance processing, etc., which can be developed as a pure backend
>> - Web: HTML + CSS + JavaScript responsible for the function of the client interface, make any interface you want, can be used as a pure front-end development
> - front-end technology: Support mainstream front-end frameworks, such as Vue, React, Angular or pure HTML+CSS
> - event driven: High performance event driven, IPC based communication, Go and Web side is very convenient function call and data interaction


### Built-in dependency&integration

- [![LCL](https://img.shields.io/badge/LCL-green)](https://github.com/energye/golcl)
- [![CEF-CEF4Delphi](https://img.shields.io/badge/CEF(Chromium%20Embedded%20Framework)%20CEF4Delphi-green)](https://github.com/salvadordf/CEF4Delphi)

### [Development environment](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/63511b14a749ba0318943f3a)

> Install automatically using the energy command line tool
#### Basic needs
> - golang >= 1.18
>
> - energy development environment
> 
> Use the energy command line tool to automatically install the development environment
>
> Get [energy](https://github.com/energye/energy) project, or use the precompiled command line tool directly [Download address](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/63511b14a749ba0318943f3a)
> 
> <p style="color:palevioletred;">If using pre compiled command-line tools, the following steps can be skipped</p>
> 
> `go get github.com/energye/energy`
>
> Enter the  [energy](https://github.com/energye/energy) command line directory
> 
> `cd energy/cmd/energy`
>
> Install command line tools
> 
> `go install`
>
> Execute the installation command
> 
> `energy install .`

### Getting Started Guide - [Transfer gate](https://energy.yanghy.cn)

* [Course](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/6350f94ca749ba0318943f25)
* [Example](https://energy.yanghy.cn/#/example/6342d986401bfe4d0cdf6067/634d3bd5a749ba0318943eb6)
* [Document](https://energy.yanghy.cn/#/document/6342d9a4401bfe4d0cdf6069/0)

### Quick Get Start
> Use [energy](https://github.com/energye/energy) Command line tool automatic installation environment dependency `energy install .`
>
> Take example/simple as an example
>
> Update latest release dependency
>
> `go mod tidy`
>
> Run `simple` in the IDE or `go run simple.go`

### example/simple Code

```go
package main

import (
	"github.com/energye/energy/v2/cef"
)

func main() {
	//Global initialization must be called by every application
	cef.GlobalInit(nil, nil)
	//Create application
	cefApp := cef.NewApplication()
	//Set URL
	cef.BrowserWindow.Config.Url = "https://www.yanghy.cn"
	//Run App
	cef.Run(cefApp)
}
```
---

### Run app
- Windows、Linux
> `go run simple.go`
- MacOS
> `go run simple.go energy_env=dev`


### Pack Project [reference](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/636e397ba749ba01d04ff595)
1. Compile: Go program compilation `go build simple.go` If you use resource built-in (HTML, CSS, JavaScript, Image, etc.), the resource will be compiled into the execution file
2. Copy: copy the execution file to the CEF directory of the ENERGY environment
3. Packaging: use the installation package tool to make it as an installation package, Or refer to the production of installation package for each system platform
4. Finally: the compiled program or installation package and CEF directory no longer need to configure the environment, and can be run directly in the CEF root directory

#### Go Compile Command
- `go build -ldflags "-H windowsgui -s -w"`
> `-ldflags`
>
>> `-H windowsgui` optional: windows hide cmd black window
>>
>> `-w` optional: Removing debugging information can reduce the size of the execution file
>>
>> `-s` optional: Removing Symbol table information can reduce the size of the execution file

---

---

### System support

![Windows](https://img.shields.io/badge/windows-supported-success.svg?logo=Windows&logoColor=blue)
![MacOS](https://img.shields.io/badge/MacOS-supported-success.svg?logo=MacOS)
![Linux](https://img.shields.io/badge/Linux-supported-success.svg?logo=Linux&logoColor=red)


|             | 32 Bit                                                                                     | 64 Bit                                                                                     | Test System Version                |
|-------------|--------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------|------------------------------------|
| Windows     | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue) | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue) | Windows 7、Windows 10               |
| MacOS       | ![MacOS](https://img.shields.io/badge/N/A-inactive.svg?logo=MacOS)                         | ![MacOS](https://img.shields.io/badge/supported-success.svg?logo=MacOS)                    | MacOSX 10.15                       |
| Linux       | ![Linux](https://img.shields.io/badge/SelfCompila-supported-success.svg?logo=Linux)        | ![Linux](https://img.shields.io/badge/supported-success.svg?logo=Linux&logoColor=red)      | Deepin20.8、Ubuntu18.04、LinuxMint21 |
| Linux ARM   | ![Linux ARM](https://img.shields.io/badge/SelfCompila-supported-success.svg?logo=Linux)    | ![Linux ARM](https://img.shields.io/badge/SelfCompila-supported-success.svg?logo=Linux)    | Kylin-V10-SP1-2107                 |

---

### Thanks Jetbrains

<p align="center">
    <a href="https://www.jetbrains.com?from=energy">
        <img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg" alt="JetBrains Logo (Main) logo.">
    </a>
</p>

Thank please give this project a star

---

### QQ Group

[![QQGroup](https://img.shields.io/badge/QQ-541258627-green.svg?logo=tencentqq&logoColor=blue)](https://jq.qq.com/?_wv=1027&k=YgFjCGJX)

<p align="center">
    <img src="https://assets.yanghy.cn/energy-doc/qq-group.jpg" width="250">
</p>
---

### Project screenshot
##### Windows-10

<img src="https://assets.yanghy.cn/energy-doc/frameless-windows-10.png">

##### Windows-7 32 & 64
<img src="https://assets.yanghy.cn/energy-doc/frameless-windows-7-64.png">
<img src="https://assets.yanghy.cn/energy-doc/frameless-windows-7-32.png">

##### Linux - Deepin
<img src="https://assets.yanghy.cn/energy-doc/frameless-deepin-20.8.png">
<img src="https://assets.yanghy.cn/energy-doc/frameless-deepin-hide-20.8.png">

##### Linux - Kylin ARM
<img src="https://assets.yanghy.cn/energy-doc/frameless-kylin-arm-V10-SP1.png">
<img src="https://assets.yanghy.cn/energy-doc/frameless-kylin-arm-hide-V10-SP1.png">

##### Linux - Ubuntu
<img src="https://assets.yanghy.cn/energy-doc/frameless-ubuntu-18.04.6.png">
<img src="https://assets.yanghy.cn/energy-doc/frameless-ubuntu-hide-18.04.6.png">

##### MacOSX
<img src="https://assets.yanghy.cn/energy-doc/frameless-macos.png">


----

### Public License

[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=green)](http://www.apache.org/licenses/LICENSE-2.0)




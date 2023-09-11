//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package command

type OS string
type Arch string

type Config struct {
	Index     int
	Wd        string
	EnergyCfg EnergyConfig
	Install   Install `command:"install" description:"install energy development dependency environment"`
	Package   Package `command:"package" description:"energy application production and installation package"`
	Version   Version `command:"version" description:"list all release version numbers of energy"`
	Setenv    Setenv  `command:"setenv" description:"set ENERGY_ HOME framework environment"`
	Env       Env     `command:"env" description:"display ENERGY_ HOME framework environment directory"`
	Init      Init    `command:"init" description:"initialize the energy application project"`
	Build     Build   `command:"build" description:"building an energy project"`
}

type Command struct {
	Run                    func(c *Config) error
	UsageLine, Short, Long string
}

type Install struct {
	Path     string `short:"p" long:"path" description:"Installation directory Default current directory"`
	Version  string `short:"v" long:"version" description:"Specifying a version number"`
	Name     string `short:"n" long:"name" description:"Name of the framework directory after installation" default:"EnergyFramework"`
	Download string `short:"d" long:"download" description:"Download Source, 0:gitee or 1:github, Default empty" default:""`
	All      bool   `short:"a" long:"all" description:"Install all, skip installation prompts (Y/n)"`
	OS       OS     `long:"os" description:"Specify install OS: [windows, linux, darwin], default current system: os"`
	Arch     Arch   `long:"arch" description:"Specify install ARCH: [386, amd64, arm64], Default current system: architecture"`
	CEF      string `long:"cef" description:"Install system supports CEF version, provide 4 options, default empty. default, windows7, gtk2, flash" default:""`
	IGolang  bool   // 是否已安装Golang
	ICEF     bool   // 是否已安装CEF
	INSIS    bool   // 是否已安装nsis
	IUPX     bool   // 是否已安装upx
	IsSame   bool   // 安装的OS和Arch是否为当前系统架构, 默认当前系统架构
}

type Package struct {
	Path  string `short:"p" long:"path" description:"Project path, default current path. Can be configured in energy.json" default:""`
	Clean bool   `short:"c" long:"clean" description:"Clear configuration and regenerate the default configuration"`
}

type Env struct {
}

type Setenv struct {
	Path string `short:"p" long:"path" description:"Energy framework dir"`
}

type Version struct {
	All bool `short:"a" long:"all" description:"show all"`
}

type Init struct {
	Name    string `short:"n" long:"name" description:"Initialized project name"`
	ResLoad string `short:"r" long:"resload" description:"Resource loading method, 1: HTTP, 2: Local Load, default 1 HTTP"`
	IGo     bool
	INSIS   bool
	IEnv    bool
	INpm    bool
}

type Build struct {
	Path    string `short:"p" long:"path" description:"Project path, default current path. Can be configured in energy.json" default:""`
	Upx     bool   `short:"u" long:"upx" description:"Set this parameter and install upx. Use upx to compress the execution file. windows"`
	UpxFlag string `long:"upxFlag" description:"Upx command line parameters" default:""`
}

type EnergyConfig struct {
	Source DownloadSource `json:"source"`
}

type DownloadSource struct {
	Golang string `json:"golang"`
	CEF    string `json:"cef"`
}

func (m OS) IsWindows() bool {
	return m == "windows"
}

func (m OS) IsLinux() bool {
	return m == "linux"
}

func (m OS) IsDarwin() bool {
	return m == "darwin"
}

func (m Arch) Is386() bool {
	return m == "386" || m == "32" // windows32
}

func (m Arch) IsAMD64() bool {
	return m == "amd64"
}

func (m Arch) IsARM64() bool {
	return m == "arm64"
}
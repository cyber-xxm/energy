//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cli

import (
	"encoding/json"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/remotecfg"
	"github.com/energye/energy/v2/cmd/internal/term"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

var remoteVersion *remotecfg.TCMDVersion

func version() error {
	var err error
	if remoteVersion == nil {
		remoteVersion, err = remotecfg.CMDVersion()
		if err != nil {
			term.Logger.Error(err.Error())
			return err
		}
	}
	return nil
}

// CheckVersion 检查版本
func CheckVersion() string {
	if err := version(); err != nil {
		return ""
	}
	term.Section.Println("CLI Current:", fmt.Sprintf("v%d.%d.%d", term.Major, term.Minor, term.Build))
	remoteVer := fmt.Sprintf("v%d.%d.%d", remoteVersion.Major, remoteVersion.Minor, remoteVersion.Build)
	term.Section.Println("CLI Latest :", remoteVer)
	cv, err := strconv.Atoi(fmt.Sprintf("%d%d%d", term.Major, term.Minor, term.Build))
	if err != nil {
		term.Logger.Error("Check cli version failed: " + err.Error())
		return ""
	}
	rv, err := strconv.Atoi(fmt.Sprintf("%d%d%d", remoteVersion.Major, remoteVersion.Minor, remoteVersion.Build))
	if err != nil {
		term.Logger.Error("Check cli version failed: " + err.Error())
		return ""
	}
	if cv < rv {
		// 下载 URL 规则:
		// https://sourceforge.net/projects/energye/files/vx.x.x/energy-[os][arch].zip
		// https://github.com/energye/energy/releases/download/vx.x.x/energy-[os][arch].zip
		// https://gitee.com/energye/energy/releases/download/vx.x.x/energy-[os][arch].zip
		cliName := CliFileName() + ".zip"
		downloadURL := remoteVersion.DownloadURL
		if strings.LastIndex(downloadURL, "/") != len(downloadURL) {
			downloadURL += "/"
		}
		downloadURL = fmt.Sprintf("%v%v/%v", downloadURL, remoteVer, cliName)
		term.Section.Println("There new version available.\n  Download:", downloadURL)
		return downloadURL
	}
	return ""
}

func CliFileName() string {
	cliName := consts.ENERGY + "-" + runtime.GOOS
	arch := runtime.GOARCH
	if arch == "amd64" {
		cliName += "64"
	} else if arch == "386" {
		cliName += "32"
	} else {
		// arm64, arm
		cliName += arch
	}
	return cliName
}

type GeoInfo struct {
	Country string `json:"country"`
}

func area() {
	resp, err := http.Get("http://ip-api.com/json/") // 使用公共API，注意实际使用时选择合适的服务
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	var geoInfo GeoInfo
	err = json.Unmarshal(body, &geoInfo)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Country:", geoInfo.Country)
}

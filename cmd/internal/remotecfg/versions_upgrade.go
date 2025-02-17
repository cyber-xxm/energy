//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package remotecfg

import (
	"encoding/json"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/env"
	"github.com/energye/energy/v2/cmd/internal/tools"
)

var upgradeList map[string]TVersionsUpgrade

type TVersionsUpgrade struct {
	Enable           int               `json:"enable"`
	Identical        string            `json:"identical"`
	DependenceModule TDependenceModule `json:"dependenceModule"`
}

type TDependenceModule struct {
	LCL map[string]string `json:"lcl"`
	CEF map[string]string `json:"cef"`
}

// 版本发布升级列表
func VersionUpgradeList() (map[string]TVersionsUpgrade, error) {
	if upgradeList == nil {
		data, err := tools.Get(env.GlobalDevEnvConfig.RemoteURL(consts.VERSIONS_UPGRADE_URL), env.GlobalDevEnvConfig.Proxy)
		if err != nil {
			return nil, err
		}
		var vu map[string]TVersionsUpgrade
		err = json.Unmarshal(data, &vu)
		if err != nil {
			return nil, err
		}
		upgradeList = vu
	}
	return upgradeList, nil
}

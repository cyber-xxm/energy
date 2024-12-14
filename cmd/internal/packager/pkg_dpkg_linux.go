//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux
// +build linux

package packager

import (
	"errors"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/assets"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/env"
	"github.com/energye/energy/v2/cmd/internal/project"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	cmd "github.com/energye/energy/v2/cmd/internal/tools/cmd"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	deb               = "DEBIAN"
	debControl        = deb + "/control"
	debPreinit        = deb + "/preinit"
	debPostinit       = deb + "/postinit"
	debPrerm          = deb + "/prerm"
	debPostrm         = deb + "/postrm"
	usrSharApps       = "usr/share/applications"
	optCompanyProduct = "opt/%s/%s"
)

const (
	linuxDebControl = "linux/control"
	linuxAppDesktop = "linux/app.desktop"
	linuxARMStartup = "linux/startup.sh"
)

func GeneraInstaller(c *command.Config, proj *project.Project) error {
	if !tools.CommandExists("dpkg") {
		return errors.New("failed to create application installation program. Could not find the dpkg command")
	}
	// 创建构建输出目录
	appRoot := fmt.Sprintf("linux/%s-%s", proj.Name, proj.Info.ProductVersion)
	buildOutDir := assets.BuildOutPath(proj)
	buildOutDir = filepath.Join(buildOutDir, appRoot)
	if !tools.IsExist(buildOutDir) {
		if err := os.MkdirAll(buildOutDir, 0755); err != nil {
			return fmt.Errorf("unable to create directory: %w", err)
		}
	}
	var err error
	// create debian/control
	if err = linuxControl(proj, appRoot); err != nil {
		return err
	}
	// create debian/copyright
	if err = linuxCopyright(proj, appRoot); err != nil {
		return err
	}
	// create app.desktop
	if err = linuxDesktop(proj, appRoot); err != nil {
		return err
	}
	// copy source
	if err = linuxOptCopy(proj, appRoot); err != nil {
		return err
	}
	// copy linux arm startup.sh
	if err = linuxARMStartupSH(proj, appRoot); err != nil {
		return err
	}
	// 7zz 压缩 CEF
	comper := proj.NSIS.Compress
	switch comper {
	case "7z", "7za":
		proj.NSIS.UseCompress = env.GlobalDevEnvConfig.Z7ZCMD() != ""
	}

	// dpkg -b
	var debName string
	if debName, err = dpkgB(proj); err != nil {
		return err
	}
	// out log
	outInstall := filepath.Join(assets.BuildOutPath(proj), "linux", debName)
	successLog := "Success \n\tInstall Package: %s\n\tInstall: sudo dpkg -i %s\n\tRemove:  sudo dpkg -r %s"
	term.Section.Println(fmt.Sprintf(successLog, outInstall, debName, proj.Dpkg.Package))
	return nil
}

func dpkgB(proj *project.Project) (string, error) {
	dir := filepath.Join(assets.BuildOutPath(proj), "linux")
	//sudo dpkg -b demo-1.0.0/ demo-[os]-[arch].deb
	app := fmt.Sprintf("%s-%s", proj.Name, proj.Info.ProductVersion)
	debName := fmt.Sprintf("%s-%s-%s.deb", proj.Name, runtime.GOOS, runtime.GOARCH)
	outFile := filepath.Join(dir, debName)
	term.Logger.Info("Generate dpkg package. Almost complete", term.Logger.Args("deb", debName))
	cmd := cmd.NewCMD()
	cmd.IsPrint = false
	cmd.Dir = dir
	var err error
	cmd.MessageCallback = func(bytes []byte, e error) {
		if e != nil {
			err = e
		}
	}
	os.Remove(outFile)
	var args = []string{"-b", app, debName}
	cmd.Command("dpkg", args...)
	cmd.Close()
	return debName, err
}

func opt(proj *project.Project) string {
	return filepath.Join("/opt", proj.Info.CompanyName, proj.Info.ProductName)
}

func linuxOptCopy(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate dpkg copy:",
		term.Logger.Args("company", proj.Info.CompanyName, "product", proj.Info.ProductName, "opt",
			fmt.Sprintf("/opt/%s/%s", proj.Info.CompanyName, proj.Info.ProductName)))
	buildOutDir := assets.BuildOutPath(proj)
	appDir := filepath.Join(buildOutDir, appRoot)
	// app/opt/[company]/[product]
	optDir := filepath.Join(appDir, fmt.Sprintf(optCompanyProduct, proj.Info.CompanyName, proj.Info.ProductName))
	if err := os.MkdirAll(optDir, 0755); err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}
	exeDir := filepath.Join(proj.ProjectPath, proj.OutputFilename)
	if !tools.IsExist(exeDir) {
		return fmt.Errorf("execution file not found: %s", exeDir)
	}
	exeIconDir := proj.Info.Icon
	if !tools.IsExist(exeIconDir) {
		return fmt.Errorf("exeIcon file not found: %s", exeIconDir)
	}

	term.Logger.Info("Generate dpkg execution " + exeDir)
	cefDir := env.GlobalDevEnvConfig.Framework
	if !tools.IsExist(cefDir) {
		return fmt.Errorf("%s not found", cefDir)
	}
	term.Logger.Info("Generate dpkg framework " + cefDir)
	var copyFiles = func(src, dst string) error {
		if srcFile, err := os.Open(src); err != nil {
			return err
		} else {
			st, err := srcFile.Stat()
			if err != nil {
				return err
			}
			if st.IsDir() {
				srcFile.Close() //close
				var pathLen = len(src)
				err := filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
					if path == src { // current root
						return nil
					}
					outPath := path[pathLen:]
					// exclude file or dir
					for _, p := range proj.Dpkg.Exclude {
						if strings.Contains(outPath, p) {
							return nil
						}
					}
					targetPath := filepath.Join(dst, outPath)
					info, _ := d.Info()
					if d.IsDir() {
						return os.MkdirAll(targetPath, info.Mode())
					} else {
						if tools.IsExistAndSize(targetPath, info.Size()) {
							//term.Logger.Info("\tcopy skip: " + outPath)
							return nil
						}
						srcFile, err := os.Open(path)
						if err != nil {
							return err
						}
						defer srcFile.Close()
						dstFile, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, info.Mode())
						if err != nil {
							return err
						}
						defer dstFile.Close()
						//term.Logger.Info("\tcopy: " + outPath)
						_, err = io.Copy(dstFile, srcFile)
						return err
					}
				})
				if err != nil {
					return err
				}
			} else {
				defer srcFile.Close()
				dstFilePath := filepath.Join(dst, st.Name())
				dstFile, err := os.OpenFile(dstFilePath, os.O_CREATE|os.O_WRONLY, st.Mode())
				if err != nil {
					return err
				}
				defer dstFile.Close()
				_, err = io.Copy(dstFile, srcFile)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	term.Logger.Info("Generate dpkg copy:", term.Logger.Args("execution", exeDir))
	if err := copyFiles(exeDir, optDir); err != nil {
		return err
	}
	term.Logger.Info("Generate dpkg copy:", term.Logger.Args("icon", exeIconDir))
	if err := copyFiles(exeIconDir, optDir); err != nil {
		return err
	}
	term.Logger.Info("Generate dpkg copy:", term.Logger.Args("framework", cefDir))
	if err := copyFiles(cefDir, optDir); err != nil {
		return err
	}
	return nil
}

func linuxARMStartupSH(proj *project.Project, appRoot string) error {
	if consts.IsLinux && consts.IsARM64 {
		term.Logger.Info("Generate dpkg startup.sh")
		buildOutDir := assets.BuildOutPath(proj)
		appDir := filepath.Join(buildOutDir, appRoot)
		if startupshData, err := assets.ReadFile(proj, assetsFSPath, linuxARMStartup); err != nil {
			return err
		} else {
			data := make(map[string]interface{})
			data["INSTALLPATH"] = opt(proj)
			data["EXECUTE"] = proj.Name
			sh := strings.NewReplacer("\r", "")
			if content, err := tools.RenderTemplate(sh.Replace(string(startupshData)), data); err != nil {
				return err
			} else {
				optDir := opt(proj)
				outFilePath := filepath.Join(appDir, optDir, fmt.Sprintf("%s.sh", proj.Name))
				outFile, err := os.OpenFile(outFilePath, os.O_CREATE|os.O_WRONLY, 0755)
				if err != nil {
					return err
				}
				defer outFile.Close()
				outFile.Write(content)
			}

		}
	}
	return nil
}

func linuxDesktop(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate dpkg desktop")
	buildOutDir := assets.BuildOutPath(proj)
	appDir := filepath.Join(buildOutDir, appRoot)
	// app/usr/share/applications
	apps := filepath.Join(appDir, usrSharApps)
	if err := os.MkdirAll(apps, 0755); err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}
	if desktopData, err := assets.ReadFile(proj, assetsFSPath, linuxAppDesktop); err != nil {
		return err
	} else {
		optDir := opt(proj)
		_, icon := filepath.Split(proj.Info.Icon)
		startup := proj.Name
		if consts.IsLinux && consts.IsARM64 {
			startup += ".sh"
		}
		data := make(map[string]interface{})
		data["Name"] = proj.Name
		data["Exec"] = filepath.Join(optDir, startup)
		data["Icon"] = filepath.Join(optDir, icon)
		data["Comments"] = proj.Info.Comments
		if content, err := tools.RenderTemplate(string(desktopData), data); err != nil {
			return err
		} else {
			debControlFile := filepath.Join(appRoot, usrSharApps, fmt.Sprintf("%s.desktop", proj.Name))
			if err = assets.WriteFile(proj, debControlFile, content); err != nil {
				return err
			}
		}
	}
	return nil
}

func linuxCopyright(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate dpkg copyright")
	return nil
}

func linuxControl(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate dpkg control")
	buildOutDir := assets.BuildOutPath(proj)
	appDir := filepath.Join(buildOutDir, appRoot)
	// DEBIAN app/DEBIAN
	debDir := filepath.Join(appDir, deb)
	if err := os.MkdirAll(debDir, 0755); err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}
	if controlData, err := assets.ReadFile(proj, assetsFSPath, linuxDebControl); err != nil {
		return err
	} else {
		data := make(map[string]interface{})
		data["Arch"] = runtime.GOARCH
		data["Info"] = proj.Info
		data["Author"] = proj.Author
		data["Dpkg"] = proj.Dpkg
		if content, err := tools.RenderTemplate(string(controlData), data); err != nil {
			return err
		} else {
			debControlFile := filepath.Join(appRoot, debControl)
			if err = assets.WriteFile(proj, debControlFile, content); err != nil {
				return err
			}
		}
	}
	return nil
}

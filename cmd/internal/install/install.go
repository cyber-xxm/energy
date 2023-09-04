//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----

package install

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/bzip2"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/env"
	progressbar "github.com/energye/energy/v2/cmd/internal/progress-bar"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type downloadInfo struct {
	fileName      string
	frameworkPath string
	downloadPath  string
	url           string
	success       bool
	isSupport     bool
	module        string
}

func Install(c *command.Config) {
	// 创建安装目录
	initInstall(c)
	// 检查Go开发环境
	//if !tools.CommandExists("go") {
	println("Golang development environment not installed!")
	// 如未安装Go开发环境，自动安装Go环境
	goRoot := installGolang(c)
	//goRoot := "C:\\go"
	//}
	// 安装CEF二进制框架
	installCEFFramework(c)
	// 设置 energy 环境变量
	env.SetEnergyHomeEnv(cefInstallPathName(c))
	// 设置 go 环境变量
	if goRoot != "" {
		env.SetGoEnv(goRoot)
	}
}

func cefInstallPathName(c *command.Config) string {
	return filepath.Join(c.Install.Path, c.Install.Name)
}

func initInstall(c *command.Config) {
	if c.Install.Path == "" {
		// current dir
		c.Install.Path = c.Wd
	}
	if c.Install.Version == "" {
		// latest
		c.Install.Version = "latest"
	}
	// 创建安装目录
	os.MkdirAll(c.Install.Path, fs.ModePerm)
	os.MkdirAll(cefInstallPathName(c), fs.ModePerm)
	os.MkdirAll(filepath.Join(c.Install.Path, command.FrameworkCache), fs.ModePerm)
}

func installGolang(c *command.Config) string {
	print("Do you want to install the Go development environment? Y/n: ")
	var s string
	fmt.Scanln(&s)
	if strings.ToLower(s) == "y" {
		s = c.Install.Path // 安装目录
		exts := map[string]string{
			"darwin":  "tar.gz",
			"linux":   "tar.gz",
			"windows": "zip",
		}
		// 开始下载并安装Go开发环境
		version := command.GolangDefaultVersion
		gos := runtime.GOOS
		arch := runtime.GOARCH
		gos = "darwin"
		arch = "amd64"
		ext := exts[gos]
		if !tools.IsExist(s) {
			println("Directory does not exist. Creating directory.", s)
			if err := os.MkdirAll(s, fs.ModePerm); err != nil {
				println("Failed to create goroot directory", err.Error())
				return ""
			}
		}
		fileName := fmt.Sprintf("go%s.%s-%s.%s", version, gos, arch, ext)
		downloadUrl := fmt.Sprintf(command.GolangDownloadURL, fileName)
		savePath := filepath.Join(s, command.FrameworkCache, fileName)
		var err error
		println("Golang Download URL:", downloadUrl)
		println("Golang Save Path:", savePath)
		if !tools.IsExist(savePath) {
			// 已经存在不再下载
			bar := progressbar.NewBar(100)
			bar.SetNotice("\t")
			bar.HideRatio()
			err = downloadFile(downloadUrl, savePath, func(totalLength, processLength int64) {
				bar.PrintBar(int((float64(processLength) / float64(totalLength)) * 100))
			})
			if err != nil {
				bar.PrintEnd("Download [" + fileName + "] failed: " + err.Error())
			} else {
				bar.PrintEnd("Download [" + fileName + "] success")
			}
		}
		if err == nil {
			// 使用 go 名字做为 go 安装目录
			targetPath := filepath.Join(s, "go")
			// 释放文件
			if !command.IsWindows {
				//zip
				ExtractUnZip(savePath, targetPath)
			} else {
				//tar
				ExtractUnTar(savePath, targetPath)
			}
			return targetPath
		}
		return ""
	}
	return ""
}

func installCEFFramework(c *command.Config) {
	// 获取提取文件配置
	extractData, err := tools.HttpRequestGET(command.DownloadExtractURL)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error(), "\n")
		os.Exit(1)
	}
	var extractConfig map[string]any
	extractData = bytes.TrimPrefix(extractData, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(extractData, &extractConfig); err != nil {
		fmt.Fprint(os.Stderr, err.Error(), "\n")
		os.Exit(1)
	}
	extractOSConfig := extractConfig[runtime.GOOS].(map[string]any)

	// 获取安装版本配置
	downloadJSON, err := tools.HttpRequestGET(command.DownloadVersionURL)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	var edv map[string]any
	downloadJSON = bytes.TrimPrefix(downloadJSON, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(downloadJSON, &edv); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

	// -c cef args value
	// default(empty), windows7, gtk2, flash
	cef := strings.ToLower(c.Install.CEF)
	if cef != command.CefEmpty && cef != command.Cef109 && cef != command.Cef106 && cef != command.Cef87 {
		fmt.Fprint(os.Stderr, "-c [cef] Incorrect args value\n")
		os.Exit(1)
	}
	installPathName := cefInstallPathName(c)
	println("Install Path", installPathName)

	println("Start downloading CEF and Energy dependency")
	// 所有版本列表
	var versionList = edv["versionList"].(map[string]any)

	// 当前安装版本
	var installVersion map[string]any
	if c.Install.Version == "latest" {
		// 默认最新版本
		if v, ok := versionList[edv["latest"].(string)]; ok {
			installVersion = v.(map[string]any)
		}
	} else {
		// 自己选择版本
		if v, ok := versionList[c.Install.Version]; ok {
			installVersion = v.(map[string]any)
		}
	}
	println("Check version")
	if installVersion == nil || len(installVersion) == 0 {
		println("Invalid version number:", c.Install.Version)
		os.Exit(1)
	}
	// 当前版本 cef 和 liblcl 版本选择
	var (
		cefModuleName, liblclModuleName string
	)
	// 使用提供的特定版本号
	if cef == command.Cef106 {
		cefModuleName = "cef-106" // CEF 106.1.1
	} else if cef == command.Cef109 {
		cefModuleName = "cef-109" // CEF 109.1.18
	} else if cef == command.Cef87 {
		// cef 87 要和 liblcl 87 配对
		cefModuleName = "cef-87"       // CEF 87.1.14
		liblclModuleName = "liblcl-87" // liblcl 87
	}
	// 如未指定CEF参数、或参数不正确，选择当前CEF模块最（新）大的版本号
	if cefModuleName == "" {
		var cefDefault string
		var number int
		for module, _ := range installVersion {
			if strings.Index(module, "cef") == 0 {
				if s := strings.Split(module, "-"); len(s) == 2 {
					// module = "cef-xxx"
					n, _ := strconv.Atoi(s[1])
					if n >= number {
						number = n
						cefDefault = module
					}
				} else {
					// module = "cef"
					cefDefault = module
					break
				}
			}
		}
		cefModuleName = cefDefault
	}
	// liblcl, 在未指定flash版本时，它是空 ""
	if liblclModuleName == "" {
		liblclModuleName = "liblcl"
	}
	// 当前安装版本的所有模块
	var modules map[string]any
	if m, ok := installVersion["modules"]; ok {
		modules = m.(map[string]any)
	}
	// 根据模块名拿到对应的模块配置
	var (
		cefModule, liblclModule map[string]any
	)
	if module, ok := modules[cefModuleName]; ok {
		cefModule = module.(map[string]any)
	}
	if module, ok := modules[liblclModuleName]; ok {
		liblclModule = module.(map[string]any)
	}
	if cefModule == nil {
		println("error: cef module", cefModuleName, "is not configured in the current version")
		os.Exit(1)
	}
	// 下载源选择
	var replaceSource = func(url, source string, sourceSelect int, module string) string {
		s := strings.Split(source, ",")
		// liblcl 如果自己选择下载源
		if module == "liblcl" && c.Install.Download != "" {
			sourceSelect = tools.ToInt(c.Install.Download)
		}
		if len(s) > sourceSelect {
			return strings.ReplaceAll(url, "{source}", s[sourceSelect])
		}
		return url
	}
	// 下载集合
	var downloads = make(map[string]*downloadInfo)
	// 根据模块名拿到版本号
	cefVersion := tools.ToRNilString(installVersion[cefModuleName], "")
	// 当前模块版本支持系统，如果支持返回下载地址
	libCEFOS, isSupport := cefOS(cefModule)
	downloadCefURL := tools.ToString(cefModule["downloadUrl"])
	downloadCefURL = replaceSource(downloadCefURL, tools.ToString(cefModule["downloadSource"]), tools.ToInt(cefModule["downloadSourceSelect"]), "cef")
	downloadCefURL = strings.ReplaceAll(downloadCefURL, "{version}", cefVersion)
	downloadCefURL = strings.ReplaceAll(downloadCefURL, "{OSARCH}", libCEFOS)
	downloads[command.CefKey] = &downloadInfo{isSupport: isSupport, fileName: urlName(downloadCefURL), downloadPath: filepath.Join(c.Install.Path, command.FrameworkCache, urlName(downloadCefURL)), frameworkPath: installPathName, url: downloadCefURL, module: cefModuleName}

	// liblcl
	// 如果选定的cef 106，在linux会指定liblcl gtk2 版本, 其它系统和版本以默认的形式区分
	// 最后根据模块名称来确定使用哪个liblcl
	liblclVersion := tools.ToRNilString(installVersion[liblclModuleName], "")
	if liblclModule != nil {
		libEnergyOS, isSupport := liblclOS(cef, liblclVersion, tools.ToString(liblclModule["buildSupportOSArch"]))
		downloadEnergyURL := tools.ToString(liblclModule["downloadUrl"])
		downloadEnergyURL = replaceSource(downloadEnergyURL, tools.ToString(liblclModule["downloadSource"]), tools.ToInt(liblclModule["downloadSourceSelect"]), "liblcl")
		module := tools.ToString(liblclModule["module"])
		downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{version}", liblclVersion)
		downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{module}", module)
		downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{OSARCH}", libEnergyOS)
		downloads[command.LiblclKey] = &downloadInfo{isSupport: isSupport, fileName: urlName(downloadEnergyURL), downloadPath: filepath.Join(c.Install.Path, command.FrameworkCache, urlName(downloadEnergyURL)), frameworkPath: installPathName, url: downloadEnergyURL, module: liblclModuleName}
	}

	// 在线下载框架二进制包
	for key, dl := range downloads {
		fmt.Printf("Download %s: %s\n", key, dl.url)
		if !dl.isSupport {
			println("Warn module is not built or configured 【", dl.module, "】")
			continue
		}
		bar := progressbar.NewBar(100)
		bar.SetNotice("\t")
		bar.HideRatio()
		err = downloadFile(dl.url, dl.downloadPath, func(totalLength, processLength int64) {
			bar.PrintBar(int((float64(processLength) / float64(totalLength)) * 100))
		})
		bar.PrintEnd("Download [" + dl.fileName + "] success")
		if err != nil {
			println("Download [", dl.fileName, "] error:", err.Error())
			os.Exit(1)
		}
		dl.success = err == nil
	}
	// 解压文件, 并根据配置提取文件
	println("Unpack files")
	for key, di := range downloads {
		if !di.isSupport {
			println("Warn module is not built or configured 【", di.module, "】")
			continue
		}
		if di.success {
			if key == command.CefKey {
				bar := progressbar.NewBar(0)
				bar.SetNotice("Unpack file " + key + ": ")
				tarName := UnBz2ToTar(di.downloadPath, func(totalLength, processLength int64) {
					bar.PrintSizeBar(processLength)
				})
				bar.PrintEnd()
				ExtractFiles(key, tarName, di, extractOSConfig)
			} else if key == command.LiblclKey {
				ExtractFiles(key, di.downloadPath, di, extractOSConfig)
			}
			println("Unpack file", key, "success\n")
		}
	}
	println("\nSUCCESS \nInstalled version:", c.Install.Version, liblclVersion)
	if liblclModule == nil {
		println("hint: liblcl module", liblclModuleName, `is not configured in the current version, You need to use built-in binary build. [go build -tags="tempdll"]`)
	}
}

func cefOS(module map[string]any) (string, bool) {
	buildSupportOSArch := tools.ToString(module["buildSupportOSArch"])
	mod := tools.ToString(module["module"])
	archs := strings.Split(buildSupportOSArch, ",")
	var isSupport = func(goarch string) bool {
		for _, v := range archs {
			if goarch == v {
				return true
			}
		}
		return false
	}
	if command.IsWindows { // windows arm for 64 bit, windows for 32/64 bit
		if runtime.GOARCH == "arm64" {
			return "windowsarm64", isSupport(command.WindowsARM64)
		}
		if strconv.IntSize == 32 {
			return fmt.Sprintf("windows%d", strconv.IntSize), isSupport(command.Windows32)
		}
		return fmt.Sprintf("windows%d", strconv.IntSize), isSupport(command.Windows64)
	} else if command.IsLinux { //linux for 64 bit
		if runtime.GOARCH == "arm64" {
			if mod == command.Cef106 {
				return "linuxarm64", isSupport(command.LinuxARM64GTK2)
			}
			return "linuxarm64", isSupport(command.LinuxARM64) || isSupport(command.LinuxARM64GTK3)
		} else if runtime.GOARCH == "amd64" {
			if mod == command.Cef106 {
				return "linux64", isSupport(command.Linux64GTK2)
			}
			return "linux64", isSupport(command.Linux64) || isSupport(command.Linux64GTK3)
		}
	} else if command.IsDarwin { // macosx for 64 bit
		//if runtime.GOARCH == "arm64" {
		//	return "macosarm64", isSupport(MacOSARM64)
		//} else if runtime.GOARCH == "amd64" {
		//	return "macosx64", isSupport(MacOSX64)
		//}
		// Mac amd64 m1 m2 架构目前使用amd64, m1,m2使用Rosetta2兼容
		return "macosx64", isSupport(command.MacOSX64)
	}
	//not support
	return fmt.Sprintf("%v %v", runtime.GOOS, runtime.GOARCH), false
}

var liblclFileNames = map[string]string{
	"windows32":       command.Windows32,
	"windows64":       command.Windows64,
	"windowsarm64":    command.WindowsARM64,
	"linuxarm64":      command.LinuxARM64,
	"linuxarm64gtk2":  command.LinuxARM64GTK2,
	"linux64":         command.Linux64,
	"linux64gtk2":     command.Linux64GTK2,
	"darwin64":        command.MacOSX64,
	"darwinarm64":     command.MacOSARM64,
	"windows32_old":   "Windows 32 bits",
	"windows64_old":   "Windows 64 bits",
	"linux64gtk2_old": "Linux GTK2 x86 64 bits",
	"linux64_old":     "Linux x86 64 bits",
	"darwin64_old":    "MacOSX x86 64 bits",
}

func liblclName(version, cef string) (string, bool) {
	var key string
	var isOld bool
	if runtime.GOARCH == "arm64" {
		if command.IsLinux && cef == command.Cef106 { // 只linux区别liblcl gtk2
			key = "linuxarm64gtk2"
		} else {
			if command.IsDarwin {
				// Mac amd64 m1 m2 架构目前使用amd64, m1,m2使用Rosetta2兼容
				key = fmt.Sprintf("%samd64", runtime.GOOS)
			} else {
				key = fmt.Sprintf("%sarm64", runtime.GOOS)
			}
		}
	} else if runtime.GOARCH == "amd64" {
		if command.IsLinux && cef == command.Cef106 { // 只linux区别liblcl gtk2
			key = "linux64gtk2"
		} else {
			key = fmt.Sprintf("%s%d", runtime.GOOS, strconv.IntSize)
		}
	}
	if tools.Compare("2.2.4", version) {
		if key != "" {
			key += "_old"
			isOld = true
		}
	}
	if key != "" {
		return liblclFileNames[key], isOld
	}
	return "", false
}

// 命名规则 OS+[ARCH]+BIT+[GTK2]
//  ARCH: 非必需, ARM 时填写, AMD为空
//  GTK2: 非必需, GTK2(Linux CEF 106) 时填写, 非Linux或GTK3时为空
func liblclOS(cef, version, buildSupportOSArch string) (string, bool) {
	archs := strings.Split(buildSupportOSArch, ",")
	noSuport := fmt.Sprintf("%v %v", runtime.GOOS, runtime.GOARCH)
	var isSupport = func(goarch string) bool {
		for _, v := range archs {
			if goarch == v {
				return true
			}
		}
		return false
	}
	if name, isOld := liblclName(version, cef); isOld {
		if name == "" {
			return noSuport, false
		}
		return name, true
	} else {
		return name, isSupport(name)
	}
}

// LibLCLName
func LibLCLName(version, buildSupportOSArch string) (string, bool) {
	return liblclOS("", version, buildSupportOSArch)
}

// 提取文件
func ExtractFiles(keyName, sourcePath string, di *downloadInfo, extractOSConfig map[string]any) {
	println("Extract", keyName, "sourcePath:", sourcePath, "targetPath:", di.frameworkPath)
	files := extractOSConfig[keyName].([]any)
	if keyName == command.CefKey {
		//tar
		ExtractUnTar(sourcePath, di.frameworkPath, files...)
	} else if keyName == command.LiblclKey {
		//zip
		ExtractUnZip(sourcePath, di.frameworkPath, files...)
	}
}

func filePathInclude(compressPath string, files ...any) (string, bool) {
	if len(files) == 0 {
		return compressPath, true
	} else {
		for _, file := range files {
			f := file.(string)
			tIdx := strings.LastIndex(f, "/*")
			if tIdx != -1 {
				f = f[:tIdx]
				if f[0] == '/' {
					if strings.Index(compressPath, f[1:]) == 0 {
						return compressPath, true
					}
				}
				if strings.Index(compressPath, f) == 0 {
					return strings.Replace(compressPath, f, "", 1), true
				}
			} else {
				if f[0] == '/' {
					if compressPath == f[1:] {
						return f, true
					}
				}
				if compressPath == f {
					f = f[strings.Index(f, "/")+1:]
					return f, true
				}
			}
		}
	}
	return "", false
}

func dir(path string) string {
	path = strings.ReplaceAll(path, "\\", string(filepath.Separator))
	lastSep := strings.LastIndex(path, string(filepath.Separator))
	return path[:lastSep]
}

func ExtractUnTar(filePath, targetPath string, files ...any) {
	reader, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error: cannot open file, error=[%v]\n", err)
		return
	}
	defer reader.Close()

	var tarReader *tar.Reader

	if filepath.Ext(filePath) == ".gz" {
		gr, err := gzip.NewReader(reader)
		if err != nil {
			fmt.Printf("error: cannot open gzip file, error=[%v]\n", err)
			return
		}
		defer gr.Close()
		tarReader = tar.NewReader(gr)
	} else {
		tarReader = tar.NewReader(reader)
	}

	bar := progressbar.NewBar(100)
	bar.SetNotice("\t")
	bar.HideRatio()
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error: cannot read tar file, error=[%v]\n", err)
			os.Exit(1)
			return
		}
		// 去除压缩包内的一级目录
		compressPath := filepath.Clean(header.Name[strings.Index(header.Name, "/")+1:])
		//compressPath := filepath.Clean(header.Name[])
		includePath, isInclude := filePathInclude(compressPath, files...)
		if !isInclude {
			continue
		}
		info := header.FileInfo()
		targetFile := filepath.Join(targetPath, includePath)
		fmt.Println("compressPath:", compressPath)
		if info.IsDir() {
			if err = os.MkdirAll(targetFile, info.Mode()); err != nil {
				fmt.Printf("error: cannot mkdir file, error=[%v]\n", err)
				os.Exit(1)
				return
			}
		} else {
			fDir := dir(targetFile)
			_, err = os.Stat(fDir)
			if os.IsNotExist(err) {
				if err = os.MkdirAll(fDir, info.Mode()); err != nil {
					fmt.Printf("error: cannot file mkdir file, error=[%v]\n", err)
					os.Exit(1)
					return
				}
			}
			file, err := os.OpenFile(targetFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
			if err != nil {
				fmt.Printf("error: cannot open file, error=[%v]\n", err)
				os.Exit(1)
				return
			}
			bar.SetCurrentValue(0)
			writeFile(tarReader, file, header.Size, func(totalLength, processLength int64) {
				bar.PrintBar(int((float64(processLength) / float64(totalLength)) * 100))
			})
			file.Sync()
			file.Close()
			bar.PrintBar(100)
			bar.PrintEnd()
			if err != nil {
				fmt.Printf("error: cannot write file, error=[%v]\n", err)
				os.Exit(1)
				return
			}
		}
	}
}

func ExtractUnZip(filePath, targetPath string, files ...any) {
	if rc, err := zip.OpenReader(filePath); err == nil {
		defer rc.Close()
		bar := progressbar.NewBar(100)
		bar.SetNotice("\t")
		bar.HideRatio()
		var createWriteFile = func(info fs.FileInfo, path string, file io.Reader) {
			targetFileName := filepath.Join(targetPath, path)
			if info.IsDir() {
				os.MkdirAll(targetFileName, info.Mode())
				return
			}
			if targetFile, err := os.Create(targetFileName); err == nil {
				fmt.Println("extract file: ", path)
				bar.SetCurrentValue(0)
				writeFile(file, targetFile, info.Size(), func(totalLength, processLength int64) {
					bar.PrintBar(int((float64(processLength) / float64(totalLength)) * 100))
				})
				bar.PrintBar(100)
				bar.PrintEnd()
				targetFile.Close()
			}
		}
		// 所有文件
		if len(files) == 0 {
			zipFiles := rc.File
			for _, f := range zipFiles {
				r, _ := f.Open()
				name := filepath.Clean(f.Name)
				createWriteFile(f.FileInfo(), name, r.(io.Reader))
				_ = r.Close()
			}
		} else {
			// 指定名字的文件
			for i := 0; i < len(files); i++ {
				if f, err := rc.Open(files[i].(string)); err == nil {
					info, _ := f.Stat()
					createWriteFile(info, files[i].(string), f)
					_ = f.Close()
				} else {
					fmt.Printf("error: cannot open file, error=[%v]\n", err)
					os.Exit(1)
					return
				}
			}
		}
	} else {
		if err != nil {
			fmt.Printf("error: cannot read zip file, error=[%v]\n", err)
			os.Exit(1)
		}
	}
}

// 释放bz2文件到tar
func UnBz2ToTar(name string, callback func(totalLength, processLength int64)) string {
	fileBz2, err := os.Open(name)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		os.Exit(1)
	}
	defer fileBz2.Close()
	dirName := fileBz2.Name()
	dirName = dirName[:strings.LastIndex(dirName, ".")]
	if !tools.IsExist(dirName) {
		r := bzip2.NewReader(fileBz2)
		w, err := os.Create(dirName)
		if err != nil {
			fmt.Errorf("%s", err.Error())
			os.Exit(1)
		}
		defer w.Close()
		writeFile(r, w, 0, callback)
	} else {
		println("File already exists")
	}
	return dirName
}

func writeFile(r io.Reader, w *os.File, totalLength int64, callback func(totalLength, processLength int64)) {
	buf := make([]byte, 1024*10)
	var written int64
	for {
		nr, err := r.Read(buf)
		if nr > 0 {
			nw, err := w.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			callback(totalLength, written)
			if err != nil {
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if err != nil {
			break
		}
	}
}

// url文件名
func urlName(downloadUrl string) string {
	if u, err := url.QueryUnescape(downloadUrl); err != nil {
		return ""
	} else {
		u = u[strings.LastIndex(u, "/")+1:]
		return u
	}
}

func isFileExist(filename string, filesize int64) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if filesize == info.Size() {
		return true
	}
	os.Remove(filename)
	return false
}

// 下载文件
func downloadFile(url string, localPath string, callback func(totalLength, processLength int64)) error {
	var (
		fsize   int64
		buf     = make([]byte, 1024*10)
		written int64
	)
	tmpFilePath := localPath + ".download"
	client := new(http.Client)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("download-error=[%v]\n", err)
		os.Exit(1)
		return err
	}
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		fmt.Printf("download-error=[%v]\n", err)
		os.Exit(1)
		return err
	}
	if isFileExist(localPath, fsize) {
		println("File already exists")
		return nil
	}
	println("Save path: [", localPath, "] file size:", fsize)
	file, err := os.Create(tmpFilePath)
	if err != nil {
		fmt.Printf("download-error=[%v]\n", err)
		os.Exit(1)
		return err
	}
	defer file.Close()
	if resp.Body == nil {
		fmt.Printf("Download-error=[body is null]\n")
		os.Exit(1)
		return nil
	}
	defer resp.Body.Close()
	for {
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			nw, err := file.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			callback(fsize, written)
			if err != nil {
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	if err == nil {
		file.Sync()
		file.Close()
		err = os.Rename(tmpFilePath, localPath)
		if err != nil {
			return err
		}
	}
	return err
}

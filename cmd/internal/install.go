//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package internal

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/bzip2"
	"encoding/json"
	"fmt"
	progressbar "github.com/energye/energy/v2/cmd/internal/progress-bar"
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

var CmdInstall = &Command{
	UsageLine: "install -p [path] -v [version] -n [name] -d [download] -c [cef]",
	Short:     "Automatically configure the CEF and Energy framework",
	Long: `
	-p Installation directory Default current directory
	-v Specifying a version number,Default latest
	-n Name of the frame after installation
	-d Download Source, gitee or github, Default gitee
	-c Install system supports CEF version, provide 4 options, default empty
		default : Automatically select support for the latest version based on the current system.
		windows7: CEF 109.1.18 is the last one to support Windows 7.
		gtk2    : CEF 106.1.1 is the last default support for GTK2 in Linux.
		flash   : CEF 87.1.14 is the last one to support Flash.
	.  Execute default command

Automatically configure the CEF and Energy framework.

During this process, CEF and Energy are downloaded.

Default framework name is "EnergyFramework".
`,
}

type downloadInfo struct {
	fileName      string
	frameworkPath string
	downloadPath  string
	url           string
	success       bool
	isSupport     bool
}

func init() {
	CmdInstall.Run = runInstall
}

const (
	GTK3 = iota + 1
	GTK2
)
const (
	CefEmpty = ""
	CefWin7  = "windows7"
	CefGtk2  = "gtk2"
	CefFlash = "flash"
)

// https://cef-builds.spotifycdn.com/cef_binary_107.1.11%2Bg26c0b5e%2Bchromium-107.0.5304.110_windows64.tar.bz2
// 运行安装
func runInstall(c *CommandConfig) error {
	// 获取提取文件配置
	extractData, err := httpRequestGET(DownloadExtractURL)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error(), "\n")
		os.Exit(1)
	}
	// 获取安装版本配置
	downloadJSON, err := httpRequestGET(DownloadVersionURL)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

	// -c cef args value
	// default(empty), windows7, gtk2, flash
	cef := strings.ToLower(c.Install.CEF)
	//if cef != CefEmpty && cef != CefWin7 && cef != CefGtk2 && cef != CefFlash {
	//	fmt.Fprint(os.Stderr, "-c [cef] Incorrect args value\n")
	//	os.Exit(1)
	//}

	if c.Install.Path == "" {
		// current dir
		c.Install.Path = c.Wd
	}
	installPathName := filepath.Join(c.Install.Path, c.Install.Name)
	println("Install Path", installPathName)
	if c.Install.Version == "" {
		// latest
		c.Install.Version = "latest"
	}
	// 创建安装目录
	os.MkdirAll(c.Install.Path, fs.ModePerm)
	os.MkdirAll(installPathName, fs.ModePerm)
	os.MkdirAll(filepath.Join(c.Install.Path, frameworkCache), fs.ModePerm)
	println("Start downloading CEF and Energy dependency")
	var edv map[string]interface{}
	downloadJSON = bytes.TrimPrefix(downloadJSON, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(downloadJSON, &edv); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	// 所有版本列表
	var versionList = edv["versionList"].(map[string]interface{})

	// 当前安装版本
	var installVersion map[string]interface{}
	if c.Install.Version == "latest" {
		// 默认最新版本
		if v, ok := versionList[edv["latest"].(string)]; ok {
			installVersion = v.(map[string]interface{})
		}
	} else {
		// 自己选择版本
		if v, ok := versionList[c.Install.Version]; ok {
			installVersion = v.(map[string]interface{})
		}
	}
	println("Check version")
	if installVersion == nil || len(installVersion) == 0 {
		println("Invalid version number:", c.Install.Version)
		os.Exit(1)
	}
	// 当前版本 cef 和 liblcl 版本选择
	var (
		cefVersion, energyVersion       string
		cefModuleName, energyModuleName string
	)
	// 使用提供的特定版本号
	if cef == CefGtk2 {
		cefModuleName = "cef-106" // CEF 106.1.1
	} else if cef == CefWin7 {
		cefModuleName = "cef-109" // CEF 109.1.18
	} else if cef == CefFlash {
		// cef 87 要和 liblcl 87 配对
		cefModuleName = "cef-87"       // CEF 87.1.14
		energyModuleName = "liblcl-87" // liblcl 87
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
	if energyModuleName == "" {
		energyModuleName = "liblcl"
	}
	// 根据模块名拿到版本号
	cefVersion = ToRNilString(installVersion[cefModuleName], "")
	energyVersion = ToRNilString(installVersion[energyModuleName], "")
	// 当前安装版本的所有模块
	var modules map[string]any
	if m, ok := installVersion["modules"]; ok {
		modules = m.(map[string]any)
	}
	// 根据模块名拿到对应的模块配置
	fmt.Println("log:", modules)
	var downloadURL map[string]interface{}
	if c.Install.Download == "gitee" {
		downloadURL = edv["gitee"].(map[string]interface{})
	} else if c.Install.Download == "github" {
		downloadURL = edv["github"].(map[string]interface{})
	} else {
		println("Invalid download source, only support github or gitee:", c.Install.Download)
		os.Exit(1)
	}
	libCEFOS, isSupport := cefOS()
	libEnergyOS, isSupport := energyOS(0)
	var downloadCefURL = downloadURL["cefURL"].(string)
	var downloadEnergyURL = downloadURL["energyURL"].(string)
	downloadCefURL = strings.ReplaceAll(downloadCefURL, "{version}", cefVersion)
	downloadCefURL = strings.ReplaceAll(downloadCefURL, "{OSARCH}", libCEFOS)
	downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{version}", energyVersion)
	downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{OSARCH}", libEnergyOS)

	// 获取安装环境配置

	var extractConfig map[string]interface{}
	extractData = bytes.TrimPrefix(extractData, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(extractData, &extractConfig); err != nil {
		fmt.Fprint(os.Stderr, err.Error(), "\n")
		os.Exit(1)
	}
	extractOSConfig := extractConfig[runtime.GOOS].(map[string]interface{})

	var downloads = make(map[string]*downloadInfo)
	downloads[cefKey] = &downloadInfo{isSupport: isSupport, fileName: urlName(downloadCefURL), downloadPath: filepath.Join(c.Install.Path, frameworkCache, urlName(downloadCefURL)), frameworkPath: installPathName, url: downloadCefURL}
	downloads[energyKey] = &downloadInfo{isSupport: isSupport, fileName: urlName(downloadEnergyURL), downloadPath: filepath.Join(c.Install.Path, frameworkCache, urlName(downloadEnergyURL)), frameworkPath: installPathName, url: downloadEnergyURL}
	for key, dl := range downloads {
		fmt.Printf("Download %s: %s\n", key, dl.url)
		if !dl.isSupport {
			println("energy command line does not support the system architecture download 【", dl.fileName, "]")
			continue
		}
		bar := progressbar.NewBar(100)
		bar.SetNotice("\t")
		bar.HideRatio()
		err = downloadFile(dl.url, dl.downloadPath, func(totalLength, processLength int64) {
			bar.PrintBar(int((float64(processLength) / float64(totalLength)) * 100))
		})
		bar.PrintEnd("Download " + dl.fileName + " success")
		if err != nil {
			println("Download", dl.fileName, "error", err)
		}
		dl.success = err == nil
	}
	println("Unpack files")
	var removeFileList = make([]string, 0, 0)
	for key, di := range downloads {
		if !di.isSupport {
			println("energy command line does not support the system architecture, continue.")
			continue
		}
		if di.success {
			if key == cefKey {
				bar := progressbar.NewBar(0)
				bar.SetNotice("Unpack file " + key + ": ")
				tarName := UnBz2ToTar(di.downloadPath, func(totalLength, processLength int64) {
					bar.PrintSizeBar(processLength)
				})
				bar.PrintEnd()
				ExtractFiles(key, tarName, di, extractOSConfig)
				removeFileList = append(removeFileList, tarName)
			} else if key == energyKey {
				ExtractFiles(key, di.downloadPath, di, extractOSConfig)
			}
			println("Unpack file", key, "success\n")
		}
	}
	for _, rmFile := range removeFileList {
		println("Remove file", rmFile)
		os.Remove(rmFile)
	}
	setEnergyHomeEnv(EnergyHomeKey, installPathName)
	println("\n", CmdInstall.Short, "SUCCESS \nVersion:", c.Install.Version, "=>", energyVersion)
	return nil
}

func cefOS() (string, bool) {
	if isWindows { // windows arm for 64 bit, windows for 32/64 bit
		if runtime.GOARCH == "arm64" {
			return "windowsarm64", true
		}
		return fmt.Sprintf("windows%d", strconv.IntSize), true
	} else if isLinux { //linux for 64 bit
		if runtime.GOARCH == "arm64" {
			return "linuxarm64", true
		} else if runtime.GOARCH == "amd64" {
			return "linux64", true
		}
	} else if isDarwin { // macosx for 64 bit
		if runtime.GOARCH == "arm64" {
			return "macosarm64", true
		} else if runtime.GOARCH == "amd64" {
			return "macosx64", true
		}
	}
	//not support
	return fmt.Sprintf("%v %v", runtime.GOOS, runtime.GOARCH), false
}

func energyOS(gtk int) (string, bool) {
	if isWindows {
		return fmt.Sprintf("Windows %d bits", strconv.IntSize), true
	} else if isLinux {
		if gtk == GTK3 {
			return "Linux x86 64 bits", true
		}
		return "Linux GTK2 x86 64 bits", true
	} else if isDarwin {
		return "MacOSX x86 64 bits", true
	}
	//not support
	return fmt.Sprintf("%v %v", runtime.GOOS, runtime.GOARCH), false
}

// 提取文件
func ExtractFiles(keyName, sourcePath string, di *downloadInfo, extractOSConfig map[string]interface{}) {
	println("Extract", keyName, "sourcePath:", sourcePath, "targetPath:", di.frameworkPath)
	files := extractOSConfig[keyName].([]interface{})
	if keyName == cefKey {
		//tar
		ExtractUnTar(sourcePath, di.frameworkPath, files...)
	} else if keyName == energyKey {
		//zip
		ExtractUnZip(sourcePath, di.frameworkPath, files...)
	}
}

func filePathInclude(compressPath string, files ...interface{}) (string, bool) {
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
	return "", false
}

func dir(path string) string {
	path = strings.ReplaceAll(path, "\\", string(filepath.Separator))
	lastSep := strings.LastIndex(path, string(filepath.Separator))
	return path[:lastSep]
}

func ExtractUnTar(filePath, targetPath string, files ...interface{}) {
	reader, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error: cannot read tar file, error=[%v]\n", err)
		return
	}
	defer reader.Close()

	tarReader := tar.NewReader(reader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error: cannot read tar file, error=[%v]\n", err)
			os.Exit(1)
			return
		}
		compressPath := header.Name[strings.Index(header.Name, "/")+1:]
		includePath, isInclude := filePathInclude(compressPath, files...)
		if !isInclude {
			continue
		}
		info := header.FileInfo()
		targetFile := filepath.Join(targetPath, includePath)
		fmt.Println("compressPath:", compressPath, "-> targetFile:", targetFile)
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
			defer file.Close()
			bar := progressbar.NewBar(100)
			bar.SetNotice("\t")
			bar.HideRatio()
			writeFile(tarReader, file, header.Size, func(totalLength, processLength int64) {
				bar.PrintBar(int((float64(processLength) / float64(totalLength)) * 100))
			})
			file.Sync()
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

func ExtractUnZip(filePath, targetPath string, files ...interface{}) {
	if rc, err := zip.OpenReader(filePath); err == nil {
		defer rc.Close()
		for i := 0; i < len(files); i++ {
			if f, err := rc.Open(files[i].(string)); err == nil {
				defer f.Close()
				st, _ := f.Stat()
				targetFileName := filepath.Join(targetPath, st.Name())
				if st.IsDir() {
					os.MkdirAll(targetFileName, st.Mode())
					continue
				}
				if targetFile, err := os.Create(targetFileName); err == nil {
					fmt.Println("extract file: ", st.Name())
					bar := progressbar.NewBar(100)
					bar.SetNotice("\t")
					bar.HideRatio()
					writeFile(f, targetFile, st.Size(), func(totalLength, processLength int64) {
						bar.PrintBar(int((float64(processLength) / float64(totalLength)) * 100))
					})
					bar.PrintBar(100)
					bar.PrintEnd()
					targetFile.Close()
				}
			} else {
				fmt.Printf("error: cannot open file, error=[%v]\n", err)
				os.Exit(1)
				return
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
	r := bzip2.NewReader(fileBz2)
	w, err := os.Create(dirName)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		os.Exit(1)
	}
	defer w.Close()
	writeFile(r, w, 0, callback)
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

package common

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// HomeDir 代表用户的主目录
var HomeDir string

// StaticPath 代表静态文件的路径
var StaticPath string

// HomePath 代表用户目录下的 'pub' 文件夹路径
var HomePath string

// 服务器配置
var Host ServerConfigMap

// 项目配置
var Project ProjectConfigMap

func init() {
	var err error
	HomeDir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	StaticPath = filepath.Join(currentDir, "static")
	HomePath = filepath.Join(HomeDir, "pub")

	jsonPath := filepath.Join(StaticPath, "host.json")
	jsonFile, _ := os.ReadFile(jsonPath)
	json.Unmarshal(jsonFile, &Host)

	jsonPath = filepath.Join(StaticPath, "project.json")
	jsonFile, _ = os.ReadFile(jsonPath)
	json.Unmarshal(jsonFile, &Project)

}

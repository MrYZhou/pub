package common

import (
	. "log/slog"

	. "github.com/MrYZhou/outil/command"
)

// WebHelper 结构体
type WebHelper struct{}

// JavaHelper 结构体
type JavaHelper struct{}

var webHelper WebHelper
var javaHelper JavaHelper

func init() {
	webHelper = WebHelper{}
	javaHelper = JavaHelper{}
}
func (web *WebHelper) ExecPub(data WebUpload) {
	web.packageCode(data)
	web.upload(data)
}
func (web *WebHelper) packageCode(data WebUpload) {
	Info("打包web应用")
	Run(data.LocalPath, "npm run build")
}
func (web *WebHelper) upload(data WebUpload) {
	Info("上传web应用")
}

func (java *JavaHelper) ExecPub(data JarUpload) {
	java.packageCode(data)
	java.upload(data)
}
func (java *JavaHelper) packageCode(data JarUpload) {
	Info("打包java应用")
	Run(data.JavaProjectPath, data.PackageCommand)
}
func (java *JavaHelper) upload(data JarUpload) {
	Info("上传java应用")
}

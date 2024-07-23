package common

import (
	. "log/slog"

	. "github.com/MrYZhou/outil/command"
)

type Publisher interface {
	ExecPub()
	packageCode()
	upload()
}


// WebHelper 结构体
type WebHelper struct{}

// 实现 Publisher 接口的 Publish 方法
func (w WebHelper) Publish() {
	// 实现细节
}

// JavaHelper 结构体
type JavaHelper struct{}

// 实现 Publisher 接口的 Publish 方法
func (j JavaHelper) Publish() {
	// 实现细节
}

// NewWebHelper 函数
func NewWebHelper() WebHelper {
	return WebHelper{}
}

// NewJavaHelper 函数
func NewJavaHelper() JavaHelper {
	return JavaHelper{}
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

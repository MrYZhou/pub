package common

import (
	. "log/slog"
)

type Publisher interface {
	ExecPub()
	packageCode()
	upload()
}

type WebHelper struct{}
type JavaHelper struct{}

func NewWebHelper() *WebHelper {
	return &WebHelper{}
}
func NewJavaHelper() *JavaHelper {
	return &JavaHelper{}
}

func (web *WebHelper) ExecPub() {
	Info("部署web应用")
	web.packageCode()
	web.upload()
}
func (web *WebHelper) packageCode() {
	Info("打包web应用")
}
func (web *WebHelper) upload() {
	Info("上传web应用")
}

func (java *JavaHelper) ExecPub() {
	Info("部署java应用")
	java.packageCode()
	java.upload()
}
func (java *JavaHelper) packageCode() {
	Info("打包java应用")
}
func (java *JavaHelper) upload() {
	Info("上传java应用")
}

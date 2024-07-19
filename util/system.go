package util

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

var staticDir = "./static"

func App() *fiber.App {
	if app == nil {
		app = fiber.New(fiber.Config{
			DisableStartupMessage: true, // 禁用启动横幅
		})
		// 注册自定义中间件以转换上下文
		app.Use(CtxMiddleware)
		// 静态文件服务
		// 打包exe文件
		if os.Getenv("GO_BUILD_TAGS") == "prod" {
			homeDir, _ := os.UserHomeDir()
			staticDir = homeDir + "/pub"
		}
		// docker部署有可能有环境变量
		staticDirEnv := os.Getenv("staticDir")
		if staticDirEnv != "" {
			staticDir = staticDirEnv
		}
		app.Static("/", staticDir)
	}

	return app
}

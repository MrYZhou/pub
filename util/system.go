package util

import (
	. "log/slog"
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

		staticDirEnv := os.Getenv("staticDir")
		if staticDirEnv != "" {
			staticDir = staticDirEnv
		}

		Info("映射目录" + staticDir)
		app.Static("/", staticDir)
	}

	return app
}

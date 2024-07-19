package util

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func App() *fiber.App {
	if app == nil {
		app = fiber.New(fiber.Config{
			DisableStartupMessage: true, // 禁用启动横幅
		})
		// 注册自定义中间件以转换上下文
		app.Use(CtxMiddleware)
		// 静态文件服务
		homeDir, _ := os.UserHomeDir()
		currentDir, _ := os.Getwd()
		staticPath := filepath.Join(currentDir, "static")
		homePath := filepath.Join(homeDir, "pub")
		app.Static("/file", staticPath)
		app.Static("/", homePath)
	}

	return app
}

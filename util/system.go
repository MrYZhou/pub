package util

import (
	. "pub/common"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func App() *fiber.App {
	if app == nil {
		app = fiber.New(fiber.Config{
			DisableStartupMessage: false, // 禁用启动横幅
		})
		// 注册自定义中间件以转换上下文
		app.Use(CtxMiddleware)

		app.Static("/file", StaticPath)
		app.Static("/", HomePath)
	}

	return app
}

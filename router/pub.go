package router

import (
	"github.com/gofiber/fiber/v2"

	. "pub/server"
	. "pub/util"
)

// 部署应用
func init() {
	app := App()

	// 创建子路由
	api := app.Group("/pub")
	api.Get("/", pub)

	app.Post("pubweb", pubweb)

	app.Post("pubjava", pubjava)
}
func pub(c *fiber.Ctx) error {
	return AppResult(c).Success()
}

func pubweb(c *fiber.Ctx) error {
	var model WebrUpload
	// 从请求体中读取JSON内容并反序列化
	if err := c.BodyParser(&model); err != nil {
		return AppResult(c).Fail("请求体数据解析错误")
	}

	if err := Pubweb(model); err != nil {
		return AppResult(c).Fail(err.Error())
	}
	return AppResult(c).Success(model, "部署web完成")
}

func pubjava(c *fiber.Ctx) error {
	var model JarUpload
	// 从请求体中读取JSON内容并反序列化
	if err := c.BodyParser(&model); err != nil {
		return AppResult(c).Fail("请求体数据解析错误")
	}

	if err := Pubjava(model); err != nil {
		return AppResult(c).Fail(err.Error())
	}
	return AppResult(c).Success(model, "部署java完成")

}

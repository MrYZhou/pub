package api
/**
控制层
*/
import (
	"github.com/gofiber/fiber/v2"

	. "pub/server"
	. "pub/util"
)
func init() {
	app := App()

	// 创建子路由
	api := app.Group("/pub")
	api.Get("/", pub)
	// 根据id发布docker部署
	api.Post("startProject", startProject)

	// 裸机部署,需要配合下使用脚本
	app.Post("pubweb", pubweb)
	// 裸机部署,需要配合下使用脚本
	app.Post("pubjava", pubjava)
	// 上传env部署文件,到static目录或者是用户的pub目录
	app.Post("uploadEnv", uploadEnv)

}
func pub(c *fiber.Ctx) error {
	return AppResult(c).Success()
}
func startProject(c *fiber.Ctx) error {

	return AppResult(c).Success()
}
func uploadEnv(c *fiber.Ctx) error {

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

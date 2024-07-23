package router

/**
控制层
*/
import (
	"path/filepath"

	. "pub/common"
	. "pub/server"
	. "pub/util"

	"github.com/gofiber/fiber/v2"
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
	var model PubInfo
	if err := c.BodyParser(&model); err != nil {
		return AppResult(c).Fail("请求体数据解析错误")
	}
	con := GetServer(Host[model.HostId])
	defer con.Client.Close()
	defer con.SftpClient.Close()
	con.Run("pwd")
	return AppResult(c).Success(con)
}
func uploadEnv(c *fiber.Ctx) error {
	// 处理文件上传
	file, err := c.FormFile("file") // 假设表单字段名为 "file"
	if err != nil {
		return AppResult(c).Fail("文件上传失败")
	}

	// 保存文件到目录
	savePath := filepath.Join(StaticPath, ".env")
	homePath := filepath.Join(HomeDir+"/pub", ".env")
	if err := c.SaveFile(file, savePath); err != nil {
		return AppResult(c).Fail("文件保存失败")
	}
	if err := c.SaveFile(file, homePath); err != nil {
		return AppResult(c).Fail("文件保存失败")
	}
	return AppResult(c).Success()
}

func pubweb(c *fiber.Ctx) error {
	var model WebUpload
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

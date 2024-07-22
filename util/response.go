package util

import (
	"reflect"
	"github.com/gofiber/fiber/v2"
)

// MyCtx 自定义上下文结构体，用于封装HTTP响应信息
type MyCtx struct {
	*fiber.Ctx // 继承Gofiber的默认上下文
	Msg        string // 响应消息
	Code       int    // HTTP状态码
}

// AppResult 返回自定义的MyCtx实例
func AppResult(c *fiber.Ctx) *MyCtx {
	myCtx := c.Locals("myctx").(*MyCtx) // 从本地变量中获取MyCtx实例
	return myCtx
}

// Success 设置响应为成功状态
func (c *MyCtx) Success(params ...interface{}) error {
	c.Code = 200
	c.Msg = "success"
	return c.judge(params)
}

// Fail 设置响应为失败状态
func (c *MyCtx) Fail(params ...interface{}) error {
	c.Code = 400
	c.Msg = "fail"
	return c.judge(params)
}

// CtxMiddleware 中间件，用于初始化MyCtx实例
func CtxMiddleware(c *fiber.Ctx) error {
	ctx := &MyCtx{Ctx: c} // 创建MyCtx实例
	c.Locals("myctx", ctx) // 将MyCtx实例存储到本地变量中
	return c.Next() // 调用下一个中间件
}

// Response 生成并发送HTTP响应
func (c *MyCtx) Response(data interface{}) error {
	res := map[string]interface{}{
		"msg":  c.Msg,   // 消息
		"code": c.Code, // 状态码
	}
	if data != nil {
		res["data"] = data // 添加数据字段
	}
	return c.JSON(res) // 发送JSON响应
}

// judge 根据参数数量和类型构建响应
func (c *MyCtx) judge(params []interface{}) error {
	// 当只有一个参数时，判断其类型
	if len(params) == 1 {
		paramType := reflect.TypeOf(params[0])
		if paramType.Kind() == reflect.String {
			c.Msg = params[0].(string) // 如果是字符串，设置消息
			return c.Response(nil)
		}
		return c.Response(params[0]) // 如果是其他类型，直接返回数据
	}
	// 当有两个参数时
	if len(params) == 2 {
		data := params[0]
		msg := params[1].(string)
		if msg != "" {
			c.Msg = msg // 设置自定义消息
		}
		return c.Response(data) // 返回数据和可能的消息
	}

	return c.Response(nil) // 默认情况，不返回任何数据
}
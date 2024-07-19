package util

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type MyCtx struct {
	*fiber.Ctx
	Msg  string
	Code int
}

func AppResult(c *fiber.Ctx) *MyCtx {
	myCtx := c.Locals("myctx").(*MyCtx)
	return myCtx
}

func (c *MyCtx) Success(params ...interface{}) error {
	c.Code = 200
	c.Msg = "success"
	return c.judge(params)
}

func (c *MyCtx) Fail(params ...interface{}) error {
	c.Code = 400
	c.Msg = "fail"
	return c.judge(params)
}

func CtxMiddleware(c *fiber.Ctx) error {
	ctx := &MyCtx{Ctx: c}
	c.Locals("myctx", ctx)
	return c.Next()
}

func (c *MyCtx) Response(data interface{}) error {
	res := map[string]interface{}{
		"msg":  c.Msg,
		"code": c.Code,
	}
	if data != nil {
		res["data"] = data
	}
	return c.JSON(res)
}
func (c *MyCtx) judge(params []interface{}) error {
	// 参数为一个的时候,判断是字符串还是对象
	if len(params) == 1 {
		paramType := reflect.TypeOf(params[0])
		if paramType.Kind() == reflect.String {
			c.Msg = params[0].(string)
			return c.Response(nil)
		}
		return c.Response(params[0])
	}
	// 参数为两个的时候
	if len(params) == 2 {
		data := params[0]
		msg := params[1].(string)
		if msg != "" {
			c.Msg = msg
		}
		return c.Response(data)
	}

	return c.Response(nil)
}

package main

import (
	"fmt"
	"testing"

	. "github.com/MrYZhou/outil/ssh"
)

// 测试连通性
func TestConnect(t *testing.T) {

	c, err := Server("121.5.68.242:22", "root", "!123qweA")
	if err != nil {
		fmt.Println("连接失败")
	} else {
		fmt.Println(c)
	}
}

// 检查端口占用
func TestPortUse(t *testing.T) {
	c, _ := Server("121.5.68.243:22", "root", "!123qweA")
	cs, _ := c.Run("lsof -i:80")
	if len(cs) > 0 {
		fmt.Println("存在")

	}
}

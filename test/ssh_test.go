package main

import (
	"io"
	"log"
	. "log/slog"
	"net"
	"os"
	"testing"

	. "github.com/MrYZhou/outil/ssh"
	"golang.org/x/crypto/ssh"
)

// 测试连通性
func TestConnect(t *testing.T) {

	c, err := Server("47.120.11.192", "root", "!123qweA")
	if err != nil {
		Info("连接失败")
	} else {
		Info(c.Host)
	}
}
func TestConnect2(t *testing.T) {

	// 读取私钥
	privateBytes, err := os.ReadFile("d:/pub/larry.pem")
	if err != nil {
		log.Fatalf("Failed to read private key file: %v", err)
	}

	// 解析PEM格式的私钥
	privateKey, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// 创建SSH认证方法
	authMethod := ssh.PublicKeys(privateKey)

	// 连接到SSH服务器
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{authMethod},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			// 检查主机密钥，这里使用一个简单的回调函数忽略检查
			return nil
		}),
	}

	// 连接SSH服务器
	conn, err := ssh.Dial("tcp", "47.110/*  */.11.197:22", config)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	// 打开一个新的会话
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	if err != nil {
		Info("连接失败")
	} else {
		Info("连接成功")
		r, err := session.StdoutPipe()
		if err != nil {
			os.Exit(1001)
		}
		go io.Copy(os.Stdout, r)

		buf, err := session.CombinedOutput("pwd")
		Info(string(buf))
	}
}

// 检查端口占用
func TestPortUse(t *testing.T) {
	c, _ := Server("121.5.68.243:22", "root", "!123qweA")
	cs, _ := c.Run("lsof -i:80")
	if len(cs) > 0 {
		Info("存在")

	}
}

package server

import (
	. "log/slog"

	. "github.com/MrYZhou/outil/command"
	. "github.com/MrYZhou/outil/ssh"
)

/*
发布web应用
*/
func Pubweb(model WebrUpload) error {
	localPath := model.LocalPath
	remotePath := model.RemotePath
	con := Myserver()
	defer con.Client.Close()
	defer con.SftpClient.Close()
	Info("开始打包")
	err := Run(localPath, "npm run build")
	if err != nil {
		return err
	}
	Info("开始上传")
	con.UploadDir(localPath+"/dist", remotePath)
	con.Run("/www/server/nginx/sbin/nginx -s reload")
	Info("上传完毕")
	return nil
}

/*
*
发布java应用
*/
func Pubjava(model JarUpload) error {
	javaProjectPath := model.JavaProjectPath
	localJarPath := model.LocalJarPath
	remotePath := model.RemotePath
	pubCommand := model.PubCommand
	execCommand := model.ExecCommand
	if pubCommand == "" {
		pubCommand = "mvn -Dfile.encoding=UTF-8 package"
	}
	con := Myserver()
	defer con.Client.Close()
	defer con.SftpClient.Close()
	Info("开始打包")
	err := Run(javaProjectPath, pubCommand)
	if err != nil {
		return err
	}
	Info("开始上传")
	con.UploadFile(localJarPath, remotePath)
	con.Run(execCommand)
	Info("上传完毕")
	return nil
}

var host = "192.168.0.62:22"
var user = "root"
var password = "YH4WfLbGPasRLVhs"

// 初始化环境
func Myserver() *Cli {
	con, _ := Server(host, user, password)
	return con
}

// java应用的模型
type JarUpload struct {
	JavaProjectPath string `json:"javaProjectPath"` // java项目根路径
	LocalJarPath    string `json:"localJarPath"`    // 生成的jar文件路径
	RemotePath      string `json:"remotePath"`      // 远程路径
	PubCommand      string `json:"pubCommand"`      // 发布命令或打包命令
	ExecCommand     string `json:"execCommand"`     // 后置发布命令
}

// web应用的模型
type WebrUpload struct {
	LocalPath  string `json:"localPath"`  // 本地web项目路径
	RemotePath string `json:"remotePath"` // 远程web项目路径
}

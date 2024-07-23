package server

/**
抽业务层是为了代码可复用。下面这些代码和框架无关都可以直接迁移走的
*/
import (
	. "log/slog"
	. "pub/common"

	. "github.com/MrYZhou/outil/command"
	. "github.com/MrYZhou/outil/ssh"
)

/*
发布web应用
*/
func Pubweb(model WebUpload) error {
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
	packageCommand := model.PackageCommand
	execCommand := model.ExecCommand
	if packageCommand == "" {
		packageCommand = "mvn -Dfile.encoding=UTF-8 package"
	}
	con := Myserver()
	defer con.Client.Close()
	defer con.SftpClient.Close()
	Info("开始打包")
	err := Run(javaProjectPath, packageCommand)
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

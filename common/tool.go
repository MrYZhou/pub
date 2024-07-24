package common

import (
	. "log/slog"
	"os"
	"path"
	"strings"

	. "github.com/MrYZhou/outil/command"
	. "github.com/MrYZhou/outil/ssh"
)

func GetServer(s ServerConfig) *Cli {
	con, _ := Server(s.Host, s.User, s.Password)
	return con
}

func PubProject(model PubInfo) {
	con := GetServer(Host[model.HostId])
	project := Project[model.ProjectId]
	defer con.Client.Close()
	defer con.SftpClient.Close()
	ExecPub(con, project)
}

// ExecPub 泛型函数
func ExecPub(con *Cli, data any) error {
	if webData, ok := data.(WebUpload); ok {
		webHelper.ExecPub(con, webData)
	}

	if javaData, ok := data.(JarUpload); ok {
		javaHelper.ExecPub(con, javaData)
	}
	return nil
}

/*
remoteJarHome  服务器jar文件所在目录

name jar文件的名字
*/
func InitDockerfile(c *Cli, remoteJarHome string, name string) bool {
	dockerFilePath := path.Join(remoteJarHome, "Dockerfile")
	init := c.IsFileExist(dockerFilePath)
	if init == false {
		// 创建dockerfile文件
		ftpFile, _ := c.CreateFile(dockerFilePath)

		version := os.Getenv("jdk")
		port := os.Getenv("port")
		if version == "" {
			version = "8"
		}

		b := []byte("FROM openjdk:" + version + "-slim" + "\n")
		ftpFile.Write(b)
		b = []byte("WORKDIR /java" + "\n")
		ftpFile.Write(b)
		// 因为使用-v映射方式,不需要直接添加进去
		// b = []byte("ADD *.jar /java/app.jar" + "\n")
		// ftpFile.Write(b)
		b = []byte(`ENTRYPOINT ["java","-jar","/java/` + name + `"]` + "\n")
		ftpFile.Write(b)
		b = []byte("EXPOSE " + port)
		ftpFile.Write(b)
		imageName := os.Getenv("imageName")
		Info("正在构建镜像")
		build := "docker build -f " + dockerFilePath + " -t  " + imageName + " " + remoteJarHome
		msg, err := c.Run(build)
		if err != nil {
			Info(err.Error())
		}
		Info(msg)
		Info("构建完成")

	}
	return init
}

/*
init 没有生成过dockerfile文件,init为false
*/
func RunContainer(init bool, c *Cli) {
	Info("运行容器")
	direct := ""
	javaContainerName := os.Getenv("javaContainerName")
	imageName := os.Getenv("imageName")
	remoteJarHome := os.Getenv("remoteJarHome")
	port := os.Getenv("port") + ":" + os.Getenv("port")

	if init == false {
		// 不需要输出,下面两行考虑到容器名可能已经存在,需要先移除
		c.RunQuiet("docker stop " + javaContainerName)
		c.RunQuiet("docker rm " + javaContainerName)
		// 需要映射目录这样restart才有意义
		direct = "docker run -d --name " + javaContainerName + " -p " + port + " -v " + remoteJarHome + ":/java " + imageName
	} else {
		direct = "docker restart " + javaContainerName
	}
	c.Run(direct)
}

// WebHelper 结构体
type WebHelper struct{}

// JavaHelper 结构体
type JavaHelper struct{}

var webHelper WebHelper
var javaHelper JavaHelper

func init() {
	webHelper = WebHelper{}
	javaHelper = JavaHelper{}
}
func (web *WebHelper) ExecPub(con *Cli, data WebUpload) {
	web.packageCode(data)
	web.upload(con, data)
}
func (web *WebHelper) packageCode(data WebUpload) {
	Info("打包web应用")
	Run(data.LocalPath, "npm run build")
}
func (web *WebHelper) upload(con *Cli, data WebUpload) {
	Info("上传web应用")
	con.UploadDir(data.LocalPath, data.RemotePath)
}

func (java *JavaHelper) ExecPub(con *Cli, data JarUpload) {
	java.packageCode(data)

	// 获取jarFilePath的jar文件名
	file, _ := os.Open(data.LocalPath)
	name := file.Name()
	data.JarName = name

	java.upload(con, data)
	// 镜像构建
	init := InitDockerfile(con, data.RemoteHome, name)
	// 运行容器
	RunContainer(init, con)
}
func (java *JavaHelper) packageCode(data JarUpload) {
	Info("打包java应用")
	Run(data.JavaProjectPath, data.PackageCommand)
}
func (java *JavaHelper) upload(con *Cli, data JarUpload) {
	Info("上传java应用")
	fileList := con.SliceUpload(data.RemoteHome, data.LocalPath, 6)
	con.ConcatRemoteFile(fileList, data.RemotePath)
	con.Run("rm -rf " + strings.Join(fileList, " "))

}

package common

import (
	// . "github.com/MrYZhou/outil/command"
	"fmt"
	"os"
	"path"

	. "github.com/MrYZhou/outil/ssh"
)

func GetServer(s ServerConfig) *Cli {
	con, _ := Server(s.Host, s.User, s.Password)
	return con
}


// ExecPub 泛型函数
func ExecPub(pubType string, data any) error {
	if webData, ok := data.(WebUpload); ok {
		webHelper.ExecPub(webData)
	}

	if javaData, ok := data.(JarUpload); ok {
		javaHelper.ExecPub(javaData)
	}
	return nil
	// switch pubType {
	// case "web":
	// 	webHelper.ExecPub(data.(WebUpload))
	// case "java":
	// 	javaHelper.ExecPub(data.(JarUpload))
	// }
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
		fmt.Println("正在构建镜像")
		build := "docker build -f " + dockerFilePath + " -t  " + imageName + " " + remoteJarHome
		msg, err := c.Run(build)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(msg)
		fmt.Println("构建完成")

	}
	return init
}

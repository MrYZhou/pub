package common

// ServerConfig 定义单个服务器的配置结构
type ServerConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Config 定义整个配置结构，使用 map 存储多个 ServerConfig
type ServerConfigMap map[string]ServerConfig

// ProjectConfig 定义单个项目配置的结构
type ProjectConfig struct {
	Type    string `json:"type"`    // 项目类型
	Name    string `json:"name"`    // 项目名称
	Content string `json:"content"` // 项目内容
}
type ProjectConfigMap map[string]ProjectConfig

// 前端接收
type PubInfo struct {
	HostId    string `json:"hostId"`
	ProjectId string `json:"projectId"`
}

// java应用的模型
type JarUpload struct {
	JavaProjectPath string `json:"javaProjectPath"` // java项目根路径
	LocalJarPath    string `json:"localJarPath"`    // 生成的jar文件路径
	RemotePath      string `json:"remotePath"`      // 远程路径
	PubCommand      string `json:"pubCommand"`      // 打包jar命令
	ExecCommand     string `json:"execCommand"`     // 后置发布命令
	ImageName       string `json:"imageName"`       // 生成镜像名
	ContainerName   string `json:"containerName"`   // 部署容器名
	JdkVersion      string `json:"jdkVersion"`      // JDK版本
	Port            string `json:"port"`            // 部署端口
}

// web应用的模型
type WebUpload struct {
	LocalPath  string `json:"localPath"`  // 本地web项目路径
	RemotePath string `json:"remotePath"` // 远程web项目路径
}

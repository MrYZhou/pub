### 项目介绍
一个基本的fiber框架，用于学习的实验性项目

### 安装依赖
代理
```
go env -w GOPROXY=https://goproxy.cn,direct
```
```
go install
```
### 启动
f5调试启动

### 打包
```
go build -tags=prod -ldflags "-w -s" . && upx pub.exe
```

### 关于静态服务器访问路径
无打包情况是当前目录的static目录
打包后默认是用户目录下的pub目录
还可以通过手动设置环境变量更换,换后重启服务
setx /M staticDir "F:\files"

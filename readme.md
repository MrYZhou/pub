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
go build -ldflags "-w -s" . && upx pub.exe
```
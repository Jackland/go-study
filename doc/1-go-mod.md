### go mod 命令
1,GO111MODULE
```text
GO111MODULE 有三个值：off, on和auto（默认值）。

GO111MODULE=off，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。
GO111MODULE=on，go命令行会使用modules，而一点也不会去GOPATH目录下查找。
GO111MODULE=auto，默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：
当前目录在GOPATH/src之外且该目录包含go.mod文件
当前文件在包含go.mod文件的目录下面。

go:
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct //七牛云 direct  表示源地址的版本信息
注： https://goproxy.io 也可替换成阿里云的 https://mirrors.aliyun.com/goproxy或https://goproxy.cn,direct


```

2,go mod init 生成go.mod文件
3,go get url 安装goo 依赖安装包

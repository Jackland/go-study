package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

//创建一个server 接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (this *Server)Handler(conn net.Conn) {
	//开始连接
	fmt.Println("服务链接成功...")
}


//启动一个服务接口
func (this *Server) Start() {
	//socket server
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.listen:", err)
		return
	}
	//close
	defer listener.Close()
	//accept
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen accept err:", err)
			continue
		}
		//do handler
		go this.Handler(conn)
	}

}

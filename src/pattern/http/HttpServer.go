package http

import (
	"base"
	"config"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type HttpServer struct {
	base.Server
}

func NewHttpServer(s base.Server) *HttpServer{
	return &HttpServer{s}
}

func (hs *HttpServer) Name() string{
	return "test"
}

func (hs *HttpServer) Serve(){
	lis,err := net.Listen("tcp","127.0.0.1:8090")
	if err != nil {
		log.Fatal("启动监听失败")
	}
	for{
		conn ,err := lis.Accept()
		if err != nil {
			log.Fatal("接受连接失败")
		}

		go hs.dealConn(conn)

	}
}

/**
处理tcp连接
 */
func (hs *HttpServer) dealConn(conn net.Conn){
	for {
		var data = make([]byte,1024)
		num,err := conn.Read(data)
		if err != nil {
			log.Fatal("读取数据失败！\n")
		}

		fmt.Println("有新数据来了")
		if num == 0{
			break
		}

		var request config.Request
		fmt.Printf("此处共接受到%d字节的数据,数据为：%s\n",num,string(data))

		err = json.Unmarshal(data[:num],&request)
		if err != nil {
			log.Fatal(fmt.Sprintf("请求解析错误：%s\n"),err.Error())
		}

		request.Conn = conn
		fmt.Printf("reques数据为:%s",request)
		hs.Handle(request.Method,request.Params,request.Conn)
	}
}
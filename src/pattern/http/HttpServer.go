package http

import (
	"base"
	"config"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"
)

type HttpServer struct {
	base.Server
}

func NewHttpServer(s base.Server) *HttpServer {
	return &HttpServer{s}
}

func (hs *HttpServer) Name() string {
	return "test"
}

func (hs *HttpServer) Serve() {
	cert ,err := myLoadX509KeyPair("/Users/zhan/cert3/server.crt","/Users/zhan/cert3/server.key")
	lis, err := net.Listen("tcp", "www.zhanhong.com:8090")

	config := &tls.Config{
		Rand:rand.Reader,
		Time:time.Now,
	}

	config.Certificates = make([]tls.Certificate,1)
	config.Certificates[0] = cert
	lis = tls.NewListener(lis,config)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal("接受连接失败")
		}

		go hs.dealConn(conn)

	}
}


func myLoadX509KeyPair(certPath ,keyPath string) (cert tls.Certificate,err error){
	certBlock := parseBlock(certPath)
	keyBlock := parseBlock(keyPath)

	cert.Certificate = append(cert.Certificate,certBlock.Bytes)
	key,err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		panic(err)
	}
	cert.PrivateKey = key

	return
}


func parseBlock(path string) (pem.Block){
	origin,_ := ioutil.ReadFile(path)
	pem ,_:= pem.Decode(origin)

	return *pem
}


/**
处理tcp连接
*/
func (hs *HttpServer) dealConn(conn net.Conn) {
	for {
		var data = make([]byte, 1024)
		num, err := conn.Read(data)
		if err != nil {
			log.Fatal("读取数据失败！\n")
		}

		fmt.Println("有新数据来了")
		if num == 0 {
			break
		}

		var request config.Request
		fmt.Printf("此处共接受到%d字节的数据,数据为：%s\n", num, string(data))

		err = json.Unmarshal(data[:num], &request)
		if err != nil {
			log.Fatal(fmt.Sprintf("请求解析错误：%s\n"), err.Error())
		}

		request.Conn = conn
		fmt.Printf("reques数据为:%s", request)
		hs.Handle(request.Method, request.Params, request.Conn)
	}
}

package main

import (
	"bufio"
	"config"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func initConn() net.Conn {
	conn, err := tls.Dial("tcp", "www.zhanhong.com:8090",&tls.Config{InsecureSkipVerify:true})
	if err != nil {
		log.Fatal("连接失败 :",err)
	}

	state := conn.ConnectionState()
	fmt.Println("client:handshake :",(state.HandshakeComplete))
	fmt.Println("client:mutual :",state.NegotiatedProtocolIsMutual)
	return conn
}

func parse(params string) []byte {
	var request config.Request

	token := strings.Split(params, " ")
	method := token[0]
	args := token[1:]

	request.Method = method
	request.Params = args

	rtn, err := json.Marshal(request)
	if err != nil {
		log.Fatal("解析请求失败！\n")
	}
	return rtn
}

func readMsg(conn net.Conn) {
	//var resp []byte
	for {
		var resp = make([]byte, 1024)
		num, err := conn.Read(resp)
		if err != nil {
			log.Fatal("读取数据失败™")
		}

		fmt.Printf("此处读取%字节数据\n", num)
		msg := string(resp)
		fmt.Printf("读取到数据:%s\n", msg)
	}
}

func main() {
	fmt.Println("开始连接\n")
	conn := initConn()
	fmt.Println("连接已建立")
	r := bufio.NewReader(os.Stdin)
	go readMsg(conn)
	for {
		msg, _, _ := r.ReadLine()
		request := parse(string(msg))
		fmt.Printf("解析后的数据为：%s\n", string(request))
		num, err := conn.Write(request)
		if err != nil {
			log.Fatal(fmt.Sprintf("发送数据错误：%s\n", err.Error()))
		}

		fmt.Printf("插入%d字节数据\n", num)
	}
}

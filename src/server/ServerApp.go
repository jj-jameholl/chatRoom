package main

import (
	"base"
	"fmt"
	"pattern/http"
)

func startService(server *base.CenterServer){
	httpServer := http.NewHttpServer(server)
	server.ServerList["http"] = httpServer
	httpServer.Serve()
}

func main() {
	centerServer := base.NewCenterServer()
	fmt.Println("启动中")
	startService(centerServer)
}

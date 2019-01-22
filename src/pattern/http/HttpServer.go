package http

import (
	"base"
	"io"
	"log"
	"net/http"
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

func (hs *HttpServer) parse(w http.ResponseWriter,r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(r.Form["cmdId"]) == 0{
		io.WriteString(w,"没有cmdId")
		return
	}
	cmdId := r.Form["cmdId"][0]
	rtn := hs.Handle(cmdId,r.Form)
	io.WriteString(w,rtn.Data)
}

func (hs *HttpServer) Serve(){
	http.HandleFunc("/" ,hs.parse)
	http.ListenAndServe(":8010",nil)
}



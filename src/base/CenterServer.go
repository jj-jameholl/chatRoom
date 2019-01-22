package base

import (
	"config"
	"element/player"
	"element/room"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
)

type CenterServer struct {
	ServerList map[string] Server
	Players    []player.Player
	Rooms      []room.Room
	mutex      sync.RWMutex
}

func (cs *CenterServer) Name() string{
	return "centerServer"
}


func (cs *CenterServer) Handle(method string ,params map[string][]string) *config.Response {
	handlers := getComandHanders()
	var status int
	var data string
	if ok := handlers[method];ok{
		switch method {
		case "send":
			status,data = cs.Send(params)
		case "login":
			status,data = cs.Login(params)
		case "listPlayer":
			status,data = cs.ListPlayer()
		}
	}else{
		fmt.Println("没有找到方法")
	}

	response := &config.Response{status,"success",data}
	return response
}


func (cs *CenterServer) Send(params map[string][]string) (code int,data string) {
	code = 200
	if ok := checkParams(params);ok{
		toId := params["To"][0]
		fromId := params["From"][0]
		content := params["Data"][0]
		if toId == "-1"{
			SendToAll(fromId,content)
		}else {
			SendToOne(fromId,toId,content)
		}
	}else{
		code = 500
		data = "参数不符合"
		return
	}

	return
}


func (cs *CenterServer) Login(params map[string][]string) (code int ,data string){
	code = 200
	if ok := checkParams(params);ok{
		player := new(player.Player)
		player.Name = params["name"][0]
		player.Id,_ = strconv.Atoi(params["id"][0])
		if ok := cs.addPlayer(*player);!ok{
			code = 500
		}
	}else{
		code = 500
	}


	data = "登录成功"
	return
}

func (cs *CenterServer) ListPlayer() (code int ,data string){
	code = 200
	players,_ := json.Marshal(cs.Players)
	data = string(players)
	return
}


func (cs *CenterServer) addPlayer(player player.Player) bool {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	cs.Players = append(cs.Players,player)
	return true
}


func SendToAll(fromId ,content string){

}


func SendToOne(toId ,fromId,content string) {

}



func checkParams(params map[string][]string) bool{
	return true
}


func NewCenterServer() *CenterServer{
	ServerList := make(map[string]Server)
	Players := make([]player.Player,0)
	Rooms := make([]room.Room,0)
	cs := &CenterServer{ServerList:ServerList,Players:Players,Rooms:Rooms}

	return cs
}


func getComandHanders() map[string]bool{
	return map[string]bool{
		"login" : true,
		"send" : true,
		"listPlayer" : true,
	}
}
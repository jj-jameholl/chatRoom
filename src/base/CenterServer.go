package base

import (
	"config"
	"element/player"
	"element/room"
	"encoding/json"
	"fmt"
	"net"
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


func (cs *CenterServer) Handle(method string ,params []string,conn net.Conn) {
	handlers := getComandHanders()
	var status int
	var data string
	if ok := handlers[method];ok{
		switch method {
		case "send":
			cs.Send(params)
			return
		case "login":
			status,data = cs.Login(params,conn)
		case "listPlayer":
			status,data = cs.ListPlayer()
		}
	}else{
		fmt.Println("没有找到方法")
	}

	response := &config.Response{status,"success",data}

	rtn,_ := json.Marshal(response)
	conn.Write(rtn)
}


func (cs *CenterServer) Send(params []string) {
	if ok := checkParams(params);ok{
		toId := params[0]
		fromId := params[1]
		content := params[2]
		if toId == "-1"{
			cs.SendToAll(fromId,content)
		}else {
			cs.SendToOne(fromId,toId,content)
		}
	}else{
		return
	}

	return
}


func (cs *CenterServer) Login(params []string,conn net.Conn) (code int ,data string){
	code = 200

	if ok := checkParams(params);ok{

		player := new(player.Player)
		player.Name = params[0]
		player.Id,_ = strconv.Atoi(params[1])
		player.Conn = conn
		fmt.Println(player)
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


func (cs *CenterServer) SendToAll(fromId ,content string){
	var msg = make([]byte,1024)
	playerList := cs.Players
	msg = []byte(content)
	fromIdInt ,_ := strconv.Atoi(fromId)
	for _,player := range playerList{
		if player.Id != fromIdInt {
			player.Conn.Write(msg)
		}
	}
}


func (cs *CenterServer) SendToOne(toId ,fromId,content string) {
	var msg = make([]byte ,1024)
	playerList := cs.Players
	msg = []byte(content)
	toIdInt ,_ := strconv.Atoi(toId)
	for _,player := range playerList{
		if player.Id == toIdInt{
			player.Conn.Write(msg)
		}
	}
}



func checkParams(params []string) bool{
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
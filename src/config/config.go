package config

type Message struct {
	From int "from who"
	To   int "to who"
	Content  string  "content"
}


type Request struct {
	Method string "method"
	Params string "params"
}


type Response struct {
	Code    int "code"
	Message string "message"
	Data    string "data"
}
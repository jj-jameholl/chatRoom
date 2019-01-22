package base

import "config"

type Server interface {
	Handle(method string ,params map[string][]string) *config.Response
	Name() string
}

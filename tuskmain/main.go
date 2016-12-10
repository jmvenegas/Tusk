package main

import (
	"os"

	"github.com/jmvenegas/tusk/model"
)

func main() {
	configPathArg := os.Args[1]
	tc := model.NewTuskConfig(configPathArg)
	webserver := model.NewWebServer(tc)
	webserver.RegisterPage("/", webserver.Welcome)
	webserver.RegisterPage("/query", webserver.Query)
	webserver.RegisterPage("/upload", webserver.Upload)
	webserver.Listen()
}

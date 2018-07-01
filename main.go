package main

import (
	"os"

	"github.com/karl-jm-huang/golang-CI/server"

	flag "github.com/spf13/pflag"
)

// PORT
const (
	PORT string = "7070" //PORT
)

func main() {
	//service.Init()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	servertt := server.NewServer()
	servertt.Run(":" + port)
}

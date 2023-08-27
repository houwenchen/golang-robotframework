package main

import "houwenchen/golang-robotframework/pkg/server"

func main() {
	server := server.NewServer()
	server.Run()
}

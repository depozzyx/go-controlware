package main

import (
	"controlware/client"
	"controlware/server"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usageError()
	}

	script := args[0]
	if script == "client" {
		go client.Run("http://138.68.102.39:4009")
		client.Run("http://192.168.1.106:4009")

	} else if script == "server" {
		server.Run("0.0.0.0:4009")
	} else {
		usageError()
	}
}

func usageError() {
	log.Println("Usage: go run . <client/server>")
	os.Exit(1)
}

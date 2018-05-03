package main

import (
	"github.com/Amniversary/wedding-plugin-game/server"
	"github.com/Amniversary/wedding-plugin-game/config"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	server.NewServer(config.NewConfig()).Run()
}




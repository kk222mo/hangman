package main

import (
	"flag"
	"fmt"
	"github.com/kk222mo/hangman/database"
	"github.com/kk222mo/hangman/server"
)

func main() {
	modePtr := flag.String("mode", "serve", "serve/initdb/getplayer")
	flag.Parse()
	if *modePtr == "serve" {
		server.Serve(8082)
	} else if *modePtr == "initdb" {
		database.InitDatabase()
	} else {
		res, err := database.ReadPlayerInfo("kek1")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}

package main

import (
	"log"

	"github.com/61lf0yl3/div-01/forum/internal/app/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := server.Start(); err != nil {
		log.Println(err)
	}
}

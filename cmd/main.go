package main

import (
	"log"

	l0 "github.com/joinusordie/Wildberries_L0"
)

func main() {
	srv := new(l0.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

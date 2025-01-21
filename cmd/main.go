package main

import (
	"log"
	"todo"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running hhtp server: %s", err.Error())
	}
}

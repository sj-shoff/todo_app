package main

import (
	"log"
	"todo"
	"todo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running hhtp server: %s", err.Error())
	}
}

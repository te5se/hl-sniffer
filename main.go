package main

import (
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	userRepo := NewUserRepo()
	tgService := NewTGService(userRepo)

	go NewSniffer(tgService).Listen()
	go tgService.Run()

	time.Sleep(time.Hour * 100)
}

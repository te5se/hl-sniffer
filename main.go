package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
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

func play() error {
	f, err := os.Open("beep.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()

	fmt.Printf("Length: %d[bytes]\n", d.Length())
	for {
		time.Sleep(time.Second)
		if !p.IsPlaying() {
			break
		}
	}

	return nil
}

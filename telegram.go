package main

import (
	"context"
	"log"
	"os"

	router "github.com/te5se/tg-mux"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgService struct {
	userRepo *UserRepo
	bot      *tgbotapi.BotAPI
	router   *router.TGRouter
}

func NewTGService(userRepo *UserRepo) *TgService {
	tgKey, exists := os.LookupEnv("HLSNIFFER_TG_KEY")
	if !exists {
		log.Fatal("token not found in env HLSNIFFER_TG_KEY")
	}

	bot, err := tgbotapi.NewBotAPI(tgKey)
	if err != nil {
		log.Fatal(err)
	}

	router, err := router.NewTGRouter(bot, func(ctx *router.TGContext) (string, error) {
		user, err := userRepo.GetUser()
		if err != nil {
			return "failed to get user", err
		}
		if user == nil {
			return "", nil
		}

		return user.State, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	startHandler := NewStartHandler(userRepo)
	startHandler.Register(router)

	return &TgService{
		userRepo: userRepo,
		bot:      bot,
		router:   router,
	}
}

func (s *TgService) Notify(text string) error {
	user, err := s.userRepo.GetUser()
	if err != nil {
		return err
	}
	_, err = s.bot.Send(tgbotapi.NewMessage(user.ChatID, text))

	return err
}

func (s *TgService) Run() error {
	s.router.Run(context.Background())
	return nil
}

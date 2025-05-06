package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgmux "github.com/te5se/tg-mux"
)

type StartHandler struct {
	UserRepo *UserRepo
}

func NewStartHandler(UserRepo *UserRepo) *StartHandler {
	return &StartHandler{
		UserRepo: UserRepo,
	}
}

func (h *StartHandler) Register(router *tgmux.TGRouter) {
	router.RegisterCommand("start", h.HandleStartCommand)
}

func (h *StartHandler) HandleStartCommand(ctx *tgmux.TGContext) (tgbotapi.MessageConfig, error) {
	h.UserRepo.SetUser(User{ChatID: ctx.Message.Chat.ID})

	return tgbotapi.NewMessage(ctx.Message.Chat.ID, "Got it"), nil
}

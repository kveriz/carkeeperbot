package app

import (
	"context"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kveriz/carkeeperbot/internal/bot/config"
	"github.com/kveriz/carkeeperbot/internal/bot/models"
)

const (
	Fuel = iota + 1
	Washing
	Service
	Duties
	Equippment
	Repairing
)

type TgBot struct {
	config   config.Config
	tgbotapi *tgbotapi.BotAPI
	RI       models.RepositoryInterface
	cancel   context.CancelFunc
}

func NewTgBot(config config.Config, RI models.RepositoryInterface) *TgBot {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Println(err)
	}
	return &TgBot{config: config, tgbotapi: bot, RI: RI}
}

func (t *TgBot) Serve(ctx context.Context) {
	if t.cancel != nil {
		fmt.Println("Bot is already running")
		return
	}
	localCtx, cancel := context.WithCancel(ctx)
	t.cancel = cancel

	log.Printf("Starting bot: %v. Version: %v", t.config.Name, t.config.Version)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.tgbotapi.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			args := update.Message.CommandArguments()
			switch update.Message.Command() {
			case "start":
				msg.Text = fmt.Sprintf(models.Messages["start"][t.config.Lang], t.config.Name)
			case "add":
				result, err := t.add(localCtx, args, msg, update)
				if err != nil {
					log.Println(err)
					continue
				}
				msg.ReplyToMessageID = update.Message.MessageID
				msg.Text = result
			case "list":
				result, err := t.list(localCtx, args, msg, update)
				if err != nil {
					log.Println(err)
					continue
				}
				msg.ReplyToMessageID = update.Message.MessageID
				msg.Text = result
			case "categories":
				msg.Text = `TODO return list of possible categories`
			case "help":
				msg.Text = models.Messages["help"][t.config.Lang]
			default:
				msg.Text = models.Messages["default"][t.config.Lang]
			}
			t.tgbotapi.Send(msg)
		}
	}
}

func (t *TgBot) Stop() {
	t.cancel()
	t.cancel = nil
	fmt.Println("Bot is stopped")
}

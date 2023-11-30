package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	bot "github.com/kveriz/carkeeperbot/internal/bot/app"
	cfg "github.com/kveriz/carkeeperbot/internal/bot/config"
	"github.com/kveriz/carkeeperbot/internal/bot/models/repository/db"
)

func main() {
	var conf string
	flag.StringVar(&conf, "conf", "", "Path to config file v2")
	flag.Parse()

	config := cfg.New(conf)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	storage := db.New(*config)
	defer storage.Close()

	tgBot := bot.NewTgBot(*config, storage)

	go tgBot.Serve(ctx)

	<-ctx.Done()
	tgBot.Stop()
}

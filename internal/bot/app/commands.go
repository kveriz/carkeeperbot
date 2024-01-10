package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kveriz/carkeeperbot/internal/bot/models"
)

func (t *TgBot) add(ctx context.Context, args string, m tgbotapi.MessageConfig, u tgbotapi.Update) (string, error) {
	p, err := t.parseAddInput(args)
	if err != nil {
		m.Text = err.Error()
		t.tgbotapi.Send(m)
		return m.Text, err
	}
	p.UserID = u.Message.From.UserName

	_, err = t.RI.Add(ctx, p)
	if err != nil {
		m.Text = fmt.Sprintf("%v", err)
	}

	m.Text = models.Messages["result"][t.config.Lang]

	return m.Text, err
}

func (t *TgBot) list(ctx context.Context, args string, m tgbotapi.MessageConfig, u tgbotapi.Update) (string, error) {
	from, to, err := t.parseListInput(args)
	if err != nil {
		m.Text = err.Error()
		t.tgbotapi.Send(m)
		return m.Text, err
	}

	rows, err := t.RI.List(ctx, u.Message.From.UserName, from, to)
	if err == nil {
		log.Println(err)
	}

	var sb strings.Builder
	var total float64

	for _, v := range rows {
		if v == (models.Stat{}) {
			continue
		}
		total += v.Amount
		sb.WriteString(v.Category + "\t\t\t" + fmt.Sprintf("%v", v.Amount) + "\n")
	}

	sb.WriteString(fmt.Sprintf(models.Messages["stats"][t.config.Lang], total))
	m.Text = sb.String()

	return m.Text, err
}

func (t *TgBot) parseAddInput(s string) (p models.Payment, err error) {
	sl := strings.Split(s, " ")

	if len(sl) != 3 {
		err = errors.New(models.Messages["args_num_error"][t.config.Lang])
		return p, err
	}

	p.ActivityType = sl[0]
	p.Amount, err = strconv.ParseFloat(sl[1], 64)
	if err != nil {
		err = errors.New(models.Messages["float_parse_error"][t.config.Lang])
		return p, err
	}
	p.Date, err = time.Parse(time.DateOnly, sl[2])
	if err != nil {
		err = errors.New(models.Messages["date_parse_error"][t.config.Lang])
		return p, err
	}

	return p, err
}

func (t *TgBot) parseListInput(s string) (from, to time.Time, err error) {
	sl := strings.Split(s, " ")

	if len(sl) != 2 {
		err = errors.New(models.Messages["args_num_error"][t.config.Lang])
		return from, to, err
	}

	from, err = time.Parse(time.DateOnly, sl[0])
	if err != nil {
		err = errors.New(models.Messages["date_parse_error"][t.config.Lang])
		return from, to, err
	}

	to, err = time.Parse(time.DateOnly, sl[1])
	if err != nil {
		err = errors.New(models.Messages["date_parse_error"][t.config.Lang])
		return from, to, err
	}

	if !inTimeSpan(from, to) {
		err = errors.New(models.Messages["date_comparasion"][t.config.Lang])
		return from, to, err
	}
	return from, to, err
}

func inTimeSpan(start, end time.Time) bool {
	return end.After(start) && start.Before(end)
}

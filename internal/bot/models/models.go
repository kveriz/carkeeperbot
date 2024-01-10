package models

import (
	"context"
	"time"
)

var (
	Messages = map[string]map[string]string{
		"start": {
			"RU": `Приветствую тебя в "%v"
Чтобы начать пользоваться отправляй команды /add и /list
Иcпользуй команду /help для получения справки по каждой из них
			`,
			"EN": `Welcome to "%v"
To start using bot call and send /add and /list commands
Use /help to get more info about them
			`,
		},
		"help": {
			"RU": `/help показать это сообщение
/add добавить новую трату на тачиллу в формате НАЧТО СУММА КОГДА. Пример "соляра 1666.66 2023-14-20"
/list получить стастику в диапазоне ОТ и ДО даты. Пример "2023-14-01 2023-14-30"
			`,
			"EN": `/help show this message
/add add new cost in format CATEGORY AMOUNT DATE, e.g. "fuel 1666.66 2023-14-20"
/list get stats between FROM and TO dates in 2000-01-01 format, e.g. "2023-14-01 2023-14-30"
			`,
		},
		"result": {
			"RU": "Запись добавлена",
			"EN": "Success",
		},
		"stats": {
			"RU": "Всего потрачено: %v",
			"EN": "Total spent this period: %v",
		},
		"args_num_error": {
			"RU": "неверное количество аргументов",
			"EN": "wrong num of args",
		},
		"date_parse_error": {
			"RU": "неверный форма даты",
			"EN": "wrong date format",
		},
		"date_comparasion": {
			"RU": "неверный порядок дат",
			"EN": "wrong date sequence",
		},
		"float_parse_error": {
			"RU": "неверный формат суммы",
			"EN": "wrong type of cost",
		},
		"default": {
			"RU": "Неизвестная команда. Используй /help для получения списка доступных",
			"EN": "I don't know that command. Use /help to get list of ones",
		},
	}
)

type Payment struct {
	UserID       string
	ActivityType string
	Amount       float64
	Date         time.Time
}

type Stat struct {
	Category string
	Amount   float64
}

type RepositoryInterface interface {
	Add(ctx context.Context, p Payment) (Payment, error)
	List(ctx context.Context, uid string, from, to time.Time) ([]Stat, error)
}

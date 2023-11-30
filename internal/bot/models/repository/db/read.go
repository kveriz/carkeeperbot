package db

import (
	"context"
	"log"
	"time"

	"github.com/kveriz/carkeeperbot/internal/bot/models"
)

func (db *DB) List(ctx context.Context, uid string, from, to time.Time) ([]models.Stat, error) {
	query := "SELECT act_type, sum(sum_total) FROM (SELECT user_id, act_type, sum(act_amount) AS sum_total, act_date FROM carkeeperbot.activities WHERE user_id = $1 AND act_date BETWEEN $2::timestamp AND $3::timestamp GROUP BY user_id, act_type, act_date) AS tmp GROUP BY act_type;"
	res, err := db.db.QueryContext(ctx, query, uid, from.Format(time.DateOnly), to.Format(time.DateOnly))
	if err != nil {
		log.Println(err)
	}
	defer res.Close()

	columns, err := res.Columns()
	if err != nil {
		log.Println(err)
	}

	result := make([]models.Stat, len(columns))

	for res.Next() {
		var row models.Stat

		if err = res.Scan(&row.Category, &row.Amount); err != nil {
			panic(err)
		}
		if row == (models.Stat{}) {
			continue
		}
		result = append(result, row)
	}

	err = res.Err()
	if err != nil {
		log.Println(err)
	}

	return result, err

}

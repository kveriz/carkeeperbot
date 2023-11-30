package db

import (
	"context"

	"github.com/kveriz/carkeeperbot/internal/bot/models"
	"github.com/pkg/errors"
)

func (db *DB) Add(ctx context.Context, payment models.Payment) (models.Payment, error) {
	query := "INSERT INTO carkeeperbot.activities (user_id, act_type, act_amount, act_date) VALUES ($1, $2, $3, $4) RETURNING id;"
	var id int64

	if err := db.db.QueryRowContext(ctx, query, payment.UserID, payment.ActivityType, payment.Amount, payment.Date).Scan(&id); err != nil {
		return payment, errors.WithMessage(err, "insert metric failed")
	}
	if id == 0 {
		return payment, errors.Errorf("insert metric failed, id=0")
	}

	return payment, nil
}

package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/kveriz/carkeeperbot/internal/bot/config"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// TODO make it usable not only postgres and sqlite3
func parseConfig(c config.Config) (string, string) {
	if c.SQL.DBType == "sqlite3" {
		return c.SQL.DBType, c.SQL.SqlitePath
	}

	tmpl := "%s://%s:%s@%s/%s?sslmode=%s"

	cs := fmt.Sprintf(tmpl, c.SQL.DBType, c.SQL.DBUser, c.SQL.DBPassword, c.SQL.DBHost, c.SQL.DB, c.SQL.Mode)

	return c.SQL.DBType, cs
}

type Config struct {
	DBType     string
	DBUser     string
	DBPassword string
	DBHost     string
	DB         string
	Mode       string
	SqlitePath string
}

type DB struct {
	db     *sql.DB
	config Config
}

func New(c config.Config) *DB {
	sql, err := sql.Open(parseConfig(c))
	if err != nil {
		log.Println(err)
	}

	config := Config(c.SQL)

	return &DB{db: sql, config: config}
}

func (db *DB) Close() {
	db.db.Close()
}

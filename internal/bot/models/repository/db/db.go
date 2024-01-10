package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/kveriz/carkeeperbot/internal/bot/config"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// TODO make it usable not only postgres and sqlite3
func parseConfig(c Config) (string, string) {
	var driver, connectionString, tmpl string
	driver = c.DBDriver
	switch driver {
	case "sqlite3":
		connectionString = c.SqlitePath
		return driver, connectionString
	case "postgres":
		tmpl = "%s://%s:%s@%s/%s?sslmode=%s"
	case "mysql":
		tmpl = "%s://%s:%s@%s/%s?tls=%s"
	default:
		return driver, connectionString
	}

	connectionString = fmt.Sprintf(tmpl, c.DBDriver, c.DBUser, c.DBPassword, c.DBHost, c.DB, c.Mode)

	return driver, connectionString
}

type Config struct {
	DBDriver   string
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
	config := Config(c.SQL)

	sql, err := sql.Open(parseConfig(config))
	if err != nil {
		log.Println(err)
	}

	return &DB{db: sql, config: config}
}

func (db *DB) Close() {
	db.db.Close()
}

func (db *DB) DoMigrate(path string) {
	_, connectionString := parseConfig(db.config)
	m, err := migrate.New("file://"+path, connectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

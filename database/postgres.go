package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rizalarfiyan/be-petang/config"
)

func Postgres() *sqlx.DB {
	config := config.Get()

	connection := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)
	db, err := sqlx.Open("postgres", connection)
	if err != nil {
		log.Fatal("Error connection:", err.Error())
	}

	db.SetConnMaxIdleTime(config.DB.ConnectionIdle)
	db.SetConnMaxLifetime(config.DB.ConnectionLifetime)
	db.SetMaxIdleConns(config.DB.MaxIdle)
	db.SetMaxOpenConns(config.DB.MaxOpen)

	return db
}

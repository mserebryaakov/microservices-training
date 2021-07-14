package handler

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const infoDB string = "user=serebryakov password=123 dbname=productServiceDB sslmode=disable"

type Products struct {
	l  *log.Logger
	db *sql.DB
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l: l}
}

type KeyProduct struct{}

func (p *Products) Open() error {
	var err error
	p.db, err = sql.Open("postgres", infoDB)
	if err != nil {
		return err
	}
	return nil
}

func (p *Products) Close() {
	p.l.Println("[SUCCESS] database close connection ")
	p.db.Close()
}

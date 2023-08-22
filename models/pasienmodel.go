package models

import (
	"database/sql"
	"github.com/jeypc/go-crud/config"
)

type PasienModel struct {
	conn *sql.DB
}

func NewPasienModel() *PasienModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	return &PasienModel {
		conn: conn,
	}
}
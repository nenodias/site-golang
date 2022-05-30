package db

import (
	"database/sql"
	"regexp"

	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() (*sql.DB, error) {
	regex := func(re, s string) (bool, error) {
		return regexp.MatchString(re, s)
	}
	sql.Register("sqlite3_extended",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				return conn.RegisterFunc("regexp", regex, true)
			},
		})
	return sql.Open("sqlite3_extended", "./banco.db")
}

func Create(conn *sql.DB) error {
	_, err := conn.Exec(`CREATE TABLE produto(
			id int64 NOT NULL,
			nome varchar(100) NOT NULL,
			descricao varchar(255) NOT NULL,
			valor float NOT NULL,
			quantidade NOT NULL
		)`, nil)
	return err
}

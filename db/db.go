package db

import (
	"database/sql"
	"fmt"
	"regexp"

	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nenodias/site-golang/model"
)

func Init() {
	regex := func(re, s string) (bool, error) {
		return regexp.MatchString(re, s)
	}
	sql.Register("sqlite3_extended",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				return conn.RegisterFunc("regexp", regex, true)
			},
		})
}

func GetConnection() (*sql.DB, error) {
	return sql.Open("sqlite3_extended", "./banco.db")
}

func Create(conn *sql.DB) error {
	_, err := conn.Exec(`CREATE TABLE produto(
			id int64 NOT NULL PRIMARY KEY,
			nome varchar(100) NOT NULL,
			descricao varchar(255) NOT NULL,
			preco float NOT NULL,
			quantidade NOT NULL
		)`, nil)
	return err
}

func Insert(conn *sql.DB, produto *model.Produto) error {
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`INSERT INTO produto(
			id,
			nome,
			descricao,
			preco,
			quantidade
		) VALUES(? , ?, ?, ?, ?) `, produto.Id, produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func Update(conn *sql.DB, produto *model.Produto) error {
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`UPDATE produto SET
			nome = ?,
			descricao = ?,
			preco = ?,
			quantidade = ?
		WHERE id = ?`, produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.Id)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func SelectById(conn *sql.DB, id int) (*model.Produto, error) {
	produto := model.Produto{}
	res := conn.QueryRow("SELECT * FROM produto WHERE id = ? ", id)
	err := res.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
	if err != nil {
		return nil, err
	}
	return &produto, nil
}

func SelectAll(conn *sql.DB) []model.Produto {
	retorno := []model.Produto{}
	res, err := conn.Query("SELECT * FROM produto", nil)
	if err == nil {
		for res.Next() {
			produto := model.Produto{}
			err := res.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
			if err != nil {
				fmt.Print(err)
				break
			}
			retorno = append(retorno, produto)
		}
	}
	return retorno
}

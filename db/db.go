package db

import (
	"database/sql"
	"log"
	"regexp"

	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nenodias/site-golang/models"
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
			id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			nome varchar(100) NOT NULL,
			descricao varchar(255) NOT NULL,
			preco float NOT NULL,
			quantidade NOT NULL
		)`, nil)
	return err
}

func Insert(conn *sql.DB, produto *models.Produto) error {
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	statement, err := tx.Prepare(`INSERT INTO produto(
		nome,
		descricao,
		preco,
		quantidade
	) VALUES(?, ?, ?, ?) `)

	if err != nil {
		return err
	}

	res, err := statement.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)

	if err != nil {
		return err
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return err
		}
		produto.Id = id
	}
	return tx.Commit()
}

func Update(conn *sql.DB, produto *models.Produto) error {
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

func Delete(conn *sql.DB, id int) error {
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	statement, err := tx.Prepare(`DELETE FROM produto WHERE id = ?`)
	if err != nil {
		return err
	} else {
		_, err = statement.Exec(id)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func SelectById(conn *sql.DB, id int) (*models.Produto, error) {
	produto := models.Produto{}
	res := conn.QueryRow("SELECT * FROM produto WHERE id = ? ", id)
	err := res.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
	if err != nil {
		return nil, err
	}
	return &produto, nil
}

func SelectAll(conn *sql.DB) []models.Produto {
	retorno := []models.Produto{}
	res, err := conn.Query("SELECT * FROM produto", nil)
	if err == nil {
		for res.Next() {
			produto := models.Produto{}
			err := res.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
			if err != nil {
				log.Print(err)
				break
			}
			retorno = append(retorno, produto)
		}
	}
	return retorno
}

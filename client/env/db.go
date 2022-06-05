package env

import (
	"context"
	"log"
	"transactions/model"

	"github.com/go-pg/pg/v10"
)

func dbConnect(connectionString string) *pg.DB {
	opt, err := pg.ParseURL(connectionString)
	if err != nil {
		log.Fatalf("failed ParseURL: %v", err)
	}

	db := pg.Connect(opt)

	return db
}

func (e Env) DBPing() error {
	return e.db.Ping(context.Background())
}

func (e *Env) GetUserByID(id int) (*model.User, error) {
	var row model.User
	err := e.db.Model(&row).Where("id = ?", id).Limit(1).Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}

	return &row, err
}

func (e *Env) GetAllUsers() ([]model.User, error) {
	rows := []model.User{}
	err := e.db.Model(&rows).Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}

	return rows, err
}

func (e *Env) CreateUser(data *model.User) error {
	_, err := e.db.Model(data).Insert()

	return err
}

func (e *Env) CreateTransaction(data *model.Transaction) error {
	_, err := e.db.Model(data).Insert()

	return err
}

func (e *Env) UpdateTransaction(data *model.Transaction) error {
	_, err := e.db.Model(data).WherePK().Update()

	return err
}

func (e *Env) UpdateUser(data *model.User) error {
	_, err := e.db.Model(data).WherePK().Update()

	return err
}

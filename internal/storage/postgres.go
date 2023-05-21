package storage

import "database/sql"

const (
	getAllAdmins = `select "ID", "name", "tg" from "admin"`
	postgresName = "postgres"
)

type (
	Postgres struct {
		handle *sql.DB
	}
)

func NewPostgres(args string) (postgres Postgres, fail error) {
	postgres.handle, fail = sql.Open(postgresName, args)
	return
}

func (p Postgres) LoadAdmins() (items []AdminModel, fail error) {
	rows, fail := p.handle.Query(getAllAdmins)
	if fail != nil {
		return
	}
	for rows.Next() {
		var model AdminModel
		if fail = rows.Scan(&model.Id, &model.Name, &model.Tg); fail != nil {
			return
		}
		items = append(items, model)
	}
	return
}

func (Postgres) LoadChats() ([]int64, error) {
	return nil, nil
}

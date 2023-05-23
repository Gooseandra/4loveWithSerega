package storage

import "database/sql"

const (
	getAllAdmins = `select "ID", "name", "tg" from "user"`
	postgresName = "postgres"
	setChatByTg  = `insert into "` + ChatModelTable + `"("` + ChatNameModelField + `", "` +
		ChatTgModelField + `") values($1,$2)` + ` on conflict ( "` + ChatTgModelField +
		`") do update set "` + ChatNameModelField + `"=$1 returning "` + ChatIDModelField + `"`
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

func (p Postgres) LoadAdmins() (items []ChatModel, fail error) {
	rows, fail := p.handle.Query(getAllAdmins)
	if fail != nil {
		return
	}
	for rows.Next() {
		var model ChatModel
		if fail = rows.Scan(&model.ID, &model.Name, &model.Tg); fail != nil {
			return
		}
		items = append(items, model)
	}
	return
}

func (Postgres) LoadChats() ([]int64, error) {
	return nil, nil
}

func (p Postgres) UpdateChatByTg(tg int64, name string) (id ChatIdModel, fail error) {
	fail = p.handle.QueryRow(setChatByTg, name, tg).Scan(&id)
	return
}

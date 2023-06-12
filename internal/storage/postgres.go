package storage

import (
	"database/sql"
	"log"
)

const (
	getAllAdmins   = `select "ID", "name", "tg" from` + UserModelTable
	postgresName   = "postgres"
	upsertChatByTg = `insert into "` + ChatModelTable + `"("` + ChatNameModelField + `","` + ChatTgModelField + `","` +
		ChatModeratedModelField + `")values($1,$2,0)on conflict("` + ChatTgModelField + `")do update set"` +
		ChatNameModelField + `"=$1 returning"` + ChatIDModelField + `","` + ChatModeratedModelField + `"`
	upsertUserByTg = `insert into"` + ChatModelTable + `"("` + ChatNameModelField + `","` + ChatTgModelField + `","` +
		ChatModeratedModelField + `")values($1,$2,0)on conflict("` + ChatTgModelField + `")do update set"` +
		ChatNameModelField + `"=$1 returning"` + ChatIDModelField + `","` + ChatModeratedModelField + `"`
	addAdmin = `insert into "` + UserModelTable + `"("` + ChatTgModelField + `","` + ChatNameModelField + `","` +
		UserAdminField + `")values($1,$2,$3)`
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

func (p Postgres) UpsertChatByTg(tg int64, name string) (result UpsertChatByTgModel, fail error) {
	fail = p.handle.QueryRow(upsertChatByTg, name, tg).Scan(&result.Id, &result.Moderated)
	return
}
func (p Postgres) UpsertUserByTg(tg int64, name string) (result UpsertUserByTgModel, fail error) {
	fail = p.handle.QueryRow(upsertUserByTg, name, tg).Scan(&result.Id, &result.Moderated)
	return
}

func (p Postgres) AddAdmins(id int64, name string) (sql.Result, error) {
	result, err := p.handle.Exec(addAdmin, id, name, 0)
	if err != nil {
		log.Println(err.Error())
	}
	return result, err
}

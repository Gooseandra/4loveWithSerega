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
	addBannedWord = `insert into "` + BannedWordModelTable + `"("` + BannedWordWordField + `","` + BannedWordDiscField +
		`")values($1,$2)`
	getPolicy = `select "word" from ` + BannedWordModelTable
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
	rows, fail := p.handle.Query(`select "tg" from "user"`)
	if fail != nil {
		return
	}
	for rows.Next() {
		var model ChatModel
		if fail = rows.Scan(&model.Tg); fail != nil {
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

func (p Postgres) AddBannedWord(word string) (sql.Result, error) {
	result, err := p.handle.Exec(addBannedWord, word)
	if err != nil {
		log.Println(err.Error())
	}
	return result, err
}

func (p Postgres) GetPolicy() (items []string, fail error) {
	rows, err := p.handle.Query(getPolicy)
	if err != nil {
		//TODO: какие то действия
	}
	for rows.Next() {
		var word string
		if fail = rows.Scan(&word); fail != nil {
			return
		}
		items = append(items, word)
	}
	return
}

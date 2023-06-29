package storage

import (
	"database/sql"
	"log"
	"strconv"
	"time"
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
	getPolicy   = `select "word" from ` + BannedWordModelTable
	getCriminal = `select "ID" from "ban" where` + ChatTgModelField + `= $1`
	addCriminal = `insert into "ban"("` + ChatTgModelField + `","` + BanWarningsField + `")values($1,$2)`
	getWarnings = `select "` + BanWarningsField + `fron "ban" where"` + ChatTgModelField + `" = $1`
	setBan      = `update "ban" set "warning" = 0, "banstart" = $1, "banendfor" = $2, banreason = $3 where
tg = $4`
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

func (p Postgres) Crime(id int64, warnings int, dur time.Duration) {
	var temp string
	row := p.handle.QueryRow(`select "ID" from "ban" where "tg" = $1`, id)
	row.Scan(&temp)
	if temp == "" {
		_, err := p.handle.Exec(addCriminal, id, 1)
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		row = p.handle.QueryRow(`select "warning" from "ban" where "tg" = $1`, id)
		row.Scan(&temp)
		warningsVal, err := strconv.Atoi(temp)
		if err != nil {
			//TODO: что то
		}
		warningsVal += 1
		if warningsVal >= warnings {
			_, err := p.handle.Exec(setBan, time.Now().Format("2006-01-02 15:04:05"), temp, "", id)
			if err != nil {
				log.Println(err.Error())
			}
			go p.Unban(id, dur)
		} else {
			p.handle.Exec(`update "ban" set "warning" = $1 where "tg" = $2`, warningsVal, id)
		}
	}
}

func (p Postgres) Unban(tg int64, dur time.Duration) {
	time.Sleep(dur)
	log.Println("Удаляю...")
	p.handle.Exec(`delete from "ban" where "tg" = $1`, tg)
	log.Println("Удалил")
}

func (p Postgres) GetBanList() []string {
	row, err := p.handle.Query(`select "tg" from "ban"`)
	if err != nil {
		//TODO: когда нибудь здесть что то будет...
	}
	var temp string
	var arr []string
	for row.Next() {
		row.Scan(&temp)
		arr = append(arr, temp)
	}
	return arr
}

func (p Postgres) SetWarnings(val int) {
	p.handle.Exec(`update "panishments" set "maxwarnings" = $1`, val)
}

func (p Postgres) SetBanTime(val int) {
	p.handle.Exec(`update "panishments" set "bantime" = $1`, val*60)
}

func (p Postgres) GetBanTime() time.Duration {
	row := p.handle.QueryRow(`select bantime from panishments`)
	var temp string
	row.Scan(&temp)
	t, err := time.Parse("01:02:03 01:02:03", temp) // тут надо получить duration
	log.Println(t)
	if err != nil {
		log.Println(err.Error())
	}
	return time.Hour
}

func (p Postgres) GetWarnings() int {
	row := p.handle.QueryRow(`select maxwarnings from panishments`)
	var temp string
	row.Scan(&temp)
	r, _ := strconv.Atoi(temp)
	return r
}

//time.Now().Local().Add(time.Hour * time.Duration(Hours) +
//time.Minute * time.Duration(Mins) +
//time.Second * time.Duration(Sec))

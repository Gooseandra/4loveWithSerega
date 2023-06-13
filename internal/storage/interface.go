package storage

import "database/sql"

type (
	Interface interface {
		LoadAdmins() ([]ChatModel, error)
		LoadChats() ([]int64, error)
		UpsertChatByTg(int64, string) (UpsertChatByTgModel, error)
		UpsertUserByTg(int64, string) (UpsertUserByTgModel, error)
		AddAdmins(int64, string) (sql.Result, error)
		AddBannedWord(string, string) (sql.Result, error)
		GetPolicy() ([]BannedWordModel, error)
	}

	UpsertChatByTgModel struct {
		Id        ChatIdModel
		Moderated bool
	}

	UpsertUserByTgModel struct {
		Id        ChatIdModel
		Moderated bool
	}
)

package storage

import (
	"database/sql"
	"time"
)

type (
	Interface interface {
		LoadAdmins() ([]ChatModel, error)
		LoadChats() ([]int64, error)
		UpsertChatByTg(int64, string) (UpsertChatByTgModel, error)
		UpsertUserByTg(int64, string) (UpsertUserByTgModel, error)
		AddAdmins(int64, string) (sql.Result, error)
		AddBannedWord(string) (sql.Result, error)
		GetPolicy() ([]string, error)
		Crime(int64, int, time.Duration)
		Unban(int64, time.Duration)
		GetBanList() []string
		SetWarnings(int)
		SetBanTime(int)
		GetWarnings() int
		GetBanTime() time.Duration
		GetUrls() []string
		GetPanishments() (string, string)
		DeleteBannedWord(string) bool
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

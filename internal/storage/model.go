package storage

const (
	ChatModelTable = "chat"

	ChatIDModelField   = "ID"
	ChatNameModelField = "name"
	ChatTgModelField   = "tg"
)

type (
	ChatModel struct {
		Name string
		Tg   int64
		ID   ChatIdModel
	}

	ChatIdModel int32
)

type UserModel struct {
	Name   string
	Id, Tg uint64
	Admin  uint32
}

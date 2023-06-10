package storage

const (
	ChatModelTable = "chat"

	ChatIDModelField        = "ID"
	ChatModeratedModelField = "moderated"
	ChatNameModelField      = "name"
	ChatTgModelField        = "tg"
)

type (
	ChatModel struct {
		Name string
		Tg   int64
		ID   ChatIdModel
	}

	ChatIdModel int32
)

type (
	UserModel struct {
		Name  string
		Tg    int64
		Id    int32
		Admin int16
	}
)

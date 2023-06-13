package storage

const (
	ChatModelTable = "chat"

	ChatIDModelField        = "ID"
	ChatModeratedModelField = "moderated"
	ChatNameModelField      = "name"
	ChatTgModelField        = "tg"

	UserModelTable = "user"
	UserAdminField = "admin"

	BannedWordModelTable = "bannedwords"
	BannedWordWordField  = "word"
	BannedWordDiscField  = "disc"
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
	BannedWordModel struct {
		id   int64
		Word string
		Disc string
	}
)

type (
	UserModel struct {
		Name  string
		Tg    int64
		Id    int32
		Admin int16
	}
)

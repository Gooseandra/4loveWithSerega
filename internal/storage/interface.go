package storage

type (
	Interface interface {
		LoadAdmins() ([]ChatModel, error)
		LoadChats() ([]int64, error)
		UpsertChatByTg(int64, string) (UpsertChatByTgModel, error)
		UpsertUserByTg(int64, string) (UpsertUserByTgModel, error)
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

package storage

type Interface interface {
	LoadAdmins() ([]ChatModel, error)
	LoadChats() ([]int64, error)
	UpdateChatByTg(int64, string) (ChatIdModel, error)
}

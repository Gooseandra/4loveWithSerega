package storage

type Interface interface {
	LoadAdmins() ([]AdminModel, error)
	LoadChats() ([]int64, error)
}

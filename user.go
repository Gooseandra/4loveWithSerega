package main

import "moderatorBot/internal/storage"

type (
	Admins struct {
		cache   map[storage.ChatIdModel]storage.ChatModel
		storage storage.Interface
	}
)

func NewAdmins(arg storage.Interface) (Admins, error) {
	slice, fail := arg.LoadAdmins()
	if fail != nil {
		return Admins{}, fail
	}
	result := Admins{cache: make(map[storage.ChatIdModel]storage.ChatModel), storage: arg}
	for _, value := range slice {
		result.cache[value.ID] = value
	}
	return result, nil
}

func (a Admins) Get(id uint64) {
	a.storage.LoadAdmins()
}

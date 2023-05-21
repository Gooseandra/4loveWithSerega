package main

import "moderatorBot/internal/storage"

type (
	Admins struct {
		cache   map[uint64]storage.AdminModel
		storage storage.Interface
	}
)

func NewAdmins(arg storage.Interface) (Admins, error) {
	slice, fail := arg.LoadAdmins()
	if fail != nil {
		return Admins{}, fail
	}
	result := Admins{cache: make(map[uint64]storage.AdminModel), storage: arg}
	for _, value := range slice {
		result.cache[value.Id] = value
	}
	return result, nil
}

func (a Admins) Get(id uint64) {
	a.storage.LoadAdmins()
}

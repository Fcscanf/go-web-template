package services

import (
	"github.com/fcant/pro/models"
	"github.com/fcant/pro/storages"
)

func ShowUserList() []models.User {
	return storages.UserList
}

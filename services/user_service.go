package services

import (
	"github.com/fcant/pro/models"
	"github.com/fcant/pro/storages"
)

func ShowUserList() models.ResponseEntity {
	return models.ResponseEntityInstance.Ok(storages.UserList)
}

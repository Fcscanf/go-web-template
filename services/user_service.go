package services

import (
	"github.com/fcant/pro/models"
	"github.com/fcant/pro/storages"
)

func ShowUserList() models.ResponseEntity {
	var responseEntityInstance = models.ResponseEntity{}
	return responseEntityInstance.Ok(storages.UserList)
}

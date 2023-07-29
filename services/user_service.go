package services

import (
	"github.com/fcant/pro/models"
	"github.com/fcant/pro/storages"
)

func ShowUserList() models.ResponseEntity {
	var ResponseEntityInstance = models.ResponseEntity{}
	return ResponseEntityInstance.Ok(storages.UserList)
}

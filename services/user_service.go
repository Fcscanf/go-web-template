package services

import (
	"github.com/fcant/pro/model"
	"github.com/fcant/pro/storage"
)

func ShowUserList() model.ResponseEntity {
	return model.ResponseEntityInstance.Ok(storage.UserList)
}

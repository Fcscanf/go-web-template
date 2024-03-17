package apis

import (
	"github.com/fcant/pro/interceptors"
	"github.com/fcant/pro/services"
	"net/http"
)

func ShowUserList(writer http.ResponseWriter, request *http.Request) {
	interceptors.ResponseFromService(writer, services.ShowUserList(), nil)
}

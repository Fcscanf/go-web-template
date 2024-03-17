package interceptors

import (
	"encoding/json"
	"github.com/Fcscanf/go-utils/httputils"
	"github.com/fcant/pro/models"
	"log"
	"net/http"
)

func HandlerRecordOptionUserIP(request *http.Request) {
	log.Printf("客户端IP为[%s]的用户发起命令执行操作-------》》》》》》", httputils.RemoteIP(request))
}

func ResponseJson(writer http.ResponseWriter, result models.ResponseEntity) {
	response, _ := json.Marshal(result)
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, _ = writer.Write(response)
}

func ResponseFromService(writer http.ResponseWriter, result any, err error) {
	if err != nil {
		log.Printf("api error of %s", err.Error())
		var responseEntityInstance = models.ResponseEntity{}
		response, _ := json.Marshal(responseEntityInstance.FailMessage("接口出错"))
		_, _ = writer.Write(response)
	} else {
		var responseEntityInstance = models.ResponseEntity{}
		ResponseJson(writer, responseEntityInstance.Ok(result))
	}

}

func ResponseVoidService(writer http.ResponseWriter, err error) {
	if err != nil {
		log.Printf("api error of %s", err.Error())
		var responseEntityInstance = models.ResponseEntity{}
		response, _ := json.Marshal(responseEntityInstance.FailMessage("接口出错"))
		_, _ = writer.Write(response)
	} else {
		var responseEntityInstance = models.ResponseEntity{}
		ResponseJson(writer, responseEntityInstance.Ok(nil))
	}

}

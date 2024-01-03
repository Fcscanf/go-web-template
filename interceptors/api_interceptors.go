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
	_, _ = writer.Write(response)
}

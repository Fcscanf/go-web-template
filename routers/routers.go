package routers

import (
	"github.com/fcant/pro/apis"
	"github.com/fcant/pro/constants"
	"net/http"
)

func AddRouters(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc(constants.ApiVersion+"user/show", cors(apis.ShowUserList))
	return mux
}

func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)                                                         // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
			w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
			w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
			w.Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json
		}
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		f(w, r)
	}
}

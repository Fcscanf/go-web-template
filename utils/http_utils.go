package utils

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// RemoteIP 通过 RemoteAddr 获取 IP 地址， 只是一个快速解析方法。
func RemoteIP(r *http.Request) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func RecordRequest(request *http.Request) {
	log.Printf("收到%s请求，对应的URL为：%s", request.Method, request.URL)
}

// RequestBodyJsonToModel 将Post请求的Body转为Struct
// 使用示例：
// var user = models.User{}
// utils.RequestBodyJsonToModel(request, &user)
func RequestBodyJsonToModel(request *http.Request, model interface{}) interface{} {
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal(body, &model)
	log.Printf("入参记录JSON转换：%#v", model)
	return model
}

func ResponseBodyJsonToModel(response *http.Response, model interface{}) interface{} {
	body, _ := io.ReadAll(response.Body)
	_ = json.Unmarshal(body, &model)
	log.Printf("出参记录JSON转换：%#v", model)
	return model
}

func RequestBodyXmlToModel(request *http.Request, model interface{}) interface{} {
	body, _ := io.ReadAll(request.Body)
	_ = xml.Unmarshal(body, &model)
	log.Printf("入参记录XML转换：%#v", model)
	return model
}

func ResponseBodyXmlToModel(response *http.Response, model interface{}) interface{} {
	body, _ := io.ReadAll(response.Body)
	_ = xml.Unmarshal(body, &model)
	log.Printf("出参记录XML转换：%#v", model)
	return model
}

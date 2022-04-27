package main

import (
	"fmt"
	"github.com/fcant/pro/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.Println("Hello, world!")
	fmt.Println("Exec Running ...")
	httpServer()
}

func httpServer() {
	server := &http.Server{
		/*设置端口地址*/
		Addr: ":4000",
		/*设置响应超时*/
		WriteTimeout: 2 * time.Second,
	}

	/*设置退出*/
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", &MHandler{})
	mux.HandleFunc("/bye", sayBye)
	mux = routers.AddRouters(mux)
	/*使当前Server指向自定义Mux*/
	server.Handler = mux

	go func() {
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close Server:", err)
		}
	}()

	log.Println("Starting Server...v3")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Fatal("Server closed under request")
		} else {
			log.Fatal("Server closed Unexpexted")
		}
	}
	log.Println("Server Exit")
}

type MHandler struct {
}

func (*MHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello,this is version 3!, the request URL is: " + r.URL.String()))
}

func sayBye(writer http.ResponseWriter, request *http.Request) {
	/*设置超时*/
	time.Sleep(1 * time.Second)
	_, _ = writer.Write([]byte("Bye bye,this is version 1!"))
}

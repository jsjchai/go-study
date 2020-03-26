package main

import (
	"fmt"
	"http-learn/router"
	"net/http"
)

type httpHandler struct {
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request url:" + r.Host + "" + r.RequestURI)
	_, _ = w.Write([]byte("hello,world"))
}

func main() {
	//http.Handle("/hello", &httpHandler{})
	//_ = http.ListenAndServe(":9898", nil)

	router.Init()

}

package router

import (
	"fmt"
	"io"
	"net/http"
)

func loginHandle(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "login success!\n")
}

func logoutHandle(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "logout success!\n")
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "WelCome!\n")
}

func errorHandle(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "error!\n")
}

func Init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandle)
	mux.HandleFunc("/index", indexHandle)
	mux.HandleFunc("/logout", logoutHandle)
	mux.HandleFunc("/", errorHandle)

	err := http.ListenAndServe(":9998", mux)
	if err != nil {
		fmt.Println(err)
	}
}

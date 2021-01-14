package server

import (
	"executor/internal/core"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	log.Printf("version: %s", core.Version)
	core.ReadConfig()
	http.HandleFunc("/", handler)

	err := startServer()
	if err != nil {
		log.Fatal("failed start server: ", err)
	}
}

func startServer() (err error) {
	log.Print("starting http server on port :8080")
	err = http.ListenAndServe(":8080", nil)
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("wrong query params err: %v", err), http.StatusBadRequest)
		return
	}

	args := make(map[string]string)
	for key, values := range r.Form {
		args[key] = values[0]
	}

	err := core.ExecuteComponent(r.FormValue("component"), r.FormValue("key"), args)
	if err != nil {
		http.Error(w, fmt.Sprintf("deploy err: %v", err), http.StatusBadRequest)
		return
	}
}

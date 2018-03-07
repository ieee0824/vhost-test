package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type request struct {
	Host string `json:"host"`
	Msg  string `json:"msg"`
}

func main() {
	http.HandleFunc("/addHost", func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		json.NewDecoder(r.Body).Decode(req)

		http.HandleFunc(req.Host, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, req.Msg)
		})
		fmt.Fprintf(w, "addHost")
	})
	http.HandleFunc("vhost.local.ieee.moe/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "vhost")
	})
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/MP_verify_32iWga2EVle6QTQm.txt", service.MpVerifyHandler)
	http.HandleFunc("/api/count", service.CounterHandler)
	http.HandleFunc("/api/callback", service.LoginHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}

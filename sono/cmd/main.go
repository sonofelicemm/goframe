package main

import (
	"flag"
	"fmt"
	"net/http"

	"sono/conf"
	"sono/log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprintf(w, "Hello world!")
}

func sono(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() err:%+v", err)
	}
	log.Init(conf.Conf.Log)
	http.HandleFunc("/hello", sayhelloName)
	http.HandleFunc("/sono", sono)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Info("ListenAndServe:%+v", err)
	}
}

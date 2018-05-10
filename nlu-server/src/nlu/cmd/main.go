package main

import (
	"fmt"
	"net/http"
	"nlu/conf"
	"nlu/log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	log.Info("Hello mm")
	fmt.Fprintf(w, "Hello world!")
}

func sono(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprintf(w, "Hello world!")
}

func main() {
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
	log.Info("hello mm")
}

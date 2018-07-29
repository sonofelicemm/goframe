package main

import (
    "fmt"
    "net/http"
    "sono/conf"
    "sono/log"
)

func mockResponse(w http.ResponseWriter, r *http.Request) {

    fmt.Fprint(w, string("test good job"))
}

func main() {
    if err := conf.Init(); err != nil {
        log.Error("conf.Init() err:%+v", err)
    }
    log.Init(conf.Conf.Log)
    http.HandleFunc("/test/", mockResponse)
    err := http.ListenAndServe(":8989", nil)
    if err != nil {
        log.Info("ListenAndServe:%+v", err)
    }
}

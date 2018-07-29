package service

import (
    "net/http"
    "fmt"
    "sono/model"
    "sono/utils"
    "encoding/json"
    "gopkg.in/tomb.v2"
    "sono/conf"
    "sono/log"
)

type Proxy struct {
    DataService *DataService
    Conf        *conf.Config
}

func (p *Proxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    bookName := utils.ExtractUuid(r.URL.Path)
    url := r.URL.Path

    if utils.IsTest(url) {
        p.Deal(rw, r, bookName, &model.PersonalBook{})
    } else if utils.IsNormal(url) {
        p.Deal(rw, r, bookName, &model.PublicBook{})
    } else {
        rw.WriteHeader(404)
        rw.Header().Add("Content-Type", "application/json; charset=utf-8")
        fmt.Fprint(rw, "Unsupported url")
    }
}

func (p *Proxy) Deal(rw http.ResponseWriter, r *http.Request, bookName string, book model.Book) {
    rw.Header().Add("Content-Type", "application/json; charset=utf-8")
    var t tomb.Tomb
    var bookInfoMap = make(model.BookInfoMap)
    localCache := p.DataService.GetLocalCache()
    t.Gos(
        func() error {
            return p.DataService.GetBookInfoByName(bookName, bookInfoMap)
        },
        func() error {
            return book.GetName(r, p.Conf)
        })
    err := t.Wait()
    if err != nil {
        rw.WriteHeader(401)
        fmt.Fprint(rw, string("Call sono error or read data error"))
    } else {
        book.GetTranslator(bookInfoMap, localCache)
        res, err := json.Marshal(book)
        if err != nil {
            log.Error("json new book error")
        }
        fmt.Fprint(rw, string(res))
    }
}

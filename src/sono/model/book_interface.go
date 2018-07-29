package model

import (
    "net/http"
    "sono/conf"
)

type Book interface {
    GetName(r *http.Request, conf *conf.Config) error
    GetTranslator(bookInfoMap BookInfoMap, localCache map[string]int) ()
}

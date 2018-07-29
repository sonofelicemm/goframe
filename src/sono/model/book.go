package model

import (
    "net/http"
    "sono/conf"
)

type PersonalBook struct {
    BookId     string `json:"book_id,omitempty"`
    BookName   int    `json:"book_name,omitempty"`
    Translator string `json:"translator,omitempty"`
}

type PublicBook struct {
    BookId     string `json:"book_id,omitempty"`
    BookName   int    `json:"book_name,omitempty"`
    Translator string `json:"translator,omitempty"`
    Other      string `json:"other"`
}

func (personalBook *PersonalBook) GetTranslator(bookInfoMap BookInfoMap, localCache map[string]int) {

}
func (personalBook *PersonalBook) GetName(r *http.Request, conf *conf.Config) error {
    return nil
}

func (publicBook *PublicBook) GetName(r *http.Request, conf *conf.Config) error {
    return nil
}

func (publicBook *PublicBook) GetTranslator(bookInfoMap BookInfoMap, localCache map[string]int) {

}

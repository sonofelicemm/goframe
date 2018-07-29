package model

import (
    "strings"
)

type PersonInfo struct {
    Id         int
    Name       string
    UserId     string
    BookIds    string
    CreateTime string
    UpdateTime string
}

func (personInfo *PersonInfo) GetBookIdList() (bookIds []string) {
    array := strings.Split(personInfo.BookIds, ",")
    return array
}

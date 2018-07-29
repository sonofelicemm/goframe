package model

type BookInfo struct {
    BookId       string
    Name         string
}

type BookInfoMap map[string]map[string]BookInfo

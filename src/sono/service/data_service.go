package service

import (
    "sono/model"
    "sono/dao"
    "sono/cache"
)

type DataService struct {
    Mysql      *dao.Mysql
    Redis      *dao.Redis
    LocalCache *cache.LocalCache
}

func (dataService *DataService) GetBookInfoByName(name string, bookInfoMap model.BookInfoMap) error {
    return nil
}

func (dataService *DataService) GetLocalCache() map[string]int {
    return nil
}

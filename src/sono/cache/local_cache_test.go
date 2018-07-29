package cache

import (
    "testing"
    "sono/dao"
)

func TestLocalCache_GetMysqlData(t *testing.T) {
    mysql := dao.NewMysql("sonofelice", "123456", "127.0.0.1:8902", "sono")
    localCache := NewCache(mysql, "@every 0.1m")
    data := localCache.GetMysqlData()
    if data["sys"] != 103 {
        t.Error(" fail")
    }
}

func TestLocalCache_SetCache(t *testing.T) {
    mysql := dao.NewMysql("sonofelice", "123456", "127.0.0.1:8902", "sono")
    localCache := NewCache(mysql, "@every 0.1m")

    cacheData := localCache.GetMysqlData()
    if cacheData["sys"] != 106 {
        t.Error("get sys fail")
    }

    localCache.SetCache(make(map[string]int))
    dataAfter := localCache.GetMysqlData()

    if len(dataAfter) != 0 {
        t.Error("set cache fail")
    }

}

package cache

import (
    "github.com/muesli/cache2go"
    "sono/constant"
    "sono/dao"
    "sono/log"
    "github.com/robfig/cron"
)

type LocalCache struct {
    Mysql *dao.Mysql
    Freq  string
}

// 之所以采用这种方式，是为了给定时器设置一个默认值，go中无法在创建struct的时候设置默认值
func NewCache(mysql *dao.Mysql, freq string) (c *LocalCache) {
    c = &LocalCache{}
    c.Mysql = mysql
    if freq == "" {
        c.Freq = "@every 30m"
    } else {
        c.Freq = freq
    }
    return c
}

func (localCache *LocalCache) GetMysqlData() (map[string]int) {
    cache := cache2go.Cache(constant.CACHE_NAME)
    res, err := cache.Value(constant.CACHE_KEY)
    if err == nil {
        data := res.Data().(*map[string]int)
        return *data
    } else {
        log.Warn("memory cache exception, read data form db ", err)
        data := localCache.Mysql.GetCacheData()
        return data
    }
}

func (localCache *LocalCache) RefreshCache() {
    c := cron.New()
    c.AddFunc(localCache.Freq, func() {
        cacheData := localCache.Mysql.GetCacheData()
        if cacheData != nil {
            localCache.SetCache(cacheData)
            log.Info("local cache refreshed")
        } else {
            log.Warn("Read from db is nil, please check db connection")
        }

    })
    c.Start()
}

func (localCache *LocalCache) Init() {
    localCache.SetCache(localCache.Mysql.GetCacheData())
    localCache.RefreshCache()
}

func (localCache *LocalCache) SetCache(cacheData map[string]int) {
    cache := cache2go.Cache(constant.CACHE_NAME)
    cache.Add(constant.CACHE_KEY, 0, &cacheData)
}

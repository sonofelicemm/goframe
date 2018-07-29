package dao

import (
    "fmt"
    "github.com/go-redis/redis"
    "sono/constant"
    "sono/log"
)

type Redis struct {
    client *redis.Client
}

func NewRedis(addr, passord string) (r *Redis) {
    r = &Redis{}
    client := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: passord, // no password set
        DB:       0,       // use default DB
    })
    r.client = client
    return
}

func (f *Redis) Set(key string, value string) error {
    err := f.client.Set(generateKey(key), value, 0).Err()
    return err
}

func (f *Redis) SetHashTable(key string, hashTable map[string]interface{}) {
    err := f.client.HMSet(generateKey(key), hashTable).Err()
    if err != nil {
        // just print log, cannot block main process
        log.Error("Set Hash Table error, please check redis status", err)
    }
}

func (f *Redis) Get(key string) (string, error) {
    val, err := f.client.Get(generateKey(key)).Result()
    if err != nil {
        log.Error("GetHashTable error", err)
        return "", err
    }
    fmt.Println("key", val)
    return val, err
}

func (f *Redis) GetHashTable(key string, field string) (string, error) {
    val, err := f.client.HMGet(generateKey(key), field).Result()
    if err != nil {
        log.Error("GetHashTable error", err)
        return "", err
    }
    var str string
    for _, param := range val {
        str += param.(string)
    }
    return str, err
}

// 主要用这个，一次从redis中读出一个id对一个的所有hashTable。但是在key对应的数据量比较大的情况下，非常不推荐使用HGetAll
func (f *Redis) GetAllHashTable(key string) (map[string]string, error) {
    hashTable, err := f.client.HGetAll(generateKey(key)).Result()
    return hashTable, err
}

func (f *Redis) Close() {
    defer f.client.Close()
}

func generateKey(key string) string {
    return constant.SONO_REDIS_KEY + key
}

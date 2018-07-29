package dao

import (
    "testing"
    "fmt"
)

func TestRedis_Get(t *testing.T) {
    redis := NewRedis("127.0.0.1:6379", "")
    redis.Set("key1", "value1")
    val1, err := redis.Get("key1")
    if err != nil {
        fmt.Println(val1)
    }
    redis.Close()
}

func TestRedis_GetEntity(t *testing.T) {
    redis := NewRedis("127.0.0.1:6379", "")
    redis.Close()
}

func TestRedis_GetAllHashTable(t *testing.T) {
}

package dao

import (
    "testing"
    "fmt"
)

func TestDao_MysqlSelect(t *testing.T) {
    db := NewMysql("sonofelice", "123456", "127.0.0.1:8902", "sono")
    var effectiveList = db.GetSonoData("sono")
    fmt.Printf("list len is  :%+v", len(effectiveList))
    db.Close() //关闭
}
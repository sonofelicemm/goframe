package dao

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "sono/log"
    "sono/model"
)

type Mysql struct {
    db *sql.DB
}

func NewMysql(userName, password, ipHost, dbName string) (d *Mysql) {
    d = &Mysql{}
    client, err := sql.Open("mysql", userName+":"+password+"@tcp("+ipHost+")/"+dbName)
    if err != nil {
        log.Error("Mysql connection failed!")
    }
    d.db = client
    return
}

func (d *Mysql) Close() {
    defer d.db.Close()
}

func (d *Mysql) GetSonoData(name string) []*model.BookInfo {
    rows, err := d.db.Query("SELECT * FROM sono_name WHERE name= \"" + name + "\"")
    if err != nil {
        log.Error("db query error:{}", err)
        return nil
    }

    var bookList []*model.BookInfo

    for rows.Next() {
        book := &model.BookInfo{}
        if err = rows.Scan(&book.BookId, &book.Name); err == nil {
            bookList = append(bookList, book)
        } else {
            log.Error("rows scan error :{}", err)
            break
        }
    }
    rows.Close()
    return bookList
}

func (d *Mysql) GetCacheData() (map[string]int) {
    rows, err := d.db.Query("select id, name FROM book")
    if err != nil {
        log.Error("db query error:{}", err)
        return nil
    }
    cacheData := make(map[string]int)
    for rows.Next() {
        var name string
        var id int
        err = rows.Scan(&id, &name)
        cacheData[name] = id
    }
    return cacheData
}

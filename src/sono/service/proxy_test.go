package service

import (
    "testing"
    "net/http/httptest"
    . "gopkg.in/check.v1"
    "net/http"
    "fmt"
    "sono/dao"
    "sono/conf"
    "sono/cache"
    "sono/log"
)


func Test(t *testing.T) { TestingT(t) }

type ServerSuite struct {
    ts *httptest.Server
}

var mysql *dao.Mysql
var redis *dao.Redis

func (s *ServerSuite) SetUpSuite(c *C) {
    h := http.NewServeMux()
    h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, GetResponse())
    })
    s.ts = httptest.NewServer(h)

    if err := conf.Init(); err != nil {
        log.Error("conf.Init() err:%+v", err)
    }
    innerConf := conf.Conf
    (*innerConf.Log).Dir = "../../../logs"
    log.Init(innerConf.Log)

    redisConf := innerConf.Redis
    mysqlConf := innerConf.Mysql

    redis = dao.NewRedis(redisConf.Addr, redisConf.Password)
    mysql = dao.NewMysql(mysqlConf.UserName, mysqlConf.Password, mysqlConf.IpHost, mysqlConf.DbName)

}

func (s *ServerSuite) TearDownSuite(c *C) {
    s.ts.Close()
}

var _ = Suite(&ServerSuite{})

func (s *ServerSuite) TestFoo(c *C) {
    if err := conf.Init(); err != nil {
        log.Error("conf.Init() err:%+v", err)
    }
    mux := http.NewServeMux()
    localCache := cache.NewCache(mysql, conf.Conf.GetCronFreq())
    localCache.Init()
    mux.Handle("/", &Proxy{DataService: &DataService{Mysql: mysql, Redis: redis, LocalCache: localCache}, Conf: conf.Conf})
    writer := httptest.NewRecorder()


    request, _ := http.NewRequest("GET", "test/sono/8c9adb35-a8fa-4c01-a2ec-c834b3b47b99?token=tWa6dO8b8qFlGvfxXQUuv7HmNwqIcKZr/test", nil)

    mux.ServeHTTP(writer, request)

    fmt.Println(writer)
}

func GetResponse() string  {
    return string("good job")
}
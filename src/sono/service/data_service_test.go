package service

import (
    "testing"
    "sono/dao"
    "sono/conf"
    "sono/log"
)

func setUp() {
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

func TestGetBookInfoByName(t *testing.T) {
    setUp()
    // test process
}

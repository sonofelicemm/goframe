package conf

import (
    "sono/log"
    "github.com/BurntSushi/toml"
    "flag"
    "os"
    "path/filepath"
    "strings"
)

var (
    confPath string
    // Conf global
    Conf = &Config{}
)

// Config .
type Config struct {
    Log        *log.Config
    Mysql      *Mysql
    Redis      *Redis
    SonoServer *SonoServer
    Cron       *Cron
}

func (config *Config) GetCronFreq() string {
    if config.Cron != nil && config.Cron.Freq != "" {
        return config.Cron.Freq
    } else {
        return ""
    }
}

type Mysql struct {
    UserName string
    Password string
    IpHost   string
    DbName   string
}

type Redis struct {
    Addr     string
    Password string
}

type SonoServer struct {
    Port string
}

type Cron struct {
    Freq string
}

func init() {
    flag.StringVar(&confPath, "conf", getConPath(), "-conf path")
}

func getConPath() string {
    s, err := os.Getwd()
    if err != nil {
        log.Info("get wd error, use default conf path")
        return "./conf/conf.toml"
    }
    if strings.HasSuffix(s, "sono-server") {
        return filepath.Join("conf", "conf.toml")
    } else {
        return filepath.Join("../../../conf", "conf.toml")
    }
}

func Init() (err error) {
    _, err = toml.DecodeFile(confPath, &Conf)
    return
}

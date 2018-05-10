package conf

import (
	"flag"

	"sono/log"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	// Conf global
	Conf = &Config{}
)

// Config .
type Config struct {
	Log *log.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init init conf
func Init() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

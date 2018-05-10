package conf

import (
	"flag"

	"nlu/log"

	"github.com/BurntSushi/toml"
)

var (
	confPath string = "./src/nlu/cmd/nlu.toml"
	// Conf global
	Conf = &Config{}
)

// Config .
type Config struct {
	Log *log.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "./src/nlu/cmd/nlu.toml", "config path")
}

// Init init conf
func Init() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

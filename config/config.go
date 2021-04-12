package config
import (
"github.com/unknwon/goconfig"
"log"
)
func GetConfig() *goconfig.ConfigFile {
	c,err := goconfig.LoadConfigFile("./config/config.ini")
	if err != nil {
		log.Fatal(err)
	}
	return c
}


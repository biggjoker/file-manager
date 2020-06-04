package g

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)



type GlobalConfig struct {

	HttpAddr            string            `json:"http_addr"`
	Debug               bool              `json:"debug"`
	BaseDir             string            `json:"path"`
	Version         	string            `json:"version"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	configContent, err := ioutil.ReadFile(cfg)

	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	configLock.Lock()
	defer configLock.Unlock()
	config = &c

	log.Println("read config file:", config, "successfully")
}

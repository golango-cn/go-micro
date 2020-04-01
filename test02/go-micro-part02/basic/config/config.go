package config

import (
	log "github.com/micro/go-micro/v2/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"sync"
)

type Config struct {
	Jwt struct {
		SecretKey string `json:"secretKey"`
	} `json:"jwt"`
	Redis struct{
		Enabled string `json:"enabled"`
		Conn string `json:"conn"`
		DbNum int `json:"dbNum"`
		Password string `json:"password"`
		Timeout int `json:"timeout"`
	} `json:"redis"`
	App struct{
		Mysql struct{
			Enabled bool `json:"enabled"`
			Url string `json:"url"`
			MaxIdleConnection int `json:"maxIdleConnection"`
			MaxOpenConnection int `json:"maxOpenConnection"`
		} `json:"mysql"`

		Etcd struct{
			Enabled bool `json:"enabled"`
			Host string `json:"host"`
			Port int32 `json:"port"`
		}
	} `json:"app"`
}

var (
	m      sync.RWMutex
	inited bool
	sp     = string(filepath.Separator)
	AppConfig *Config
)

func Init() {

	m.Lock()
	defer m.Unlock()

	if inited {
		log.Info("[Init] 配置已经初始化过")
		return
	}

	path, _ := filepath.Abs(filepath.Dir("."))
	path = filepath.Join(path, "conf", "application.yml")

	bytes, err := ioutil.ReadFile(path)
	if err != nil{
		log.Fatal(err)
	}

	err = yaml.Unmarshal(bytes, &AppConfig)
	if err != nil{
		log.Fatal(err)
	}

}

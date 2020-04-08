package config

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"sync"

	"github.com/micro/go-micro/util/log"
)

var AppConfig *Config

type Config struct {
	App struct{

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

		Zap struct{
			zap.Config
			LogFileDir    string `json:logFileDir`
			AppName       string `json:"appName"`
			ErrorFileName string `json:"errorFileName"`
			WarnFileName  string `json:"warnFileName"`
			InfoFileName  string `json:"infoFileName"`
			DebugFileName string `json:"debugFileName"`
			MaxSize       int    `json:"maxSize"` // megabytes
			MaxBackups    int    `json:"maxBackups"`
			MaxAge        int    `json:"maxAge"` // days
		} `json:"zap"`

	} `json:"app"`
}

var (
	m      sync.RWMutex
	inited bool
)

// Configurator 配置器
type Configurator interface {
	App(name string, config interface{}) (err error)
	Path(path string, config interface{}) (err error)
}

// Init 初始化配置
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
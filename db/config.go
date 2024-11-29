package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

var conf *Config

type Config struct {
	Mysql      Mysql                  `yaml:"mysql"`
	Redis      []Redis                `yaml:"redis"`
	Mongod     Mongodb                `yaml:"mongodb"`
	RabitMQ    RabitMQ                `yaml:"rabitmq"`
	GlobalConf map[string]interface{} `yaml:"globalconf"`
}

type Mysql struct {
	Username string `yaml:"username"`
	Pwd      string `yaml:"pwd"`
	Host     string `yaml:"host"`
	DbName   string `yaml:"db-name"`
}

type Redis struct {
	Pwd  string `yaml:"pwd"`
	Host string `yaml:"host"`
	Db   int    `yaml:"db"`
}
type Mongodb struct {
	Host string `yaml:"host"`
}

type RabitMQ struct {
	Host string `yaml:"host"`
}

func init2() {
	//gin.SetMode(gin.ReleaseMode)
	yamlPath := "./config-dev.yaml"
	if gin.Mode() == gin.ReleaseMode {
		yamlPath = "./config-prod.yaml"
	}
	fmt.Println("yamlPath" + yamlPath)
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--config=") {
			yamlPath = arg[len("--config="):]
		}
	}
	res, err := os.ReadFile(yamlPath)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(res, &conf)
	if err != nil {
		log.Fatalln(err)
	}

	initRedis()
}

func GetConfig() *Config {
	return conf
}

func GetGlobalConf() map[string]interface{} {
	return conf.GlobalConf
}
func GetGlobalConfVal(k string) interface{} {
	if val, ok := conf.GlobalConf[k]; ok {
		return val
	}
	return nil
}

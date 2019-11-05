package config

import (
	"flag"
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
)

type AppConfig struct {
	AppMode string `yaml:"app_mode"`
	Version string `yaml:"version"`
	DB      string `yaml:"db"`
	Port    int    `yaml:"port"`
	GinMode string `yaml:"gin_mode"`
}

var AppCfg AppConfig

var env = flag.String("env", "dev", "app mode")

//var env = flag.String("env", "prod", "app mode")

func init() {
	flag.Parse()

	log.Println(*env)

	f := fmt.Sprintf("config/app.%s.yaml", *env)

	config, err := ioutil.ReadFile(f)

	if err != nil {
		log.Println(err)
	}
	yaml.Unmarshal(config, &AppCfg)
}

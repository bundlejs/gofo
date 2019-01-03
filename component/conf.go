package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Mysql struct {
	User        string `yaml:"USER"`
	Password    string `yaml:"PASSWORD"`
	Host        string `yaml:"HOST"`
	Port        string `yaml:"PORT"`
	Name        string `yaml:"NAME"`
	TablePrefix string `yaml:"TABLE_PREFIX"`
}

type Redis struct {
	User        string `yaml:"USER"`
	Password    string `yaml:"PASSWORD"`
	Host        string `yaml:"HOST"`
	Port        string `yaml:"PORT"`
	Name        string `yaml:"NAME"`
	TablePrefix string `yaml:"TABLE_PREFIX"`
}

type Database struct {
	Mysql Mysql `yaml:"MYSQL"`
	Redis Redis `yaml:"REDIS"`
}

type Jwt struct {
	Secret string `yaml:"SECRET"`
}

type Server struct {
	Port         string `yaml:"PORT"`
	ReadTimeout  string `yaml:"READ_TIMEOUT"`
	WriteTimeout string `yaml:"WRITE_TIMEOUT"`
}

type Conf struct {
	RunMode  string   `yaml:"RUN_MODE"`
	Database Database `yaml:"DATABASE"`
	Jwt      Jwt      `yaml:"JWT"`
	Server   Server   `yaml:"SERVER"`
}

func Config() Conf {
	var conf Conf
	config, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		fmt.Print(err)
	}
	if err := yaml.Unmarshal(config, &conf); err != nil {
		fmt.Print(err)
	}
	return conf
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"gopkg.in/yaml.v3"
)

type GinConfig1 struct {
	Port        string `yaml:"port"`
	DebugMode   string `yaml:"debug_mode"`
	SwaggerMode string `yaml:"swagger_mode"`
}

type Config1 struct {
	Config struct {
		GinConfig GinConfig1 `yaml:"http_config"`
	} `yaml:"config"`
}

func main() {
	fmt.Println("Hello, world!")
	path, err := filepath.Abs("demo.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println("path: ", path)
	fmt.Println(reflect.ValueOf(path).String())

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(err)
	}

	// 第一種方法
	fun1 := Config1{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &fun1)
	if err != nil {
		panic(err)
	}
	fmt.Println(fun1.Config.GinConfig)

}

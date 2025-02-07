package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func yamlv3(path string) {
	getfile := Config1{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &getfile)
	if err != nil {
		panic(err)
	}
	fmt.Println(getfile.Config.GinConfig)
}

func viper1(path string) *Config2 {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	getdata := &Config2{}
	if err := v.UnmarshalKey("config", getdata); err != nil {
		panic(err)
	}

	return getdata
}

func viper2(path string) *Config2 {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	getdata := &Config2{}
	setDefault(getdata, "") // parse nested structure `config`
	if err := v.UnmarshalKey("config", getdata); err != nil {
		panic(err)
	}

	return getdata
}

// parse nested structure `config` and set to default value by tag `default`
func setDefault(config any, rootKey string) {
	keyStack := []string{rootKey}
	pointer := reflect.ValueOf(config)
	if pointer.Kind() != reflect.Pointer {
		return
	}
	pointer = pointer.Elem()
	if pointer.Kind() != reflect.Struct {
		return
	}
	stack := []reflect.Value{pointer}

	for len(stack) > 0 {
		pointer = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		key := keyStack[len(keyStack)-1]
		keyStack = keyStack[:len(keyStack)-1]

		for i := 0; i < pointer.NumField(); i++ {
			field := pointer.Field(i)
			switch field.Kind() {
			case reflect.Struct:
				stack = append(stack, field)
				keyStack = append(keyStack, fmt.Sprintf("%s.%s", key, field.Type().Name()))
				break
			default:
				tag := pointer.Type().Field(i).Tag.Get("default")
				if tag == "" {
					continue
				}
				field.Set(reflect.ValueOf(tag))
			}
		}
	}
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
	yamlv3(path)

	// 第二種方法
	val1 := viper1(path)
	fmt.Println(val1)
	val2 := viper2(path)
	fmt.Println(val2)
}

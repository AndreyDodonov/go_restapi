package main

import (
	"flag"
	"fmt"
	"go_restapi/internal/app/apiserver"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

// специальная функция, выполняет перед main. init можут быть не один
func init()  {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

// точка входа в программу
func main() {
	flag.Parse()
	fmt.Println("1") //TODO debug
	fmt.Println(configPath) //TODO debug

	fmt.Println("newConfig") //TODO debug
	config := apiserver.NewConfig()
	fmt.Println(config)
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
 fmt.Println("start") //TODO debug
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}

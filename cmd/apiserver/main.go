package main

import (
	"flag"
	"fmt"
	"go_restapi/internal/app/apiserver"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

// специальная функция, выполняет перед main. init можут быть не один
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

// точка входа в программу
func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}


	// graceful shutdown

	go func() {
		fmt.Print("server shutdown")
		shutdown := make(chan struct{})
		signals := make(chan os.Signal, 1)
		signal.Notify(signals,
			syscall.SIGTERM,
			syscall.SIGINT)
			<- signals

			close(shutdown)
	}()

}

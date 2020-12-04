package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/env"
	"github.com/micro/go-plugins/config/source/configmap/v2"
)

var (
	DefaultNamespace  = "go-micro"
	DefaultConfigName = "go-micro-config"
)

func main() {
	if cfg, err := config.NewConfig(); err == nil {
		err = cfg.Load(
			env.NewSource(),
			configmap.NewSource(
				configmap.WithName(DefaultConfigName),
				configmap.WithNamespace(DefaultNamespace)),
		)
		if err == nil {
			fmt.Println(cfg.Map())
		}
		fmt.Println(err)
	}
}

package cmd

import (
	"fmt"
	"github.com/arzonus/alertmetrics/pkg/appc"
	"github.com/arzonus/alertmetrics/pkg/config"
	"github.com/arzonus/alertmetrics/pkg/context"
)

const (
	configPath = "./config.yaml"
)

// Start point for app
func Start() {
	fmt.Println("Starting API")
	err := Init(configPath)
	if err != nil {
		fmt.Println(err)
	}
}

func Init(configPath string) (err error) {
	// Init configuration - read file
	cfg, err := config.Init(configPath)
	if err != nil {
		return
	}

	if err = cfg.Validate(); err != nil {
		return
	}

	_, err = context.Init()
	if err != nil {
		return
	}

	app := appc.New()
	app.Init()
	app.Run()

	return
}

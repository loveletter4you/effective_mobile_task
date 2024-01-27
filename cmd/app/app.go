package main

import "github.com/loveletter4you/effective_mobile_task/internal/app"

const configPath = "config/config.yaml"

func main() {
	app.Run(configPath)
}

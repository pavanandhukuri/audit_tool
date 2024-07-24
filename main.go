package main

import (
	"security_audit_tool/app"
	"security_audit_tool/services/config"
)

func main() {
	loadConfig()
	app.Run()
}

func loadConfig() {
	_, err := config.Load()
	if err != nil {
		panic(err)
	}
}

package main

import (
	"security_audit_tool/app"
	"security_audit_tool/services/config"
)

func main() {
	_, err := config.Load()
	if err != nil {
		return
	}
	app.Run()
}

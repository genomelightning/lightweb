package main

import (
	"github.com/astaxie/beego"

	"github.com/genomelightning/lightweb/routers"
)

const (
	APP_VER = "0.0.0.1220"
)

func main() {
	beego.Info("Lightning Web", APP_VER)

	// Register routers.
	beego.Router("/", &routers.HomeRouter{})

	beego.Run()
}

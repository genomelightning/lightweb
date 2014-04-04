package main

import (
	"github.com/astaxie/beego"

	"github.com/genomelightning/lightweb/models"
	"github.com/genomelightning/lightweb/routers"
)

const (
	APP_VER = "0.0.0.0403"
)

func main() {
	models.CallMe()
	beego.Info("Lightning Web", APP_VER)

	// Register routers.
	beego.Router("/", &routers.HomeRouter{})
	beego.Router("/query", &routers.QueryRouter{})
	beego.Router("/tools/genomebrowser", &routers.GenomeBrowserRouter{})

	beego.Run()
}

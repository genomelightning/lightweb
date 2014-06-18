package main

import (
	"github.com/astaxie/beego"

	// "github.com/genomelightning/lightweb/models"
	"github.com/genomelightning/lightweb/routers"
)

const (
	APP_VER = "0.0.1.0618"
)

func main() {
	beego.Info("Lightning Web", APP_VER)

	// Register routers.
	beego.Router("/", &routers.HomeRouter{})
	beego.Router("/query", &routers.QueryRouter{})
	beego.Router("/tools/genomebrowser", &routers.GenomeBrowserRouter{})

	beego.Router("/api/v1/pngs", &routers.HomeRouter{}, "get:Pngs")
	beego.Run()
}

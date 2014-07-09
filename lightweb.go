package main

import (
	"github.com/astaxie/beego"

	// "github.com/genomelightning/lightweb/models"
	"github.com/genomelightning/lightweb/routers"
	"github.com/genomelightning/lightweb/routers/api"
)

const (
	APP_VER = "0.0.2.0708"
)

func main() {
	beego.Info("Lightning Web", APP_VER)

	// Register routers.
	beego.Router("/", &routers.HomeRouter{})
	beego.Router("/query", &routers.QueryRouter{})
	beego.Router("/tools/genomebrowser", &routers.GenomeBrowserRouter{})

	beego.AddNamespace(beego.NewNamespace("/api/v1",
		beego.NSNamespace("/gms",
			beego.NSRouter("/tiles", &api.ApiV1Router{}, "get:FetchGMSTiles"),
		),
	))
	beego.Run()
}

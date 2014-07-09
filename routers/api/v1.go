package api

import (
	"fmt"
	"path"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
)

type ApiV1Router struct {
	beego.Controller
}

func (c *ApiV1Router) URLMapping() {
	c.Mapping("FetchGMSTiles", c.FetchGMSTiles)
}

func (this *ApiV1Router) FetchGMSTiles() {
	band, _ := this.GetInt("band")
	pos, _ := this.GetInt("pos")
	pngUrls := make([]string, 9)

	curPos := pos - 1
	curBand := band - 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			pngName := fmt.Sprintf("%d/%d.png", curBand, curPos)
			if com.IsFile(path.Join("static/pngs", pngName)) {
				pngUrls[i*3+j] = fmt.Sprintf("http://localhost:8085/static/pngs/%s", pngName)
				pngUrls[i*3+j] = fmt.Sprintf("http://lightning-dev.curoverse.com/static/pngs/%s", pngName)
			}
			curBand++
			// url:=fmt.Sprintf("http://lightning-dev.curoverse.com/api/v1/", ...)
		}
		curBand = band - 1
		curPos++
	}
	this.Data["json"] = pngUrls
	this.ServeJson(true)
}

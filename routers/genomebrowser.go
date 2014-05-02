package routers

import (
// "fmt"

// "github.com/ajstarks/svgo"

// "github.com/genomelightning/lightning/cytomap"

// "github.com/genomelightning/lightweb/models"
)

type GenomeBrowserRouter struct {
	baseRouter
}

var colorMap = map[string]string{
	"gneg":    "none",
	"gpos25":  "gray",
	"gpos50":  "#424242",
	"gpos75":  "#1C1C1C",
	"gpos100": "black",
	"acen":    "red",
}

func (this *GenomeBrowserRouter) Get() {
	// rules := make([]*cytomap.CytoRule, 0, 10)
	// var totalLength int64 = 0
	// for _, rule := range models.HgCytoMaps[19].Rules {
	// 	if rule.Chr == "chr1" {
	// 		continue
	// 	} else if rule.Chr != "chr2" {
	// 		break
	// 	}
	// 	totalLength = rule.End
	// 	rules = append(rules, rule)
	// }

	// width := 200
	// maxHeight := 800
	// unitLength := float64(totalLength) / float64(maxHeight)

	// canvas := svg.New(this.Ctx.ResponseWriter)
	// canvas.Start(width, maxHeight)

	// blockHeight := 0.0
	// heightStart := 30.0
	// curHeightStart := heightStart
	// curHeight := 0.0
	// for i, rule := range rules {
	// 	if rule.Section[0] == 'q' && blockHeight < 1 {
	// 		blockHeight = curHeightStart
	// 		canvas.Rect(50, 30, 30, int(curHeightStart-heightStart), "fill:none; stroke:black")
	// 	}
	// 	curHeight = float64(rule.End-rule.Start) / unitLength
	// 	canvas.Rect(50, int(curHeightStart), 30, int(curHeight), "fill:"+colorMap[rule.Color])
	// 	fmt.Println(i, int(curHeightStart-heightStart), int(curHeight), int((curHeightStart-heightStart+curHeight)*unitLength), colorMap[rule.Color])
	// 	curHeightStart += curHeight
	// }
	// canvas.Rect(50, int(blockHeight), 30, int(curHeightStart-heightStart-blockHeight), "fill:none; stroke:black")
	// fmt.Println(int(blockHeight), curHeightStart-heightStart, int(curHeightStart-heightStart-blockHeight))
	// // canvas.Rect(50, 45, 30, 7, "fill:gray")
	// // canvas.Rect(50, 70, 30, 13, "fill:#424242")
	// // canvas.Rect(50, 90, 30, 10)
	// // canvas.Rect(50, 110, 30, 5, "fill:#585858")
	// // canvas.Rect(50, 122, 30, 15)
	// // canvas.Rect(50, 144, 30, 10, "fill:gray")
	// // canvas.Rect(50, 172, 30, 11)
	// // canvas.Rect(50, 188, 30, 12)
	// // canvas.Rect(50, 210, 30, 7, "fill:gray")
	// // canvas.Rect(50, 223, 30, 5, "fill:gray")
	// // canvas.Rect(50, 231, 30, 22)
	// // canvas.Rect(50, 260, 30, 10, "fill:red")
	// // canvas.Rect(50, 30, 30, 240, "fill:none; stroke:black")
	// canvas.End()
	this.TplNames = "genomebrowser.html"
}

package routers

import (
	"fmt"
	"os"
	"path"

	"github.com/Unknwon/com"
)

type HomeRouter struct {
	baseRouter
}

func (this *HomeRouter) Get() {
	this.Data["IsHome"] = true
	this.TplNames = "home.html"
}

func (this *HomeRouter) Post() {
	this.TplNames = "home.html"

	// Get form data.
	start, err := this.GetInt("startPoint")
	if err != nil {
		this.Data["Result"] = "Invalid start position."
		return
	}

	end, err := this.GetInt("endPoint")
	if err != nil {
		this.Data["Result"] = "Invalid end position."
		return
	}

	// Load raw genome data.
	f, err := os.Open("data/chr17.fa")
	if err != nil {
		this.Data["Result"] = "Fail to open file: " + err.Error()
		return
	}
	defer f.Close()

	if start <= 0 || end <= 0 {
		this.Data["Result"] = fmt.Sprintf("Invalid range: %v~%v", start, end)
		return
	}

	p := make([]byte, end-start+1)
	f.ReadAt(p, start+6)
	this.Data["Result"] = string(p)
}

func (this *HomeRouter) Pngs() {
	band, _ := this.GetInt("band")
	pos, _ := this.GetInt("pos")
	pngUrls := make([]string, 9)

	curPos := pos - 1
	curBand := band - 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			pngName := fmt.Sprintf("%d-%d.png", curBand, curPos)
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

// func (this *HomeRouter) Get() {
// 	gs1 := &genome.Sequence{}
// 	gs1.Blocks = append(gs1.Blocks, &genome.Block{
// 		Valid: false,
// 		Data:  []byte("GGGGGGGGAAAANAAACCACCCCCC"),
// 	})
// 	gs1.Blocks = append(gs1.Blocks, &genome.Block{
// 		Valid: true,
// 		Data:  []byte("GGCGGGGGAAAAAAAACCCCCCCCC"),
// 	})
// 	gs1.Blocks = append(gs1.Blocks, &genome.Block{
// 		Valid: false,
// 		Data:  []byte("GGGGNNGGAAAGAAAACCCCCCCCC"),
// 	})

// 	gs2 := &genome.Sequence{}
// 	gs2.Blocks = append(gs2.Blocks, &genome.Block{
// 		Valid: true,
// 		Data:  []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
// 	})
// 	gs2.Blocks = append(gs2.Blocks, &genome.Block{
// 		Valid: false,
// 		Data:  []byte("GGGGGGGGAAAAAAAACCNNCCCCC"),
// 	})
// 	gs2.Blocks = append(gs2.Blocks, &genome.Block{
// 		Valid: false,
// 		Data:  []byte("GGGGGGGNNAAAAAAACCCCCCCCC"),
// 	})
// 	bs, err := lightning.ComputeDiffSeq(gs1, gs2)
// 	if err != nil {
// 		this.Ctx.WriteString("error: " + err.Error())
// 		return
// 	}

// 	this.Ctx.WriteString("DumpAsBits:\n" + bs.DumpAsBits() +
// 		"\nDumpAsType:\n" + bs.DumpAsType())
// }

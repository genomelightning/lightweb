package routers

import (
	"fmt"
	"os"
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

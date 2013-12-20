package routers

import (
	"github.com/astaxie/beego"
	"github.com/genomelightning/lightning"
	"github.com/genomelightning/lightning/genome"
)

type baseRouter struct {
	beego.Controller
}

type HomeRouter struct {
	baseRouter
}

func (this *HomeRouter) Get() {
	gs1 := &genome.Sequence{}
	gs1.Blocks = append(gs1.Blocks, &genome.Block{
		Valid: false,
		Data:  []byte("GGGGGGGGAAAANAAACCACCCCCC"),
	})
	gs1.Blocks = append(gs1.Blocks, &genome.Block{
		Valid: true,
		Data:  []byte("GGCGGGGGAAAAAAAACCCCCCCCC"),
	})
	gs1.Blocks = append(gs1.Blocks, &genome.Block{
		Valid: false,
		Data:  []byte("GGGGNNGGAAAGAAAACCCCCCCCC"),
	})

	gs2 := &genome.Sequence{}
	gs2.Blocks = append(gs2.Blocks, &genome.Block{
		Valid: true,
		Data:  []byte("GGGGGGGGAAAAAAAACCCCCCCCC"),
	})
	gs2.Blocks = append(gs2.Blocks, &genome.Block{
		Valid: false,
		Data:  []byte("GGGGGGGGAAAAAAAACCNNCCCCC"),
	})
	gs2.Blocks = append(gs2.Blocks, &genome.Block{
		Valid: false,
		Data:  []byte("GGGGGGGNNAAAAAAACCCCCCCCC"),
	})
	bs, err := lightning.ComputeDiffSeq(gs1, gs2)
	if err != nil {
		this.Ctx.WriteString("error: " + err.Error())
		return
	}

	this.Ctx.WriteString("DumpAsBits:\n" + bs.DumpAsBits() +
		"\nDumpAsType:\n" + bs.DumpAsType())
}

package routers

type QueryRouter struct {
	baseRouter
}

func (this *QueryRouter) Get() {
	this.TplNames = "query.html"
}

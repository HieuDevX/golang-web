package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Handle("GET", "/", func(ctx iris.Context) {
		htmlH1String := "<h1 style='color:red'>Web server chạy bằng IRIS</h1>"
		htmlH2String := "<h2 style='font-style: italic; font-size: 20px; color:green'>Học lập trình back-end</h2>"
		returnHTML := htmlH1String + htmlH2String
		ctx.HTML(returnHTML)
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

package main

import (
	"Blog/logs"
	"github.com/kataras/iris"
)

func main() {
	f := logs.NewLogFile()
	defer f.Close()

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Logger().AddOutput(f)
	app.Handle("GET", "/", func(ctx iris.Context) {
		app.Logger().Error("fuck")
		ctx.Gzip(true)
		ctx.HTML("<h1>fuck you</h1>")

	})

	app.Run(iris.Addr(":8080"))

}



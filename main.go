package main

import (
	_ "gf-eshop/boot"
	_ "gf-eshop/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}

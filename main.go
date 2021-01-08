package main

import (
	_ "Branches/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("yiyue")
	beego.Run()
}


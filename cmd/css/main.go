package main

import (
	"github.com/mlctrez/goappnew/goapp/css"
	"github.com/mlctrez/goappnew/goapp/global"
	"log"
)

func main() {

	log.Println("Generate css style file to ", global.AppCss)

	if err := css.Gen(); err != nil {
		panic(err)
	}

	log.Println("Css generated ok")
}

package main

import (
	"log"

	"github.com/Grayson/lko/pkg/lko"
)

func main() {
	app := lko.InitApp()
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

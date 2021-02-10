package main

import (
	"github.com/niconosenzo/devopsapi/pkg/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}

package main

import (
	"github.com/microservice_template/src/core"
	_ "net/http/pprof"
)

func main() {
	app := core.New()
	app.Start()
}

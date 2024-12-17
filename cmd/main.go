package main

import (
	"github.com/veronicashkarova/CalcServer/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}

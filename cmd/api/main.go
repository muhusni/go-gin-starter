package main

import (
	"log"

	"github.com/muhusni/go-gin-starter/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

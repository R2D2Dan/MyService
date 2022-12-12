package main

import (
	"ServiceR2/internal/pkg/app"
	"log"
)

func main() {

	a, err := app.New()
	if err != nil {
		log.Fatal("init app return error: ", err)
	}

	if err := a.Run(); err != nil {
		log.Fatal("service return error: ", err)
	}

}

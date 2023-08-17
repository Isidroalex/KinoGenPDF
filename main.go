package main

import (
	app "Breeding/internal/app"
	"log"
)

func main() {

	application, err := app.New()
	if err != nil {
		log.Fatalln(err)
	}
	if err = application.Run(); err != nil {
		log.Fatalln(err)
	}

}

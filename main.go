// main is the entrypoint for the app
package main

import (
	"errors"
	"log"

	_ "go.uber.org/mock/mockgen/model"
)

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}

func Main() error {
	log.Println("Starting app")
	return errors.New("error")
}

package main

import (
	"go-ddd-example/api"
	"log"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	api.Init()
}

package main

import (
	"log"
	"web-form-go/api"
)

func main() {
	err := api.InitAPI()
	if err != nil {
		log.Fatal("Error initializing server: ", err)
	}

}

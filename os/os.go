package main

import (
	"log"
	"os"
)

func main() {

	if err := os.Mkdir("toniya", 0700); err != nil {
		log.Fatal(err)
	}

}

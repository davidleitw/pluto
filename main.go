package main

// github.com/davidleitw/pluto

import (
	"log"

	"github.com/davidleitw/pluto/pkg/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Println(err)
	}
}

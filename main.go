package main

import (
	"log"
	"main/app"
	"os"
)

func main() {
	ipfas := app.CreateApp()
	err := ipfas.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"cli-app/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting point")

	// Creating the application
	application := app.Generate()
	// Running the application and checking for errors
	if error := application.Run(os.Args); error != nil {
		// If an error occurs, this will stop the app execution
		log.Fatal(error)
	}
}

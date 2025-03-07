package main

import (
	"file/organizer"
	"flag"
	"log"
)

func main() {
	Dirpath := flag.String("path", "", "Used for sorting files in a directory")
	flag.Parse()
	err := organizer.Organize(*Dirpath)
	if err != nil {
		log.Fatalf("Error:/%v", err)
	}
}

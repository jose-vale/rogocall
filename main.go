package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jose-vale/rogocall/rogo"
)

func main() {

	helpPtr := flag.Bool("h", false, "shows usage example")
	flag.Parse()

	if *helpPtr {
		fmt.Println(rogo.Example)
		os.Exit(0)
	}

	rogo.LoadConfig()

	if len(os.Args) != 3 {
		rogo.ErrorWithExample("Insuficient or invalid arguments")
	}

	rogo.Number = os.Args[1]
	rogo.Transcription = os.Args[2]

	if !rogo.IsValidNumber() {
		rogo.ErrorWithExample("Invalid phone number")
	}

	if !rogo.HasCorrectExtension() {
		rogo.ErrorWithExample("Invalid transcription file")
	}

	rogo.SendTranscription()
}

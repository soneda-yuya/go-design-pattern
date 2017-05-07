package main

import (
	"flag"
	"os"
	"log"

	"./shape"
)

func main() {
	var output = flag.String("output", "text", "The output to use between 'text' and 'image' file")
 	flag.Parse()

	activeStrategy, err := shape.NewPrinter(*output)
	if err != nil {
	  log.Fatal(err)
	}

	switch *output {
	case shape.TEXT_STRATEGY:
	  activeStrategy.SetWriter(os.Stdout)
	case shape.IMAGE_STRATEGY:
	  w, err := os.Create("/tmp/image.jpg")
	  if err != nil {
	    log.Fatal("Error opening image")
	  }
	  defer w.Close()
	  activeStrategy.SetWriter(w)
	}

	err = activeStrategy.Print()
	if err != nil {
	  log.Fatal(err)
	}
}

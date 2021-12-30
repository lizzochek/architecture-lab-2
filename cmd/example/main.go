package main

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"strings"

	lab2 "software-architecture-2"
)

var expression = flag.String("e", "", "Expression")
var inputPath = flag.String("f", "", "Path to file containing expression")
var outputPath = flag.String("o", "", "Path to output file")

func main() {
	flag.Parse()

	var reader io.Reader
	var writer *os.File

	if *expression != "" {
		reader = strings.NewReader(*expression)
	} else if *inputPath != "" {
		data, err := ioutil.ReadFile(*inputPath)
		if err != nil {
			os.Stderr.WriteString(err.Error())
			return
		}
		reader = strings.NewReader(string(data))
	} else {
		os.Stderr.WriteString("Expression not provided")
	}

	if *outputPath != "" {
		var err error
		writer, err = os.Create(*outputPath)
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
	} else {
		writer = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Reader: reader,
		Writer: writer,
	}

	err := handler.Compute()
	writer.Close()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}

package main

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"strings"

	lab2 "github.com/inovarka/lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Read the expressions from file")
	resultFile      = flag.String("o", "", "Write result to file")
)

func main() {
	flag.Parse()

	var (
		inp  io.Reader
		outp io.Writer
		err  error
	)

	if *inputFile != "" {
		str, err := ioutil.ReadFile(*inputFile)
		if err != nil {
			os.Stderr.WriteString("unable to open the file")
			return
		}
		inp = strings.NewReader(string(str))
	} else {
		inp = strings.NewReader(*inputExpression)
	}

	if *resultFile != "" {
		outp, err = os.Create(*resultFile)
		if err != nil {
			os.Stderr.WriteString("Unable to create the file")
			return
		}
	} else {
		outp = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  inp,
		Output: outp,
	}

	err = handler.Compute()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

}

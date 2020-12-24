package lab2

import (
	"fmt"
	"io"
	"io/ioutil"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.

//ComputeHandler io.Reader io.Writer necessary struct
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	expr, err := ioutil.ReadAll(ch.Input)
	if err != nil {
		return err
	}

	if res, err := ComputePrefix(string(expr)); err != nil {
		return err
	} else {
		buf := []byte(fmt.Sprintf("%f", res))
		if _, e := ch.Output.Write(buf); e != nil {
			return e
		}
	}

	return nil
}

package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"strings"
)

type outputTest struct {
	called bool
}

func (o *outputTest) Write(p []byte) (n int, err error) {
	o.called = true
	return 0, nil
}

func TestHandlCompute(t *testing.T) {
	exp := "+ 2 3"
	in := strings.NewReader(exp)
	out := outputTest{}

	handler := ComputeHandler{
		Input:  in,
		Output: &out,
	}
	handler.Compute()

	assert.Equal(t, true, out.called)
}

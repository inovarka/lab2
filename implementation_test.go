package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Syntax test
func TestPrefixComputeIncorrect(t *testing.T) {
	// - - syntax error
	_, err := ComputePrefix("-")
	assert.NotNil(t, err)

	// ++ 2 - syntax error
	_, err = ComputePrefix("++ 2")
	assert.NotNil(t, err)

	// 2 + - syntax error
	_, err = ComputePrefix("2 +")
	assert.NotNil(t, err)

	// "" - syntax error
	_, err = ComputePrefix("")
	assert.NotNil(t, err)

	// "zxc" - syntax error
	_, err = ComputePrefix("zxc")
	assert.NotNil(t, err)
}

//Compute test
func TestPrefixComputeCorrect(t *testing.T) {
	// 3 = 3
	res, err := ComputePrefix("3")
	if assert.Nil(t, err) {
		assert.Equal(t, 3.0, res)
	}

	// (7 - 5) + 3 * 8 = 26
	res, err = ComputePrefix("+ * 3 8 - 7 5")
	if assert.Nil(t, err) {
		assert.Equal(t, 26.0, res)
	}

	// 130 - ( 120 * (81 - ( 80 + 2 ^ 2 ) ) )
	res, err = ComputePrefix("- 130 * 120 - 81 + 80 ^ 2 2")
	if assert.Nil(t, err) {
		assert.Equal(t, 490.0, res)
	}

	//(((((2 * 20) / 2) + ((3 + 4) * (3 ^ 2))) - 6) + 15)
	res, err = ComputePrefix("+ - + / * 2 20 2 * + 3 4 ^ 3 2 6 15")
	if assert.Nil(t, err) {
		assert.Equal(t, 92.0, res)
	}
}


 func ExampleComputePrefix() {
	 // example of ComputePrefix
		res, _ := ComputePrefix("^ 10 2")
		fmt.Println(res)
 
	// 100
}

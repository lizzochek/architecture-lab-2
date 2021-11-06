package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToPrefixShort1(t *testing.T) {
	res, err := PostfixToPrefix("A B * C D / +")
	assert.Equal(t, res, "+ * A B / C D")
	assert.Nil(t, err)
}

func TestPostfixToPrefixShort2(t *testing.T) {
	res, err := PostfixToPrefix("A B C + * D /")
	assert.Equal(t, res, "/ * A + B C D")
	assert.Nil(t, err)
}

func TestPostfixToPrefixShort3(t *testing.T) {
	res, err := PostfixToPrefix("A B C D / + *")
	assert.Equal(t, res, "* A + B / C D")
	assert.Nil(t, err)
}

func TestPostfixToPrefixError(t *testing.T) {
	_, err := PostfixToPrefix("A B C D E")
	assert.Error(t, err)
}

func ExamplePostfixToPrefix() {
	res, _ := PostfixToPrefix("5 10 * 2 /")
	fmt.Println(res)
}

package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerValid(t *testing.T) {
	output := new(bytes.Buffer)
	handler := ComputeHandler{
		Reader: strings.NewReader("A B * C D / +"),
		Writer: output,
	}
	err := handler.Compute()
	assert.Nil(t, err)
}

func TestHandlerError(t *testing.T) {
	output := new(bytes.Buffer)
	handler := ComputeHandler{
		Reader: strings.NewReader("* H J L K R Y"),
		Writer: output,
	}
	err := handler.Compute()
	assert.NotNil(t, err)
}

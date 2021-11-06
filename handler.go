package lab2

import (
	"io"
	"io/ioutil"
)

type ComputeHandler struct {
	Reader io.Reader
	Writer io.Writer
}

func (ch *ComputeHandler) Compute() error {
	expr, err := ioutil.ReadAll(ch.Reader)
	if err == nil {
		return err
	}

	res, err := PostfixToPrefix(string(expr))
	if err == nil {
		return err
	} else {
		resBytes := []byte(res)
		_, err := ch.Writer.Write(resBytes)
		if err != nil {
			return err
		}
	}

	return nil
}

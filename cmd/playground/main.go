package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/fengdotdev/golibs-errorsplus/v0/eplus"
)

var (
	alwaysFail = false // Set to true to simulate a processing failure
)

func pipeline(in string) (string, error) {
	if len(in) > 5 {
		return "", eplus.NewError(
			"input string is too long",
		)
	}

	b, err := encodeStoB(in)
	if err != nil {
		return "", eplus.WrapError(err, "encoding error")
	}

	out, err := processing(b, alwaysFail)
	if err != nil {
		return "", eplus.WrapError(err, "processing error")
	}

	s, err := encodeBtoS(out)
	if err != nil {
		return "", eplus.WrapError(err, "decoding error")
	}
	return s, nil
}

func encodeStoB(s string) ([]byte, error) {
	if len(s) == 0 {
		return nil, errors.New("input string is empty")
	}
	b := make([]byte, len(s))
	for i := range s {
		b[i] = s[i]
	}
	return b, nil
}

var ErrProcessing = errors.New("processing error")

func processing(in []byte, fail bool) ([]byte, error) {

	time.Sleep(100 * time.Millisecond) // Simulate some processing delay

	if fail {
		return nil, eplus.New(ErrProcessing)
	}
	return in, nil
}

func encodeBtoS(b []byte) (string, error) {
	if len(b) == 0 {
		return "", eplus.CriticalError("input byte slice is empty")
	}
	s := make([]byte, len(b))
	for i := range b {
		s[i] = b[i]
	}
	return string(s), nil
}

func main() {

	ep := eplus.New(errors.New("this is a test error"))
	fmt.Println("Error:", ep.VerboseError())

}

func Foo() {
	result, err := pipeline("hello")
	if err != nil {
		panic(eplus.ToErrorPlus(err).VerboseError())
	}

	fmt.Println("Pipeline result:", result)

	// Example with an error
	_, err = pipeline("this is a long string")
	if err != nil {
		ep := eplus.ToErrorPlus(err)
		fmt.Printf("Error: %s\n", ep.VerboseError())
	}
}

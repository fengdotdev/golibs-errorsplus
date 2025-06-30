package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/fengdotdev/golibs-errorsplus/sandbox/draft2/errorplus"
)

var (
	alwaysFail = false // Set to true to simulate a processing failure
)

func pipeline(in string) (string, error) {
	if len(in) > 5 {
		return "", errorplus.NewError(
			errors.New("input string is too long"),
			"pipeline error",
			[]string{"pipeline", "input"},
		)
	}

	b, err := encodeStoB(in)
	if err != nil {
		return "", errorplus.NewError(err, "encoding error", []string{"pipeline", "encodeStoB"})
	}

	out, err := processing(b, alwaysFail)
	if err != nil {
		return "", errorplus.NewError(err, "processing error", []string{"pipeline", "processing"})
	}

	s, err := encodeBtoS(out)
	if err != nil {
		return "", errorplus.NewError(err, "decoding error", []string{"pipeline", "encodeBtoS"})
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

func processing(in []byte, fail bool) ([]byte, error) {

	time.Sleep(100 * time.Millisecond) // Simulate some processing delay

	if fail {
		return nil, errorplus.New(errors.New("processing failed"))
	}
	return in, nil
}

func encodeBtoS(b []byte) (string, error) {
	if len(b) == 0 {
		return "", errors.New("input byte slice is empty")
	}
	s := make([]byte, len(b))
	for i := range b {
		s[i] = b[i]
	}
	return string(s), nil
}

func main() {

	// Example usage of the pipeline function
	_, err := pipeline("hello")
	if err != nil {
		ep := errorplus.ToErrorPlus(err)
		fmt.Printf("Error: %s\n", ep.VerboseError())
	}
	//println("Pipeline result:", result)

	// Example with an error
	_, err = pipeline("this is a long string")
	if err != nil {
		ep := errorplus.ToErrorPlus(err)
		fmt.Printf("Error: %s\n", ep.VerboseError())
	}

	alwaysFail = false // Simulate a processing failure
	_, err = pipeline("hello")
	if err != nil {
		ep := errorplus.ToErrorPlus(err)
		fmt.Printf("Error: %s\n", ep.VerboseError())
	}
}

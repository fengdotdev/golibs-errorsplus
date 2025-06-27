package main

import (
	"errors"
	"time"
)

var (
	alwaysFail = false // Set to true to simulate a processing failure
)

func pipeline(in string) (string, error) {
	if len(in) > 5 {
		return "", errors.New("input string is too long")
	}

	b, err := encodeStoB(in)
	if err != nil {
		return "", err
	}

	out, err := processing(b, alwaysFail)
	if err != nil {
		return "", err
	}

	s, err := encodeBtoS(out)
	if err != nil {
		return "", err
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
		return nil, errors.New("processing failed")
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
	result, err := pipeline("hello")
	if err != nil {
		panic(err) // Handle error appropriately in real applications
	}
	println("Pipeline result:", result)

	// Example with an error
	_, err = pipeline("this is a long string")
	if err != nil {
		println("Error:", err.Error())
	}

	alwaysFail = true // Simulate a processing failure
	_, err = pipeline("hello")
	if err != nil {
		println("Error:", err.Error())
	}
}

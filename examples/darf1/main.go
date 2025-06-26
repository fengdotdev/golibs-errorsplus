package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/fengdotdev/golibs-errorsplus/sandbox/draft1/e"
)

func AssertEtoPlus(err error) *e.ErrorPlus {
	var zero e.ErrorPlus

	if err == nil {
		return &zero
	}

	if errPlus, ok := err.(*e.ErrorPlus); ok {
		return errPlus
	}

	return e.E(err, "assert error to ErrorPlus failed", []string{"assert-err"}, AssertEtoPlus)
}

func SomeProcess() (err error) {

	data, err := OpenFile()
	if err != nil {
		return e.E(err, "open file failed", []string{"file-err"}, SomeProcess)
	}

	if len(data) == 0 {
		return e.E(nil, "file is empty", []string{"file-err"}, SomeProcess, "file.txt")
	}

	fmt.Println("file content:", string(data))
	return nil

}

func OpenFile() ([]byte, error) {

	file, err := os.Open("file.txt")
	if err != nil {
		return nil, e.E(err, "open file failed", []string{"file-err"}, OpenFile)
	}

	data, err := os.ReadFile(file.Name())
	if err != nil {
		return nil, e.E(err, "read file failed", []string{"file-err"}, OpenFile, file.Name())
	}
	return data, nil
}

func division(n1, n2 int) (val float64, err error) {

	defer func() {
		if r := recover(); r != nil {
			val = 0
			err = fmt.Errorf("recover: %s", r)
		}
	}()

	if n2 == 0 {
		panic("div by zero")
	}
	return float64(n1) / float64(n2), nil
}

func main() {

	err := SomeProcess()
	if false {
		errPlus := AssertEtoPlus(err)
		fmt.Println("ErrorPlus:", errPlus.VerboseError())
		return
	}
	err2 := errors.New("test error")

	if err != nil {
		errPlu2 := AssertEtoPlus(err2)
		fmt.Println("ErrorPlus1:", errPlu2.VerboseError())
	}

}

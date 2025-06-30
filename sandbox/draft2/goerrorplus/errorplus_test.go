package goerrorplus_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/fengdotdev/golibs-errorsplus/sandbox/draft2/goerrorplus"
)

func TestErrorPlus(t *testing.T) {

	err := errors.New("test error")

	errPlus := goerrorplus.New(err)

	fmt.Printf("ErrorPlus: %s\n", errPlus.Error())
	fmt.Printf("VerboseError: %s\n", errPlus.VerboseError())

}

package eplus_test

import (
	"errors"
	"testing"

	"github.com/fengdotdev/golibs-errorsplus/v0/eplus"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestErrorPlus(t *testing.T) {

	err := eplus.New(errors.New("native error"))

	assert.NotNilWithMessage(t, err, "ErrorPlus should not be nil")

	assert.EqualWithMessage(t, "native error", err.Error(), "ErrorPlus should return the correct error message")

}

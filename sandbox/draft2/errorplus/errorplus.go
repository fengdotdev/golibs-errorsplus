package errorplus

import (
	"github.com/fengdotdev/golibs-errorsplus/sandbox/draft2/goerrorplus"
	"github.com/fengdotdev/golibs-errorsplus/sandbox/draft2/interfaces"
)

func New(err error) interfaces.ErrorPlus {
	if err == nil {
		return nil
	}
	return goerrorplus.New(err)
}

func NewError(err error, message string, tags []string) interfaces.ErrorPlus {
	if err == nil {
		return nil
	}
	return goerrorplus.NewError(err, message, tags)
}

type ErrorPlus = interfaces.ErrorPlus

func IsErrorPlus(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(interfaces.ErrorPlus)
	return ok
}

func ToErrorPlus(err error) interfaces.ErrorPlus {
	if err == nil {
		return nil
	}
	if ep, ok := err.(interfaces.ErrorPlus); ok {
		return ep
	}
	return goerrorplus.New(err)
}

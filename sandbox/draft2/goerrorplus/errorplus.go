package goerrorplus

import (
	"time"

	"github.com/fengdotdev/golibs-errorsplus/sandbox/draft2/interfaces"
)

var _ interfaces.ErrorPlus = (*GoErrorPlus)(nil)

type GoErrorPlus struct {
	err          error
	message      string
	tags         []string
	fN           string
	args         []interface{}
	trace        string
	time         time.Time
	runtimeGoVer string
}




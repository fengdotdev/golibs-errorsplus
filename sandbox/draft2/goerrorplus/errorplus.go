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




// transitive version of error plus
type TErrorPlus struct {
	Err     error  `json:"err"`
	Message string `json:"message"`
	Tags    []string `json:"tags"`
	FN      string `json:"fn"`
	Args    []interface{} `json:"args"`
	Trace   string `json:"trace"`
	Time    time.Time `json:"time"`
	RuntimeGoVer string `json:"runtimeGoVer"`
}

package eplus

import (
	"github.com/fengdotdev/golibs-errorsplus/v0/customtypes/safebool"
	"github.com/fengdotdev/golibs-errorsplus/v0/customtypes/safeint"
	"github.com/fengdotdev/golibs-errorsplus/v0/customtypes/safestring"
)

// safebool is a concurrent safe boolean type

var showTrace = safebool.True()
var showFile = safebool.True()
var showFunc = safebool.True()
var showLine = safebool.True()
var showTimestamp = safebool.True()
var showTags = safebool.True()
var showMessage = safebool.True()
var showFn = safebool.True()
var showFnArgs = safebool.True()
var obscureArgs = safebool.True()

//skip first 3 frames in the trace

var traceSkipFirst = safeint.New(3)

// 15 maximum depth of the trace
var traceMaxDepth = safeint.New(15)

// skip last 2 frames in the trace for runtime.main and runtime.goexit
var traceSkipLast = safeint.New(2)

// Time format used in VerboseError "2006-01-02 15:04:05"
var timeFormat = safestring.New("2006-01-02 15:04:05")

// Separator used in VerboseError  " | "
var separator = safestring.New(" | ")

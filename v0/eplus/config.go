package eplus

import (
	safebolean "github.com/fengdotdev/golibs-errorsplus/v0/customtypes/safeboolean"
	"github.com/fengdotdev/golibs-errorsplus/v0/customtypes/safeint"
)

// safebolean is a concurrent safe boolean type

var showTrace = safebolean.True()
var showFile = safebolean.True()
var showFunc = safebolean.True()
var showLine = safebolean.True()
var showTimestamp = safebolean.True()
var showTags = safebolean.True()
var showMessage = safebolean.True()
var showFn = safebolean.True()
var showFnArgs = safebolean.True()
var obscureArgs = safebolean.True()

//skip first 3 frames in the trace

var traceSkipFirst = safeint.New(3)

// 15 maximum depth of the trace
var traceMaxDepth = safeint.New(15)

// skip last 2 frames in the trace for runtime.main and runtime.goexit
var traceSkipLast = safeint.New(2)

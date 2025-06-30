package goerrorplus

type GlobalOptions struct {
	TimeFormat       string // Time format used in VerboseError
	ShowRuntimeGoVer bool   // Show runtime Go version in VerboseError
	ShowTrace        bool   // Show trace in VerboseError
	ShowArgs         bool   // Show args in VerboseError
	ShowTags         bool   // Show tags in VerboseError
	ShowFunctionName bool   // Show function name in VerboseError
	ShowTimestamp    bool   // Show timestamp in VerboseError
	ShowMessage      bool   // Show message in VerboseError
}

type Options map[string]interface{}

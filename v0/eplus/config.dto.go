package eplus

// Config holds the configuration for the errorplus package.
type Config struct {
	ShowTimestamp bool `json:"show_timestamp"`
	ShowTrace     bool `json:"show_trace"`
	ShowFile      bool `json:"show_file"`
	ShowFunc      bool `json:"show_func"`
	ShowLine      bool `json:"show_line"`
	ShowTags      bool `json:"show_tags"`
	ShowMessage   bool `json:"show_message"`
	ShowFn        bool `json:"show_fn"`
	ShowFnArgs    bool `json:"show_fn_args"`
	ObscureArgs   bool `json:"obscure_args"`
	TraceSkipFirst int  `json:"trace_skip_first"` // skip first n frames in the trace
	TraceMaxDepth  int  `json:"trace_max_depth"`  // maximum depth of the trace
	TraceSkipLast  int  `json:"trace_skip_last"`  // skip last n frames in the trace
}

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
}

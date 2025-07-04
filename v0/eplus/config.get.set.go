package eplus

import "sync"

// UpdateConfig updates the configuration for the errorplus package.
// It sets the values of the boolean flags based on the provided Config struct.
// This function is safe to call concurrently.
// in the config omitted fields are false
// the values of the config are copied to the package variables
// CAREFULL: THIS CHANGES THE GLOBAL STATE OF THE PACKAGE
func UpdateConfig(cfg Config) {

	saveDefaultConfig() // ensure default config is saved

	showTrace.Set(cfg.ShowTrace)
	showFile.Set(cfg.ShowFile)
	showFunc.Set(cfg.ShowFunc)
	showLine.Set(cfg.ShowLine)
	showTimestamp.Set(cfg.ShowTimestamp)
	showTags.Set(cfg.ShowTags)
	showMessage.Set(cfg.ShowMessage)
	showFn.Set(cfg.ShowFn)
	showFnArgs.Set(cfg.ShowFnArgs)
	obscureArgs.Set(cfg.ObscureArgs)
	traceSkipFirst.Set(cfg.TraceSkipFirst)
	traceMaxDepth.Set(cfg.TraceMaxDepth)
	traceSkipLast.Set(cfg.TraceSkipLast)
}

// GetConfig returns the current configuration for the errorplus package.
// It retrieves the values of the boolean flags and returns them in a Config struct.
// This function is safe to call concurrently.
// the config obj is a copy of the current config
func GetConfig() Config {
	saveDefaultConfig() // ensure default config is saved

	return Config{
		ShowTimestamp:  showTimestamp.Get(),
		ShowTrace:      showTrace.Get(),
		ShowFile:       showFile.Get(),
		ShowFunc:       showFunc.Get(),
		ShowLine:       showLine.Get(),
		ShowTags:       showTags.Get(),
		ShowMessage:    showMessage.Get(),
		ShowFn:         showFn.Get(),
		ShowFnArgs:     showFnArgs.Get(),
		ObscureArgs:    obscureArgs.Get(),
		TraceSkipFirst: traceSkipFirst.Get(),
		TraceMaxDepth:  traceMaxDepth.Get(),
		TraceSkipLast:  traceSkipLast.Get(),
	}
}

func GetDefaultConfig() Config {
	saveDefaultConfig() // ensure default config is saved
	return defaultConfig
}

var once sync.Once

func saveDefaultConfig() {

	once.Do(func() {
		conf := Config{
			ShowTimestamp:  showTimestamp.Get(),
			ShowTrace:      showTrace.Get(),
			ShowFile:       showFile.Get(),
			ShowFunc:       showFunc.Get(),
			ShowLine:       showLine.Get(),
			ShowTags:       showTags.Get(),
			ShowMessage:    showMessage.Get(),
			ShowFn:         showFn.Get(),
			ShowFnArgs:     showFnArgs.Get(),
			ObscureArgs:    obscureArgs.Get(),
			TraceSkipFirst: traceSkipFirst.Get(),
			TraceMaxDepth:  traceMaxDepth.Get(),
			TraceSkipLast:  traceSkipLast.Get(),
		}

		defaultConfig = conf

	})

}

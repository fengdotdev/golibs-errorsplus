package eplus

import (
	"sync"

	"github.com/fengdotdev/golibs-errorsplus/v0/customtypes/try"
)

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

	// ensure traceSkipFirst is not negative and in range of traceMaxDepth
	try.If(cfg.TraceSkipFirst >= 0 && cfg.TraceSkipFirst < traceMaxDepth.Get()).Try(func() {
		traceSkipFirst.Set(cfg.TraceSkipFirst)
	})

	// ensure traceMaxDepth is not negative
	try.If(cfg.TraceMaxDepth >= 0).Try(func() {
		traceMaxDepth.Set(cfg.TraceMaxDepth)
	})
	// ensure traceSkipLast is not negative and in range of traceMaxDepth
	try.If(cfg.TraceSkipLast >= 0 && cfg.TraceSkipLast < traceMaxDepth.Get()).Try(func() {
		traceSkipLast.Set(cfg.TraceSkipLast)
	})

	// ensure time format is not empty
	try.If(cfg.TimeFormat != "").Try(func() {
		timeFormat.Set(cfg.TimeFormat)
	})

	// ensure separator is not empty
	try.If(cfg.Separator != "").Try(func() {
		separator.Set(cfg.Separator)
	})
}

// GetConfig returns the current configuration for the errorplus package.
// It retrieves the values of the boolean flags and returns them in a Config struct.
// This function is safe to call concurrently.
// the config obj is a copy of the current config
func GetConfig() Config {
	saveDefaultConfig() // ensure default config is saved
	return getConfig()
}

func getConfig() Config {
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
		TimeFormat:     timeFormat.Get(),
		Separator:      separator.Get(),
	}
}

func GetDefaultConfig() Config {
	saveDefaultConfig() // ensure default config is saved
	return defaultConfig
}

var once sync.Once

func saveDefaultConfig() {

	once.Do(func() {
		conf := getConfig() // get the current config

		defaultConfig = conf

	})

}

func ResetConfig() {

	showTrace.Reset()
	showFile.Reset()
	showFunc.Reset()
	showLine.Reset()
	showTimestamp.Reset()
	showTags.Reset()
	showMessage.Reset()
	showFn.Reset()
	showFnArgs.Reset()
	obscureArgs.Reset()
	traceSkipFirst.Reset()
	traceMaxDepth.Reset()
	traceSkipLast.Reset()
	timeFormat.Reset()
	separator.Reset()

}

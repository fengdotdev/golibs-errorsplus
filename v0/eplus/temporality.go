package eplus

type Temporality int

const (
	// Permanent indicates an error that is unlikely to be resolved.
	Permanent Temporality = iota
	// Temporary indicates an error that may be resolved in the future.
	Temporary
)

func OnTemporality(errorplus error, retryFn func() error) error {

	panic("OnTemporality is not implemented yet")
	return nil
}

func OverrideErrorPlusWithErr(errorplus error, newError error) error {
	panic("OverrideErrorPlusWithErr is not implemented yet")
}

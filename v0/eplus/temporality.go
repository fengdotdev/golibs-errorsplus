package eplus

type Temporality bool

const (
	// Permanent indicates an error that is unlikely to be resolved.
	Permanent Temporality = false // default temporality
	// Temporary indicates an error that may be resolved in the future.
	Temporary Temporality = true
)

func OnTemporality(errorplus error, retryFn func() error) error {

	panic("OnTemporality is not implemented yet")
	return nil
}

func OverrideErrorPlusWithErr(errorplus error, newError error) error {
	panic("OverrideErrorPlusWithErr is not implemented yet")
}

package errorplus

type Temporality int

const (
	// Permanent indicates an error that is unlikely to be resolved.
	Permanent Temporality = iota
	// Temporary indicates an error that may be resolved in the future.
	Temporary
)

func OnTemporality(errorplus error, retryFn func() error) error {

	// note to ia the return err is from the retryFn not the original error

	if errorplus == nil {
		return nil
	}
	// exec retry function only if the error is temporary
	if gerr, ok := errorplus.(*GoError); ok {

	}

}

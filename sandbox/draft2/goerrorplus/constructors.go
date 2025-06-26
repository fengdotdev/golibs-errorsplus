package goerrorplus

func New(err error) *GoErrorPlus {
	if err == nil {
		return nil
	}
	return &GoErrorPlus{
		err: err,
	}
}

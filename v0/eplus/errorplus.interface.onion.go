package eplus

type OnionError interface {
	Surface() error
	Origin() error
	Wrap(err error) error
	Unwrap() error
	Deep() int
	DeepLevel(int)
}

type OnOnionError interface {
	IsOrigin() bool
	IsSurface() bool
	HasChild() bool
}

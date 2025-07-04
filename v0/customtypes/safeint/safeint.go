package safeint

type SafeInt interface {
	String() string
	Get() int
	Value() int
	GetDefault() int // returns the default value
	Set(value int)
	Increment()
	Decrement()
	Reset() // to default value not 0
}

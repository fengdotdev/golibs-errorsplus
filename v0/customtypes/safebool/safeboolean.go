package safebool

type SafeBoolean interface {
	String() string
	Get() bool
	GetDefault() bool // returns the default value

	Value() bool
	Set(value bool)
	Toggle()
	Reset() // to default value not false necessarily
}

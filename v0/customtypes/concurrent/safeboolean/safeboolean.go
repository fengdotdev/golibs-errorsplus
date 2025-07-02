package safebolean

type SafeBoolean interface {
	IsTrue() bool
	IsFalse() bool
	String() string
	Get() bool
	Value() bool
	Set(value bool)
	Toggle()
}

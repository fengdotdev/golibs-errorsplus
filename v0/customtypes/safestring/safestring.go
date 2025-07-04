package safestring

type SafeString interface { // String returns the string representation of the SafeString.
	String() string
	Get() string        // Get returns the current value of the SafeString.
	GetDefault() string // GetDefault returns the default value of the SafeString.
	Set(value string)   // Set sets the value of the SafeString.
	Reset()             // Reset resets the SafeString to its default value.not ""
}

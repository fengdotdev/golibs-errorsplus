package safebolean

func New() *GoBoolean {
	return &GoBoolean{
		value: false,
	}
}

func NewValue(value bool) *GoBoolean {
	return &GoBoolean{
		value: value,
	}
}
func True() *GoBoolean {
	return &GoBoolean{
		value: true,
	}
}

func False() *GoBoolean {
	return &GoBoolean{
		value: false,
	}
}

package ontry

type GoTry struct {
	condition bool
}

func (gt *GoTry) Try(fn func()) {
	if gt.condition {
		fn()
	}
}
func If(condition bool) *GoTry {
	return &GoTry{
		condition: condition,
	}
}

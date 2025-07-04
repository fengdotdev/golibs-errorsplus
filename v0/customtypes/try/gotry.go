package try

var _ IfTry = (*GoTry)(nil)

type GoTry struct {
	condition bool
}

func (gt *GoTry) Try(fn func()) {
	if gt.condition {
		fn()
	}
}

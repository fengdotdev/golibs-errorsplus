package try

func If(condition bool) IfTry {
	return &GoTry{
		condition: condition,
	}
}

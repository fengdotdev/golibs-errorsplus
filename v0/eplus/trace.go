package eplus

import "fmt"

type Trace struct {
	File     string
	Function string
	Line     int
}

func (t *Trace) String() string {
	return "func: " + t.Function + " at: " + t.File + " line: " + t.LineStr()
}

func (t *Trace) LineStr() string {
	return fmt.Sprintf("%d", t.Line)
}

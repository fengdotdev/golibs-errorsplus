package try_test

import (
	"testing"

	"github.com/fengdotdev/golibs-errorsplus/v0/customtypes/try"
)

func TestTry(t *testing.T) {

	t.Run("IfTry", func(t *testing.T) {

		try.If(false).Try(func() {
			t.Error("This should not be executed")
		})

		try.If(true).Try(func() {
			t.Log("This should be executed")
		})
	})
}

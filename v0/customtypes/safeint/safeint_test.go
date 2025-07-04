package safeint_test

import (
	"testing"

	"github.com/fengdotdev/golibs-errorsplus/v0/customtypes/safeint"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestSafeInt(t *testing.T) {

	t.Run("concurrent access", func(t *testing.T) {
		t.Parallel()

		si:= safeint.New(5)

		done := make(chan struct{})

		for i := 0; i < 100; i++ {
			go func(si safeint.SafeInt) {
				defer func() { done <- struct{}{} }()
				si.Increment()
			}(si)
		}

		for i := 0; i < 100; i++ {
			<-done
		}

		assert.EqualWithMessage(t, si.Get(), 105, "expected SafeInt to be 105 after 100 increments")

	})

}

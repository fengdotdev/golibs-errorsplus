package safebool_test

import (
	"testing"

	"github.com/fengdotdev/golibs-errorsplus/v0/customtypes/safebool"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestSafeBoolean(t *testing.T) {
	t.Run("concurrent access", func(t *testing.T) {
		t.Parallel()

		sb := safebool.True()

		done := make(chan struct{})

		for i := 0; i < 100; i++ {
			go func(sb safebool.SafeBoolean) {
				defer func() { done <- struct{}{} }()
				sb.Set(false)
			}(sb)
		}

		for i := 0; i < 100; i++ {
			<-done
		}

		assert.EqualWithMessage(t, sb.Get(), false, "expected SafeBoolean to be false after 100 concurrent sets")
	})
}

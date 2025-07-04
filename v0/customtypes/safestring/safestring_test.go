package safestring_test

import (
	"testing"

	"github.com/fengdotdev/golibs-errorsplus/v0/customtypes/safestring"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestSafeString(t *testing.T) {

	t.Run("concurrent access", func(t *testing.T) {
		t.Parallel()

		ss := safestring.New("test")

		done := make(chan struct{})

		for i := 0; i < 100; i++ {
			go func(ss safestring.SafeString) {
				defer func() { done <- struct{}{} }()
				ss.Set("test")
			}(ss)
		}

		for i := 0; i < 100; i++ {
			<-done
		}

		assert.EqualWithMessage(t, ss.Get(), "test", " expected SafeString to be 'test' after 100 concurrent sets")
	})
}

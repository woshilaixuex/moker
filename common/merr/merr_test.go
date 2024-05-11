package merr

import (
	"github.com/pkg/errors"
	"testing"
)

func TestCause(t *testing.T) {
	innerMerr := NewCauser(123, "inner error")
	merr := NewMError(456, "outer error", innerMerr)
	// Now you can use errors.Cause to unwrap MError and access the inner MError
	cause := errors.Cause(merr)
	t.Log(cause.Error())
}

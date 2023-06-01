package shims

import (
	"os"
	"testing"
)

func Test_SingleValueDiscardError(t *testing.T) {
	t.Logf("%#v", SingleValueDiscardError(os.Open("/tmp/test")))
}

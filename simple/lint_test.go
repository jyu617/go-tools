package simple

import (
	"testing"

	"github.com/jyu617/go-tools/lint/testutil"
)

func TestAll(t *testing.T) {
	testutil.TestAll(t, NewChecker(), "")
}

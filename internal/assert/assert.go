package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, excepted T) {
	t.Helper()

	if actual != excepted {
		t.Errorf("got: %v; want: %v", actual, excepted)
	}
}

func StringContains(t *testing.T, actual, exceptedSubstring string) {
	t.Helper()

	if !strings.Contains(actual, exceptedSubstring) {
		t.Errorf("got: %q; excepted to contain: %q", actual, exceptedSubstring)
	}
}

func NilError(t *testing.T, actual error) {
	t.Helper()

	if actual != nil {
		t.Errorf("got: %v; expected: nil", actual)
	}
}

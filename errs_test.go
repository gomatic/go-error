package errs_test

import (
	"errors"
	"fmt"
	"testing"

	errs "github.com/skykernel/go-error"
	"github.com/stretchr/testify/assert"
)

// Consumer-side sentinels: the library ships no values, so the test declares its
// own exactly as a real consumer would.
const (
	errSentinel errs.Error = "sentinel failed"
	errOther    errs.Error = "other"
)

func TestErrorImplementsError(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	want.Equal("sentinel failed", errSentinel.Error())
	// Usable anywhere an error is expected, and still matchable once wrapped.
	want.EqualError(fmt.Errorf("%w", errSentinel), "sentinel failed")
}

func TestErrorIs(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	want.ErrorIs(fmt.Errorf("%w: %q", errSentinel, "x"), errSentinel)
	want.NotErrorIs(errOther, errSentinel)
}

func TestErrorWith(t *testing.T) {
	t.Parallel()
	cause := errors.New("disk full")

	tests := []struct {
		name        string
		cause       error
		args        []any
		wantIs      []error
		wantMessage string
	}{
		{
			name:        "sentinel only",
			cause:       nil,
			args:        nil,
			wantIs:      []error{errSentinel},
			wantMessage: "sentinel failed",
		},
		{
			name:        "wraps cause",
			cause:       cause,
			args:        nil,
			wantIs:      []error{errSentinel, cause},
			wantMessage: "sentinel failed: disk full",
		},
		{
			name:        "appends args",
			cause:       nil,
			args:        []any{"key", "app.name"},
			wantIs:      []error{errSentinel},
			wantMessage: "sentinel failed: key app.name",
		},
		{
			name:        "wraps cause and appends args",
			cause:       cause,
			args:        []any{"path", "/tmp/x"},
			wantIs:      []error{errSentinel, cause},
			wantMessage: "sentinel failed: disk full: path /tmp/x",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := assert.New(t)
			err := errSentinel.With(tt.cause, tt.args...)
			for _, target := range tt.wantIs {
				want.ErrorIs(err, target)
			}
			want.Equal(tt.wantMessage, err.Error())
		})
	}
}

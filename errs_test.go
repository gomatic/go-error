package errs_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	errs "github.com/gomatic/go-error"
)

// Consumer-side sentinels: the library ships no values, so the test declares its
// own exactly as a real consumer would.
const (
	errSentinel errs.Const = "sentinel failed"
	errOther    errs.Const = "other"
)

func TestErrorImplementsError(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	want.Equal("sentinel failed", errSentinel.Error())
	want.EqualError(fmt.Errorf("%w", errSentinel), "sentinel failed")
}

func TestErrorIs(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	want.ErrorIs(fmt.Errorf("%w: %q", errSentinel, "x"), errSentinel)
	want.NotErrorIs(errOther, errSentinel)
}

func TestConstWith(t *testing.T) {
	t.Parallel()
	cause := errors.New("disk full")

	tests := []struct {
		cause       error
		name        string
		wantMessage string
		args        []any
		wantIs      []error
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

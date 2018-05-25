package pointer

import (
	"context"

	"github.com/pkg/errors"
)

type ctxKeytypePointer int

const (
	ctxKeyPointer ctxKeytypePointer = iota
)

// WithContext -
func WithContext(ctx context.Context, rs ...ReferenceToken) context.Context {
	ptr := Empty
	if len(rs) > 0 {
		ptr = FromContext(ctx)
		ptr = Append(ptr, rs...)
	}
	return context.WithValue(ctx, ctxKeyPointer, ptr)
}

// FromContext -
func FromContext(ctx context.Context) Pointer {
	ret := Empty
	if val := ctx.Value(ctxKeyPointer); val != nil {
		ret = val.(Pointer)
	}
	return ret
}

// WrapError -
func WrapError(ctx context.Context, err error) error {
	if ptr := FromContext(ctx); ptr.ReferenceToken() != nil {
		err = errors.Wrap(err, ptr.String())
	}
	return err
}

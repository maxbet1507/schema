package patch

import (
	"context"
	"fmt"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
)

// -
var (
	ErrInvalidPointer  = fmt.Errorf("Invalid Pointer")
	ErrInvalidJSONType = fmt.Errorf("Invalid JSON Type")
)

type ctxRefer struct {
	left []pointer.ReferenceToken
}

func (s *ctxRefer) Refer(ctx context.Context, raw []byte) ([]byte, error) {
	if len(s.left) == 0 {
		return raw, nil
	}

	ctx = pointer.WithContext(ctx, s.left[0])
	s.left = s.left[1:]

	aux, err := types.Unmarshal(raw)
	if err != nil {
		return nil, pointer.WrapError(ctx, err)
	}

	switch aux.(type) {
	case types.Object:
		aux := aux.(types.Object)
		key := pointer.FromContext(ctx).ReferenceToken().MemberName()
		val, ok := aux[key]
		if !ok {
			return nil, pointer.WrapError(ctx, ErrInvalidPointer)
		}

		return s.Refer(ctx, val)

	case types.Array:
		aux := aux.(types.Array)
		idx, err := pointer.FromContext(ctx).ReferenceToken().ArrayIndex()
		if err != nil {
			return nil, pointer.WrapError(ctx, ErrInvalidPointer)
		}
		if idx < 0 || len(aux) <= idx {
			return nil, pointer.WrapError(ctx, ErrInvalidPointer)
		}

		return s.Refer(ctx, aux[idx])

	default:
		return nil, pointer.WrapError(ctx, ErrInvalidJSONType)
	}
}

func rawRefer(ptr pointer.Pointer, raw []byte) ([]byte, error) {
	ctxRefer := ctxRefer{
		left: ptr.ReferenceTokens(),
	}

	return ctxRefer.Refer(context.Background(), raw)
}

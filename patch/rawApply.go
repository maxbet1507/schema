package patch

import (
	"context"
	"encoding/json"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
)

type ctxApply struct {
	left []pointer.ReferenceToken
	last pointer.ReferenceToken
	fn   func(pointer.ReferenceToken, []byte) ([]byte, error)
}

func (s *ctxApply) Apply(ctx context.Context, raw []byte) ([]byte, error) {
	if len(s.left) == 0 {
		ret, err := s.fn(s.last, raw)
		return ret, pointer.WrapError(ctx, err)
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

		ret, err := s.Apply(ctx, val)
		if err != nil {
			return nil, err
		}
		aux[key] = ret

		return json.Marshal(aux)

	case types.Array:
		aux := aux.(types.Array)
		idx, err := pointer.FromContext(ctx).ReferenceToken().ArrayIndex()
		if err != nil {
			return nil, pointer.WrapError(ctx, ErrInvalidPointer)
		}
		if idx < 0 || len(aux) <= idx {
			return nil, pointer.WrapError(ctx, ErrInvalidPointer)
		}

		ret, err := s.Apply(ctx, aux[idx])
		if err != nil {
			return nil, err
		}
		aux[idx] = ret

		return json.Marshal(aux)

	default:
		return nil, pointer.WrapError(ctx, ErrInvalidJSONType)
	}
}

type rawApplyFunc func(pointer.ReferenceToken, []byte) ([]byte, error)

func rawApply(ptr pointer.Pointer, raw []byte, fn rawApplyFunc) ([]byte, error) {
	refs := ptr.ReferenceTokens()
	if len(refs) < 1 {
		return nil, ErrInvalidPointer
	}

	ctxApply := ctxApply{
		left: refs[:len(refs)-1],
		last: refs[len(refs)-1],
		fn:   fn,
	}

	return ctxApply.Apply(context.Background(), raw)
}

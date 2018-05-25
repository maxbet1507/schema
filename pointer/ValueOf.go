package pointer

import (
	"context"
	"fmt"

	"github.com/maxbet1507/schema/types"
)

// -
var (
	ErrInvalidJSONType = fmt.Errorf("Invalid JSON Type")
)

type ctxValueOf struct {
	Pointers []Pointer
}

func (s *ctxValueOf) ValueOf(ctx context.Context, raw []byte) error {
	path := FromContext(ctx)

	aux, err := types.Unmarshal(raw)
	if err != nil {
		return WrapError(ctx, err)
	}

	switch aux.(type) {
	case types.Object:
		if path.ReferenceToken() != nil {
			s.Pointers = append(s.Pointers, path)
		}

		for key, raw := range aux.(types.Object) {
			sub := WithContext(ctx, MemberNameOf(key))
			s.ValueOf(sub, raw) // always success
		}

		return nil

	case types.Array:
		if path.ReferenceToken() != nil {
			s.Pointers = append(s.Pointers, path)
		}

		for idx, raw := range aux.(types.Array) {
			sub := WithContext(ctx, ArrayIndexOf(idx))
			s.ValueOf(sub, raw) // always success
		}

		return nil

	default:
		if path.ReferenceToken() == nil {
			return ErrInvalidJSONType
		}

		s.Pointers = append(s.Pointers, path)
		return nil
	}
}

// ValueOf -
func ValueOf(raw []byte) ([]Pointer, error) {
	ctxValueOf := ctxValueOf{}
	if err := ctxValueOf.ValueOf(context.Background(), raw); err != nil {
		return nil, err
	}
	return ctxValueOf.Pointers, nil
}

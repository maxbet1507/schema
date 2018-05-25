package patch

import (
	"context"
	"reflect"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
)

type ctxDiff struct {
	ret Patches
}

func (s *ctxDiff) DiffObject(ctx context.Context, aux1, aux2 types.Object) error {
	keys := map[string]struct{}{}
	for key := range aux1 {
		keys[key] = struct{}{}
	}
	for key := range aux2 {
		keys[key] = struct{}{}
	}

	for key := range keys {
		ctx := pointer.WithContext(ctx, pointer.MemberNameOf(key))

		if _, ok := aux1[key]; !ok {
			val, _ := Add(pointer.FromContext(ctx).String(), aux2[key]) // always success
			s.ret = append(s.ret, val)
			continue
		}

		if _, ok := aux2[key]; !ok {
			val, _ := Remove(pointer.FromContext(ctx).String()) // always success
			s.ret = append(s.ret, val)
			continue
		}

		s.Diff(ctx, aux1[key], aux2[key]) // always success
	}

	return nil
}

func (s *ctxDiff) DiffArray(ctx context.Context, aux1, aux2 types.Array) error {
	for i := 0; i < len(aux1) || i < len(aux2); i++ {
		ctx := pointer.WithContext(ctx, pointer.ArrayIndexOf(i))

		if i >= len(aux1) {
			val, _ := Add(pointer.FromContext(ctx).String(), aux2[i]) // always success
			s.ret = append(s.ret, val)
			continue
		}

		if i >= len(aux2) {
			val, _ := Remove(pointer.FromContext(ctx).String()) // always success
			s.ret = append(s.ret, val)
			continue
		}

		s.Diff(ctx, aux1[i], aux2[i]) // always success
	}

	return nil
}

func (s *ctxDiff) DiffValue(ctx context.Context, aux1, aux2 interface{}) error {
	if pointer.FromContext(ctx).ReferenceToken() == nil {
		return ErrInvalidJSONType
	}

	if reflect.DeepEqual(aux1, aux2) {
		return nil
	}

	val, _ := Replace(pointer.FromContext(ctx).String(), aux2) // always success
	s.ret = append(s.ret, val)
	return nil
}

func (s *ctxDiff) Diff(ctx context.Context, raw1, raw2 []byte) error {
	aux1, err := types.Unmarshal(raw1)
	if err != nil {
		return err
	}

	aux2, err := types.Unmarshal(raw2)
	if err != nil {
		return err
	}

	switch aux1.(type) {
	case types.Object:
		if aux2, ok := aux2.(types.Object); ok {
			return s.DiffObject(ctx, aux1.(types.Object), aux2)
		}

	case types.Array:
		if aux2, ok := aux2.(types.Array); ok {
			return s.DiffArray(ctx, aux1.(types.Array), aux2)
		}
	}

	return s.DiffValue(ctx, aux1, aux2)
}

// Diff -
func Diff(v1, v2 []byte) (Patches, error) {
	ctxDiff := ctxDiff{
		ret: Patches{},
	}

	if err := ctxDiff.Diff(context.Background(), v1, v2); err != nil {
		return nil, err
	}
	return ctxDiff.ret, nil
}

package patch

import (
	"encoding/json"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
)

type rawRemove struct {
	Path pointer.Pointer
}

func (s *rawRemove) fnApply(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
	aux, err := types.Unmarshal(raw)
	if err != nil {
		return nil, err
	}

	switch aux.(type) {
	case types.Object:
		aux := aux.(types.Object)
		delete(aux, ref.MemberName())

		return json.Marshal(aux)

	case types.Array:
		aux := aux.(types.Array)
		idx, err := ref.ArrayIndex()
		if err != nil {
			return nil, err
		}

		if idx < 0 || len(aux) <= idx {
			return nil, ErrInvalidPointer
		}

		ret := append(aux[:idx], aux[idx+1:]...)

		return json.Marshal(ret)

	default:
		return nil, ErrInvalidJSONType
	}
}

func (s *rawRemove) Apply(raw []byte) ([]byte, error) {
	return rawApply(s.Path, raw, s.fnApply)
}

func (s *rawRemove) MarshalJSON() ([]byte, error) {
	aux := struct {
		Op   string `json:"op"`
		Path string `json:"path"`
	}{
		Op:   "remove",
		Path: s.Path.String(),
	}

	return json.Marshal(aux)
}

// Remove -
func Remove(path string) (Patch, error) {
	var err error

	ret := &rawRemove{}
	if ret.Path, err = pointer.StringOf(path); err != nil {
		return nil, err
	}

	return ret, nil
}

package patch

import (
	"encoding/json"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
)

type rawAdd struct {
	Path  pointer.Pointer
	Value json.RawMessage
}

func (s *rawAdd) fnApply(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
	aux, err := types.Unmarshal(raw)
	if err != nil {
		return nil, err
	}

	switch aux.(type) {
	case types.Object:
		aux := aux.(types.Object)
		aux[ref.MemberName()] = s.Value

		return json.Marshal(aux)

	case types.Array:
		aux := aux.(types.Array)

		idx, err := ref.ArrayIndex()
		if err != nil {
			return nil, err
		}

		if idx > len(aux) {
			return nil, ErrInvalidPointer
		}

		if idx < 0 {
			idx = len(aux)
		}

		ret := make([]json.RawMessage, len(aux)+1)
		copy(ret, aux[:idx])
		copy(ret[idx+1:], aux[idx:])
		ret[idx] = s.Value

		return json.Marshal(ret)

	default:
		return nil, ErrInvalidJSONType
	}
}

func (s *rawAdd) Apply(raw []byte) ([]byte, error) {
	return rawApply(s.Path, raw, s.fnApply)
}

func (s *rawAdd) MarshalJSON() ([]byte, error) {
	aux := struct {
		Op    string          `json:"op"`
		Path  string          `json:"path"`
		Value json.RawMessage `json:"value"`
	}{
		Op:    "add",
		Path:  s.Path.String(),
		Value: s.Value,
	}

	return json.Marshal(aux)
}

// Add -
func Add(path string, value interface{}) (Patch, error) {
	var err error

	ret := &rawAdd{}
	if ret.Path, err = pointer.StringOf(path); err != nil {
		return nil, err
	}

	if ret.Value, err = json.Marshal(value); err != nil {
		return nil, err
	}

	return ret, nil
}

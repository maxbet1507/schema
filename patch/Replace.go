package patch

import (
	"encoding/json"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
)

type rawReplace struct {
	Path  pointer.Pointer
	Value json.RawMessage
}

func (s *rawReplace) fnApply(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
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

		if idx < 0 || len(aux) <= idx {
			return nil, ErrInvalidPointer
		}
		aux[idx] = s.Value

		return json.Marshal(aux)

	default:
		return nil, ErrInvalidJSONType
	}
}

func (s *rawReplace) Apply(raw []byte) ([]byte, error) {
	return rawApply(s.Path, raw, s.fnApply)
}

func (s *rawReplace) MarshalJSON() ([]byte, error) {
	aux := struct {
		Op    string          `json:"op"`
		Path  string          `json:"path"`
		Value json.RawMessage `json:"value"`
	}{
		Op:    "replace",
		Path:  s.Path.String(),
		Value: s.Value,
	}

	return json.Marshal(aux)
}

// Replace -
func Replace(path string, value interface{}) (Patch, error) {
	var err error

	ret := &rawReplace{}
	if ret.Path, err = pointer.StringOf(path); err != nil {
		return nil, err
	}

	if ret.Value, err = json.Marshal(value); err != nil {
		return nil, err
	}

	return ret, nil
}

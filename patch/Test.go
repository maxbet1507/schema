package patch

import (
	"encoding/json"
	"fmt"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
)

// -
var (
	ErrTestNotEqual = fmt.Errorf("Not Equal")
)

type rawTest struct {
	Path  pointer.Pointer
	Value json.RawMessage
}

func (s *rawTest) Apply(raw []byte) ([]byte, error) {
	value, err := rawRefer(s.Path, raw)
	if err != nil {
		return nil, err
	}

	if !types.DeepEqual(s.Value, value) {
		return nil, ErrTestNotEqual
	}
	return raw, nil
}

func (s *rawTest) MarshalJSON() ([]byte, error) {
	aux := struct {
		Op    string          `json:"op"`
		Path  string          `json:"path"`
		Value json.RawMessage `json:"value"`
	}{
		Op:    "test",
		Path:  s.Path.String(),
		Value: s.Value,
	}

	return json.Marshal(aux)
}

// Test -
func Test(path string, value interface{}) (Patch, error) {
	var err error

	ret := &rawTest{}
	if ret.Path, err = pointer.StringOf(path); err != nil {
		return nil, err
	}

	if ret.Value, err = json.Marshal(value); err != nil {
		return nil, err
	}

	return ret, nil
}

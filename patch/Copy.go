package patch

import (
	"encoding/json"

	"github.com/maxbet1507/schema/pointer"
)

type rawCopy struct {
	Path pointer.Pointer
	From pointer.Pointer
}

func (s *rawCopy) Apply(raw []byte) ([]byte, error) {
	value, err := rawRefer(s.From, raw)
	if err != nil {
		return nil, err
	}

	rawAdd := rawAdd{
		Path:  s.Path,
		Value: value,
	}
	return rawAdd.Apply(raw)
}

func (s *rawCopy) MarshalJSON() ([]byte, error) {
	aux := struct {
		Op   string `json:"op"`
		Path string `json:"path"`
		From string `json:"from"`
	}{
		Op:   "copy",
		Path: s.Path.String(),
		From: s.From.String(),
	}

	return json.Marshal(aux)
}

// Copy -
func Copy(path, from string) (Patch, error) {
	var err error

	ret := &rawCopy{}
	if ret.Path, err = pointer.StringOf(path); err != nil {
		return nil, err
	}

	if ret.From, err = pointer.StringOf(from); err != nil {
		return nil, err
	}

	return ret, nil
}

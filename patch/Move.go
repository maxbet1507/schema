package patch

import (
	"encoding/json"

	"github.com/maxbet1507/schema/pointer"
)

type rawMove struct {
	Path pointer.Pointer
	From pointer.Pointer
}

func (s *rawMove) Apply(raw []byte) ([]byte, error) {
	value, err := rawRefer(s.From, raw)
	if err != nil {
		return nil, err
	}

	rawRemove := rawRemove{
		Path: s.From,
	}
	raw, _ = rawRemove.Apply(raw) // always success

	rawAdd := rawAdd{
		Path:  s.Path,
		Value: value,
	}
	return rawAdd.Apply(raw)
}

func (s *rawMove) MarshalJSON() ([]byte, error) {
	aux := struct {
		Op   string `json:"op"`
		Path string `json:"path"`
		From string `json:"from"`
	}{
		Op:   "move",
		Path: s.Path.String(),
		From: s.From.String(),
	}

	return json.Marshal(aux)
}

// Move -
func Move(path, from string) (Patch, error) {
	var err error

	ret := &rawMove{}
	if ret.Path, err = pointer.StringOf(path); err != nil {
		return nil, err
	}

	if ret.From, err = pointer.StringOf(from); err != nil {
		return nil, err
	}

	return ret, nil
}

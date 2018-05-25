package patch

import (
	"encoding/json"
	"fmt"
)

// -
var (
	ErrInvalidJSONPatch = fmt.Errorf("ErrInvalidJSONPatch")
)

// Patch -
type Patch interface {
	json.Marshaler
	Apply([]byte) ([]byte, error)
}

// BytesOf -
func BytesOf(raw []byte) (Patch, error) {
	aux := struct {
		Op    *string         `json:"op"`
		Path  *string         `json:"path"`
		From  *string         `json:"from"`
		Value json.RawMessage `json:"value"`
	}{}

	if err := json.Unmarshal(raw, &aux); err != nil {
		return nil, err
	}

	var value interface{}
	if aux.Value != nil {
		json.Unmarshal(aux.Value, &value) // always success
	}

	if aux.Op != nil {
		switch *aux.Op {
		case "add":
			if aux.Path != nil && aux.Value != nil {
				return Add(*aux.Path, value)
			}
		case "remove":
			if aux.Path != nil {
				return Remove(*aux.Path)
			}
		case "replace":
			if aux.Path != nil && aux.Value != nil {
				return Replace(*aux.Path, value)
			}
		case "move":
			if aux.Path != nil && aux.From != nil {
				return Move(*aux.Path, *aux.From)
			}
		case "copy":
			if aux.Path != nil && aux.From != nil {
				return Copy(*aux.Path, *aux.From)
			}
		case "test":
			if aux.Path != nil && aux.Value != nil {
				return Test(*aux.Path, value)
			}
		}
	}

	return nil, ErrInvalidJSONPatch
}

// Patches -
type Patches []Patch

// Apply -
func (s *Patches) Apply(raw json.RawMessage) (json.RawMessage, error) {
	for _, p := range *s {
		var err error

		raw, err = p.Apply(raw)
		if err != nil {
			return nil, err
		}
	}

	return raw, nil
}

// UnmarshalJSON -
func (s *Patches) UnmarshalJSON(data []byte) error {
	raws := []json.RawMessage{}
	if err := json.Unmarshal(data, &raws); err != nil {
		return err
	}

	ret := []Patch{}
	for _, raw := range raws {
		val, err := BytesOf(raw)
		if err != nil {
			return err
		}
		ret = append(ret, val)
	}

	*s = ret
	return nil
}

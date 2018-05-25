package types

import (
	"encoding/json"
	"reflect"
)

// DeepEqual -
func DeepEqual(v1, v2 []byte) bool {
	var a1, a2 interface{}

	if err := json.Unmarshal(v1, &a1); err != nil {
		return false
	}
	if err := json.Unmarshal(v2, &a2); err != nil {
		return false
	}

	return reflect.DeepEqual(a1, a2)
}

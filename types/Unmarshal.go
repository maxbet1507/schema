package types

import (
	"encoding/json"
	"fmt"
)

// -
var (
	ErrUnknownType = fmt.Errorf("unknown")
)

// Object -
type Object map[string]json.RawMessage

func asObject(v []byte) (interface{}, error) {
	var ret Object
	err := json.Unmarshal(v, &ret)
	return ret, err
}

// Array -
type Array []json.RawMessage

func asArray(v []byte) (interface{}, error) {
	var ret Array
	err := json.Unmarshal(v, &ret)
	return ret, err
}

// Number -
type Number float64

func asNumber(v []byte) (interface{}, error) {
	var ret Number
	err := json.Unmarshal(v, &ret)
	return ret, err
}

// Boolean -
type Boolean bool

func asBoolean(v []byte) (interface{}, error) {
	var ret Boolean
	err := json.Unmarshal(v, &ret)
	return ret, err
}

// String -
type String string

func asString(v []byte) (interface{}, error) {
	var ret String
	err := json.Unmarshal(v, &ret)
	return ret, err
}

// Null -
type Null struct{}

func asNull(v []byte) (interface{}, error) {
	var ret Null
	err := json.Unmarshal(v, &ret)
	return ret, err
}

var (
	convs = map[byte]func([]byte) (interface{}, error){
		'{': asObject,
		'[': asArray,
		'"': asString,
		't': asBoolean,
		'f': asBoolean,
		'n': asNull,
		'-': asNumber,
		'0': asNumber,
		'1': asNumber,
		'2': asNumber,
		'3': asNumber,
		'4': asNumber,
		'5': asNumber,
		'6': asNumber,
		'7': asNumber,
		'8': asNumber,
		'9': asNumber,
	}
	noops = map[byte]struct{}{
		' ':  struct{}{},
		'\n': struct{}{},
		'\r': struct{}{},
		'\t': struct{}{},
	}
)

// Unmarshal -
func Unmarshal(v []byte) (interface{}, error) {
	for _, w := range v {
		if _, ok := noops[w]; ok {
			continue
		}

		fn := convs[w]
		if fn == nil {
			break
		}
		return fn(v)
	}
	return nil, ErrUnknownType
}

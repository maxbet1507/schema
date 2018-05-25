package patch

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
	"github.com/pkg/errors"
)

func TestRawApplyObject(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"key1": "val1",
	})
	c, _ := json.Marshal(map[string]interface{}{
		"key1": "val1",
		"key2": "val2",
	})
	p, _ := pointer.StringOf("/key2")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		if ref.String() != "key2" {
			t.Fatal(ref)
		}

		aux := map[string]interface{}{}
		if err := json.Unmarshal(raw, &aux); err != nil {
			t.Fatal(err)
		}
		aux["key2"] = "val2"

		return json.Marshal(aux)
	}

	if r, err := rawApply(p, v, fn); err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestRawApplyArray(t *testing.T) {
	v, _ := json.Marshal([]interface{}{"val1"})
	c, _ := json.Marshal([]interface{}{"val2"})
	p, _ := pointer.StringOf("/0")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		if ref.String() != "0" {
			t.Fatal(ref)
		}

		aux := []interface{}{}
		if err := json.Unmarshal(raw, &aux); err != nil {
			t.Fatal(err)
		}
		aux[0] = "val2"

		return json.Marshal(aux)
	}

	if r, err := rawApply(p, v, fn); err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestRawApplyObjectObject(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"key1": map[string]interface{}{
			"key2": "val2",
		},
	})
	c, _ := json.Marshal(map[string]interface{}{
		"key1": map[string]interface{}{
			"key2": "val2",
			"key3": "val3",
		},
	})
	p, _ := pointer.StringOf("/key1/key3")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		if ref.String() != "key3" {
			t.Fatal(ref)
		}

		aux := map[string]interface{}{}
		if err := json.Unmarshal(raw, &aux); err != nil {
			t.Fatal(err)
		}
		aux["key3"] = "val3"

		return json.Marshal(aux)
	}

	if r, err := rawApply(p, v, fn); err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestRawApplyObjectArray(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"key1": []interface{}{
			"val1",
		},
	})
	c, _ := json.Marshal(map[string]interface{}{
		"key1": []interface{}{
			"val2",
		},
	})
	p, _ := pointer.StringOf("/key1/0")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		if ref.String() != "0" {
			t.Fatal(ref)
		}

		aux := []interface{}{}
		if err := json.Unmarshal(raw, &aux); err != nil {
			t.Fatal(err)
		}
		aux[0] = "val2"

		return json.Marshal(aux)
	}

	if r, err := rawApply(p, v, fn); err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestRawApplyArrayObject(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"val1",
		map[string]interface{}{
			"key2": "val2",
		},
	})
	c, _ := json.Marshal([]interface{}{
		"val1",
		map[string]interface{}{
			"key2": "val2",
			"key3": "val3",
		},
	})
	p, _ := pointer.StringOf("/1/key3")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		if ref.String() != "key3" {
			t.Fatal(ref)
		}

		aux := map[string]interface{}{}
		if err := json.Unmarshal(raw, &aux); err != nil {
			t.Fatal(err)
		}
		aux["key3"] = "val3"

		return json.Marshal(aux)
	}

	if r, err := rawApply(p, v, fn); err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestRawApplyArrayArray(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"val1",
		[]interface{}{
			"val2",
		},
	})
	c, _ := json.Marshal([]interface{}{
		"val1",
		[]interface{}{
			"val3",
		},
	})
	p, _ := pointer.StringOf("/1/0")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		if ref.String() != "0" {
			t.Fatal(ref)
		}

		aux := []interface{}{}
		if err := json.Unmarshal(raw, &aux); err != nil {
			t.Fatal(err)
		}
		aux[0] = "val3"

		return json.Marshal(aux)
	}

	if r, err := rawApply(p, v, fn); err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestRawApplyInvalidPointer(t *testing.T) {
	v, _ := json.Marshal([]interface{}{})
	p, _ := pointer.StringOf("")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		return raw, nil
	}

	if _, err := rawApply(p, v, fn); errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(err)
	}
}

func TestRawApplyInvalidReferenceToken1(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"key1": map[string]interface{}{
			"key2": "val2",
		},
	})
	p, _ := pointer.StringOf("/key2/key3")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		return raw, nil
	}

	if _, err := rawApply(p, v, fn); errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(err)
	}
}

func TestRawApplyInvalidReferenceToken2(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"val1",
		map[string]interface{}{
			"key2": "val2",
		},
	})
	p, _ := pointer.StringOf("/key2/key3")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		return raw, nil
	}

	if _, err := rawApply(p, v, fn); errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(err)
	}
}

func TestRawApplyInvalidReferenceToken3(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"val1",
		map[string]interface{}{
			"key2": "val2",
		},
	})
	p, _ := pointer.StringOf("/-/key3")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		return raw, nil
	}

	if _, err := rawApply(p, v, fn); errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(err)
	}
}

func TestRawApplyError1(t *testing.T) {
	mockerr := fmt.Errorf("mock")
	v, _ := json.Marshal(map[string]interface{}{
		"key1": map[string]interface{}{},
	})
	p, _ := pointer.StringOf("/key1/key2")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		return raw, mockerr
	}

	if _, err := rawApply(p, v, fn); errors.Cause(err) != mockerr {
		t.Fatal(err)
	}
}

func TestRawApplyError2(t *testing.T) {
	mockerr := fmt.Errorf("mock")
	v, _ := json.Marshal([]interface{}{
		[]interface{}{},
	})
	p, _ := pointer.StringOf("/0/0")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		return raw, mockerr
	}

	if _, err := rawApply(p, v, fn); errors.Cause(err) != mockerr {
		t.Fatal(err)
	}
}

func TestRawApplyUnmarshalError1(t *testing.T) {
	v := []byte(`{"key1":{dummy}}`)
	p, _ := pointer.StringOf("/key1/key2")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		return raw, nil
	}

	if _, err := rawApply(p, v, fn); !strings.Contains(errors.Cause(err).Error(), "invalid character") {
		t.Fatal(err)
	}
}

func TestRawApplyUnmarshalError2(t *testing.T) {
	v := []byte(`[[dummy]`)
	p, _ := pointer.StringOf("/0/0")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		return raw, nil
	}

	if _, err := rawApply(p, v, fn); !strings.Contains(errors.Cause(err).Error(), "invalid character") {
		t.Fatal(err)
	}
}

func TestRawApplyUnmarshalError3(t *testing.T) {
	v := []byte(`true`)
	p, _ := pointer.StringOf("/0/0")

	fn := func(ref pointer.ReferenceToken, raw []byte) ([]byte, error) {
		return raw, nil
	}

	if _, err := rawApply(p, v, fn); errors.Cause(err) != ErrInvalidJSONType {
		t.Fatal(err)
	}
}

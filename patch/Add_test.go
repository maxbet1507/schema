package patch

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
	"github.com/pkg/errors"
)

func TestAdd(t *testing.T) {
	p, err := Add("/abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	c := []byte(`{
		"op": "add",
		"path": "/abc",
		"value": null
	}`)

	b, err := json.Marshal(p)
	if err != nil || !types.DeepEqual(b, c) {
		t.Fatal(string(b), err)
	}
}

func TestAddError1(t *testing.T) {
	_, err := Add("abc", nil)
	if errors.Cause(err) != pointer.ErrInvalidSyntax {
		t.Fatal(err)
	}
}

func TestAddError2(t *testing.T) {
	_, err := Add("/abc", make(chan struct{}))
	if !strings.Contains(err.Error(), "unsupported type") {
		t.Fatal(err)
	}
}

func TestAddApplyInvalidJSONType(t *testing.T) {
	v, _ := json.Marshal(true)
	p, _ := Add("/abc", nil)

	r, err := p.Apply(v)
	if errors.Cause(err) != ErrInvalidJSONType {
		t.Fatal(string(r), err)
	}
}

func TestAddApplyObject(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"abc": 1,
	})
	c, _ := json.Marshal(map[string]interface{}{
		"abc": 1,
		"def": nil,
	})
	p, _ := Add("/def", nil)

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestAddApplyObjectUnmarshalError(t *testing.T) {
	v := []byte(`{"abc":{"def":dummy}}`)
	p, _ := Add("/ghi", nil)

	r, err := p.Apply(v)
	if !strings.Contains(errors.Cause(err).Error(), "invalid character") {
		t.Fatal(string(r), err)
	}
}

func TestAddApplyArray1(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"key1",
	})
	c, _ := json.Marshal([]interface{}{
		nil,
		"key1",
	})
	p, _ := Add("/0", nil)

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestAddApplyArray2(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"key1",
	})
	c, _ := json.Marshal([]interface{}{
		"key1",
		nil,
	})
	p, _ := Add("/-", nil)

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestAddApplyArrayUnmarshalError(t *testing.T) {
	v := []byte(`[[dummy]]`)
	p, _ := Add("/0", nil)

	r, err := p.Apply(v)
	if !strings.Contains(errors.Cause(err).Error(), "invalid character") {
		t.Fatal(string(r), err)
	}
}

func TestAddApplyArrayInvalidPointer1(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"key1",
	})
	p, _ := Add("/dummy", nil)

	r, err := p.Apply(v)
	if errors.Cause(err) != pointer.ErrInvalidSyntax {
		t.Fatal(string(r), err)
	}
}

func TestAddApplyArrayInvalidPointer2(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"key1",
	})
	p, _ := Add("/2", nil)

	r, err := p.Apply(v)
	if errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(string(r), err)
	}
}

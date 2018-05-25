package patch

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
	"github.com/pkg/errors"
)

func TestReplace(t *testing.T) {
	p, err := Replace("/abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	c := []byte(`{
		"op": "replace",
		"path": "/abc",
		"value": null
	}`)

	b, err := json.Marshal(p)
	if err != nil || !types.DeepEqual(b, c) {
		t.Fatal(string(b), err)
	}
}

func TestReplaceError1(t *testing.T) {
	_, err := Replace("abc", nil)
	if errors.Cause(err) != pointer.ErrInvalidSyntax {
		t.Fatal(err)
	}
}

func TestReplaceError2(t *testing.T) {
	_, err := Replace("/abc", make(chan struct{}))
	if !strings.Contains(err.Error(), "unsupported type") {
		t.Fatal(err)
	}
}

func TestReplaceApplyInvalidJSONType(t *testing.T) {
	v, _ := json.Marshal(true)
	p, _ := Replace("/abc", nil)

	r, err := p.Apply(v)
	if errors.Cause(err) != ErrInvalidJSONType {
		t.Fatal(string(r), err)
	}
}

func TestReplaceApplyObject(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"abc": 1,
	})
	c, _ := json.Marshal(map[string]interface{}{
		"abc": 1,
		"def": nil,
	})
	p, _ := Replace("/def", nil)

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestReplaceApplyObjectUnmarshalError(t *testing.T) {
	v := []byte(`{"abc":{"def":dummy}}`)
	p, _ := Replace("/ghi", nil)

	r, err := p.Apply(v)
	if !strings.Contains(errors.Cause(err).Error(), "invalid character") {
		t.Fatal(string(r), err)
	}
}

func TestReplaceApplyArray1(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"key1",
	})
	c, _ := json.Marshal([]interface{}{
		nil,
	})
	p, _ := Replace("/0", nil)

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestReplaceApplyArrayUnmarshalError(t *testing.T) {
	v := []byte(`[[dummy]]`)
	p, _ := Replace("/0", nil)

	r, err := p.Apply(v)
	if !strings.Contains(errors.Cause(err).Error(), "invalid character") {
		t.Fatal(string(r), err)
	}
}

func TestReplaceApplyArrayInvalidPointer1(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"key1",
	})
	p, _ := Replace("/dummy", nil)

	r, err := p.Apply(v)
	if errors.Cause(err) != pointer.ErrInvalidSyntax {
		t.Fatal(string(r), err)
	}
}

func TestReplaceApplyArrayInvalidPointer2(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"key1",
	})
	p, _ := Replace("/2", nil)

	r, err := p.Apply(v)
	if errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(string(r), err)
	}
}

func TestReplaceApplyArrayInvalidPointer3(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"key1",
	})
	p, _ := Replace("/-", nil)

	r, err := p.Apply(v)
	if errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(string(r), err)
	}
}

package patch

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
	"github.com/pkg/errors"
)

func TestRemove(t *testing.T) {
	p, err := Remove("/abc")
	if err != nil {
		t.Fatal(err)
	}

	c := []byte(`{
		"op": "remove",
		"path": "/abc"
	}`)

	b, err := json.Marshal(p)
	if err != nil || !types.DeepEqual(b, c) {
		t.Fatal(string(b), err)
	}
}

func TestRemoveError(t *testing.T) {
	_, err := Remove("abc")
	if errors.Cause(err) != pointer.ErrInvalidSyntax {
		t.Fatal(err)
	}
}

func TestRemoveApplyInvalidJSONType(t *testing.T) {
	v, _ := json.Marshal(true)
	p, _ := Remove("/abc")

	r, err := p.Apply(v)
	if errors.Cause(err) != ErrInvalidJSONType {
		t.Fatal(string(r), err)
	}
}

func TestRemoveApplyObject(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"abc": 0,
		"def": 1,
	})
	c, _ := json.Marshal(map[string]interface{}{
		"def": 1,
	})
	p, _ := Remove("/abc")

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestRemoveApplyObjectUnmarshalError(t *testing.T) {
	v := []byte(`{"abc":{"def":dummy}}`)
	p, _ := Remove("/abc")

	r, err := p.Apply(v)
	if !strings.Contains(errors.Cause(err).Error(), "invalid character") {
		t.Fatal(string(r), err)
	}
}

func TestRemoveApplyArray(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"key1",
		"key2",
	})
	c, _ := json.Marshal([]interface{}{
		"key2",
	})
	p, _ := Remove("/0")

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestRemoveApplyArrayUnmarshalError(t *testing.T) {
	v := []byte(`[[dummy]]`)
	p, _ := Remove("/0")

	r, err := p.Apply(v)
	if !strings.Contains(errors.Cause(err).Error(), "invalid character") {
		t.Fatal(string(r), err)
	}
}

func TestRemoveApplyArrayInvalidPointer1(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"key1",
	})
	p, _ := Remove("/dummy")

	r, err := p.Apply(v)
	if errors.Cause(err) != pointer.ErrInvalidSyntax {
		t.Fatal(string(r), err)
	}
}

func TestRemoveApplyArrayInvalidPointer2(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"key1",
	})
	p, _ := Remove("/1")

	r, err := p.Apply(v)
	if errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(string(r), err)
	}
}

func TestRemoveApplyArrayInvalidPointer3(t *testing.T) {
	v, _ := json.Marshal([]interface{}{
		"key1",
	})
	p, _ := Remove("/-")

	r, err := p.Apply(v)
	if errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(string(r), err)
	}
}

package patch

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
	"github.com/pkg/errors"
)

func TestTest(t *testing.T) {
	p, err := Test("/abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	c := []byte(`{
		"op": "test",
		"path": "/abc",
		"value": null
	}`)

	b, err := json.Marshal(p)
	if err != nil || !types.DeepEqual(b, c) {
		t.Fatal(string(b), err)
	}
}

func TestTestError1(t *testing.T) {
	_, err := Test("abc", nil)
	if errors.Cause(err) != pointer.ErrInvalidSyntax {
		t.Fatal(err)
	}
}

func TestTestError2(t *testing.T) {
	_, err := Test("/abc", make(chan struct{}))
	if !strings.Contains(err.Error(), "unsupported type") {
		t.Fatal(err)
	}
}

func TestTestApply(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"abc": 1,
	})
	c, _ := json.Marshal(map[string]interface{}{
		"abc": 1,
	})
	p, _ := Test("/abc", 1)

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestTestApplyNotEqual(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"abc": 1,
	})
	p, _ := Test("/abc", 2)

	r, err := p.Apply(v)
	if errors.Cause(err) != ErrTestNotEqual {
		t.Fatal(string(r), err)
	}
}

func TestTestApplyError1(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"abc": 1,
	})
	p, _ := Test("/def", 1)

	r, err := p.Apply(v)
	if errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(string(r), err)
	}
}

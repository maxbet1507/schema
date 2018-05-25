package patch

import (
	"encoding/json"
	"testing"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
	"github.com/pkg/errors"
)

func TestCopy(t *testing.T) {
	p, err := Copy("/abc", "/def")
	if err != nil {
		t.Fatal(err)
	}

	c := []byte(`{
		"op": "copy",
		"path": "/abc",
		"from": "/def"
	}`)

	b, err := json.Marshal(p)
	if err != nil || !types.DeepEqual(b, c) {
		t.Fatal(string(b), err)
	}
}

func TestCopyError1(t *testing.T) {
	_, err := Copy("abc", "/def")
	if errors.Cause(err) != pointer.ErrInvalidSyntax {
		t.Fatal(err)
	}
}

func TestCopyError2(t *testing.T) {
	_, err := Copy("/abc", "def")
	if errors.Cause(err) != pointer.ErrInvalidSyntax {
		t.Fatal(err)
	}
}

func TestCopyApply(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"def": 1,
	})
	c, _ := json.Marshal(map[string]interface{}{
		"abc": 1,
		"def": 1,
	})
	p, _ := Copy("/abc", "/def")

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestCopyApplyError(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"ghi": 1,
	})
	p, _ := Copy("/abc", "/def")

	r, err := p.Apply(v)
	if errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(string(r), err)
	}
}

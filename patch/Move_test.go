package patch

import (
	"encoding/json"
	"testing"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
	"github.com/pkg/errors"
)

func TestMove(t *testing.T) {
	p, err := Move("/abc", "/def")
	if err != nil {
		t.Fatal(err)
	}

	c := []byte(`{
		"op": "move",
		"path": "/abc",
		"from": "/def"
	}`)

	b, err := json.Marshal(p)
	if err != nil || !types.DeepEqual(b, c) {
		t.Fatal(string(b), err)
	}
}

func TestMoveError1(t *testing.T) {
	_, err := Move("abc", "/def")
	if errors.Cause(err) != pointer.ErrInvalidSyntax {
		t.Fatal(err)
	}
}

func TestMoveError2(t *testing.T) {
	_, err := Move("/abc", "def")
	if errors.Cause(err) != pointer.ErrInvalidSyntax {
		t.Fatal(err)
	}
}

func TestMoveApply(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"def": 1,
	})
	c, _ := json.Marshal(map[string]interface{}{
		"abc": 1,
	})
	p, _ := Move("/abc", "/def")

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestMoveApplyError(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"ghi": 1,
	})
	p, _ := Move("/abc", "/def")

	r, err := p.Apply(v)
	if errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(string(r), err)
	}
}

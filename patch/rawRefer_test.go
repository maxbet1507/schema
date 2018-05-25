package patch

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/maxbet1507/schema/pointer"
	"github.com/maxbet1507/schema/types"
	"github.com/pkg/errors"
)

var (
	testReferValue, _ = json.Marshal(map[string]interface{}{
		"abc~123": map[string]interface{}{
			"abc/123": 1,
		},
		"abc/123": []interface{}{
			2, 3,
		},
	})
)

func TestRawRefer(t *testing.T) {
	p := pointer.Empty
	c := testReferValue
	if r, err := rawRefer(p, testReferValue); err != nil || !types.DeepEqual(r, c) {
		t.Fatal(string(r), err)
	}
}

func TestRawReferError(t *testing.T) {
	p := pointer.Append(pointer.Empty, pointer.MemberNameOf("abc~123"))
	if _, err := rawRefer(p, []byte("true")); errors.Cause(err) != ErrInvalidJSONType {
		t.Fatal(err)
	}
}

func TestRawReferObject(t *testing.T) {
	p := pointer.Append(pointer.Empty, pointer.MemberNameOf("abc~123"))
	c, _ := json.Marshal(map[string]interface{}{"abc/123": 1})
	if r, err := rawRefer(p, testReferValue); err != nil || !types.DeepEqual(r, c) {
		t.Fatal(r, err)
	}
}

func TestRawReferObjectMember(t *testing.T) {
	p := pointer.Append(pointer.Empty, pointer.MemberNameOf("abc~123"), pointer.MemberNameOf("abc/123"))
	c, _ := json.Marshal(1)
	if r, err := rawRefer(p, testReferValue); err != nil || !types.DeepEqual(r, c) {
		t.Fatal(r, err)
	}
}

func TestRawReferObjectUnmarshalError(t *testing.T) {
	p := pointer.Append(pointer.Empty, pointer.MemberNameOf("abc~123"), pointer.MemberNameOf("abc/124"))
	v := []byte(`{dummy}`)
	if _, err := rawRefer(p, v); !strings.Contains(errors.Cause(err).Error(), "invalid character") {
		t.Fatal(err)
	}
}

func TestRawReferObjectMemberError(t *testing.T) {
	p := pointer.Append(pointer.Empty, pointer.MemberNameOf("abc~123"), pointer.MemberNameOf("abc/124"))
	if _, err := rawRefer(p, testReferValue); errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(err)
	}
}

func TestRawReferArray(t *testing.T) {
	p := pointer.Append(pointer.Empty, pointer.MemberNameOf("abc/123"))
	c, _ := json.Marshal([]interface{}{2, 3})
	if r, err := rawRefer(p, testReferValue); err != nil || !types.DeepEqual(r, c) {
		t.Fatal(r, err)
	}
}

func TestRawReferArrayIndex1(t *testing.T) {
	p := pointer.Append(pointer.Empty, pointer.MemberNameOf("abc/123"), pointer.ArrayIndexOf(0))
	c, _ := json.Marshal(2)
	if r, err := rawRefer(p, testReferValue); err != nil || !types.DeepEqual(r, c) {
		t.Fatal(r, err)
	}
}

func TestRawReferArrayIndex2(t *testing.T) {
	p := pointer.Append(pointer.Empty, pointer.MemberNameOf("abc/123"), pointer.ArrayIndexOf(1))
	c, _ := json.Marshal(3)
	if r, err := rawRefer(p, testReferValue); err != nil || !types.DeepEqual(r, c) {
		t.Fatal(r, err)
	}
}

func TestRawReferArrayUnmarshalError(t *testing.T) {
	p := pointer.Append(pointer.Empty, pointer.MemberNameOf("abc~123"))
	v := []byte(`[dummy]`)
	if _, err := rawRefer(p, v); !strings.Contains(errors.Cause(err).Error(), "invalid character") {
		t.Fatal(err)
	}
}

func TestRawReferArrayIndexError1(t *testing.T) {
	p := pointer.Append(pointer.Empty, pointer.MemberNameOf("abc/123"), pointer.ArrayIndexOf(2))
	if _, err := rawRefer(p, testReferValue); errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(err)
	}
}

func TestRawReferArrayIndexError2(t *testing.T) {
	p := pointer.Append(pointer.Empty, pointer.MemberNameOf("abc/123"), pointer.MemberNameOf("-"))
	if _, err := rawRefer(p, testReferValue); errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(err)
	}
}

func TestRawReferArrayIndexError3(t *testing.T) {
	p := pointer.Append(pointer.Empty, pointer.MemberNameOf("abc/123"), pointer.MemberNameOf("a"))
	if _, err := rawRefer(p, testReferValue); errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(err)
	}
}

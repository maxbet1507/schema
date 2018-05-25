package patch

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/maxbet1507/schema/types"
	"github.com/pkg/errors"
)

func TestBytesOfAdd(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":    "add",
		"path":  "/abc",
		"value": nil,
	})

	p, err := BytesOf(v)
	if err != nil {
		t.Fatal(err)
	}

	r, err := json.Marshal(p)
	if err != nil || !types.DeepEqual(r, v) {
		t.Fatal(string(r), err)
	}
}

func TestBytesOfAddError1(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":    "add",
		"value": nil,
	})

	_, err := BytesOf(v)
	if errors.Cause(err) != ErrInvalidJSONPatch {
		t.Fatal(err)
	}
}

func TestBytesOfAddError2(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":   "add",
		"path": "/abc",
	})

	_, err := BytesOf(v)
	if errors.Cause(err) != ErrInvalidJSONPatch {
		t.Fatal(err)
	}
}

func TestBytesOfRemove(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":   "remove",
		"path": "/abc",
	})

	p, err := BytesOf(v)
	if err != nil {
		t.Fatal(err)
	}

	r, err := json.Marshal(p)
	if err != nil || !types.DeepEqual(r, v) {
		t.Fatal(string(r), err)
	}
}

func TestBytesOfRemoveError(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op": "remove",
	})

	_, err := BytesOf(v)
	if errors.Cause(err) != ErrInvalidJSONPatch {
		t.Fatal(err)
	}
}

func TestBytesOfReplace(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":    "replace",
		"path":  "/abc",
		"value": nil,
	})

	p, err := BytesOf(v)
	if err != nil {
		t.Fatal(err)
	}

	r, err := json.Marshal(p)
	if err != nil || !types.DeepEqual(r, v) {
		t.Fatal(string(r), err)
	}
}

func TestBytesOfReplaceError1(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":    "replace",
		"value": nil,
	})

	_, err := BytesOf(v)
	if errors.Cause(err) != ErrInvalidJSONPatch {
		t.Fatal(err)
	}
}

func TestBytesOfReplaceError2(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":   "replace",
		"path": "/abc",
	})

	_, err := BytesOf(v)
	if errors.Cause(err) != ErrInvalidJSONPatch {
		t.Fatal(err)
	}
}

func TestBytesOfMove(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":   "move",
		"path": "/abc",
		"from": "/def",
	})

	p, err := BytesOf(v)
	if err != nil {
		t.Fatal(err)
	}

	r, err := json.Marshal(p)
	if err != nil || !types.DeepEqual(r, v) {
		t.Fatal(string(r), string(v), err)
	}
}

func TestBytesOfMoveError1(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":   "move",
		"from": "/def",
	})

	_, err := BytesOf(v)
	if errors.Cause(err) != ErrInvalidJSONPatch {
		t.Fatal(err)
	}
}

func TestBytesOfMoveError2(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":   "move",
		"path": "/abc",
	})

	_, err := BytesOf(v)
	if errors.Cause(err) != ErrInvalidJSONPatch {
		t.Fatal(err)
	}
}

func TestBytesOfCopy(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":   "copy",
		"path": "/abc",
		"from": "/def",
	})

	p, err := BytesOf(v)
	if err != nil {
		t.Fatal(err)
	}

	r, err := json.Marshal(p)
	if err != nil || !types.DeepEqual(r, v) {
		t.Fatal(string(r), string(v), err)
	}
}

func TestBytesOfCopyError1(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":   "copy",
		"from": "/def",
	})

	_, err := BytesOf(v)
	if errors.Cause(err) != ErrInvalidJSONPatch {
		t.Fatal(err)
	}
}

func TestBytesOfCopyError2(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":   "copy",
		"path": "/abc",
	})

	_, err := BytesOf(v)
	if errors.Cause(err) != ErrInvalidJSONPatch {
		t.Fatal(err)
	}
}

func TestBytesOfTest(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":    "test",
		"path":  "/abc",
		"value": nil,
	})

	p, err := BytesOf(v)
	if err != nil {
		t.Fatal(err)
	}

	r, err := json.Marshal(p)
	if err != nil || !types.DeepEqual(r, v) {
		t.Fatal(string(r), err)
	}
}

func TestBytesOfTestError1(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":    "test",
		"value": nil,
	})

	_, err := BytesOf(v)
	if errors.Cause(err) != ErrInvalidJSONPatch {
		t.Fatal(err)
	}
}

func TestBytesOfTestError2(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"op":   "test",
		"path": "/abc",
	})

	_, err := BytesOf(v)
	if errors.Cause(err) != ErrInvalidJSONPatch {
		t.Fatal(err)
	}
}

func TestBytesOfError(t *testing.T) {
	v := []byte(`dummy`)

	_, err := BytesOf(v)
	if !strings.Contains(err.Error(), "invalid character") {
		t.Fatal(err)
	}
}

func TestPatches(t *testing.T) {
	j, _ := json.Marshal([]interface{}{
		map[string]interface{}{
			"op":    "test",
			"path":  "/abc",
			"value": 1,
		},
		map[string]interface{}{
			"op":    "test",
			"path":  "/def",
			"value": 2,
		},
	})

	p := Patches{}
	if err := json.Unmarshal(j, &p); err != nil {
		t.Fatal(err)
	}

	v1, _ := json.Marshal(map[string]interface{}{
		"abc": 1,
		"def": 2,
	})
	c1, _ := json.Marshal(map[string]interface{}{
		"abc": 1,
		"def": 2,
	})
	r1, err := p.Apply(v1)
	if err != nil || !types.DeepEqual(r1, c1) {
		t.Fatal(string(r1), err)
	}

	v2, _ := json.Marshal(map[string]interface{}{
		"abc": 1,
		"def": 3,
	})
	r2, err := p.Apply(v2)
	if errors.Cause(err) != ErrTestNotEqual {
		t.Fatal(string(r2), err)
	}
}

func TestPatchesUnmarshalError1(t *testing.T) {
	j, _ := json.Marshal(true)

	p := Patches{}
	if err := json.Unmarshal(j, &p); !strings.Contains(err.Error(), "cannot unmarshal") {
		t.Fatal(err)
	}
}

func TestPatchesUnmarshalError2(t *testing.T) {
	j, _ := json.Marshal([]interface{}{
		map[string]interface{}{
			"op":    "dummy",
			"path":  "/abc",
			"value": 1,
		},
	})

	p := Patches{}
	if err := json.Unmarshal(j, &p); err != ErrInvalidJSONPatch {
		t.Fatal(err)
	}
}

func TestRFC6902_Examples_A1_Adding_an_Object_Member(t *testing.T) {
	var p Patches

	v := []byte(`{ "foo": "bar"}`)
	s := []byte(`{ "baz": "qux", "foo": "bar" }`)

	err := json.Unmarshal([]byte(`[{ "op": "add", "path": "/baz", "value": "qux" }]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, s) {
		t.Fatal(r, err)
	}
}

func TestRFC6902_Examples_A2_Adding_an_Array_Element(t *testing.T) {
	var p Patches

	v := []byte(`{ "foo": [ "bar", "baz" ] }`)
	s := []byte(`{ "foo": [ "bar", "qux", "baz" ] }`)

	err := json.Unmarshal([]byte(`[{ "op": "add", "path": "/foo/1", "value": "qux" }]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, s) {
		t.Fatal(r, err)
	}
}

func TestRFC6902_Examples_A3_Removing_an_Object_Member(t *testing.T) {
	var p Patches

	v := []byte(`{ "baz": "qux", "foo": "bar" }`)
	s := []byte(`{ "foo": "bar" }`)

	err := json.Unmarshal([]byte(`[{ "op": "remove", "path": "/baz" }]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, s) {
		t.Fatal(r, err)
	}
}

func TestRFC6902_Examples_A4_Removing_an_Array_Element(t *testing.T) {
	var p Patches

	v := []byte(`{ "foo": [ "bar", "qux", "baz" ] }`)
	s := []byte(`{ "foo": [ "bar", "baz" ] }`)

	err := json.Unmarshal([]byte(`[{ "op": "remove", "path": "/foo/1" }]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, s) {
		t.Fatal(r, err)
	}
}

func TestRFC6902_Examples_A5_Replacing_a_Value(t *testing.T) {
	var p Patches

	v := []byte(`{ "baz": "qux", "foo": "bar" }`)
	s := []byte(`{ "baz": "boo", "foo": "bar" }`)

	err := json.Unmarshal([]byte(`[{ "op": "replace", "path": "/baz", "value": "boo" }]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, s) {
		t.Fatal(r, err)
	}
}

func TestRFC6902_Examples_A6_Moving_a_Value(t *testing.T) {
	var p Patches

	v := []byte(`   {
		"foo": {
			"bar": "baz",
			"waldo": "fred"
		},
		"qux": {
			"corge": "grault"
		}
	}`)
	s := []byte(`{
		"foo": {
			"bar": "baz"
		},
		"qux": {
			"corge": "grault",
			"thud": "fred"
		}
	}`)

	err := json.Unmarshal([]byte(`[{ "op": "move", "from": "/foo/waldo", "path": "/qux/thud" }]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, s) {
		t.Fatal(r, err)
	}
}

func TestRFC6902_Examples_A7_Moving_an_Array_Element(t *testing.T) {
	var p Patches

	v := []byte(`{ "foo": [ "all", "grass", "cows", "eat" ] }`)
	s := []byte(`{ "foo": [ "all", "cows", "eat", "grass" ] }`)

	err := json.Unmarshal([]byte(`[{ "op": "move", "from": "/foo/1", "path": "/foo/3" }]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, s) {
		t.Fatal(r, err)
	}
}

func TestRFC6902_Examples_A8_Testing_a_Value_Success(t *testing.T) {
	var p Patches

	v := []byte(`{ "baz": "qux", "foo": [ "a", 2, "c" ] }`)
	s := []byte(`{ "baz": "qux", "foo": [ "a", 2, "c" ] }`)

	err := json.Unmarshal([]byte(`[
		{ "op": "test", "path": "/baz", "value": "qux" },
		{ "op": "test", "path": "/foo/1", "value": 2 }
	]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, s) {
		t.Fatal(r, err)
	}
}

func TestRFC6902_Examples_A9_Testing_a_Value_Error(t *testing.T) {
	var p Patches

	v := []byte(`{ "baz": "qux" }`)

	err := json.Unmarshal([]byte(`[{ "op": "test", "path": "/baz", "value": "bar" }]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	_, err = p.Apply(v)
	if errors.Cause(err) != ErrTestNotEqual {
		t.Fatal(err)
	}
}

func TestRFC6902_Examples_A10_Adding_a_Nested_Member_Object(t *testing.T) {
	var p Patches

	v := []byte(`{ "foo": "bar" }`)
	s := []byte(`{
		"foo": "bar",
		"child": {
			"grandchild": {}
		}
	}`)

	err := json.Unmarshal([]byte(`[{ "op": "add", "path": "/child", "value": { "grandchild": { } } }]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, s) {
		t.Fatal(r, err)
	}
}

func TestRFC6902_Examples_A11_Ignoring_Unrecognized_Elements(t *testing.T) {
	var p Patches

	v := []byte(`{ "foo": "bar" }`)
	s := []byte(`{ "foo": "bar", "baz": "qux" }`)

	err := json.Unmarshal([]byte(`[{ "op": "add", "path": "/baz", "value": "qux", "xyz": 123 }]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, s) {
		t.Fatal(r, err)
	}
}

func TestRFC6902_Examples_A12_Adding_to_a_Nonexistent_Target(t *testing.T) {
	var p Patches

	v := []byte(`{ "foo": "bar" }`)

	err := json.Unmarshal([]byte(`[{ "op": "add", "path": "/baz/bat", "value": "qux" }]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	_, err = p.Apply(v)
	if errors.Cause(err) != ErrInvalidPointer {
		t.Fatal(err)
	}
}

func TestRFC6902_Examples_A13_Invalid_JSON_Patch_Document(t *testing.T) {
	var p Patches

	err := json.Unmarshal([]byte(`[{ "op": "add", "path": "/baz", "value": "qux", "op": "remove" }]`), &p)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRFC6902_Examples_A14_Escape_Ordering(t *testing.T) {
	var p Patches

	v := []byte(`{ "/": 9, "~1": 10 }`)
	s := []byte(`{ "/": 9, "~1": 10 }`)

	err := json.Unmarshal([]byte(`[{"op": "test", "path": "/~01", "value": 10}]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, s) {
		t.Fatal(r, err)
	}
}

func TestRFC6902_Examples_A15_Comparing_Strings_and_Numbers(t *testing.T) {
	var p Patches

	v := []byte(`{ "/": 9, "~1": 10 }`)

	err := json.Unmarshal([]byte(`[{"op": "test", "path": "/~01", "value": "10"}]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	_, err = p.Apply(v)
	if errors.Cause(err) != ErrTestNotEqual {
		t.Fatal(err)
	}
}

func TestRFC6902_Examples_A16_Adding_an_Array_Value(t *testing.T) {
	var p Patches

	v := []byte(`{ "foo": ["bar"] }`)
	s := []byte(`{ "foo": ["bar", ["abc", "def"]] }`)

	err := json.Unmarshal([]byte(`[{ "op": "add", "path": "/foo/-", "value": ["abc", "def"] }]`), &p)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v)
	if err != nil || !types.DeepEqual(r, s) {
		t.Fatal(r, err)
	}
}

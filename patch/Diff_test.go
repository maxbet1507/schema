package patch

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/maxbet1507/schema/types"
	"github.com/pkg/errors"
)

func TestDiffObject(t *testing.T) {
	v1, _ := json.Marshal(map[string]interface{}{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
	})

	v2, _ := json.Marshal(map[string]interface{}{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
	})

	p, err := Diff(v1, v2)
	if err != nil {
		t.Fatal(err)
	}

	if len(p) != 0 {
		t.Fatal(p)
	}
}

func TestDiffObjectPatch1(t *testing.T) {
	v1, _ := json.Marshal(map[string]interface{}{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
	})

	v2, _ := json.Marshal(map[string]interface{}{
		"key1": "val1",
		"key2": "val4",
		"key3": "val3",
	})

	p, err := Diff(v1, v2)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v1)
	if err != nil || !types.DeepEqual(r, v2) {
		t.Fatal(string(r), err)
	}
}

func TestDiffObjectPatch2(t *testing.T) {
	v1, _ := json.Marshal(map[string]interface{}{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
	})

	v2, _ := json.Marshal(map[string]interface{}{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
		"key4": "val4",
	})

	p, err := Diff(v1, v2)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v1)
	if err != nil || !types.DeepEqual(r, v2) {
		t.Fatal(string(r), err)
	}
}

func TestDiffObjectPatch3(t *testing.T) {
	v1, _ := json.Marshal(map[string]interface{}{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
		"key4": "val4",
	})

	v2, _ := json.Marshal(map[string]interface{}{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
	})

	p, err := Diff(v1, v2)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v1)
	if err != nil || !types.DeepEqual(r, v2) {
		t.Fatal(string(r), err)
	}
}

func TestDiffObjectPatch4(t *testing.T) {
	v1, _ := json.Marshal(map[string]interface{}{
		"key1": "val1",
		"key2": "val2",
		"key3": true,
	})

	v2, _ := json.Marshal(map[string]interface{}{
		"key1": "val1",
		"key2": "val4",
		"key3": 10,
	})

	p, err := Diff(v1, v2)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v1)
	if err != nil || !types.DeepEqual(r, v2) {
		t.Fatal(string(r), err)
	}
}

func TestDiffObjectPatchError1(t *testing.T) {
	v1 := []byte(`{dummy}`)
	v2 := []byte(`{}`)

	_, err := Diff(v1, v2)
	if !strings.Contains(err.Error(), "invalid character") {
		t.Fatal(err)
	}
}

func TestDiffObjectPatchError2(t *testing.T) {
	v1 := []byte(`{}`)
	v2 := []byte(`{dummy}`)

	_, err := Diff(v1, v2)
	if !strings.Contains(err.Error(), "invalid character") {
		t.Fatal(err)
	}
}

func TestDiffArray(t *testing.T) {
	v1, _ := json.Marshal([]interface{}{
		"val1", "val2", "val3",
	})

	v2, _ := json.Marshal([]interface{}{
		"val1", "val2", "val3",
	})

	p, err := Diff(v1, v2)
	if err != nil {
		t.Fatal(err)
	}

	if len(p) != 0 {
		t.Fatal(p)
	}
}

func TestDiffArrayPatch1(t *testing.T) {
	v1, _ := json.Marshal([]interface{}{
		"val1", "val2", "val3",
	})

	v2, _ := json.Marshal([]interface{}{
		"val1", "val4", "val3",
	})

	p, err := Diff(v1, v2)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v1)
	if err != nil || !types.DeepEqual(r, v2) {
		t.Fatal(string(r), err)
	}
}

func TestDiffArrayPatch2(t *testing.T) {
	v1, _ := json.Marshal([]interface{}{
		"val1", "val2", "val3",
	})

	v2, _ := json.Marshal([]interface{}{
		"val1", "val2", "val3", "val4",
	})

	p, err := Diff(v1, v2)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v1)
	if err != nil || !types.DeepEqual(r, v2) {
		t.Fatal(string(r), err)
	}
}

func TestDiffArrayPatch3(t *testing.T) {
	v1, _ := json.Marshal([]interface{}{
		"val1", "val2", "val3", "val4",
	})

	v2, _ := json.Marshal([]interface{}{
		"val1", "val2", "val3",
	})

	p, err := Diff(v1, v2)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v1)
	if err != nil || !types.DeepEqual(r, v2) {
		t.Fatal(string(r), err)
	}
}

func TestDiffArrayPatch4(t *testing.T) {
	v1, _ := json.Marshal([]interface{}{
		"val1", "val2", true,
	})

	v2, _ := json.Marshal([]interface{}{
		"val1", "val4", 10,
	})

	p, err := Diff(v1, v2)
	if err != nil {
		t.Fatal(err)
	}

	r, err := p.Apply(v1)
	if err != nil || !types.DeepEqual(r, v2) {
		t.Fatal(string(r), err)
	}
}

func TestDiffArrayPatchError1(t *testing.T) {
	v1 := []byte(`[dummy]`)
	v2 := []byte(`[]`)

	_, err := Diff(v1, v2)
	if !strings.Contains(err.Error(), "invalid character") {
		t.Fatal(err)
	}
}

func TestDiffArrayPatchError2(t *testing.T) {
	v1 := []byte(`[]`)
	v2 := []byte(`[dummy]`)

	_, err := Diff(v1, v2)
	if !strings.Contains(err.Error(), "invalid character") {
		t.Fatal(err)
	}
}

func TestDiffRootType1(t *testing.T) {
	v1 := []byte(`"val1"`)
	v2 := []byte(`"val2"`)

	_, err := Diff(v1, v2)
	if errors.Cause(err) != ErrInvalidJSONType {
		t.Fatal(err)
	}
}

func TestDiffRootType2(t *testing.T) {
	v1 := []byte(`{dummy}`)
	v2 := []byte(`[dummy]`)

	_, err := Diff(v1, v2)
	if !strings.Contains(err.Error(), "invalid character") {
		t.Fatal(err)
	}
}

func TestDiffRootType3(t *testing.T) {
	v1 := []byte(`dummy`)
	v2 := []byte(`dummy`)

	_, err := Diff(v1, v2)
	if errors.Cause(err) != types.ErrUnknownType {
		t.Fatal(err)
	}
}

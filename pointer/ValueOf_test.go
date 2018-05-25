package pointer_test

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/maxbet1507/schema/pointer"
	"github.com/pkg/errors"
)

func TestValueOf(t *testing.T) {
	v, _ := json.Marshal(map[string]interface{}{
		"foo":  []interface{}{"bar", "baz"},
		"":     0,
		"a/b":  1,
		"c%d":  2,
		"e^f":  3,
		"g|h":  4,
		"i\\j": 5,
		"k\"l": 6,
		" ":    7,
		"m~n":  map[string]interface{}{"key1": "val1"},
	})

	c := []string{
		"/",
		"/ ",
		"/a~1b",
		"/c%d",
		"/e^f",
		"/foo", "/foo/0",
		"/foo/1",
		"/g|h",
		"/i\\j",
		"/k\"l",
		"/m~0n",
		"/m~0n/key1",
	}

	r, err := pointer.ValueOf(v)
	if err != nil {
		t.Fatal(err)
	}

	s := []string{}
	for _, r := range r {
		s = append(s, r.String())
	}
	sort.Strings(s)

	if !reflect.DeepEqual(s, c) {
		t.Fatal(s)
	}
}

func TestValueOfError1(t *testing.T) {
	v := []byte(`{dummy}`)

	_, err := pointer.ValueOf(v)
	if !strings.Contains(err.Error(), "invalid character") {
		t.Fatal(err)
	}
}

func TestValueOfError2(t *testing.T) {
	v := []byte(`[dummy]`)

	_, err := pointer.ValueOf(v)
	if !strings.Contains(err.Error(), "invalid character") {
		t.Fatal(err)
	}
}

func TestValueOfError3(t *testing.T) {
	v, _ := json.Marshal(nil)

	_, err := pointer.ValueOf(v)
	if errors.Cause(err) != pointer.ErrInvalidJSONType {
		t.Fatal(err)
	}
}

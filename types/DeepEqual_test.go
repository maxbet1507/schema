package types_test

import (
	"testing"

	"github.com/maxbet1507/schema/types"
)

func TestDeepEqual(t *testing.T) {
	if r := types.DeepEqual([]byte(``), []byte(`{}`)); r != false {
		t.Fatal(r)
	}

	if r := types.DeepEqual([]byte(`{}`), []byte(``)); r != false {
		t.Fatal(r)
	}

	if r := types.DeepEqual([]byte(`{}`), []byte(`{}`)); r != true {
		t.Fatal(r)
	}

	if r := types.DeepEqual([]byte(`{"key1":"val1"}`), []byte(`{"key1":"val1"}`)); r != true {
		t.Fatal(r)
	}

	if r := types.DeepEqual([]byte(`{"key1":"val1","key2":"val2"}`), []byte(`{"key1":"val1"}`)); r != false {
		t.Fatal(r)
	}

	if r := types.DeepEqual([]byte(`{"key1":"val1","key2":"val2"}`), []byte(`{"key1":"val1","key2":"val2"}`)); r != true {
		t.Fatal(r)
	}

	if r := types.DeepEqual([]byte(`{"key2":"val2","key1":"val1"}`), []byte(`{"key1":"val1","key2":"val2"}`)); r != true {
		t.Fatal(r)
	}
}

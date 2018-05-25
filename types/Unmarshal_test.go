package types_test

import (
	"testing"

	"github.com/maxbet1507/schema/types"
)

func TestObject(t *testing.T) {
	ret, err := types.Unmarshal([]byte(` {} `))
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := ret.(types.Object); !ok {
		t.Fatal(ok)
	}
}

func TestArray(t *testing.T) {
	ret, err := types.Unmarshal([]byte(` [] `))
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := ret.(types.Array); !ok {
		t.Fatal(ok)
	}
}

func TestString(t *testing.T) {
	ret, err := types.Unmarshal([]byte(` "" `))
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := ret.(types.String); !ok {
		t.Fatal(ok)
	}
}

func TestBoolean1(t *testing.T) {
	ret, err := types.Unmarshal([]byte(` true `))
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := ret.(types.Boolean); !ok {
		t.Fatal(ok)
	}
}

func TestBoolean2(t *testing.T) {
	ret, err := types.Unmarshal([]byte(` false `))
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := ret.(types.Boolean); !ok {
		t.Fatal(ok)
	}
}

func TestNull(t *testing.T) {
	ret, err := types.Unmarshal([]byte(` null `))
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := ret.(types.Null); !ok {
		t.Fatal(ok)
	}
}

func TestNumber1(t *testing.T) {
	ret, err := types.Unmarshal([]byte(` -0 `))
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := ret.(types.Number); !ok {
		t.Fatal(ok)
	}
}

func TestNumber2(t *testing.T) {
	ret, err := types.Unmarshal([]byte(` 0 `))
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := ret.(types.Number); !ok {
		t.Fatal(ok)
	}
}

func TestUnknown1(t *testing.T) {
	_, err := types.Unmarshal([]byte(` // `))
	if err != types.ErrUnknownType {
		t.Fatal(err)
	}
}

func TestUnknown2(t *testing.T) {
	_, err := types.Unmarshal([]byte(`  `))
	if err != types.ErrUnknownType {
		t.Fatal(err)
	}
}

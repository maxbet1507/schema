package pointer_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/maxbet1507/schema/pointer"
)

func TestContext(t *testing.T) {
	ctx1 := context.Background()
	ctx2 := pointer.WithContext(ctx1, pointer.MemberNameOf("abc"))
	ctx3 := pointer.WithContext(ctx1, pointer.MemberNameOf("def"))
	ctx4 := pointer.WithContext(ctx2, pointer.MemberNameOf("ghi"))

	if ptr1r := pointer.FromContext(ctx1); ptr1r.String() != "" {
		t.Fatal(ptr1r)
	}
	if ptr2r := pointer.FromContext(ctx2); ptr2r.String() != "/abc" {
		t.Fatal(ptr2r)
	}
	if ptr3r := pointer.FromContext(ctx3); ptr3r.String() != "/def" {
		t.Fatal(ptr3r)
	}
	if ptr4r := pointer.FromContext(ctx4); ptr4r.String() != "/abc/ghi" {
		t.Fatal(ptr4r)
	}

	err := fmt.Errorf("dummy")
	if err := pointer.WrapError(ctx1, err); err.Error() != "dummy" {
		t.Fatal(err)
	}
	if err := pointer.WrapError(ctx2, err); err.Error() != "/abc: dummy" {
		t.Fatal(err)
	}
	if err := pointer.WrapError(ctx3, err); err.Error() != "/def: dummy" {
		t.Fatal(err)
	}
	if err := pointer.WrapError(ctx4, err); err.Error() != "/abc/ghi: dummy" {
		t.Fatal(err)
	}
}

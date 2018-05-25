package pointer_test

import (
	"strings"
	"testing"

	"github.com/maxbet1507/schema/pointer"
)

func TestStringOf(t *testing.T) {
	if p, _ := pointer.StringOf(""); p.String() != "" ||
		p.ReferenceToken() != nil {
		t.Fatal(p.String(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.StringOf("/foo"); p.String() != "/foo" ||
		p.ReferenceToken().MemberName() != "foo" {
		t.Fatal(p.String(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.StringOf("/foo/0"); p.String() != "/foo/0" ||
		p.ReferenceToken().MemberName() != "0" {
		t.Fatal(p.String(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.StringOf("/"); p.String() != "/" ||
		p.ReferenceToken().MemberName() != "" {
		t.Fatal(p.String(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.StringOf("/a~1b"); p.String() != "/a~1b" ||
		p.ReferenceToken().MemberName() != "a/b" {
		t.Fatal(p.String(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.StringOf("/c%d"); p.String() != "/c%d" ||
		p.ReferenceToken().MemberName() != "c%d" {
		t.Fatal(p.String(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.StringOf("/e^f"); p.String() != "/e^f" ||
		p.ReferenceToken().MemberName() != "e^f" {
		t.Fatal(p.String(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.StringOf("/g|h"); p.String() != "/g|h" ||
		p.ReferenceToken().MemberName() != "g|h" {
		t.Fatal(p.String(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.StringOf("/i\\j"); p.String() != "/i\\j" ||
		p.ReferenceToken().MemberName() != "i\\j" {
		t.Fatal(p.String(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.StringOf("/k\"l"); p.String() != "/k\"l" ||
		p.ReferenceToken().MemberName() != "k\"l" {
		t.Fatal(p.String(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.StringOf("/ "); p.String() != "/ " ||
		p.ReferenceToken().MemberName() != " " {
		t.Fatal(p.String(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.StringOf("/m~0n"); p.String() != "/m~0n" ||
		p.ReferenceToken().MemberName() != "m~n" {
		t.Fatal(p.String(), p.ReferenceToken().MemberName())
	}

	if p, err := pointer.StringOf("#"); err != pointer.ErrInvalidSyntax {
		t.Fatal(p.String(), err)
	}
}

func TestURIFragmentOf(t *testing.T) {
	if p, _ := pointer.URIFragmentOf(""); p.String() != "" ||
		p.URIFragment() != "" ||
		p.ReferenceToken() != nil {
		t.Fatal(p.String(), p.URIFragment(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.URIFragmentOf("/foo"); p.String() != "/foo" ||
		p.URIFragment() != "/foo" ||
		p.ReferenceToken().MemberName() != "foo" {
		t.Fatal(p.String(), p.URIFragment(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.URIFragmentOf("/foo/0"); p.String() != "/foo/0" ||
		p.URIFragment() != "/foo/0" ||
		p.ReferenceToken().MemberName() != "0" {
		t.Fatal(p.String(), p.URIFragment(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.URIFragmentOf("/"); p.String() != "/" ||
		p.URIFragment() != "/" ||
		p.ReferenceToken().MemberName() != "" {
		t.Fatal(p.String(), p.URIFragment(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.URIFragmentOf("/a~1b"); p.String() != "/a~1b" ||
		p.URIFragment() != "/a~1b" ||
		p.ReferenceToken().MemberName() != "a/b" {
		t.Fatal(p.String(), p.URIFragment(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.URIFragmentOf("/c%25d"); p.String() != "/c%d" ||
		p.URIFragment() != "/c%25d" ||
		p.ReferenceToken().MemberName() != "c%d" {
		t.Fatal(p.String(), p.URIFragment(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.URIFragmentOf("/e%5Ef"); p.String() != "/e^f" ||
		p.URIFragment() != "/e%5Ef" ||
		p.ReferenceToken().MemberName() != "e^f" {
		t.Fatal(p.String(), p.URIFragment(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.URIFragmentOf("/g%7Ch"); p.String() != "/g|h" ||
		p.URIFragment() != "/g%7Ch" ||
		p.ReferenceToken().MemberName() != "g|h" {
		t.Fatal(p.String(), p.URIFragment(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.URIFragmentOf("/i%5Cj"); p.String() != "/i\\j" ||
		p.URIFragment() != "/i%5Cj" ||
		p.ReferenceToken().MemberName() != "i\\j" {
		t.Fatal(p.String(), p.URIFragment(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.URIFragmentOf("/k%22l"); p.String() != "/k\"l" ||
		p.URIFragment() != "/k%22l" ||
		p.ReferenceToken().MemberName() != "k\"l" {
		t.Fatal(p.String(), p.URIFragment(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.URIFragmentOf("/%20"); p.String() != "/ " ||
		p.URIFragment() != "/%20" ||
		p.ReferenceToken().MemberName() != " " {
		t.Fatal(p.String(), p.URIFragment(), p.ReferenceToken().MemberName())
	}
	if p, _ := pointer.URIFragmentOf("/m~0n"); p.String() != "/m~0n" ||
		p.URIFragment() != "/m~0n" ||
		p.ReferenceToken().MemberName() != "m~n" {
		t.Fatal(p.String(), p.URIFragment(), p.ReferenceToken().MemberName())
	}

	if p, err := pointer.URIFragmentOf("#"); err != pointer.ErrInvalidSyntax {
		t.Fatal(p.String(), err)
	}
	if p, err := pointer.URIFragmentOf("/%zz"); !strings.HasPrefix(err.Error(), "invalid URL escape") {
		t.Fatal(p.String(), err)
	}
}

func TestArrayIndex(t *testing.T) {
	if r := pointer.ArrayIndexOf(-1); r.String() != "-" {
		t.Fatal(r.String())
	}
	if v, err := pointer.ArrayIndexOf(-1).ArrayIndex(); v != -1 || err != nil {
		t.Fatal(v, err)
	}

	if r := pointer.ArrayIndexOf(0); r.String() != "0" {
		t.Fatal(r.String())
	}
	if v, err := pointer.ArrayIndexOf(0).ArrayIndex(); v != 0 || err != nil {
		t.Fatal(v, err)
	}

	if v, err := pointer.MemberNameOf("01").ArrayIndex(); err != pointer.ErrInvalidSyntax {
		t.Fatal(v, err)
	}
}

func TestAppend(t *testing.T) {
	p, _ := pointer.StringOf("/root")
	r1 := pointer.MemberNameOf("key")
	r2 := pointer.ArrayIndexOf(0)

	if np := pointer.Append(p, r1, r2); np.String() != "/root/key/0" || p.String() != "/root" {
		t.Fatal(np.String(), p.String())
	}
}

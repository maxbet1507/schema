package pointer

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// -
var (
	ErrInvalidSyntax         = fmt.Errorf("invalid syntax")
	Empty            Pointer = rawPointer{}

	repRFC6901Decoder = strings.NewReplacer("~1", "/", "~0", "~")
	repRFC6901Encoder = strings.NewReplacer("~", "~0", "/", "~1")
	rexArrayIndex     = regexp.MustCompile("^(0|[1-9][0-9]*)$")
)

// ReferenceToken -
type ReferenceToken interface {
	fmt.Stringer

	URIFragment() string
	ArrayIndex() (int, error)
	MemberName() string
}

// Pointer -
type Pointer interface {
	fmt.Stringer

	ReferenceToken() ReferenceToken
	ReferenceTokens() []ReferenceToken
	URIFragment() string
}

type rawReferenceToken string

// URIFragment -
func (s rawReferenceToken) URIFragment() string {
	ret := repRFC6901Encoder.Replace(string(s))
	return url.PathEscape(ret)
}

// String -
func (s rawReferenceToken) String() string {
	return repRFC6901Encoder.Replace(string(s))
}

// ArrayIndex -
func (s rawReferenceToken) ArrayIndex() (int, error) {
	k := string(s)
	if k == "-" {
		return -1, nil
	}

	if !rexArrayIndex.MatchString(k) {
		return 0, ErrInvalidSyntax
	}
	return strconv.Atoi(k)
}

// MemberName -
func (s rawReferenceToken) MemberName() string {
	return string(s)
}

// ArrayIndexOf -
func ArrayIndexOf(v int) ReferenceToken {
	if v < 0 {
		return rawReferenceToken("-")
	}
	return rawReferenceToken(strconv.Itoa(v))
}

// MemberNameOf -
func MemberNameOf(v string) ReferenceToken {
	return rawReferenceToken(v)
}

type rawPointer []ReferenceToken

// ReferenceToken -
func (s rawPointer) ReferenceToken() ReferenceToken {
	if len(s) == 0 {
		return nil
	}
	return s[len(s)-1]
}

func (s rawPointer) ReferenceTokens() []ReferenceToken {
	return s
}

func (s rawPointer) URIFragment() string {
	ret := []string{""}
	for _, v := range s {
		ret = append(ret, v.URIFragment())
	}

	if len(ret) == 1 {
		return ""
	}
	return strings.Join(ret, "/")
}

func (s rawPointer) String() string {
	ret := []string{""}
	for _, v := range s {
		ret = append(ret, v.String())
	}

	if len(ret) == 1 {
		return ""
	}
	return strings.Join(ret, "/")
}

// Append -
func Append(p Pointer, rs ...ReferenceToken) Pointer {
	return append(rawPointer(p.ReferenceTokens()), rs...)
}

// URIFragmentOf -
func URIFragmentOf(v string) (Pointer, error) {
	w := strings.Split(v, "/")
	if w[0] != "" {
		return nil, ErrInvalidSyntax
	}

	r := rawPointer{}
	for _, v := range w[1:] {
		var err error
		if v, err = url.PathUnescape(v); err != nil {
			return nil, err
		}
		v = repRFC6901Decoder.Replace(v)
		r = append(r, MemberNameOf(v))
	}
	return r, nil
}

// StringOf -
func StringOf(v string) (Pointer, error) {
	w := strings.Split(v, "/")
	if w[0] != "" {
		return nil, ErrInvalidSyntax
	}

	r := rawPointer{}
	for _, v := range w[1:] {
		v = repRFC6901Decoder.Replace(v)
		r = append(r, MemberNameOf(v))
	}
	return r, nil
}

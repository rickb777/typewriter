package typewriter

import (
	"fmt"
	"go/types"
	"regexp"
	"strings"
)

type Type struct {
	Pointer          Pointer
	Name             string
	Tags             TagSlice
	Comparable       bool // any type that can be compared using '==' etc (which includes all pointers)
	FullyComparable  bool // as above, provided it is not a pointer to a non-comparable
	Numeric, Ordered bool
	test             test
	types.Type
}

func (t Type) String() (result string) {
	return fmt.Sprintf("%s%s", t.Pointer.String(), t.Name)
}

func (t Type) Ptr() (result string) {
	return t.Pointer.String()
}

// ValueOf inverts the star symbol, supplying * where a pointer value needs to be dereferenced.
func (t Type) ValueOf() (result string) {
	return t.Pointer.ValueOf()
}

// AddrOf returns & when this t a pointer type, allowing a template to take the address of a plain value.
func (t Type) AddrOf() string {
	return t.Pointer.AddrOf()
}

// IsBasic returns true for basic types (string, int etc)
func (t Type) IsBasic() bool {
	_, ok := t.Type.(*types.Basic)
	return ok
}

// Underlying returns the underlying Go type as a Type
func (t Type) Underlying() Type {
	u := t.Type.Underlying()
	return Type{Name: u.String(), Type: u}
}

// LongName provides a name that may be useful for generated names.
// For example, map[string]Foo becomes MapStringFoo.
func (t Type) LongName() string {
	s := strings.Replace(t.String(), "[]", "Slice[]", -1) // hacktastic

	r := regexp.MustCompile(`[\[\]{}*]`)
	els := r.Split(s, -1)

	var parts []string

	for _, s := range els {
		parts = append(parts, strings.Title(s))
	}

	return strings.Join(parts, "")
}

func (t Type) FindTag(name string) (Tag, bool) {
	for _, tag := range t.Tags {
		if tag.Name == name {
			return tag, true
		}
	}
	return Tag{}, false
}

// Pointer exists as a type to allow simple use as bool or as String, which returns *
type Pointer bool

func (p Pointer) String() string {
	if p {
		return "*"
	}
	return ""
}

// ValueOf inverts the star symbol, supplying * where a pointer value needs to be dereferenced.
func (p Pointer) ValueOf() string {
	if p {
		return ""
	}
	return "*"
}

// AddrOf returns & when this is a pointer, allowing a template to take the address of a plain value.
func (p Pointer) AddrOf() string {
	if p {
		return "&"
	}
	return ""
}

type test bool

// a convenience for using bool in file name, see WriteAll
func (t test) String() string {
	if t {
		return "_test"
	}
	return ""
}

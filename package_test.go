package typewriter

import (
	"go/types"
	"testing"
)

func TestEval(t *testing.T) {
	// this'll create a real package with types from this, um, package
	a, err := NewApp("+test")

	if err != nil {
		t.Error(err)
		return // we got problems, continuing will panic
	}

	p := a.Packages[0]

	s1 := "App"
	t1, err1 := p.Eval(s1)

	if err1 != nil {
		t.Error(err1)
	}

	if t1.Pointer {
		t.Errorf("'app' is not a pointer type")
	}

	if t1.Comparable {
		t.Errorf("'app' is not a comparable type")
	}

	if t1.FullyComparable {
		t.Errorf("'app' is not a fully-comparable type")
	}

	if t1.Numeric {
		t.Errorf("'app' is not a numeric type")
	}

	if t1.Ordered {
		t.Errorf("'app' is not an ordered type")
	}

	// embedded types.Type should be accessible via type assertion
	tt1, ok1 := t1.Type.Underlying().(*types.Struct)
	if !ok1 {
		t.Errorf("unable to assert %s as a *types.Struct", t1)
	}

	if tt1.NumFields() != 4 {
		t.Errorf("%s should have 4 fields", tt1)
	}

	s2 := "*App"
	t2, err2 := p.Eval(s2)

	if err2 != nil {
		t.Error(err2)
	}

	if !t2.Pointer {
		t.Errorf("'*app' is a pointer type")
	}

	if !t2.Comparable {
		t.Errorf("'*app' is a comparable type")
	}

	if t2.FullyComparable {
		t.Errorf("'*app' is not a fully-comparable type")
	}

	if t2.Numeric {
		t.Errorf("'*app' is not a numeric type")
	}

	if t2.Ordered {
		t.Errorf("'*app' is not an ordered type")
	}

	s3 := "int"
	t3, err3 := p.Eval(s3)

	if err3 != nil {
		t.Error(err3)
	}

	if t3.Pointer {
		t.Errorf("int is not a pointer type")
	}

	if !t3.Comparable {
		t.Errorf("int is a comparable type")
	}

	if !t3.FullyComparable {
		t.Errorf("int is a fully-comparable type")
	}

	if !t3.Numeric {
		t.Errorf("int is a numeric type")
	}

	if !t3.Ordered {
		t.Errorf("int is an ordered type")
	}

	s4 := "notreal"
	_, err4 := p.Eval(s4)

	if err4 == nil {
		t.Error("'notreal' should not successfully evaluate as a type")
	}
}

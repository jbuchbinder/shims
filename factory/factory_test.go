package factory

import "testing"

type Inst interface {
	InstType() string
}

type InstDummy struct {
}

func (i InstDummy) InstType() string {
	return "dummy"
}

type InstA struct {
}

func (i InstA) InstType() string {
	return "a"
}

type InstB struct {
}

func (i InstB) InstType() string {
	return "b"
}

func Test_Factory(t *testing.T) {
	f := New[Inst](InstDummy{})
	f.Register("a", func() Inst { return new(InstA) })
	f.Register("b", func() Inst { return new(InstB) })

	a, err := f.Instantiate("a")
	if err != nil {
		t.Errorf("could not instantiate InstA")
	}
	if a.InstType() != "a" {
		t.Errorf("failed to instantiate InstA")
	}

	b, err := f.Instantiate("b")
	if err != nil {
		t.Errorf("could not instantiate InstB")
	}
	if b.InstType() != "b" {
		t.Errorf("failed to instantiate InstB")
	}

	_, err = f.Instantiate("c")
	if err == nil {
		t.Errorf("should have failed to instantiate InstC")
	}

}

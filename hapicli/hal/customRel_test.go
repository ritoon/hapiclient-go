package hal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewCustomRel(t *testing.T) {
	data := []struct {
		title  string
		in     string
		outCR  *customRel
		outErr error
	}{
		{"A", "test", &customRel{"test"}, nil},
		{"B", "", nil, ErrNameEmpty},
		{"C", " ", nil, ErrNameEmpty},
		{"D", " another test", &customRel{"another test"}, nil},
	}

	for _, v := range data {
		cr, err := NewCustomRel(v.in)

		if !reflect.DeepEqual(err, v.outErr) {
			t.Error("for", v.title, "waiting", v.outErr, "got", err)
		}

		if err == nil && cr.Name() != v.outCR.Name() {
			t.Error("for", v.title, "waiting", v.outCR, "got", cr)
		}
	}
}

func TestName(t *testing.T) {
	data := []struct {
		title string
		in    *customRel
		out   string
	}{
		{"A", &customRel{"Leonardo"}, "Leonardo"},
		{"B", &customRel{"Donatello"}, "Donatello"},
		{"C", &customRel{"Michelangelo"}, "Michelangelo"},
		{"D", &customRel{"Raphael"}, "Raphael"},
	}
	for _, v := range data {
		name := v.in.Name()
		if name != v.out {
			t.Error("for", v.title, "waiting", v.out, "got", name)
		}
	}
}

func TestString(t *testing.T) {
	data := []struct {
		title string
		in    *customRel
		out   string
	}{
		{"A", &customRel{"Mister Fantastic"}, "Mister Fantastic"},
		{"B", &customRel{"Invisible Woman"}, "Invisible Woman"},
		{"C", &customRel{"Human Torch"}, "Human Torch"},
		{"D", &customRel{"The Thing"}, "The Thing"},
	}
	for _, v := range data {
		name := fmt.Sprint(v.in)
		if name != v.out {
			t.Error("for", v.title, "waiting", v.out, "got", name)
		}
	}
}

func TestTrimSpace(t *testing.T) {
	data := []struct {
		title  string
		in     string
		outCR  *customRel
		outErr error
	}{
		{"A", "test", &customRel{"test"}, nil},
		{"B", "", nil, ErrNameEmpty},
		{"C", " ", nil, ErrNameEmpty},
		{"D", " another test", &customRel{"another test"}, nil},
	}

	for _, v := range data {
		cr, err := trimSpace(v.in)
		if !reflect.DeepEqual(err, v.outErr) {
			t.Error("for", v.title, "waiting", v.outErr, "got", err)
		}

		if err == nil && cr != v.outCR.Name() {
			t.Error("for", v.title, "waiting", v.outCR.Name(), "got", cr)
		}
	}
}

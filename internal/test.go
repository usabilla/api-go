package internal

import (
	"testing"
)

// S contains a testing type T.
type S struct {
	t *testing.T
}

// SR contains a testing type T and the expected result.
type SR struct {
	t        *testing.T
	expected []interface{}
}

// Spec returns a base testing structure.
func Spec(t *testing.T) *S {
	return &S{t: t}
}

// Expect returns the exected result.
func (s *S) Expect(expected ...interface{}) *SR {
	return &SR{t: s.t, expected: expected}
}

// ToEqual tests for equality between actuals and expected.
func (sr *SR) ToEqual(actuals ...interface{}) {
	for index, actual := range actuals {
		if sr.expected[index] != actual {
			sr.t.Errorf("expected %+v to equal %+v", sr.expected[index], actual)
		}
	}
}

// ToNotEqual tests for inequality between actuals and expected.
func (sr *SR) ToNotEqual(actuals ...interface{}) {
	for index, actual := range actuals {
		if sr.expected[index] == actual {
			sr.t.Errorf("expected %+v to not equal %+v", sr.expected[index], actual)
		}
	}
}

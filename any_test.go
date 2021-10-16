package any

import (
	"reflect"
	"testing"
)

var anytests = []struct {
	in  Any
	out string
}{
	{42, "int"},
	{"hello", "string"},
	{3.14159, "float64"},
}

func TestAnyType(t *testing.T) {
	for _, tt := range anytests {
		var i Any = tt.in
		if reflect.TypeOf(i).String() != tt.out {
			t.Errorf("got %T, want %s", i, tt.out)
		}
	}
}

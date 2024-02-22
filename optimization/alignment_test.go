package optimization

import "testing"

type (
	Struct1 struct {
		FieldA int
		FieldB string
		FieldC bool
		C      byte
		FieldD []string
		B      int8
		FieldE interface{}
		A      int8
		FieldF float32
		T      struct{}
	}

	Struct2 struct {
		FieldX float64
		FieldY []int
		FieldZ string
		FieldW map[string]int
		FieldV struct{}
		FieldU byte
	}

	Struct3 struct {
		FieldOne   complex128
		FieldTwo   []bool
		FieldThree rune
		FieldFour  interface{}
		FieldFive  uint
		FieldSix   int64
	}
)

func TestOptimization(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{value: Struct1{}}},
		{name: "2", args: args{value: Struct2{}}},
		{name: "3", args: args{value: Struct3{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Optimization(tt.args.value)
		})
	}
}

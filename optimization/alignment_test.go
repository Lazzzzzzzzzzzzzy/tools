package optimization

import "testing"

type (
	Struct1 struct {
		FieldA int    `json:"field_a" validate:"min=0,max=100"`
		FieldB string `json:"field_b" validate:"required"`
		FieldC bool   `json:"-"`
		C      byte
		FieldD []string    `json:"field_d"`
		B      int8        `json:"-"`
		FieldE interface{} `json:"field_e"`
		A      int8        "asdasddasdsd "
		FieldF float32     `json:"field_f" example:"3.14"`
		T      struct{}    `json:"-"`
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

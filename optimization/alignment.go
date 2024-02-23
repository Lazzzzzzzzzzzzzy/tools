package optimization

import (
	"fmt"
	"reflect"
	"sort"
)

type t struct {
	Name string
	Type string
	Size int
}

type ts []t

func (o ts) Len() int {
	return len(o)
}

func (o ts) Less(i, j int) bool {
	if o[i].Size <= o[j].Size {
		return true
	}
	return false
}

func (o ts) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

// Optimization 用于结构体内存对齐，将结构体字段按照大小从小到大排列。
func Optimization(value interface{}) {
	v := reflect.TypeOf(value)
	if v.Kind() != reflect.Struct {
		fmt.Println("value not struct")
		return
	}
	s := []t{}
	for i := 0; i < v.NumField(); i++ {
		s = append(s, t{
			Name: v.Field(i).Name,
			Type: v.Field(i).Type.String(),
			Size: int(v.Field(i).Type.Size()),
		})
	}
	sort.Sort(ts(s))
	fmt.Println()
	s1 := []reflect.StructField{}
	for _, info := range s {
		name, _ := v.FieldByName(info.Name)
		s1 = append(s1, name)
	}
	for _, field := range s1 {
		fmt.Println(field.Name, field.Type.String())
	}
}

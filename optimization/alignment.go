package optimization

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

const (
	quotationMarks = "\""
	backtick       = "`"
)

// Optimization 用于结构体内存对齐，将结构体字段按照大小从小到大排列。
func Optimization(value interface{}) {

	v := reflect.TypeOf(value)
	if v.Kind() != reflect.Struct {
		fmt.Println("value not struct")
		return
	}
	// 获取结构体字段并排序
	s := getFieldAndSort(v)
	// 输出结果可直接替换原有结构体
	printResult(s, v.Name())

}

// getFieldAndSort 获取结构体字段并按照大小排序
func getFieldAndSort(v reflect.Type) ts {
	s := make(ts, 0, 8)

	for i := 0; i < v.NumField(); i++ {
		sf := v.Field(i)
		s = append(s, t{
			Name: sf.Name,
			Type: sf.Type.String(),
			Size: int(sf.Type.Size()),
		})
		if len(sf.Tag) == 0 {
			continue
		}
		if strings.Contains(string(sf.Tag), quotationMarks) {
			s[i].Tag = backtick + string(sf.Tag) + backtick
		} else {
			s[i].Tag = quotationMarks + string(sf.Tag) + quotationMarks
		}
	}
	sort.Sort(s)

	return s
}

func printResult(s ts, name string) {

	fmt.Println("\ntype " + name + " struct{")
	for _, field := range s {
		fmt.Println(field.Name, field.Type, field.Tag)
	}
	fmt.Println("}")
	fmt.Println()
}

type t struct {
	Name string
	Type string
	Size int
	Tag  string
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

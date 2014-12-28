package shuffler

import (
	"math/rand"
	"reflect"
	"time"
)

// case bool:
// case float32:
// case float64:
// case int:
// case int8:
// case int16:
// case int32:
// case int64:
// case uint:
// case uint8:
// case uint16:
// case uint32:
// case uint64:
// case string:

func Shuffle(ref interface{}) {
	switch ref.(type) {
	case []bool:
		ary := ref.([]bool)
		t := Shuffler(ary).(reflect.Value)
		for i := 0; i < t.Len(); i++ {
			ary[i] = bool(t.Index(i).Bool())
		}
	case []float32:
		ary := ref.([]float32)
		t := Shuffler(ary).(reflect.Value)
		for i := 0; i < t.Len(); i++ {
			ary[i] = float32(t.Index(i).Float())
		}
	case []float64:
		ary := ref.([]float64)
		t := Shuffler(ary).(reflect.Value)
		for i := 0; i < t.Len(); i++ {
			ary[i] = float64(t.Index(i).Float())
		}
	case []int:
		ary := ref.([]int)
		t := Shuffler(ary).(reflect.Value)
		for i := 0; i < t.Len(); i++ {
			ary[i] = int(t.Index(i).Int())
		}
	case []int8:
		ary := ref.([]int8)
		t := Shuffler(ary).(reflect.Value)
		for i := 0; i < t.Len(); i++ {
			ary[i] = int8(t.Index(i).Int())
		}
	case []int16:
		ary := ref.([]int16)
		t := Shuffler(ary).(reflect.Value)
		for i := 0; i < t.Len(); i++ {
			ary[i] = int16(t.Index(i).Int())
		}
	case []int32:
		ary := ref.([]int32)
		t := Shuffler(ary).(reflect.Value)
		for i := 0; i < t.Len(); i++ {
			ary[i] = int32(t.Index(i).Int())
		}
	case []int64:
		ary := ref.([]int64)
		t := Shuffler(ary).(reflect.Value)
		for i := 0; i < t.Len(); i++ {
			ary[i] = int64(t.Index(i).Int())
		}
	case []uint:
		ary := ref.([]uint)
		t := Shuffler(ary).(reflect.Value)
		for i := 0; i < t.Len(); i++ {
			ary[i] = uint(t.Index(i).Uint())
		}
	case []uint8:
		ary := ref.([]uint8)
		t := Shuffler(ary).(reflect.Value)
		for i := 0; i < t.Len(); i++ {
			ary[i] = uint8(t.Index(i).Uint())
		}
	case []uint16:
		ary := ref.([]uint16)
		t := Shuffler(ary).(reflect.Value)
		for i := 0; i < t.Len(); i++ {
			ary[i] = uint16(t.Index(i).Uint())
		}
	case []uint32:
		ary := ref.([]uint32)
		t := Shuffler(ary).(reflect.Value)
		for i := 0; i < t.Len(); i++ {
			ary[i] = uint32(t.Index(i).Uint())
		}
	case []uint64:
		ary := ref.([]uint64)
		t := Shuffler(ary).(reflect.Value)
		for i := 0; i < t.Len(); i++ {
			ary[i] = uint64(t.Index(i).Uint())
		}
	case []string:
		ary := ref.([]string)
		t := Shuffler(ary).(reflect.Value)
		for i := 0; i < t.Len(); i++ {
			ary[i] = string(t.Index(i).String())
		}
	}
}

func Shuffler(src interface{}) interface{} {
	rand.Seed(time.Now().UTC().UnixNano())

	s := reflect.ValueOf(src)
	t := reflect.TypeOf(src)

	dest := reflect.MakeSlice(reflect.SliceOf(t.Elem()), 0, 0)
	for _, v := range rand.Perm(s.Len()) {
		dest = reflect.Append(dest, s.Index(v))
	}

	return dest
}

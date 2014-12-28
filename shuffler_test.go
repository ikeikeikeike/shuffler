package shuffler

import (
	"reflect"
	"testing"
)

func TestIntSlice(t *testing.T) {
	a := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	Shuffle(a)

	if reflect.DeepEqual(a, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Errorf("Error could not shuffle []int")
	}
}

func TestStringSlice(t *testing.T) {
	a := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}

	Shuffle(a)

	if reflect.DeepEqual(a, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}) {
		t.Errorf("Error could not shuffle []string")
	}
}

func TestFloat32Slice(t *testing.T) {
	a := []float32{1.0, 1.1, 1.2, 1.3, 1.4, 2.0, 2.1, 2.2, 2.3}

	Shuffle(a)

	if reflect.DeepEqual(a, []float32{1.0, 1.1, 1.2, 1.3, 1.4, 2.0, 2.1, 2.2, 2.3}) {
		t.Errorf("Error could not shuffle []float32")
	}
}

func TestFloat64Slice(t *testing.T) {
	a := []float64{1.0, 1.1, 1.2, 1.3, 1.4, 2.0, 2.1, 2.2, 2.3}

	Shuffle(a)

	if reflect.DeepEqual(a, []float64{1.0, 1.1, 1.2, 1.3, 1.4, 2.0, 2.1, 2.2, 2.3}) {
		t.Errorf("Error could not shuffle []float64")
	}
}

func TestBoolSlice(t *testing.T) {
	a := []bool{
		true, true, true, true, true, true,
		false, false, false, false, false, false,
	}

	Shuffle(a)

	if reflect.DeepEqual(a, []bool{true, true, true, true, true, true, false, false, false, false, false, false}) {
		t.Errorf("Error could not shuffle []bool")
	}
}

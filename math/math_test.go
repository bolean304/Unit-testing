package math

import "testing"

type AddData struct {
	name   string
	x, y   int
	result int
}

var testdata = []AddData{
	{"a", 1, 2, 3},
	{"b", 2, 3, 5},
	{"c", 7, 4, 13},
}

func TestAdd(t *testing.T) {

	for _, data := range testdata {
		//result := Add(data.x, data.y)
		//if result != data.result {
		//	t.Errorf("add(%d,%d) is failed. expected %d , got %d", data.x, data.y, data.result, result)
		//} else {
		//	t.Logf("add(%d,%d) is passed. expected %d , got %d", data.x, data.y, data.result, result)
		//}
		t.Run(data.name, func(t *testing.T) {
			result := Add(data.x, data.y)
			if result != data.result {
				t.Errorf("add(%d,%d) is failed. expected %d , got %d", data.x, data.y, data.result, result)
			} else {
				t.Logf("add(%d,%d) is passed. expected %d , got %d", data.x, data.y, data.result, result)
			}
		})

	}
}

//
//func TestDivide(t *testing.T) {
//
//	result := Divide(5, 0)
//
//	if result != 0 {
//		t.Logf("%f", result)
//		t.Errorf("divide(5,4) is failed. expected %f, got %f", 1.25, result)
//	} else {
//		t.Logf("divide(5,4) is passed. expected %f, got %f", 1.25, result)
//	}
//}

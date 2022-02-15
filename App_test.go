package main
import 
	"testing"
type AddData struct {
	x, y int
	result float64
}

func TestPlus(t *testing.T){
var x, y, result = 6, 7, 13

realResult := Plus(x, y)

	if realResult != result {
	t.Errorf("result %d != check %d", realResult, result)
	}
}
func TestMinus(t *testing.T){
var x, y, result = 7, 6, 1
	
realResult := Minus(x, y)
	
	if realResult != result {	
	t.Errorf("result %d != check %d", realResult, result)
	}
}
func TestMult(t *testing.T){
var x, y, result = 6, 7, 42
		
realResult := Mult(x, y)
		
	if realResult != result {
	t.Errorf("result %d != check %d", realResult, result)
	}
}		
func TestDiv(t *testing.T){
testData := []AddData {
	{1, 2, 3},
	{3, 5, 8},
	{7, -4,3},
}
	for _, datum := range testData{
		result := Div(datum.x, datum.y)
	if result != datum.result{
		t.Errorf("Add(%d, %d) Failed. Eksected %v got %v", datum.x, datum.y, datum.result, result)
	}
	}
} 
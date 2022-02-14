package main
import 
	"testing"

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
var x, y, result = 54, 6, 9
	
realResult := Div(x, y)
	
	if realResult == 0 {
	t.Errorf("realResult: %d, result: %d", realResult, result)
	}	
	if realResult != result {
	t.Errorf("result %d != check %d", realResult, result)
	}
}
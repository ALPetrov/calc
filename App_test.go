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


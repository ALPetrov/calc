package main

import (
	"testing"

	"github.com/go-playground/assert"
)
type Data struct {
	x, y int
	result int
}

func TestPlus(t *testing.T){
testPlusData := [] Data {
	{1, 2, 3},
	{-1, 2, 1},
	{5, 4, 9},
}
	for _, datum := range testPlusData{
		result := Plus(datum.x, datum.y)
	
	assert.Equal(t, result, datum.result)
	}
}
func TestMinus(t *testing.T){
	testMinusData := [] Data {
		{3, 2, 1},
		{-1, 2, -3},
		{12, 4, 8},
	}
		for _, datum := range testMinusData{
			result := Minus(datum.x, datum.y)
		
		assert.Equal(t, result, datum.result)
	}
}
func TestMult(t *testing.T){
	testMultData := [] Data {
		{1, 2, 2},
		{-1, 2, -2},
		{5, 4, 20},
	}
		for _, datum := range testMultData{
			result := Mult(datum.x, datum.y)
		
		assert.Equal(t, result, datum.result)
	}
}		
func TestDiv(t *testing.T){
testDivData := [] Data {
	{4, 2, 2},
	{-42, 7, -6},
	{12, 4, 3},
}
	for _, datum := range testDivData{
		result := Div(datum.x, datum.y)
	
	assert.Equal(t, result, datum.result)
	}
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	
	gpaSlice := []float64{4.1, 4.5, 3.9} // slice literal
	fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

	// gpaSlice := []float64{4.1, 4.5, 3.9} // slice literal
	// fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

	// 
	// 	gpas := [5]float64{3.5, 4.1, 4.5, 3.9, 4.23} // array := array literal
	// 	fmt.Println(gpas, reflect.TypeOf(gpas))
	// 	gpaSlice := gpas[1:4] // slice := slicing array
	// 	fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))
}

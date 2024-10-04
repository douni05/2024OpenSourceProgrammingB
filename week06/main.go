package main

import (
	"fmt" // c 언어 #include <stdio.h>
	"reflect"
)

func main() {
	i := 13
	var f float64 = 12.9 // f := 12.9

	fmt.Printf("value i : %d, value f : %f\n", i, f)
	//fmt.Printf("%d * %f = %f\n", i, f, i*f)  // invalid operation: i * f (mismatched types int and float64) : go에서 묵시적 형변환 X
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f) // conversion
	fmt.Printf("%d * %f = %d\n", i, f, i*int(f))
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
}

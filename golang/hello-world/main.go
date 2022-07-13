package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func main() {
	//hello()
	//variables()
	//typeConversion()
	//usingMyParseBool()
	//handleErrors()
	//pointers()
	doc()
}

func hello() {
	fmt.Println("Hello, Go!")
}

func variables() {
	var myInteger int
	myInteger = 1

	var myOtherInteger = 2

	shortInteger := 3

	fmt.Println("I'm an integer ", reflect.TypeOf(myInteger))
	fmt.Println("Simple sum ", myOtherInteger+shortInteger)
}

func typeConversion() {
	var length float64 = 1.2
	var width int = 2
	// But you can if you do type conversions!
	fmt.Println("Area is", length*float64(width))
}

func myParseBool(myString string) (bool, error) {
	if myString == "true" {
		return true, nil
	} else if myString == "false" {
		return false, nil
	} else {
		return false, fmt.Errorf("bad string %s", myString)
	}
}

func usingMyParseBool() {
	myBool, err := myParseBool("false")
	if err != nil {
		log.Fatal("A fatal error usingMyParseBool")
	}
	fmt.Println(myBool)
}

func handleErrors() {
	quotient, divideError := divide(5.6, 2)
	if divideError != nil {
		log.Fatal("A Fatal error useDivide")
	}
	fmt.Printf("%0.2f\n", quotient) // => 2.80
}

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by 0")
	}

	return a / b, nil
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func pointers() {
	truth := true
	negate(&truth)
	fmt.Println(truth)

	lies := false
	negate(&lies)
	fmt.Println(lies)
}

func doc() {
	//  go doc strconv Parsefloat
	string1 := "12.345"
	string2 := "1.234"

	number1, err := strconv.ParseFloat(string1, 64)

	if err != nil {
		log.Fatal("Could not parse string")
	}

	number2, err := strconv.ParseFloat(string2, 64)
	if err != nil {
		log.Fatal("Could not parse string")
	}

	fmt.Println(number1 - number2)
}

package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func main() {
	hello()
	variables()
	typeConversion()
	usingMyParseBool()
	handleErrors()
	pointers()
	doc()
	slices()
	useAnonymousFieldsOnTypes()
	useMethods()
	useInterfaces()
	useHandlingErrors()
	useChannels()
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

func slices() {
	// Create a variable to hold a slice of ints.
	var primes []int
	// Create a slice with 2 elements.
	primes = make([]int, 2)
	// Assign values to the first 2 elements.
	primes[0] = 2
	primes[1] = 3
	// Add a third element to the end of the slice.
	primes = append(primes, 5)
	fmt.Println(primes) // => [2 3 5]

	// Write a map literal with int keys and string values.
	elements := map[int]string{1: "H", 2: "He", 3: "Li"}
	// Loop over each key/value pair in the map.
	for atomicNumber, symbol := range elements {
		fmt.Println(atomicNumber, symbol)
	}
	// => 1 H
	// => 2 He
	// => 3 Li
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}
type Employee struct {
	Name   string
	Salary float64
	Address
}
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

// YOUR CODE HERE:
// Define a struct type named Address that has Street, City, State,
// and PostalCode fields, each with a type of "string".
// Then embed the Address type within the Subscriber and Employee
// types using anonymous fields, so that the code in "main" will
// compile, run, and produce the output shown.

func useAnonymousFieldsOnTypes() {
	var subscriber Subscriber
	subscriber.Name = "Aman Singh"
	subscriber.Street = "123 Oak St"
	subscriber.City = "Omaha"
	subscriber.State = "NE"
	subscriber.PostalCode = "68111"
	fmt.Println(subscriber)

	var employee Employee
	employee.Name = "Joy Carr"
	employee.Street = "456 Elm St"
	employee.City = "Portland"
	employee.State = "OR"
	employee.PostalCode = "97222"
	fmt.Println(employee)
}

// YOUR CODE HERE:
// Define a Rectangle struct type with Length and Width
// fields, each of which has a type of float64.
type Rectangle struct {
	Length float64
	Width  float64
}

// YOUR CODE HERE:
// Define an Area method on the Rectangle type. It should
// accept no parameters (other than the receiver parameter).
// It should return a float64 value calculated by multiplying
// the receiver's Length by its Width.
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// YOUR CODE HERE:
// Define a Perimeter method on the Rectangle type. It should
// accept no parameters. It should return a float64 value
// representing the receiver's perimeter (2 times its Length
// plus 2 times its Width)
func (r Rectangle) Perimeter() float64 {
	return (2 * r.Length) + (2 * r.Width)
}

func useMethods() {
	var myRectangle Rectangle
	myRectangle.Length = 2
	myRectangle.Width = 3
	fmt.Println("Area:", myRectangle.Area())           // => Area: 6
	fmt.Println("Perimeter:", myRectangle.Perimeter()) // => Perimeter: 10
}

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}
func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

// YOUR CODE HERE:
// Define a NoiseMaker interface type, which the above
// Whistle, Horn, and Robot types will all satisfy.
// It should require one method, MakeSound, which has
// no parameters and no return values.
type NoiseMaker interface {
	MakeSound()
}

// YOUR CODE HERE:
// Define a Play function that accepts a parameter with
// the NoiseMaker interface. Play should call MakeSound
// on the parameter it receives.
func Play(n NoiseMaker) {
	n.MakeSound()
}

func useInterfaces() {
	Play(Whistle("Toyco Canary")) // => Tweet!
	Play(Horn("Toyco Blaster"))   // => Honk!
	Play(Robot("Botco Ambler"))   // => Beep Boop
}

type Refrigerator struct {
	Brand string
}
type Food string

func (r Refrigerator) Open() {
	fmt.Println("Opening refrigerator")
}
func (r Refrigerator) Close() {
	fmt.Println("Closing refrigerator")
}
func (r Refrigerator) FindFood(food string) (Food, error) {
	// Food storage not implemented yet; always return error!
	// Note: don't change FindFood as part of this exercise!
	return Food(""), fmt.Errorf("%s not found", food)
}

// YOUR CODE HERE:
// Modify the code in the Eat function so that fridge.Close will
// always be called at the end, even if fridge.FindFood returns
// an error. Once you've figured the solution out, your changes
// will actually be quite small! Note: it wouldn't be appropriate
// to use either "panic" or "recover" in this exercise; we won't
// be using either one.
func Eat(fridge Refrigerator) error {
	fridge.Open()
	defer fridge.Close()

	food, err := fridge.FindFood("bananas")
	if err != nil {
		return err
	}
	fmt.Println("Eating", food)

	return nil
}

func useHandlingErrors() {
	var fridge Refrigerator
	err := Eat(fridge)
	if err != nil {
		fmt.Println(err)
	}
}

// This program should call the "repeat" function twice, using two
// separate goroutines. The first goroutine should print the string
// "x" repeatedly, and the second goroutine should print "y"
// repeatedly. You'll also need to create a channel that carries
// boolean values to pass to "repeat", so the goroutine can signal
// when it's done.
//
// Output will vary, but here's one possible result:
// yyyyyyyyyyyyyyyyyyyyxxxxxxxxxxxxxxxxxxxxxyyyyyyyyyxxxxxxxxxy
// repeat prints a string multiple times, then writes "true" to  the
// provided channel to signal it's done.
func repeat(s string, channel chan bool) {
	for i := 0; i < 30; i++ {
		fmt.Print(s)
	}
	channel <- true
}

func useChannels() {
	channel := make(chan bool)
	go repeat("x", channel)
	go repeat("y", channel)
	x := <-channel
	<-channel
	fmt.Println(x)
}

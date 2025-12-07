// Syntax & Basic Types
package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

// Config represents application configuration
type Config struct {
	Name  string
	Value string
}

func main() {
	name := "John Paul"
	age := 28
	isActive := true

	// type inference with var
	city := "Valencia City"

	// := short declaration (only inside functions)
	country := "Philippines"
	score := 95.5

	// Multiple declarations
	x, y, z := 1, 2, 3
	a, b := "Hello", 42

	email := "john@gmail.com"

	i := 42
	var i8 int8 = 82

	const pi float32 = 12312313212132133.14
	const pi2 float32 = 12312313212132133.1412321313133123123123123122222222

	const (
		StatusOk       = 200
		StatusNotFound = 404
	)

	// iota — auto-incrementing constant generator (no JS equivalent)
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	// Common pattern: bit flags
	const (
		ReadPermission    = 1 << iota // 1
		WritePermission               // 2
		ExecutePermission             // 4
	)

	// Skip values
	const (
		_  = iota             // 0 discarded
		KB = 1 << (10 * iota) // 1 << 20
		MB                    // 1 << 20
		GB                    // 2 << 30
	)

	fmt.Println(name, age, isActive, city, country, score)
	fmt.Println(x, y, z)
	fmt.Println(a, b)
	fmt.Println(email)
	fmt.Println(i)
	fmt.Println(i8)
	fmt.Println(pi)
	fmt.Println(pi2)
	fmt.Println(StatusOk)
	fmt.Println(StatusNotFound)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	fmt.Println(ReadPermission, WritePermission, ExecutePermission)

	// Go is strict - no magic conversions
	ci := 42
	cf := 3.14

	// this won't compile:
	// sum := ci + cf // Mismatch  types int and float64

	// You must explicitly convert:
	sum := float64(ci) + cf

	// String concatenation - no auto-conversion
	cAge := 28
	// cMessage := "Age: " + age // WON'T compile
	cMessage := "Age " + fmt.Sprint(cAge)
	fmt.Println(sum, cMessage)

	// Control Structures (if statements)
	dAge := 28

	if dAge >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	// if with initialization statement (very Go-idiomatic)
	// if score := calculateScore(); score > 90 {
	// 	fmt.Println("Excellent!")
	// } else if score > 70 {
	// 	fmt.Println("Good")
	// } else {
	// 	fmt.Println("Keep trying.")
	// }

	// // score is not accessible here - scope to the if block
	// // Common pattern: error checking
	// if err := doSomething(); err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }

	// for loop (Go's only Loop)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While-style loop (for with just condition)
	fCount := 0
	for fCount < 5 {
		fmt.Println(fCount)
		fCount++
	}

	// Infinite loop
	// for {
	//     fmt.Println("Forever")
	//     break  // Use break to exit
	// }

	fruits := []string{"Apple", "Banana", "Cherry"}

	for index, value := range fruits {
		fmt.Printf("%d %s \n", index, value)
	}

	// Ignore index with _
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// Range over map
	gAges := map[string]int{"alice": 30, "bob": 20}
	for name, age := range gAges {
		fmt.Printf("%s is %d\n", name, age)
	}

	// Range over string (iterates over runes, not bytes)
	for i, char := range "Hello" {
		fmt.Printf("%d %c\n", i, char)
	}

	// Switch Statement
	// Much more powerful than JS switch.
	day := "Monday"

	// Basic switch (no break needed! No fall-through by default)
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("TGIF!")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Midweek")
	}

	// Switch with initialization
	switch today := time.Now().Weekday(); today {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// Switch without expression (cleaner than if-else chains)
	gScore := 85

	switch {
	case gScore >= 90:
		fmt.Println("A")
	case gScore >= 80:
		fmt.Println("B")
	case gScore >= 70:
		fmt.Println("C")
	default:
		fmt.Println("F")
	}

	// Explicit fallthrough (rare, must be intentional)
	switch num1 := 1; num1 {
	case 1:
		fmt.Println("one")
		fallthrough // continues to the next case
	case 2:
		fmt.Println("two")
	}
	// Output: One, Two

	// Type switch (for interfaces)
	var xx interface{} = 42
	switch v := xx.(type) {
	case int:
		fmt.Println("Integer: ", v)
	case string:
		fmt.Println("String: ", v)
	default:
		fmt.Println("Unknown type")
	}

	// defer Statement
	// Unique to Go — schedules a function call to run when the surrounding function returns.
	// Executes in LIFO order (last defer runs first).

	// defer fmt.Println("This run last [defer]")
	// fmt.Println("This run first")

	// xX := 10
	// defer fmt.Println(xX)

	// xX = 20
	// fmt.Println(xX)
	// Output:
	// 20
	// 10  (deferred call used the captured value)

	// To use the final value, wrap in a closure:
	// defer func() { fmt.Println(xX) }() // Closure captures variable, not value // output: 20
	// xX = 20

	// Load another main
	loadAnotherMain()
}

// defer - common use cases:

// 1. Closing resources (like finally in JS)
func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	// ...read file...
	return nil
}

// 2. Unlocking mutexes
func safeIncrement(mu *sync.Mutex, counter *int) {
	mu.Lock()
	defer mu.Unlock() // always unlocks even if panic
	*counter++
}

// 3. Timing functions
func expensiveOperation() {
	defer timeTrack(time.Now(), "expensiveOperation")
	// ... do work ...
}

func timeTrack(start time.Time, name string) {
	fmt.Printf("%s took %v\n", name, time.Since(start))
}

// Functions
// Basic Functions & Multiple Returns
func greet(name string) string {
	return "Hello, " + name
}

// Multiple parameters of same type
func add(a, b int) int {
	return a + b
}

// Multiple return values (VERY common in GO)
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}

	return a / b, nil
}

// Named Return Values

// Named returns — useful for documentation and naked returns
func rectangleProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // Naked return - returns area and perimeter
}

// better for clarity in complex functions
func parseConfig(data []byte) (config Config, err error) {
	// config and err are already declared
	// useful with defer for error handling
	defer func() {
		if err != nil {
			err = fmt.Errorf("parseConfig: %w", err)
		}
	}()

	// ...parsing logic...
	return
}

// Variadic Functions

// Accepts zero or more ints
func sum(nums ...int) int {
	total := 0

	for _, n := range nums {
		total += n
	}

	return total
}

// usage
// fmt.Println(sum())           // 0
// fmt.Println(sum(1, 2))       // 3
// fmt.Println(sum(1, 2, 3, 4)) // 10

// Spread a slice with ...
// numbers := []int{1, 2, 3, 4, 5}
// fmt.Println(sum(numbers...)) // 15

// Common: variadic with other params (variadic must be last)
// func printf(format string, args ...interface{}) {
// ...
// }

// Functions as First-Class Citizens
// Just like JS, functions are values.

func loadAnotherMain() {
	fmt.Println("\n\n\n\n\n\n\n\n\n")
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(2, 3)) // 5

	result := operate(10, 5, add)
	fmt.Println(result) // expect: 15

	// Return function from function
	multiplier := makeMultiplier(3)
	fmt.Println(multiplier(4)) // 12

	next := counter()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	// Another call to counter() creates a new closure
	another := counter()
	fmt.Println(another())

	// Immediately Invoked Function Expression (IIFE)

	resultN := func(x int) int {
		return x * x
	}(5)

	fmt.Println(resultN)

	// Useful for goroutines
	go func(msg string) {
		fmt.Println(msg)
	}("Hello from goroutine")
}

// Functions as parameter
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// Functions returning function (closure)
func makeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor // closes over 'factor'
	}
}

// Closures
// Work just like in JS.
func counter() func() int {
	count := 0
	return func() int {
		count++

		return count
	}
}

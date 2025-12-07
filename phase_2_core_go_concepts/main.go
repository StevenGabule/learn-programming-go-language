package main

import (
	"encoding/json"
	"fmt"
	"math"
	"slices"
)

func compositeTypesArrayFixSize() {
	// Declare array with size
	var numbers [5]int // [0, 0, 0,0,0] - zero values
	numbers[0] = 10
	numbers[1] = 20
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	// Array literal
	fruits := [3]string{"apple", "banana", "cherry"}
	colors := [...]string{"red", "green", "blue"} // size inferred as 3

	// Important: array size is part of the type!
	var a [3]int
	var b [4]int
	// a = b  // WON'T compile — different types!

	// Arrays are VALUES (copied on assignment)
	original := [3]int{1, 2, 3}
	copy := original // Creates a full copy
	copy[0] = 999
	fmt.Println(original) // [1 2 3] — unchanged!
	fmt.Println(copy)     // [999, 2, 3]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(fruits)
	fmt.Println(colors)
}

// Slices (What You'll Actually Use)
// Slices are dynamic, flexible views into arrays. This is Go's answer to JS arrays.
func compositeTypesSlices() {
	// Create slice with make(type, length, capacity)
	s := make([]int, 3, 5)         // len 3, cap=5
	fmt.Println(s, len(s), cap(s)) // [0 0 0] 3 5

	// Slice literal
	fruits := []string{"apple", "banana", "cherry"}

	// Nil slice (zero value)
	var empty []int
	fmt.Println(empty == nil) // true

	// Append (like JS push but returns new slice)
	fruits = append(fruits, "date")
	fruits = append(fruits, "elderberry", "fig")

	// append another slice
	more := []string{"grape", "honeydew"}
	fruits = append(fruits, more...)

	// Slicing (like JS Slice)
	nums := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(nums[1:4]) // [1 2 3] — index 1 to 3 (exclusive end)
	fmt.Println(nums[:3])  // [0 1 2] — from start
	fmt.Println(nums[3:])  // [3 4 5] — to end
	fmt.Println(nums[:])   // [0 1 2 3 4 5] — full copy reference

	// IMPORTANT: Slices share underlying array!
	original := []int{1, 2, 3, 4, 5}
	slice := original[1:4]
	slice[0] = 999

	fmt.Println(original) // [1 999 3 4 5]
	fmt.Println(fruits)   // [apple banana cherry date elderberry fig grape honeydew]
}

func compositeTypesSlicesInternal() {
	s := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// len=3 cap=5 [0 0 0]

	s = append(s, 1) // Fits in capacity
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// len=4 cap=5 [0 0 0 1]

	s = append(s, 2) // Still fits
	s = append(s, 3) // Exceeds capacity — Go allocates new, larger array!
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// len=6 cap=10 [0 0 0 1 2 3] — capacity doubled
}

func compositeTypesSliceOperations() {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}

	// copy (to avoid shared backing array)
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums) // built-in-copy function

	// Or with Go 1.21+
	numsCopy2 := slices.Clone(nums)

	// Delete element at index (no built-in, common pattern)
	i := 2 // delete index 2
	nums = append(nums[:i], nums[i+1:]...)
	fmt.Println(nums) // [3 1 1 5 9 2 6]

	// Insert at index
	nums = []int{1, 2, 4, 5}
	i = 2
	nums = append(nums[:i], append([]int{3}, nums[i:]...)...)
	fmt.Println(nums) // [1 2 3 4 5]

	// With Go 1.21+ slices package
	nums = slices.Insert(nums, 2, 99)

	// Sort (Go 1.21+)
	data := []int{3, 1, 4, 1, 5}
	slices.Sort(data)
	fmt.Println(data) // [1,1,3,4,5]

	// Contains
	fmt.Println("Contains: ", slices.Contains(data, 4)) // true
	fmt.Println("nums: ", nums)                         //[1 2 99 3 4 5]
	fmt.Println("numsCopy2: ", numsCopy2)               // [3 1 4 1 5 9 2 6]

	slices.Reverse(data)
}

// Maps
// Go maps are like JS objects/Maps — key-value pairs with O(1) lookup.
func compositeTypesMaps() {
	// Create with make
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25

	scores := map[string]int{
		"Alice": 95,
		"Bob":   85,
		"Carol": 92, // trailing comma required!
	}

	// Access
	fmt.Println(scores["Alice"])   // 95
	fmt.Println(scores["Unknown"]) // 0 (zero value, NOT an error)

	// Check if key exists (comma-ok idiom)
	score, exists := scores["Dave"]

	if exists {
		fmt.Println("Dave's score: ", score)
	} else {
		fmt.Println("Dave not found.")
	}

	// Common pattern: check and use in one line
	if score, ok := scores["Alice"]; ok {
		fmt.Println("Alice's score: ", score)
	}

	// Delete
	delete(scores, "Bob")

	// Iterate (order is random!)
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}

	// Nil map gotcha
	var nilMap map[string]int
	fmt.Println(nilMap["key"]) // 0 (reading is ok)
	// nilMap["key"] = 1	   // PANIC! Can't write to nil map

	// Nested maps
	users := map[string]map[string]string{
		"user1": {
			"name":  "Alice",
			"email": "someone@gmail.com",
		},
	}

	fmt.Println(users["user1"]["name"])
}

/*
 * Structs
 * Structs are Go's way to create custom types — similar to classes but without inheritance.
 */

type Person struct {
	Name   string
	Age    int
	Email  string
	Active bool
}

type Company struct {
	Name      string
	CEO       Person
	Employees []Person
}

func composeStructs() {
	// Create struct instances

	// 1. Zero value (all fields are zero values)
	var p1 Person
	fmt.Println(p1) // {0 false}

	// 2. Named fields (preferred - order doesn't matter)
	p2 := Person{
		Name:   "John Paul Gabule",
		Age:    28,
		Email:  "jp@example.com",
		Active: true,
	}

	// 3. Position (fragile, avoid for structs with many fields)
	p3 := Person{
		"Alice",
		30,
		"alice@example.com",
		true,
	}

	fmt.Println(p2.Name)
	p2.Age = 29

	// Struct comparison (if all fields are comparable)
	p4 := Person{Name: "Alice", Age: 30, Email: "alice@example.com", Active: true}
	fmt.Println(p3 == p4) // true

	// Anonymous struct (useful for one-off data)
	point := struct {
		X, Y int
	}{10, 20}
	fmt.Println(point.X, point.Y)

	// Nested structs
	company := Company{
		Name: "TechCorp",
		CEO:  Person{Name: "Bob", Age: 25},
		Employees: []Person{
			{Name: "Alice", Age: 30},
			{Name: "Carol", Age: 25},
		},
	}

	fmt.Println(company.CEO.Name)
}

// ***Struct Embedding (Composition over Inheritance)

type Address struct {
	Street  string
	City    string
	Country string
}

type Employee struct {
	Name string
	Age  int
	Address
}

func structEmbedding() {
	e := Employee{
		Name: "John Paul",
		Age:  28,
		Address: Address{
			Street:  "123 Main st.",
			City:    "Valencia City",
			Country: "Philippines",
		},
	}

	// Access embedded fields directly (promoted)
	fmt.Println(e.City)
	fmt.Println(e.Address.City)

	// This is composition, NOT inheritance
	// Employee "has an" Address, not "is an" Address
}

type User struct {
	ID        int    `json:"id" db:"user_id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email" validate:"required,email"`
	Password  string `json:"-" db:"password_hash"` // "-" = omit from JSON
}

func compositionStructTags() {
	user := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Gabule",
		Email:     "jp@example.com",
		Password:  "secret",
	}

	// JSON encoding respects tags
	data, _ := json.Marshal(user)
	fmt.Println(string(data))
}

// ***Pointers
// This is likely new territory coming from JS. JavaScript hides memory management;
// Go makes it explicit.

func compositionPointers() {
	x := 42

	// & = "address of" - get pointer to x
	p := &x

	fmt.Println(p)  // 0xc0000120c8 (memory address)
	fmt.Println(*p) // 42 - dereference to get value

	// Modify through pointer
	*p = 100
	fmt.Println(x) // 100 - x changed!

	// Pointer Type
	var ptr *int     // Pointer to int (nil by default)
	fmt.Println(ptr) // <nil>

	ptr = &x
	fmt.Println(*ptr) // 100

	// new() creates a pointer to zero value
	numPtr := new(int) // *int pointing to 0
	*numPtr = 42
}

// ***Why Pointers Matter: Value vs Reference

type Person1 struct {
	Name string
	Age  int
}

func pointersMatter() {
	x := 10

	doubleValue(x)
	fmt.Println(x) // 10 — unchanged!

	doublePointer(&x)
	fmt.Println(x) // 20 — modified!
}

func pointersWithStructs() {
	person := Person1{Name: "Alice", Age: 30}

	birthdayValue(person)
	fmt.Println(person.Age) // 30 — unchanged

	birthdayPointer(&person)
	fmt.Println(person.Age) // 31 — modified

	// Creating pointer to struct
	p := &Person1{Name: "Bob", Age: 25}
	fmt.Println(p.Name)    // "Bob" — auto-dereference
	fmt.Println((*p).Name) // Same thing, explicit
}

// WITHOUT pointer — function receives a COPY
func doubleValue(n int) {
	n = n * 2
}

// WITH pointer — function receives address, can modify original
func doublePointer(n *int) {
	*n = *n * 2 // Modifies original
}

// Without pointer — receives copy, changes don't persist
func birthdayValue(p Person1) {
	p.Age++ // Modifies copy
}

// With pointer — modifies original
func birthdayPointer(p *Person1) {
	p.Age++ // Go automatically dereferences for struct fields
	// Same as: (*p).Age++
}

// ***When to Use Pointers
// 1. When you need to modify the original
// func updateUser(u *User) {
//     u.Name = "Updated"
// }

// 2. For large structs (avoid copying overhead)
// type LargeStruct struct {
//     Data [1000000]int  // 8MB!
// }

// func processLarge(ls *LargeStruct) {
// Only passes 8-byte pointer, not 8MB copy
// }

// 3. To represent "no value" (nil)
// func findUser(id int) *User {
// 		...
//     if notFound {
//         return nil  // Can't return nil for value type
//     }
//     return &user
// }

// Pointer Gotchas
// 1. Nil pointer dereference = panic
// var p *int
// fmt.Println(*p)  // PANIC!

// Always check for nil
// if p != nil {
//     fmt.Println(*p)
// }

// 2. Can't take address of literals or constants
// p := &42        // WON'T compile
// p := &"hello"   // WON'T compile

// Workaround: helper function
// func intPtr(i int) *int {
//     return &i
// }
// p := intPtr(42)

// 3. No pointer arithmetic (unlike C)
// Go doesn't allow p++ or p + 4

// Methods on Structs
// Methods are functions with a "receiver" — the type they belong to.

type Rectangle0 struct {
	Width, Height float64
}

func (r Rectangle0) Area() float64 {
	// Value receiver - receives copy
	return r.Width * r.Height
}

func (r *Rectangle0) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func (r Rectangle0) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func composeMethodsInterfaces() {
	rect := Rectangle0{Width: 10, Height: 5}

	fmt.Println(rect.Area())      // 50
	fmt.Println(rect.Perimeter()) // 30

	rect.Scale(2)
	fmt.Println(rect.Area()) // 200 (scaled)

	// Go automatically handles &/* for method calls
	rectPtr := &rect
	fmt.Println(rectPtr.Area()) // Works! Go auto-dereferences

	rect2 := Rectangle0{Width: 3, Height: 4}
	rect2.Scale(2)
}

// Use POINTER receiver when:
// 1. Method modifies the receiver
// func (u *User) UpdateEmail(email string) {
//     u.Email = email
// }

// 2. Receiver is large struct (avoid copy overhead)
// func (ls *LargeStruct) Process() { }

// 3. Consistency — if one method needs pointer, use pointer for all
// type User struct { /* ... */ }
// func (u *User) Save() error { }
// func (u *User) Delete() error { }
// func (u *User) String() string { }  // Pointer for consistency

// Use VALUE receiver when:
// 1. Method doesn't modify receiver
// 2. Receiver is small (int, small struct)
// 3. Receiver is immutable by design

// type Point struct { X, Y int }
// func (p Point) Distance(other Point) float64 {
//     dx := float64(p.X - other.X)
//     dy := float64(p.Y - other.Y)
//     return math.Sqrt(dx*dx + dy*dy)
// }

// Interfaces
// Interfaces define behavior. Types implement interfaces implicitly — no implements keyword needed.

// Interface definition — just method signatures
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle1 implements Shape (implicitly)
type Rectangle1 struct {
	Width, Height float64
}

func (r Rectangle1) Area() float64 {
	return r.Width + r.Height
}

func (r Rectangle1) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle also implements Shape (implicitly)
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Function that accepts any Shape
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func composeInterfaces() {
	rect := Rectangle1{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	// Both work because both implement shape
	printShapeInfo(rect)
	printShapeInfo(circle)

	// Slice of interfaces
	shapes := []Shape{rect, circle}
	for _, shape := range shapes {
		printShapeInfo(shape)
	}
}

// Common Standard Library Interfaces
// Stringer — like JS toString()
// type Stringer interface {
//     String() string
// }

// type Person struct {
//     Name string
//     Age  int
// }

// func (p Person) String() string {
//     return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
// }

// func main() {
//     p := Person{Name: "Alice", Age: 30}
//     fmt.Println(p)  // "Alice (30 years old)" — String() called automatically
// }

// io.Reader and io.Writer — everywhere in Go
// type Reader interface {
//     Read(p []byte) (n int, err error)
// }

// type Writer interface {
//     Write(p []byte) (n int, err error)
// }

// error interface — the foundation of error handling
// type error interface {
//     Error() string
// }

// Empty Interface interface{} / any
// The empty interface has no methods, so ALL types satisfy it — like any in TypeScript.

// func main() {
// Can hold any type
//     var anything interface{}
//     anything = 42
//     anything = "hello"
//     anything = Person{Name: "Alice"}

// Go 1.18+ alias
//     var anything2 any = "same thing"

// Common use: generic containers (before generics)
//     data := []any{1, "two", 3.0, true}

// Type assertion — extract concrete type
//     var val interface{} = "hello"

//     str := val.(string)  // Panics if wrong type!
//     fmt.Println(str)

// Safe type assertion with comma-ok
//     str, ok := val.(string)
//     if ok {
//         fmt.Println("It's a string:", str)
//     }

// Type switch — handle multiple types
//     printType(42)
//     printType("hello")
//     printType(3.14)
// }

// func printType(val any) {
//     switch v := val.(type) {
//     case int:
//         fmt.Println("Integer:", v)
//     case string:
//         fmt.Println("String:", v)
//     case float64:
//         fmt.Println("Float:", v)
//     default:
//         fmt.Printf("Unknown type: %T\n", v)
//     }
// }

// Interface Composition
// Small, focused interfaces

// type Reader interface {
//     Read(p []byte) (n int, err error)
// }

// type Writer interface {
//     Write(p []byte) (n int, err error)
// }

// type Closer interface {
//     Close() error
// }

// Composed interfaces
// type ReadWriter interface {
//     Reader
//     Writer
// }

// type ReadWriteCloser interface {
//     Reader
//     Writer
//     Closer
// }
// Real example from standard library: io.ReadWriteCloser
// Files, network connections, etc. implement this

// Interface Best Practices

// 1. Keep interfaces small (1-3 methods)
// Good
// type Saver interface {
//     Save() error
// }

// Avoid large interfaces unless necessary

// 2. Accept interfaces, return concrete types
// func ProcessData(r io.Reader) (*Result, error) {
// Accept interface — flexible
// Return concrete type — clear
// }

// 3. Define interfaces where they're used (consumer side)
// Not in the package that implements them

// 4. Verify interface compliance at compile time
// var _ Shape = (*Rectangle)(nil)  // Compile error if Rectangle doesn't implement Shape

// Error Handling
// Go has no exceptions. Errors are values returned from functions.
// This is a major paradigm shift from JS.

// The Basics
// func main() {
// 	   Functions return error as last return value
//     file, err := os.Open("nonexistent.txt")
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     defer file.Close()

// 	   strconv example
//     num, err := strconv.Atoi("not a number")
//     if err != nil {
//         fmt.Println("Parse error:", err)
//     }
//     fmt.Println(num)  // 0 (zero value on error)
// }

// ***Creating Errors

// import (
//     "errors"
//     "fmt"
// )

// Simple error
// func divide(a, b float64) (float64, error) {
//     if b == 0 {
//         return 0, errors.New("division by zero")
//     }
//     return a / b, nil
// }

// Formatted error with fmt.Errorf
// func validateAge(age int) error {
//     if age < 0 {
//         return fmt.Errorf("invalid age: %d (must be non-negative)", age)
//     }
//     if age > 150 {
//         return fmt.Errorf("invalid age: %d (too high)", age)
//     }
//     return nil
// }

// func main() {
//     result, err := divide(10, 0)
//     if err != nil {
//         fmt.Println(err)  // "division by zero"
//         return
//     }
//     fmt.Println(result)
// }

// ***Custom Error Types

// Custom error type — implement error interface
// type ValidationError struct {
//     Field   string
//     Message string
// }

// func (e *ValidationError) Error() string {
//     return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
// }

// type NotFoundError struct {
//     Resource string
//     ID       int
// }

// func (e *NotFoundError) Error() string {
//     return fmt.Sprintf("%s with ID %d not found", e.Resource, e.ID)
// }

// func findUser(id int) (*User, error) {
// ... database lookup ...
//     if notFound {
//         return nil, &NotFoundError{Resource: "User", ID: id}
//     }
//     return &user, nil
// }

// func validateUser(u *User) error {
//     if u.Email == "" {
//         return &ValidationError{Field: "email", Message: "required"}
//     }
//     return nil
// }

// ***Error Wrapping (Go 1.13+)
// ***Wrap errors to add context while preserving the original.

// import (
//     "errors"
//     "fmt"
//     "os"
// )

// func readConfig(path string) ([]byte, error) {
//     data, err := os.ReadFile(path)
//     if err != nil {
//         // Wrap with %w verb — preserves original error
//         return nil, fmt.Errorf("readConfig: failed to read %s: %w", path, err)
//     }
//     return data, nil
// }

// func loadApp() error {
//     _, err := readConfig("config.json")
//     if err != nil {
//         return fmt.Errorf("loadApp: %w", err)
//     }
//     return nil
// }

// func main() {
//     err := loadApp()
//     if err != nil {
//         fmt.Println(err)
//         // "loadApp: readConfig: failed to read config.json: open config.json: no such file or directory"

//         // Unwrap to get original error
//         fmt.Println(errors.Unwrap(err))
//     }
// }

// ***errors.Is() and errors.As()
// ***Check for specific errors in a chain of wrapped errors.

// var ErrNotFound = errors.New("not found")
// var ErrPermissionDenied = errors.New("permission denied")

// func getResource(id int) error {
//     // ...
//     return fmt.Errorf("getResource: %w", ErrNotFound)
// }

// func main() {
//     err := getResource(123)

//     // errors.Is — check if error matches (works through wrapping)
//     if errors.Is(err, ErrNotFound) {
//         fmt.Println("Resource not found!")
//     }

//     // Check for os.PathError
//     _, err = os.Open("nonexistent.txt")

//     // errors.As — extract specific error type
//     var pathErr *os.PathError
//     if errors.As(err, &pathErr) {
//         fmt.Println("Path:", pathErr.Path)
//         fmt.Println("Op:", pathErr.Op)
//         fmt.Println("Err:", pathErr.Err)
//     }
// }

// ***Sentinel Errors Pattern
// package database

// import "errors"

// // Sentinel errors — package-level, reusable
// var (
//     ErrNotFound     = errors.New("record not found")
//     ErrDuplicate    = errors.New("duplicate entry")
//     ErrInvalidInput = errors.New("invalid input")
// )

// func FindUser(id int) (*User, error) {
//     // ...
//     if notFound {
//         return nil, ErrNotFound
//     }
//     return &user, nil
// }

// // Usage
// func main() {
//     user, err := database.FindUser(123)
//     if errors.Is(err, database.ErrNotFound) {
//         // Handle not found case
//     }
// }

// ***Error Handling Patterns

// 1. Early return (most common)
// func processFile(path string) error {
//     file, err := os.Open(path)
//     if err != nil {
//         return fmt.Errorf("open file: %w", err)
//     }
//     defer file.Close()

//     data, err := io.ReadAll(file)
//     if err != nil {
//         return fmt.Errorf("read file: %w", err)
//     }

//     if err := process(data); err != nil {
//         return fmt.Errorf("process: %w", err)
//     }

//     return nil
// }

// // 2. Inline error handling with if-init
// if err := doSomething(); err != nil {
//     return err
// }

// // 3. Defer with named return for error wrapping
// func complexOperation() (result *Result, err error) {
//     defer func() {
//         if err != nil {
//             err = fmt.Errorf("complexOperation: %w", err)
//         }
//     }()

//     // ... multiple operations that might fail ...
//     return result, nil
// }

// // 4. Multiple cleanup with defer
// func copyFile(src, dst string) error {
//     source, err := os.Open(src)
//     if err != nil {
//         return err
//     }
//     defer source.Close()

//     dest, err := os.Create(dst)
//     if err != nil {
//         return err
//     }
//     defer dest.Close()

//     _, err = io.Copy(dest, source)
//     return err
// }

// ** Panic and Recover (Use Sparingly)
// ** Panic is for unrecoverable errors — similar to throwing uncaught exceptions.
// ** Avoid in library code.

// func main() {
//     // panic stops normal execution
//     // panic("something went terribly wrong")

//     // recover catches panics (must be in defer)
//     defer func() {
//         if r := recover(); r != nil {
//             fmt.Println("Recovered from:", r)
//         }
//     }()

//     panic("oh no!")
//     fmt.Println("This won't print")
// }

// **Appropriate uses of panic:
// **1. Programming errors (should never happen)
// **2. Initialization failures
// **3. Internal invariant violations

// func mustCompileRegex(pattern string) *regexp.Regexp {
//     re, err := regexp.Compile(pattern)
//     if err != nil {
//         panic(fmt.Sprintf("invalid regex %q: %v", pattern, err))
//     }
//     return re
// }

// // "Must" prefix convention signals it panics on error
// var emailRegex = mustCompileRegex(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func main() {
	// compositeTypesArrayFixSize()
	// compositeTypesSlices()
	// compositeTypesSlicesInternal()
	// compositeTypesSliceOperations()
	// compositeTypesMaps()
	// composeStructs()
	// structEmbedding()
	// compositionStructTags()
	// compositionPointers()
	// pointersMatter()
	// pointersWithStructs()
	// composeMethodsInterfaces()
	composeInterfaces()
}

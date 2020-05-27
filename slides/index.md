![Go logo](images/go-logo.png)

# Introduction to Go

Jean Khechfe, Jakub Jasiński, Paweł Kuffel

---

<!-- .slide: data-background="images/slide-background/goopher.png" -->

---

<!-- .slide: class="text-left" -->

## GO's aims

In the design process of Golang, the following priorities were established:

-   Programming ease compared to dynamic languages
-   Performance as close to C/C++
-   Quick compile times
-   Good developer experience
<!-- you can add more -->

---

<!-- .slide: class="text-left" -->

## What is GO

Go is a new, general-purpose programming language which is:

-   Compiled
-   Statically typed
-   Memory safe
-   Garbage-collected
<!-- you can add more -->

---

<!-- .slide: class="text-left" -->

## History of GO

1. Project starts at Google in 2007 (by Griesemer, Pike, Thompson)
2. Open source release in November 2009
3. More than 250 contributors join the project
4. Version 1.0 released in March 2012

---

<!-- .slide: class="text-left" -->

## Install GO

go to [golang.org](https://golang.org), you will find a link to download and Install GO for your operation system, and once you have done that you will find the GO command-line tools available.

---

<!-- .slide: class="text-left" -->

## Test your GO Installation

Open your terminal and type:

```bash
go version
```

and you should see your go version.

---

<!-- .slide: class="text-left" -->

## Create your first GO application

Create a file named _hello.go_ that looks like:

```go
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```

Then build it with the go tool:

```bash
> go build hello.go
```

The command above will build an executable in the current directory alongside your source code and its name will be **hello.exe** in windows or **hello** in macOS. Execute it to see the greeting:

```bash
> hello
hello, world
```

OR

```bash
> go run hello.go
```

---

<!-- .slide: class="text-left" -->

## Declaring variables

```go
	// Using the var keyword
	var x int // initiated with zero
	var name string = "Bob"
	const isCool = true
	var y float32 = 2.3
	var x, y, z int
	var a, b int = 4, 8 //a = 4, b = 8

	//implicit declaration
	var name = "Bob"
	var y = 2.3
	var a, b = 4, 8

	// Shorthand
	name := "Bob"
	y := 2.3
	a := 4
	b := 8
	sum := a + b
```

---

<!-- .slide: class="text-left" -->

## Packages

```go
package main

import (
	"fmt"
	"math"

	"jean/go_course/packages/strutil"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("hello"))
}
```

---

<!-- .slide: class="text-left" -->

## Functions

```go
package main

import (
	"fmt"
)

func greeting() {
	fmt.Println("Hello TSD")
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func multiply(num1, num2 int) int {
	return num1 * num2
}

func main() {
	greeting()
	fmt.Println(add(3, 4))
	fmt.Println(multiply(3, 4))
}
```

---

<!-- .slide: class="text-left" -->

## Functions (Multiple return values)

```go
package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

func main() {
	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
```

---

<!-- .slide: class="text-left" -->

## Functions (Named return values)

```go
package main

import (
	"fmt"
)

func function1() (i int, s string) {
    i = 17
    s = "abc"
    return // same as return i, s
}

func function2() (i int) {
    return
}

func main() {
	fmt.Println(function1()) // 17 abc
	fmt.Println(function2()) // 0
}
```

---

<!-- .slide: class="text-left" -->

## Conditionals

```go
	x := 10
	y := 10

	if x < y {
		fmt.Printf("%d is less than to %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal to %d\n", y, x)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "red"

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is NOT blue or red")
	}
```

---

<!-- .slide: class="text-left" -->

## Arrays

```go
func main() {

	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fruitArr[2] = "Cherry" // error


	// Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
}
```

---

<!-- .slide: class="text-left" -->

## Maps

```go
func main() {
	vertices := make(map[string]int)

	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["dodecagon"] = 12

	fmt.Println(vertices)
	fmt.Println(len(vertices))
	fmt.Println(vertices["triangle"])

	delete(vertices, "square")

	fmt.Println(vertices)

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	emails["Mike"] = "mike@gmail.com"
}
```

---

<!-- .slide: class="text-left" -->

## Loops

```go
func main() {
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// For each loop
	arr := [3]string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}
}
```

---

<!-- .slide: class="text-left" -->

## Range and blank identifier

```go
func main() {
	// For each loop
	arr := [3]string{"a", "b", "c"}

	// ERROR: index declared and not used
	for index, value := range arr {
		fmt.Println("value:", value)
	}

	for _, value := range arr { // Running without errors
		fmt.Println("value:", value)
	}
}
```

---

<!-- .slide: class="text-left" -->

## Slices

-   Go uses an abstraction built on top of array type called a slice
-   It provides a dynamic view into the elements of an array
-   A slice consists of a pointer to an array, segment length and it's capacity

```go
numbers := [6]int{2, 3, 5, 7, 11, 13}

var numbersSlice []int = numbers[2:4]

fmt.Println(numbersSlice) // [5 7]
```

---

<!-- .slide: class="text-left" -->

## Slices

```go
q := []int{2, 3, 5, 7, 11, 13}

fmt.Println(cap(q)) // 6
q = append(q, 1)
q = append(q, 2, 3, 4)
fmt.Println(cap(q)) // 12

q = q[:5]
fmt.Println(len(q))
//panic: runtime error: slice bounds out of range [:20] with capacity 12

board := [][]string{
	[]string{"x", "x", "x"},
	[]string{"x", "x", "x"},
	[]string{"x", "x", "x"},
}

board[0][2] = "X"
```

---

<!-- .slide: class="text-left" -->

## Structs

```go
type Animal struct {
	Age  int
	Name string
}
func (a *Animal) Move() {
	fmt.Println("Animal moved")
}
func (a *Animal) PrintAge() {
	fmt.Printf("Animal age: %d\n", a.Age)
}
func newAnimal(age int, name string) *Animal {
	return &Animal{Age: age, Name: name}
}

func main() {
	animal1 := Animal{10, "Animal1"}
	animal1.PrintAge()
	animal1.Move()
	animal2 := newAnimal(20, "Animal2")
	animal2.PrintAge() // animal2 is a pointer (*Animal)
```

---

<!-- .slide: class="text-left" -->

## Nesting structs

```go
type Dog struct {
	CollarColor string
	*Animal
}

func newDog(age int, name string, collarColor string) *Dog {
	return &Dog{collarColor, newAnimal(age, name)}
}

func (d *Dog) Move() {
	d.Animal.Move()
	fmt.Println("And that animal was a dog")
}

func main() {
	barky := newDog(5, "Barky", "Red")
	barky.Move() //animal moved and that animal was a dog
	barky.PrintAge() //prints age: 5
}
```

---

<!-- .slide: class="text-left" -->

## Interfaces

```go
type IAnimal interface {
	Move()
}

func DoSomething(a IAnimal) {
	a.Move()
}

func main() {
	/// ...
	animals := []IAnimal{&animal1, animal2, barky}
	// we can't add animal1 unless we provide a pointer to it
	for _, a := range animals {
		a.Move()
		DoSomething(a)
	}
}
```

---

<!-- .slide: class="text-left" -->

## Pointers

```go
func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	//  Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)

	// Another example
	i := 7
	fmt.Println("Memory address:", &i)

	inc1(i)
	fmt.Println("i not modified:", i)

	inc2(&i)
	fmt.Println("i incremented successfully:", i)
}

func inc1(x int) {
	x++
}

func inc2(x *int) {
	*x++
}
```

---

## Memory safety in go

```
func main() {
   items := []int{1, 2, 3}

   printList(items)
}

func printList(list []int) {
   println(list[2])
   println(list[3])
}
```

`go tool compile -S main.go`

```
0x0057 00087 (main.go:11)  MOVQ   "".list+48(SP), CX
0x005c 00092 (main.go:11)  CMPQ   CX, $3
0x0060 00096 (main.go:11)  JLS    151
// ...
0x0096 00150 (main.go:12)  RET
0x0097 00151 (main.go:11)  MOVL   $3, AX
0x009c 00156 (main.go:11)  CALL   runtime.panicIndex(SB)
0x00a1 00161 (main.go:10)  MOVL   $2, AX
0x00a6 00166 (main.go:10)  CALL   runtime.panicIndex(SB)
```

---

<!-- .slide: class="text-left" -->

## Examples

```go
package main

// go get "github.com/bradtraversy/go_crash_course/03_packages/strutil"

import (
    "fmt"
    "github.com/bradtraversy/go_crash_course/03_packages/strutil"
    "strconv"
)

func isPalindromeNumber(number int) bool {
    var numberAsString = strconv.Itoa(number)
    var reversedNumber = strutil.Reverse(numberAsString)
    return reversedNumber == numberAsString
}

func main() {
    fmt.Println(isPalindromeNumber(12344321)) // true
    fmt.Println(isPalindromeNumber(10111)) // false
}
```

---

<!-- .slide: class="text-left" -->

## Examples

```go
package main

import (
    "fmt"
    "strings"
    "strconv"
)

func digitCount(number int) map[int]int {
	numberAsString := strconv.Itoa(number)
	counts := make(map[int]int, 10)
	for i := 0; i <= 9; i++ {
		counts[i] = strings.Count(numberAsString, strconv.Itoa(i))
	}
	return counts
}

func main() {
    count := digitCount(123444321)
    printMap(count) // 0:0 1:2 2:2 3:2 4:3 5:0 6:0 7:0 8:0 9:0
}

func printMap(resultMap map[int]int) {
	for i := 0; i <= 9; i++ {
    	fmt.Print(i,":", resultMap[i], " ")
    }
}
```

---

<!-- .slide: class="text-left" -->

## Examples

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println(primesSum([]int{1, 2, 3})) // 10
    fmt.Println(primesSum([]int{1, 10, 20})) // 102
}

func primesSum(positionsToSum []int) int {

    var sum int = 0
    var primeNumbers = generateFirst900PrimeNumbers()

    for _, position := range positionsToSum {
    	sum += primeNumbers[position - 1]
    }

    return sum
}

func generateFirst900PrimeNumbers() (primes []int) {
    N := 7000
    b := make([]bool, N)
    for i := 2; i < N; i++ {
        if b[i] == true { continue }
        primes = append(primes, i)
        for k := i * i; k < N; k += i {
            b[k] = true
        }
    }
    return
}
```

---

<!-- .slide: class="text-left" -->

## Examples

```go
package main

import (
	"fmt"
)

func main() {
	var fruit = NewFruit("blue", "medium", 30)
	fruit.showDetails()
}

type Fruit struct {
	color, size string
	price int
}

func NewFruit(newColor string, newSize string, newPrice int) *Fruit {
	return &Fruit{color:newColor, size:newSize, price:newPrice}
}

func (fruit Fruit) showDetails() {
	fmt.Println("I have", fruit.color, "color, my size is", fruit.size, "and I cost", fruit.price)
}
```

---

<!-- .slide: class="text-left" -->

## Examples

```go
package main

import (
	"fmt"
)

func main() {
	var apple = NewApple("red", "big", 50, "Poland", "medium")
	apple.showDetails()
	apple.showAdditionalInfo()
}

type Apple struct {
	*Fruit
	originCountry, sweetness string
}

func NewApple(newColor string, newSize string, newPrice int, newCountry string, newSweetness string) *Apple {
	return &Apple{Fruit:&Fruit{color:newColor, size:newSize, price:newPrice}, originCountry:newCountry, sweetness:newSweetness}
}

func (apple Apple) showAdditionalInfo() {
	fmt.Println("I am from", apple.originCountry, "and my sweetness is", apple.sweetness)
}

type Fruit struct {
	color, size string
	price int
}

func NewFruit(newColor string, newSize string, newPrice int) *Fruit {
	return &Fruit{color:newColor, size:newSize, price:newPrice}
}

func (fruit Fruit) showDetails() {
	fmt.Println("I have", fruit.color, "color, my size is", fruit.size, "and I cost", fruit.price)
}
```

---

<!-- .slide: class="text-left" -->

## Examples

```go
package main

import (
	"fmt"
)

func main() {
	var apple = NewApple("red", "big", 50, "Poland", "medium")
	apple.showDetails()
}

func (apple Apple) showDetails() {
	fmt.Println("I have", apple.color, "color, my size is", apple.size, "and I cost", apple.price, ". Also I am from", apple.originCountry, "and my sweetness is", apple.sweetness, ".")
}

type Apple struct {
	*Fruit
	originCountry, sweetness string
}
```

---

<!-- .slide: class="text-left" -->

## Examples

```go
package main

import (
	"fmt"
)

func main() {
	var yellowApple = NewApple("yellow", "medium", 70, "Italy", "high")
	var fruit = NewFruit("blue", "medium", 30)
	var redApple = NewApple("red", "big", 50, "Poland", "medium")

	sliceOfFruits := [3]*Fruit{yellowApple.Fruit, fruit, redApple.Fruit}

	for _, fruit := range sliceOfFruits {
		fruit.showDetails()
	}
}
```

---

## Resources

-   [Go by example](https://gobyexample.com/)
-   [Tour of go](https://tour.golang.org/welcome/1)

<style>
::selection {
    background: #02283e;
}
.reveal {
    font-size: 30px;
}
.reveal p {
    font-size: 0.8em;
}
.dark_back {
  background-color: rgba(0, 0, 0, 0.9);
  color: #fff;
  padding: 20px;
}
.text-left {
	text-align: left;
}
</style>

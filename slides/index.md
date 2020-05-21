![Go logo](images/go-logo.png)

# Introduction to Go

Jean Khechfe, Jakub Jasiński, Paweł Kuffel

---

<!-- .slide: data-background="images/slide-background/slide-background.jpg" -->

## Second slide <!-- .element: class="dark_back" -->

> Best quote ever. <!-- .element: class="dark_back" -->

---

<!-- .slide: class="text-left" -->
## What is GO

Go is a new, general-purpose programming language which is:
* Compiled
* statically typed
* syntactically similar to C
* memory safe
* garbage collection
<!-- you can add more -->

---

<!-- .slide: class="text-left" -->
## History of GO

1. Project starts at Google in 2007 (by Griesemer, Pike, Thompson)
2. Open source release in November 2009
3. More than 250 contributors join the project
4. Version 1.0 release in March 2012

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

1. Create a file named *hello.go* that looks like:

```go
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```

2. Then build it with the go tool:

```bash
> go build hello.go
```

3. The command above will build an executable in the current directory alongside your source code and its name will be **hello.exe** in windows or **hello** in macOS. Execute it to see the greeting:

```bash
> hello
hello, world
```

---

<!-- .slide: class="text-left" -->
## Declaring variables

```go
	var x int
	const isCool = true
	var y float32 = 2.3

	// Shorthand
	name := "Bob"
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

	"github.com/bradtraversy/go_crash_course/03_packages/strutil"
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
	"errors"
	"fmt"
	"math"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(getSum(3, 4))

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


	// Decalre and assign
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
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
## Struct

```go
import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func main() {

	person := Person{"Bob", "Johnson", "New York", "m", 30}
	fmt.Println(person.greet())
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

<!-- .slide: class="text-left" -->
## Web

```go
import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
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

## Yet another slide

1. A bulletpoint
2. Another one
3. Yet another one

---

### Another slide

Some Go code

```go
type Dog struct {
	Animal
}
type Animal struct {
	Age int
}

func (a *Animal) Move() {
	fmt.Println("Animal moved")
}
func (a *Animal) SayAge() {
	fmt.Printf("Animal age: %d\n", a.Age)
}
func main() {
	d := Dog{}
	d.Age = 3
	d.Move()
	d.SayAge()
}
```

---

## Summary

-   Lorem ipsum <!-- .element: class="fragment" data-fragment-index="1" -->
-   Dolor sit amet <!-- .element: class="fragment" data-fragment-index="2" -->
-   Yet another point <!-- .element: class="fragment" data-fragment-index="3" -->

---

## Final slide

Some text

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

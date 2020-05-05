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

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(getSum(3, 4))
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

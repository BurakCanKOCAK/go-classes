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

## Create your first GO project

1. Step 1
2. Step 2
3. Step 3

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

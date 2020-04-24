![Go logo](images/go-logo.png)

# Introduction to Go

Jean Khechfe, Jakub Jasiński, Paweł Kuffel

---

<!-- .slide: data-background="images/slide-background/slide-background.jpg" -->

## Second slide <!-- .element: class="dark_back" -->

> Best quote ever. <!-- .element: class="dark_back" -->

----

## Yet another slide

1. A bulletpoint
2. Another one
3. Yet another one

----

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
</style>

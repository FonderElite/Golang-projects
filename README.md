# <img src="https://img.icons8.com/color/48/000000/golang.png"/> Golang-projects & Guide
> Golang projects and fundamentals.

## What is Go?
Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency

## Go lang syntax
```go
Important Terms
//Main Types
//string
//bool
//int
//int uint int8 int16 int32 int64
//uint uint8 uint16 uint32 uint64
//byte - alias for uint8
//rune - alias for int32
//float32 float64
//complex64 complex128
```

```go
//Variables
var a int = 1337
b:= 1386
```
```go
//Arrays
var a [2]string
a[0] = "Hello"
a[1] = "World"
fmt.Println(a[0], a[1])
//Output:Hello World
primes := [6]int{2, 3, 5, 7, 11, 13}
fmt.Println(primes)
//Output:[2, 3, 5, 7, 11, 13]
```
```go
//Maps and slices
language := make(map[string]int)
language["JavaScript"] = 0
language["Python"] = 1
fmt.Println(language["JavaScript"],language["Python"])
//Output: 0 1

container := []int{1,2,3,4,5}
var s []int = container[1:3]
fmt.Println(s)
//Output: [2,3]
```
```go
//Range
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
for i, v := range pow {
fmt.Printf("%d = %d\n", i, v)
/*Output:
0 = 1
1 = 2
2 = 4
3 = 8
4 = 16
5 = 32
6 = 64
7 = 128
*/

//Structs
type Employee struct {
    firstName, lastName string
    age,salary int
}
employee1 = &Employee{"Sam", "Anderson", 55, 6000}
fmt.Println("First-Name(*employee1).firstName)
//Output: Sam
```

```go
//Functions
func plus(a int, b int) int {
   return a + b
}
func main(){
 res := plus(1, 2)
    fmt.Println("1+2 =", res)
}
//Output: 1+2 = 3

func plusPlus(a, b, c int) int {
    return a + b + c
}
func main(){
 res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}
//Output: 1+2+3= 6
```

```go
//Pointers
i, j := 42, 2701

p := &i         // point to i
fmt.Println(*p) // read i through the pointer
*p = 21         // set i through the pointer
fmt.Println(i)  // see the new value of i

p = &j         // point to j
*p = *p / 37   // divide j through the pointer
fmt.Println(j) // see the new value of j
//Output:
42
21
73
```

```golang
//Channels
func myfunc(ch chan int){
fmt.Println(394 + <- ch)
//Output: 788
}
func main() {
v := 394
f := make(chan int)
go myfunc(f)
f  <- v
fmt.Println(f)
//Output:0xc00006a060
}
```
## Why use Go?
<img src="https://softensy.com/wp-content/uploads/2020/06/why-use-google-go-language.png" width=500px>
Due to the simple fact, Go is an open-source initiative. Go Offers easy support: Aside from having code that easy to learn, Go offers developers to use several tools to work with. Go contains a simple API allowing developers the environment for easy testing, profiling, and much more.

## Golang Speed
<img src="https://objectbox.io/wordpress/wp-content/uploads/2018/11/ObjectBox-Go-Performance.png" width=500px>

## Resources:

[https://gobyexample.com/](https://gobyexample.com/)
[https://golang.org/](https://golang.org/)
[https://golangcode.com](https://golangcode.com)


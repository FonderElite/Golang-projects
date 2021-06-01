# <img src="https://img.icons8.com/color/48/000000/golang.png"/> Golang-projects & Guide
> Golang projects and fundamentals.

## What is Go?
Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency

## Go lang syntax
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
```

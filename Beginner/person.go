package main

import "fmt"

type Person struct {
	name string
	age  int
}

func myfuncall(name string, age int) *Person {
	info := Person{name: name, age: age}
	fmt.Printf("Person's Name: %v\n", info.name)
	fmt.Printf("Person's Age: %v\n", info.age)
	return &info
}

func myfunc(name chan string) {
	new := make(map[int]string)
	new[1] = "Jonathan"
	fmt.Println("New name: ", new[1]+<-name)
}
func main() {
	name := "Blue"
	newch := make(chan string)
	go myfunc(newch)
	newch <- name
    myfuncall("John", 29)
}

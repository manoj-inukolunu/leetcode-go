package main

import "fmt"

type Worker interface {
	Work()
}
type Person struct {
	name string
}

func (p Person) Work() {
	fmt.Println(p.name, "is working")
}
func describe(w Worker) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

func assert(i interface{}) {
	s := i.(string)
	fmt.Println(s)
}

func main() {
	p := Person{"Manoj"}
	var w Worker = p
	describe(w)
	w.Work()
	var s interface{} = "asdf"
	assert(s)
}

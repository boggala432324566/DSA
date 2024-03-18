package main

import (
	"fmt"
	"sync"
)

type single struct {
	Val int
}

var inst *single
var once sync.Once

func singleton() *single {
	once.Do(func() {
		inst = &single{Val: 10}
	})
	return inst
}

func main() {
	sing1 := singleton()
	sing2 := singleton()
	fmt.Println(sing1.Val)
	fmt.Println(sing2.Val)
}

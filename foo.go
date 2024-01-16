package lib1

import (
	"fmt"
	"github.com/jkblume/lib2"
)

func Foo() {
	lib2.Bar()
}

func Foo2() {
	fmt.Println("foo2")
}

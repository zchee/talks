package main

/*
#cgo pkg-config: python3
#include <stdlib.h>
#include <Python.h>
*/
import "C"
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getpid())
	C
}

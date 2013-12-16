package main

/*
#cgo LDFLAGS: -lhello
#include "hello.h"
*/
import "C"
import (
    "fmt"
    //"unsafe"
)

func main() {

    cStr := C.hello()
    str := C.GoString(cStr)
    //C.free(unsafe.Pointer(cStr))
    fmt.Println(str)
}

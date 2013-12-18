package main

/*
#cgo LDFLAGS: -lsay
#include <stdlib.h>
#include "say.h"
*/
import "C"
import (
    "fmt"
    "unsafe"
)

func main() {

    cName := C.CString("fukata")
    cStr := C.say_cstring(cName)
    str := C.GoString(cStr)
    C.free(unsafe.Pointer(cName))
    C.free(unsafe.Pointer(cStr))
    fmt.Println(str)
}

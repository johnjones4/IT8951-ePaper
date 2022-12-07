package it8951

// #cgo LDFLAGS: -lbcm2835 -lm -lrt -lpthread
// #include "IT8951.h"
// #include <stdlib.h>
import "C"
import "unsafe"

func Init() {
	C.ext_IT8951_init()
}

func Width() int {
	return int(C.ext_IT8951_width())
}

func Height() int {
	return int(C.ext_IT8951_height())
}

func Draw(b []byte) {
	a := C.CString(string(b))
	C.ext_IT8951_draw(a)
	C.free(unsafe.Pointer(a))
}

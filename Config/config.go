package config

// #cgo LDFLAGS: -lbcm2835 -lm -lrt -lpthread
// #include "DEV_Config.h"
import "C"

func Init() {
	C.DEV_Module_Init()
}

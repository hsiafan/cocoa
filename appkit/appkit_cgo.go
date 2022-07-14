package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// void importAppKitProtocols();
import "C"

func init() {
	C.importAppKitProtocols()
}

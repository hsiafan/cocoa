package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #import <stdlib.h>
// void importFoundationProtocols();
import "C"

func init() {
	C.importFoundationProtocols()
}

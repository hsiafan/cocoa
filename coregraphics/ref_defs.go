package coregraphics

import "unsafe"

type ImageRef unsafe.Pointer
type ContextRef unsafe.Pointer
type ColorRef unsafe.Pointer
type ColorSpaceRef unsafe.Pointer
type EventRef unsafe.Pointer
type PathRef unsafe.Pointer

type Glyph = uint16
type FontIndex = uint16

type Error uint32

const (
	ErrorSuccess           Error = 0
	ErrorFailure           Error = 1000
	ErrorIllegalArgument   Error = 1001
	ErrorInvalidConnection Error = 1002
	ErrorInvalidContext    Error = 1003
	ErrorCannotComplete    Error = 1004
	ErrorNotImplemented    Error = 1006
	ErrorRangeCheck        Error = 1007
	ErrorTypeCheck         Error = 1008
	ErrorInvalidOperation  Error = 1010
	ErrorNoneAvailable     Error = 1011
)

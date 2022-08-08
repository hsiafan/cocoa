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

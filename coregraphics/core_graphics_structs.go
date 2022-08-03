package coregraphics

// struct def should be sync with struct in <CoreGraphics/CGGeometry.h> <CoreGraphics/CGAffineTransform.h>

type Float = float64

type AffineTransform struct {
	A  Float
	B  Float
	C  Float
	D  Float
	TX Float
	TY Float
}

type Point struct {
	X Float
	Y Float
}

type Rect struct {
	Origin Point
	Size   Size
}

type Size struct {
	Width  Float
	Height Float
}
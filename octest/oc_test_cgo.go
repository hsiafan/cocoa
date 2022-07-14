package octest

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreGraphics -framework Foundation -framework AppKit
// #import <CoreGraphics/CGGeometry.h>
// #import <CoreGraphics/CGAffineTransform.h>
// #import <Foundation/NSRange.h>
// #import <Foundation/NSAffineTransform.h>
// #import <Foundation/NSDecimal.h>
// #import <AppKit/NSCollectionViewCompositionalLayout.h>
import "C"

type CInt C.int
type CLong C.long
type CBOOL C.BOOL
type CBool C.bool

type Integer C.NSInteger

type AffineTransform C.CGAffineTransform
type Point C.CGPoint
type Size C.CGSize
type Rect C.CGRect
type EdgeInsets C.NSEdgeInsets
type Range C.NSRange
type AffineTransformStruct C.NSAffineTransformStruct
type Decimal C.NSDecimal
type DirectionalEdgeInsets C.NSDirectionalEdgeInsets

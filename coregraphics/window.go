package coregraphics

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import <CoreGraphics/CGGeometry.h>
// void ArrayGetValues(void* array, long len, void* values);
// long ArrayCount(void* array);
// void* ArrayCreate(void *values, long numValues);
// void* WindowListCreate(uint32_t option, uint32_t relativeToWindow);
// void* WindowListCreateImage(CGRect screenBounds, uint32_t listOption, uint32_t windowID, uint32_t imageOption);
// void* WindowListCreateImageFromArray(CGRect screenBounds, void* windowArray, uint32_t imageOption);
import "C"
import (
	"unsafe"

	"github.com/hsiafan/cocoa/internal"
)

type WindowID uint32

const NullWindowID WindowID = 0

type WindowListOption uint32

const (
	WindowListOptionAll                 WindowListOption = 0
	WindowListOptionOnScreenOnly        WindowListOption = (1 << 0)
	WindowListOptionOnScreenAboveWindow WindowListOption = (1 << 1)
	WindowListOptionOnScreenBelowWindow WindowListOption = (1 << 2)
	WindowListOptionIncludingWindow     WindowListOption = (1 << 3)
	WindowListExcludeDesktopElements    WindowListOption = (1 << 4)
)

type WindowImageOption uint32

const (
	WindowImageDefault             WindowImageOption = 0
	WindowImageBoundsIgnoreFraming WindowImageOption = (1 << 0)
	WindowImageShouldBeOpaque      WindowImageOption = (1 << 1)
	WindowImageOnlyShadows         WindowImageOption = (1 << 2)
	WindowImageBestResolution      WindowImageOption = (1 << 3)
	WindowImageNominalResolution   WindowImageOption = (1 << 4)
)

type WindowSharingType uint32

const (
	WindowSharingNone      WindowSharingType = 0
	WindowSharingReadOnly  WindowSharingType = 1
	WindowSharingReadWrite WindowSharingType = 2
)

type WindowBackingType uint32

const (
	BackingStoreRetained    WindowBackingType = 0
	BackingStoreNonretained WindowBackingType = 1
	BackingStoreBuffered    WindowBackingType = 2
)

func WindowListCreate(option WindowListOption, relativeToWindow WindowID) []WindowID {
	p := C.WindowListCreate(C.uint32_t(option), C.uint32_t(relativeToWindow))
	if p == nil {
		return nil
	}
	len := int(C.ArrayCount(p))
	ids := make([]uintptr, len)
	if len > 0 {
		C.ArrayGetValues(p, C.long(len), unsafe.Pointer(&ids[0]))
	}
	windowIds := make([]WindowID, len)
	for i, id := range ids {
		windowIds[i] = WindowID(id)
	}
	return windowIds
}

func WindowListCreateImage(screenBounds Rect, listOption WindowListOption, windowID WindowID, imageOption WindowImageOption) ImageRef {
	p := C.WindowListCreateImage(internal.ForceCast[Rect, C.CGRect](screenBounds),
		C.uint32_t(listOption), C.uint32_t(windowID), C.uint32_t(imageOption))
	return ImageRef(p)
}

func WindowListCreateImageFromArray(screenBounds Rect, windowIDs []WindowID, imageOption WindowImageOption) ImageRef {
	ids := make([]uintptr, len(windowIDs))
	for i, id := range windowIDs {
		ids[i] = uintptr(id)
	}
	windowIDArrayRef := C.ArrayCreate(unsafe.Pointer(&ids[0]), C.long(len(ids)))
	p := C.WindowListCreateImageFromArray(internal.ForceCast[Rect, C.CGRect](screenBounds), windowIDArrayRef, C.uint32_t(imageOption))
	return ImageRef(p)
}

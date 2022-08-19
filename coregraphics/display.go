package coregraphics

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import <CoreGraphics/CGGeometry.h>
// uint32_t MainDisplayID();
// uint32_t GetOnlineDisplayList(uint32_t maxDisplays, void* onlineDisplays, uint32_t *displayCount);
// uint32_t GetActiveDisplayList(uint32_t maxDisplays, void* activeDisplays, uint32_t *displayCount);
// uint32_t GetDisplaysWithPoint(CGPoint point, uint32_t maxDisplays, void* activeDisplays, uint32_t *displayCount);
// uint32_t GetDisplaysWithRect(CGRect rect, uint32_t maxDisplays, void* activeDisplays, uint32_t *displayCount);
// void* DisplayCreateImage(uint32_t displayID);
// void* DisplayCreateImageForRect(uint32_t displayID, CGRect rect);
// CGSize DisplayScreenSize(uint32_t displayID);
// bool DisplayIsActive(uint32_t displayID);
// bool DisplayIsAsleep(uint32_t displayID);
// bool DisplayIsBuiltin(uint32_t displayID);
// bool DisplayIsMain(uint32_t displayID);
// bool DisplayIsOnline(uint32_t displayID);
// size_t DisplayPixelsHigh(uint32_t displayID);
// size_t DisplayPixelsWide(uint32_t displayID);
// CGRect DisplayBounds(uint32_t displayID);
import "C"
import (
	"unsafe"

	"github.com/hsiafan/cocoa/internal"
)

// A unique identifier for an attached display.
type DirectDisplayID uint32

// Returns the display ID of the main display.
func MainDisplayID() DirectDisplayID {
	r := C.MainDisplayID()
	return DirectDisplayID(r)
}

// Provides a list of displays that are online (active, mirrored, or sleeping).
func GetOnlineDisplayList(maxDisplays uint32) ([]DirectDisplayID, Error) {
	displayIds := make([]DirectDisplayID, maxDisplays)
	var count C.uint32_t
	errCode := Error(C.GetOnlineDisplayList(C.uint32_t(maxDisplays), unsafe.Pointer(&displayIds[0]), &count))
	if errCode != ErrorSuccess {
		return nil, errCode
	}
	return displayIds[:int(count)], ErrorSuccess

}

// Provides a list of displays that are online (active, mirrored, or sleeping).
func GetActiveDisplayList(maxDisplays uint32) ([]DirectDisplayID, Error) {
	displayIds := make([]DirectDisplayID, maxDisplays)
	var count C.uint32_t
	errCode := Error(C.GetActiveDisplayList(C.uint32_t(maxDisplays), unsafe.Pointer(&displayIds[0]), &count))
	if errCode != ErrorSuccess {
		return nil, errCode
	}
	return displayIds[:int(count)], ErrorSuccess

}

// Provides a list of online displays with bounds that include the specified point.
func GetDisplaysWithPoint(point Point, maxDisplays uint32) ([]DirectDisplayID, Error) {
	displayIds := make([]DirectDisplayID, maxDisplays)
	var count C.uint32_t
	errCode := Error(C.GetDisplaysWithPoint(internal.ForceCast[Point, C.CGPoint](point), C.uint32_t(maxDisplays), unsafe.Pointer(&displayIds[0]), &count))
	if errCode != ErrorSuccess {
		return nil, errCode
	}
	return displayIds[:int(count)], ErrorSuccess

}

func GetDisplaysWithRect(rect Rect, maxDisplays uint32) ([]DirectDisplayID, Error) {
	displayIds := make([]DirectDisplayID, maxDisplays)
	var count C.uint32_t
	errCode := Error(C.GetDisplaysWithRect(internal.ForceCast[Rect, C.CGRect](rect), C.uint32_t(maxDisplays), unsafe.Pointer(&displayIds[0]), &count))
	if errCode != ErrorSuccess {
		return nil, errCode
	}
	return displayIds[:int(count)], ErrorSuccess

}

// Returns an image containing the contents of the specified display.
func DisplayCreateImage(displayID DirectDisplayID) ImageRef {
	p := C.DisplayCreateImage(C.uint32_t(displayID))
	return ImageRef(p)
}

func DisplayCreateImageForRect(displayID DirectDisplayID, rect Rect) ImageRef {
	p := C.DisplayCreateImageForRect(C.uint32_t(displayID), internal.ForceCast[Rect, C.CGRect](rect))
	return ImageRef(p)
}

func DisplayScreenSize(displayID DirectDisplayID) Size {
	return internal.ForceCast[C.CGSize, Size](C.DisplayScreenSize(C.uint32_t(displayID)))
}

func DisplayIsActive(displayID DirectDisplayID) bool {
	return bool(C.DisplayIsActive(C.uint32_t(displayID)))
}

func DisplayIsAsleep(displayID DirectDisplayID) bool {
	return bool(C.DisplayIsAsleep(C.uint32_t(displayID)))
}

func DisplayIsBuiltin(displayID DirectDisplayID) bool {
	return bool(C.DisplayIsBuiltin(C.uint32_t(displayID)))
}

func DisplayIsMain(displayID DirectDisplayID) bool {
	return bool(C.DisplayIsMain(C.uint32_t(displayID)))
}

func DisplayIsOnline(displayID DirectDisplayID) bool {
	return bool(C.DisplayIsOnline(C.uint32_t(displayID)))
}

func DisplayPixelsHigh(displayID DirectDisplayID) uint {
	return uint(C.DisplayPixelsHigh(C.uint32_t(displayID)))
}

func DisplayPixelsWide(displayID DirectDisplayID) uint {
	return uint(C.DisplayPixelsWide(C.uint32_t(displayID)))
}

func DisplayBounds(displayID DirectDisplayID) Rect {
	return internal.ForceCast[C.CGRect, Rect](C.DisplayBounds(C.uint32_t(displayID)))
}

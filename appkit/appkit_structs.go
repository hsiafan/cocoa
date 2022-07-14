package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
)

// stuct define should be synced with <AppKit/NSCollectionViewCompositionalLayout.h>

type DirectionalEdgeInsets struct {
	Top      coregraphics.Float
	Leading  coregraphics.Float
	Bottom   coregraphics.Float
	Trailing coregraphics.Float
}

type ModalSession unsafe.Pointer

// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type CollectionViewPrefetching interface {
	ImplementsCollectionView_CancelPrefetchingForItemsAtIndexPaths() bool
	// optional
	CollectionView_CancelPrefetchingForItemsAtIndexPaths(collectionView CollectionView, indexPaths []foundation.IndexPath)
	// required
	CollectionView_PrefetchItemsAtIndexPaths(collectionView CollectionView, indexPaths []foundation.IndexPath)
}

type CollectionViewPrefetchingWrapper struct {
	objc.Object
}

func (c_ *CollectionViewPrefetchingWrapper) ImplementsCollectionView_CancelPrefetchingForItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:cancelPrefetchingForItemsAtIndexPaths:"))
}

func (c_ CollectionViewPrefetchingWrapper) CollectionView_CancelPrefetchingForItemsAtIndexPaths(collectionView ICollectionView, indexPaths []foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:cancelPrefetchingForItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), indexPaths)
}

func (c_ CollectionViewPrefetchingWrapper) CollectionView_PrefetchItemsAtIndexPaths(collectionView ICollectionView, indexPaths []foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:prefetchItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), indexPaths)
}

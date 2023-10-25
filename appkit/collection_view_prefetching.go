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

func WrapCollectionViewPrefetching(v CollectionViewPrefetching) objc.Object {
	return objc.WrapAsProtocol("NSCollectionViewPrefetching", v)
}

type CollectionViewPrefetchingBase struct {
}

func (p *CollectionViewPrefetchingBase) ImplementsCollectionView_CancelPrefetchingForItemsAtIndexPaths() bool {
	return false
}

func (p *CollectionViewPrefetchingBase) CollectionView_CancelPrefetchingForItemsAtIndexPaths(collectionView CollectionView, indexPaths []foundation.IndexPath) {
	panic("unimpemented protocol method")
}

// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type CollectionViewDataSource interface {
	ImplementsNumberOfSectionsInCollectionView() bool
	// optional
	NumberOfSectionsInCollectionView(collectionView CollectionView) int
	// required
	CollectionView_NumberOfItemsInSection(collectionView CollectionView, section int) int
	// required
	CollectionView_ItemForRepresentedObjectAtIndexPath(collectionView CollectionView, indexPath foundation.IndexPath) ICollectionViewItem
	ImplementsCollectionView_ViewForSupplementaryElementOfKind_AtIndexPath() bool
	// optional
	CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath(collectionView CollectionView, kind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) IView
}

func WrapCollectionViewDataSource(v CollectionViewDataSource) objc.Object {
	return objc.WrapAsProtocol("NSCollectionViewDataSource", v)
}

type CollectionViewDataSourceBase struct {
}

func (p *CollectionViewDataSourceBase) ImplementsNumberOfSectionsInCollectionView() bool {
	return false
}

func (p *CollectionViewDataSourceBase) NumberOfSectionsInCollectionView(collectionView CollectionView) int {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDataSourceBase) ImplementsCollectionView_ViewForSupplementaryElementOfKind_AtIndexPath() bool {
	return false
}

func (p *CollectionViewDataSourceBase) CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath(collectionView CollectionView, kind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) IView {
	panic("unimpemented protocol method")
}

type CollectionViewDataSourceWrapper struct {
	objc.Object
}

func (c_ CollectionViewDataSourceWrapper) NumberOfSectionsInCollectionView(collectionView ICollectionView) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("numberOfSectionsInCollectionView:"), objc.ExtractPtr(collectionView))
	return rv
}

func (c_ CollectionViewDataSourceWrapper) CollectionView_NumberOfItemsInSection(collectionView ICollectionView, section int) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("collectionView:numberOfItemsInSection:"), objc.ExtractPtr(collectionView), section)
	return rv
}

func (c_ CollectionViewDataSourceWrapper) CollectionView_ItemForRepresentedObjectAtIndexPath(collectionView ICollectionView, indexPath foundation.IIndexPath) CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](c_, objc.GetSelector("collectionView:itemForRepresentedObjectAtIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionViewDataSourceWrapper) CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath(collectionView ICollectionView, kind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) View {
	rv := objc.CallMethod[View](c_, objc.GetSelector("collectionView:viewForSupplementaryElementOfKind:atIndexPath:"), objc.ExtractPtr(collectionView), kind, objc.ExtractPtr(indexPath))
	return rv
}

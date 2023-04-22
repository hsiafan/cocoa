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

type CollectionViewDataSourceWrapper struct {
	objc.Object
}

func (c_ *CollectionViewDataSourceWrapper) ImplementsNumberOfSectionsInCollectionView() bool {
	return c_.RespondsToSelector(objc.GetSelector("numberOfSectionsInCollectionView:"))
}

func (c_ CollectionViewDataSourceWrapper) NumberOfSectionsInCollectionView(collectionView ICollectionView) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("numberOfSectionsInCollectionView:"), collectionView)
	return rv
}

func (c_ CollectionViewDataSourceWrapper) CollectionView_NumberOfItemsInSection(collectionView ICollectionView, section int) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("collectionView:numberOfItemsInSection:"), collectionView, section)
	return rv
}

func (c_ CollectionViewDataSourceWrapper) CollectionView_ItemForRepresentedObjectAtIndexPath(collectionView ICollectionView, indexPath foundation.IIndexPath) CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](c_, objc.GetSelector("collectionView:itemForRepresentedObjectAtIndexPath:"), collectionView, indexPath)
	return rv
}

func (c_ *CollectionViewDataSourceWrapper) ImplementsCollectionView_ViewForSupplementaryElementOfKind_AtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:viewForSupplementaryElementOfKind:atIndexPath:"))
}

func (c_ CollectionViewDataSourceWrapper) CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath(collectionView ICollectionView, kind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) View {
	rv := objc.CallMethod[View](c_, objc.GetSelector("collectionView:viewForSupplementaryElementOfKind:atIndexPath:"), collectionView, kind, indexPath)
	return rv
}

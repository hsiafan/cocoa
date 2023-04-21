// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionViewDiffableDataSourceClass = _CollectionViewDiffableDataSourceClass{objc.GetClass("NSCollectionViewDiffableDataSource")}

type _CollectionViewDiffableDataSourceClass struct {
	objc.Class
}

type ICollectionViewDiffableDataSource interface {
	objc.IObject
	ItemIdentifierForIndexPath(indexPath foundation.IIndexPath) objc.Object
	IndexPathForItemIdentifier(identifier objc.IObject) foundation.IndexPath
	Snapshot() DiffableDataSourceSnapshot
	ApplySnapshot_AnimatingDifferences(snapshot IDiffableDataSourceSnapshot, animatingDifferences bool)
	SupplementaryViewProvider() func(param1 ICollectionView, param2 string, param3 foundation.IIndexPath) View
	SetSupplementaryViewProvider(value func(param1 CollectionView, param2 string, param3 foundation.IndexPath) IView)
}

type CollectionViewDiffableDataSource struct {
	objc.Object
}

func MakeCollectionViewDiffableDataSource(ptr unsafe.Pointer) CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSource{
		Object: objc.MakeObject(ptr),
	}
}

func (c_ CollectionViewDiffableDataSource) InitWithCollectionView_ItemProvider(collectionView ICollectionView, itemProvider func(param1 CollectionView, param2 foundation.IndexPath, param3 objc.Object) ICollectionViewItem) CollectionViewDiffableDataSource {
	rv := objc.CallMethod[CollectionViewDiffableDataSource](c_, "initWithCollectionView:itemProvider:", collectionView, itemProvider)
	return rv
}

func (cc _CollectionViewDiffableDataSourceClass) Alloc() CollectionViewDiffableDataSource {
	rv := objc.CallMethod[CollectionViewDiffableDataSource](cc, "alloc")
	return rv
}

func (cc _CollectionViewDiffableDataSourceClass) New() CollectionViewDiffableDataSource {
	rv := objc.CallMethod[CollectionViewDiffableDataSource](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewDiffableDataSource() CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSourceClass.New()
}

func (c_ CollectionViewDiffableDataSource) Init() CollectionViewDiffableDataSource {
	rv := objc.CallMethod[CollectionViewDiffableDataSource](c_, "init")
	return rv
}

func (c_ CollectionViewDiffableDataSource) ItemIdentifierForIndexPath(indexPath foundation.IIndexPath) objc.Object {
	rv := objc.CallMethod[objc.Object](c_, "itemIdentifierForIndexPath:", indexPath)
	return rv
}

func (c_ CollectionViewDiffableDataSource) IndexPathForItemIdentifier(identifier objc.IObject) foundation.IndexPath {
	rv := objc.CallMethod[foundation.IndexPath](c_, "indexPathForItemIdentifier:", identifier)
	return rv
}

func (c_ CollectionViewDiffableDataSource) Snapshot() DiffableDataSourceSnapshot {
	rv := objc.CallMethod[DiffableDataSourceSnapshot](c_, "snapshot")
	return rv
}

func (c_ CollectionViewDiffableDataSource) ApplySnapshot_AnimatingDifferences(snapshot IDiffableDataSourceSnapshot, animatingDifferences bool) {
	objc.CallMethod[objc.Void](c_, "applySnapshot:animatingDifferences:", snapshot, animatingDifferences)
}

func (c_ CollectionViewDiffableDataSource) SupplementaryViewProvider() func(param1 ICollectionView, param2 string, param3 foundation.IIndexPath) View {
	rv := objc.CallMethod[func(param1 ICollectionView, param2 string, param3 foundation.IIndexPath) View](c_, "supplementaryViewProvider")
	return rv
}

func (c_ CollectionViewDiffableDataSource) SetSupplementaryViewProvider(value func(param1 CollectionView, param2 string, param3 foundation.IndexPath) IView) {
	objc.CallMethod[objc.Void](c_, "setSupplementaryViewProvider:", value)
}

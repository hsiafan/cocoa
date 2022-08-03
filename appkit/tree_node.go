// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TreeNodeClass = _TreeNodeClass{objc.GetClass("NSTreeNode")}

type _TreeNodeClass struct {
	objc.Class
}

type ITreeNode interface {
	objc.IObject
	DescendantNodeAtIndexPath(indexPath foundation.IIndexPath) TreeNode
	SortWithSortDescriptors_Recursively(sortDescriptors []foundation.ISortDescriptor, recursively bool)
	RepresentedObject() objc.Object
	IndexPath() foundation.IndexPath
	IsLeaf() bool
	ChildNodes() []TreeNode
	MutableChildNodes() foundation.MutableArray
	ParentNode() TreeNode
}

type TreeNode struct {
	objc.Object
}

func MakeTreeNode(ptr unsafe.Pointer) TreeNode {
	return TreeNode{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TreeNodeClass) TreeNodeWithRepresentedObject(modelObject objc.IObject) TreeNode {
	rv := ffi.CallMethod[TreeNode](tc, "treeNodeWithRepresentedObject:", modelObject)
	return rv
}

func (t_ TreeNode) InitWithRepresentedObject(modelObject objc.IObject) TreeNode {
	rv := ffi.CallMethod[TreeNode](t_, "initWithRepresentedObject:", modelObject)
	rv.Autorelease()
	return rv
}

func (tc _TreeNodeClass) Alloc() TreeNode {
	rv := ffi.CallMethod[TreeNode](tc, "alloc")
	return rv
}

func (t_ TreeNode) Init() TreeNode {
	rv := ffi.CallMethod[TreeNode](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TreeNodeClass) New() TreeNode {
	rv := ffi.CallMethod[TreeNode](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTreeNode() TreeNode {
	return TreeNodeClass.New()
}

func (t_ TreeNode) DescendantNodeAtIndexPath(indexPath foundation.IIndexPath) TreeNode {
	rv := ffi.CallMethod[TreeNode](t_, "descendantNodeAtIndexPath:", indexPath)
	return rv
}

func (t_ TreeNode) SortWithSortDescriptors_Recursively(sortDescriptors []foundation.ISortDescriptor, recursively bool) {
	ffi.CallMethod[ffi.Void](t_, "sortWithSortDescriptors:recursively:", sortDescriptors, recursively)
}

func (t_ TreeNode) RepresentedObject() objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "representedObject")
	return rv
}

func (t_ TreeNode) IndexPath() foundation.IndexPath {
	rv := ffi.CallMethod[foundation.IndexPath](t_, "indexPath")
	return rv
}

func (t_ TreeNode) IsLeaf() bool {
	rv := ffi.CallMethod[bool](t_, "isLeaf")
	return rv
}

func (t_ TreeNode) ChildNodes() []TreeNode {
	rv := ffi.CallMethod[[]TreeNode](t_, "childNodes")
	return rv
}

func (t_ TreeNode) MutableChildNodes() foundation.MutableArray {
	rv := ffi.CallMethod[foundation.MutableArray](t_, "mutableChildNodes")
	return rv
}

func (t_ TreeNode) ParentNode() TreeNode {
	rv := ffi.CallMethod[TreeNode](t_, "parentNode")
	return rv
}

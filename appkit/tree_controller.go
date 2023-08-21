// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TreeControllerClass = _TreeControllerClass{objc.GetClass("NSTreeController")}

type _TreeControllerClass struct {
	objc.Class
}

type ITreeController interface {
	IObjectController
	RearrangeObjects()
	SetSelectionIndexPath(indexPath foundation.IIndexPath) bool
	SetSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool
	AddSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool
	RemoveSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool
	AddChild(sender objc.IObject)
	Insert(sender objc.IObject)
	InsertChild(sender objc.IObject)
	InsertObject_AtArrangedObjectIndexPath(object objc.IObject, indexPath foundation.IIndexPath)
	InsertObjects_AtArrangedObjectIndexPaths(objects []objc.IObject, indexPaths []foundation.IIndexPath)
	RemoveObjectAtArrangedObjectIndexPath(indexPath foundation.IIndexPath)
	RemoveObjectsAtArrangedObjectIndexPaths(indexPaths []foundation.IIndexPath)
	MoveNode_ToIndexPath(node ITreeNode, indexPath foundation.IIndexPath)
	MoveNodes_ToIndexPath(nodes []ITreeNode, startingIndexPath foundation.IIndexPath)
	ChildrenKeyPathForNode(node ITreeNode) string
	CountKeyPathForNode(node ITreeNode) string
	LeafKeyPathForNode(node ITreeNode) string
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.ISortDescriptor)
	ArrangedObjects() TreeNode
	SelectionIndexPath() foundation.IndexPath
	SelectionIndexPaths() []foundation.IndexPath
	SelectedNodes() []TreeNode
	SelectsInsertedObjects() bool
	SetSelectsInsertedObjects(value bool)
	AvoidsEmptySelection() bool
	SetAvoidsEmptySelection(value bool)
	PreservesSelection() bool
	SetPreservesSelection(value bool)
	AlwaysUsesMultipleValuesMarker() bool
	SetAlwaysUsesMultipleValuesMarker(value bool)
	CanAddChild() bool
	CanInsert() bool
	CanInsertChild() bool
	ChildrenKeyPath() string
	SetChildrenKeyPath(value string)
	CountKeyPath() string
	SetCountKeyPath(value string)
	LeafKeyPath() string
	SetLeafKeyPath(value string)
}

type TreeController struct {
	ObjectController
}

func MakeTreeController(ptr unsafe.Pointer) TreeController {
	return TreeController{
		ObjectController: MakeObjectController(ptr),
	}
}

func (t_ TreeController) InitWithContent(content objc.IObject) TreeController {
	rv := objc.CallMethod[TreeController](t_, objc.GetSelector("initWithContent:"), objc.ExtractPtr(content))
	return rv
}

func (t_ TreeController) Init() TreeController {
	rv := objc.CallMethod[TreeController](t_, objc.GetSelector("init"))
	return rv
}

func (tc _TreeControllerClass) Alloc() TreeController {
	rv := objc.CallMethod[TreeController](tc, objc.GetSelector("alloc"))
	return rv
}

func (tc _TreeControllerClass) New() TreeController {
	rv := objc.CallMethod[TreeController](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTreeController() TreeController {
	return TreeControllerClass.New()
}

func (t_ TreeController) RearrangeObjects() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("rearrangeObjects"))
}

func (t_ TreeController) SetSelectionIndexPath(indexPath foundation.IIndexPath) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("setSelectionIndexPath:"), objc.ExtractPtr(indexPath))
	return rv
}

func (t_ TreeController) SetSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("setSelectionIndexPaths:"), indexPaths)
	return rv
}

func (t_ TreeController) AddSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("addSelectionIndexPaths:"), indexPaths)
	return rv
}

func (t_ TreeController) RemoveSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("removeSelectionIndexPaths:"), indexPaths)
	return rv
}

func (t_ TreeController) AddChild(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("addChild:"), objc.ExtractPtr(sender))
}

func (t_ TreeController) Insert(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("insert:"), objc.ExtractPtr(sender))
}

func (t_ TreeController) InsertChild(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("insertChild:"), objc.ExtractPtr(sender))
}

func (t_ TreeController) InsertObject_AtArrangedObjectIndexPath(object objc.IObject, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("insertObject:atArrangedObjectIndexPath:"), objc.ExtractPtr(object), objc.ExtractPtr(indexPath))
}

func (t_ TreeController) InsertObjects_AtArrangedObjectIndexPaths(objects []objc.IObject, indexPaths []foundation.IIndexPath) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("insertObjects:atArrangedObjectIndexPaths:"), objects, indexPaths)
}

func (t_ TreeController) RemoveObjectAtArrangedObjectIndexPath(indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("removeObjectAtArrangedObjectIndexPath:"), objc.ExtractPtr(indexPath))
}

func (t_ TreeController) RemoveObjectsAtArrangedObjectIndexPaths(indexPaths []foundation.IIndexPath) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("removeObjectsAtArrangedObjectIndexPaths:"), indexPaths)
}

func (t_ TreeController) MoveNode_ToIndexPath(node ITreeNode, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("moveNode:toIndexPath:"), objc.ExtractPtr(node), objc.ExtractPtr(indexPath))
}

func (t_ TreeController) MoveNodes_ToIndexPath(nodes []ITreeNode, startingIndexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("moveNodes:toIndexPath:"), nodes, objc.ExtractPtr(startingIndexPath))
}

func (t_ TreeController) ChildrenKeyPathForNode(node ITreeNode) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("childrenKeyPathForNode:"), objc.ExtractPtr(node))
	return rv
}

func (t_ TreeController) CountKeyPathForNode(node ITreeNode) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("countKeyPathForNode:"), objc.ExtractPtr(node))
	return rv
}

func (t_ TreeController) LeafKeyPathForNode(node ITreeNode) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("leafKeyPathForNode:"), objc.ExtractPtr(node))
	return rv
}

func (t_ TreeController) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.CallMethod[[]foundation.SortDescriptor](t_, objc.GetSelector("sortDescriptors"))
	return rv
}

func (t_ TreeController) SetSortDescriptors(value []foundation.ISortDescriptor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSortDescriptors:"), value)
}

func (t_ TreeController) ArrangedObjects() TreeNode {
	rv := objc.CallMethod[TreeNode](t_, objc.GetSelector("arrangedObjects"))
	return rv
}

func (t_ TreeController) SelectionIndexPath() foundation.IndexPath {
	rv := objc.CallMethod[foundation.IndexPath](t_, objc.GetSelector("selectionIndexPath"))
	return rv
}

func (t_ TreeController) SelectionIndexPaths() []foundation.IndexPath {
	rv := objc.CallMethod[[]foundation.IndexPath](t_, objc.GetSelector("selectionIndexPaths"))
	return rv
}

func (t_ TreeController) SelectedNodes() []TreeNode {
	rv := objc.CallMethod[[]TreeNode](t_, objc.GetSelector("selectedNodes"))
	return rv
}

func (t_ TreeController) SelectsInsertedObjects() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("selectsInsertedObjects"))
	return rv
}

func (t_ TreeController) SetSelectsInsertedObjects(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectsInsertedObjects:"), value)
}

func (t_ TreeController) AvoidsEmptySelection() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("avoidsEmptySelection"))
	return rv
}

func (t_ TreeController) SetAvoidsEmptySelection(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAvoidsEmptySelection:"), value)
}

func (t_ TreeController) PreservesSelection() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("preservesSelection"))
	return rv
}

func (t_ TreeController) SetPreservesSelection(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setPreservesSelection:"), value)
}

func (t_ TreeController) AlwaysUsesMultipleValuesMarker() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("alwaysUsesMultipleValuesMarker"))
	return rv
}

func (t_ TreeController) SetAlwaysUsesMultipleValuesMarker(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAlwaysUsesMultipleValuesMarker:"), value)
}

func (t_ TreeController) CanAddChild() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("canAddChild"))
	return rv
}

func (t_ TreeController) CanInsert() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("canInsert"))
	return rv
}

func (t_ TreeController) CanInsertChild() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("canInsertChild"))
	return rv
}

func (t_ TreeController) ChildrenKeyPath() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("childrenKeyPath"))
	return rv
}

func (t_ TreeController) SetChildrenKeyPath(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setChildrenKeyPath:"), value)
}

func (t_ TreeController) CountKeyPath() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("countKeyPath"))
	return rv
}

func (t_ TreeController) SetCountKeyPath(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setCountKeyPath:"), value)
}

func (t_ TreeController) LeafKeyPath() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("leafKeyPath"))
	return rv
}

func (t_ TreeController) SetLeafKeyPath(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLeafKeyPath:"), value)
}

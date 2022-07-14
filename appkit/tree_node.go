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
	rv := ffi.CallMethod[TreeController](t_, "initWithContent:", content)
	rv.Autorelease()
	return rv
}

func (t_ TreeController) Init() TreeController {
	rv := ffi.CallMethod[TreeController](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TreeControllerClass) Alloc() TreeController {
	rv := ffi.CallMethod[TreeController](tc, "alloc")
	return rv
}

func (tc _TreeControllerClass) New() TreeController {
	rv := ffi.CallMethod[TreeController](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTreeController() TreeController {
	return TreeControllerClass.New()
}

func (t_ TreeController) RearrangeObjects() {
	ffi.CallMethod[ffi.Void](t_, "rearrangeObjects")
}

func (t_ TreeController) SetSelectionIndexPath(indexPath foundation.IIndexPath) bool {
	rv := ffi.CallMethod[bool](t_, "setSelectionIndexPath:", indexPath)
	return rv
}

func (t_ TreeController) SetSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool {
	rv := ffi.CallMethod[bool](t_, "setSelectionIndexPaths:", indexPaths)
	return rv
}

func (t_ TreeController) AddSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool {
	rv := ffi.CallMethod[bool](t_, "addSelectionIndexPaths:", indexPaths)
	return rv
}

func (t_ TreeController) RemoveSelectionIndexPaths(indexPaths []foundation.IIndexPath) bool {
	rv := ffi.CallMethod[bool](t_, "removeSelectionIndexPaths:", indexPaths)
	return rv
}

func (t_ TreeController) AddChild(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "addChild:", sender)
}

func (t_ TreeController) Insert(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "insert:", sender)
}

func (t_ TreeController) InsertChild(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "insertChild:", sender)
}

func (t_ TreeController) InsertObject_AtArrangedObjectIndexPath(object objc.IObject, indexPath foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](t_, "insertObject:atArrangedObjectIndexPath:", object, indexPath)
}

func (t_ TreeController) InsertObjects_AtArrangedObjectIndexPaths(objects []objc.IObject, indexPaths []foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](t_, "insertObjects:atArrangedObjectIndexPaths:", objects, indexPaths)
}

func (t_ TreeController) RemoveObjectAtArrangedObjectIndexPath(indexPath foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](t_, "removeObjectAtArrangedObjectIndexPath:", indexPath)
}

func (t_ TreeController) RemoveObjectsAtArrangedObjectIndexPaths(indexPaths []foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](t_, "removeObjectsAtArrangedObjectIndexPaths:", indexPaths)
}

func (t_ TreeController) MoveNode_ToIndexPath(node ITreeNode, indexPath foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](t_, "moveNode:toIndexPath:", node, indexPath)
}

func (t_ TreeController) MoveNodes_ToIndexPath(nodes []ITreeNode, startingIndexPath foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](t_, "moveNodes:toIndexPath:", nodes, startingIndexPath)
}

func (t_ TreeController) ChildrenKeyPathForNode(node ITreeNode) string {
	rv := ffi.CallMethod[string](t_, "childrenKeyPathForNode:", node)
	return rv
}

func (t_ TreeController) CountKeyPathForNode(node ITreeNode) string {
	rv := ffi.CallMethod[string](t_, "countKeyPathForNode:", node)
	return rv
}

func (t_ TreeController) LeafKeyPathForNode(node ITreeNode) string {
	rv := ffi.CallMethod[string](t_, "leafKeyPathForNode:", node)
	return rv
}

func (t_ TreeController) SortDescriptors() []foundation.SortDescriptor {
	rv := ffi.CallMethod[[]foundation.SortDescriptor](t_, "sortDescriptors")
	return rv
}

func (t_ TreeController) SetSortDescriptors(value []foundation.ISortDescriptor) {
	ffi.CallMethod[ffi.Void](t_, "setSortDescriptors:", value)
}

func (t_ TreeController) ArrangedObjects() TreeNode {
	rv := ffi.CallMethod[TreeNode](t_, "arrangedObjects")
	return rv
}

func (t_ TreeController) SelectionIndexPath() foundation.IndexPath {
	rv := ffi.CallMethod[foundation.IndexPath](t_, "selectionIndexPath")
	return rv
}

func (t_ TreeController) SelectionIndexPaths() []foundation.IndexPath {
	rv := ffi.CallMethod[[]foundation.IndexPath](t_, "selectionIndexPaths")
	return rv
}

func (t_ TreeController) SelectedNodes() []TreeNode {
	rv := ffi.CallMethod[[]TreeNode](t_, "selectedNodes")
	return rv
}

func (t_ TreeController) SelectsInsertedObjects() bool {
	rv := ffi.CallMethod[bool](t_, "selectsInsertedObjects")
	return rv
}

func (t_ TreeController) SetSelectsInsertedObjects(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setSelectsInsertedObjects:", value)
}

func (t_ TreeController) AvoidsEmptySelection() bool {
	rv := ffi.CallMethod[bool](t_, "avoidsEmptySelection")
	return rv
}

func (t_ TreeController) SetAvoidsEmptySelection(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAvoidsEmptySelection:", value)
}

func (t_ TreeController) PreservesSelection() bool {
	rv := ffi.CallMethod[bool](t_, "preservesSelection")
	return rv
}

func (t_ TreeController) SetPreservesSelection(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setPreservesSelection:", value)
}

func (t_ TreeController) AlwaysUsesMultipleValuesMarker() bool {
	rv := ffi.CallMethod[bool](t_, "alwaysUsesMultipleValuesMarker")
	return rv
}

func (t_ TreeController) SetAlwaysUsesMultipleValuesMarker(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAlwaysUsesMultipleValuesMarker:", value)
}

func (t_ TreeController) CanAddChild() bool {
	rv := ffi.CallMethod[bool](t_, "canAddChild")
	return rv
}

func (t_ TreeController) CanInsert() bool {
	rv := ffi.CallMethod[bool](t_, "canInsert")
	return rv
}

func (t_ TreeController) CanInsertChild() bool {
	rv := ffi.CallMethod[bool](t_, "canInsertChild")
	return rv
}

func (t_ TreeController) ChildrenKeyPath() string {
	rv := ffi.CallMethod[string](t_, "childrenKeyPath")
	return rv
}

func (t_ TreeController) SetChildrenKeyPath(value string) {
	ffi.CallMethod[ffi.Void](t_, "setChildrenKeyPath:", value)
}

func (t_ TreeController) CountKeyPath() string {
	rv := ffi.CallMethod[string](t_, "countKeyPath")
	return rv
}

func (t_ TreeController) SetCountKeyPath(value string) {
	ffi.CallMethod[ffi.Void](t_, "setCountKeyPath:", value)
}

func (t_ TreeController) LeafKeyPath() string {
	rv := ffi.CallMethod[string](t_, "leafKeyPath")
	return rv
}

func (t_ TreeController) SetLeafKeyPath(value string) {
	ffi.CallMethod[ffi.Void](t_, "setLeafKeyPath:", value)
}

// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var OperationQueueClass = _OperationQueueClass{objc.GetClass("NSOperationQueue")}

type _OperationQueueClass struct {
	objc.Class
}

type IOperationQueue interface {
	objc.IObject
	AddOperation(op IOperation)
	AddOperations_WaitUntilFinished(ops []IOperation, wait bool)
	AddOperationWithBlock(block func())
	AddBarrierBlock(barrier func())
	CancelAllOperations()
	WaitUntilAllOperationsAreFinished()
	// deprecated
	Operations() []Operation
	// deprecated
	OperationCount() uint
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	MaxConcurrentOperationCount() int
	SetMaxConcurrentOperationCount(value int)
	Progress() Progress
	IsSuspended() bool
	SetSuspended(value bool)
	Name() string
	SetName(value string)
}

type OperationQueue struct {
	objc.Object
}

func MakeOperationQueue(ptr unsafe.Pointer) OperationQueue {
	return OperationQueue{
		Object: objc.MakeObject(ptr),
	}
}

func (oc _OperationQueueClass) Alloc() OperationQueue {
	rv := objc.CallMethod[OperationQueue](oc, "alloc")
	return rv
}

func (oc _OperationQueueClass) New() OperationQueue {
	rv := objc.CallMethod[OperationQueue](oc, "new")
	rv.Autorelease()
	return rv
}

func NewOperationQueue() OperationQueue {
	return OperationQueueClass.New()
}

func (o_ OperationQueue) Init() OperationQueue {
	rv := objc.CallMethod[OperationQueue](o_, "init")
	return rv
}

func (o_ OperationQueue) AddOperation(op IOperation) {
	objc.CallMethod[objc.Void](o_, "addOperation:", op)
}

func (o_ OperationQueue) AddOperations_WaitUntilFinished(ops []IOperation, wait bool) {
	objc.CallMethod[objc.Void](o_, "addOperations:waitUntilFinished:", ops, wait)
}

func (o_ OperationQueue) AddOperationWithBlock(block func()) {
	objc.CallMethod[objc.Void](o_, "addOperationWithBlock:", block)
}

func (o_ OperationQueue) AddBarrierBlock(barrier func()) {
	objc.CallMethod[objc.Void](o_, "addBarrierBlock:", barrier)
}

func (o_ OperationQueue) CancelAllOperations() {
	objc.CallMethod[objc.Void](o_, "cancelAllOperations")
}

func (o_ OperationQueue) WaitUntilAllOperationsAreFinished() {
	objc.CallMethod[objc.Void](o_, "waitUntilAllOperationsAreFinished")
}

func (oc _OperationQueueClass) MainQueue() OperationQueue {
	rv := objc.CallMethod[OperationQueue](oc, "mainQueue")
	return rv
}

func (oc _OperationQueueClass) CurrentQueue() OperationQueue {
	rv := objc.CallMethod[OperationQueue](oc, "currentQueue")
	return rv
}

// deprecated
func (o_ OperationQueue) Operations() []Operation {
	rv := objc.CallMethod[[]Operation](o_, "operations")
	return rv
}

// deprecated
func (o_ OperationQueue) OperationCount() uint {
	rv := objc.CallMethod[uint](o_, "operationCount")
	return rv
}

func (o_ OperationQueue) QualityOfService() QualityOfService {
	rv := objc.CallMethod[QualityOfService](o_, "qualityOfService")
	return rv
}

func (o_ OperationQueue) SetQualityOfService(value QualityOfService) {
	objc.CallMethod[objc.Void](o_, "setQualityOfService:", value)
}

func (o_ OperationQueue) MaxConcurrentOperationCount() int {
	rv := objc.CallMethod[int](o_, "maxConcurrentOperationCount")
	return rv
}

func (o_ OperationQueue) SetMaxConcurrentOperationCount(value int) {
	objc.CallMethod[objc.Void](o_, "setMaxConcurrentOperationCount:", value)
}

func (o_ OperationQueue) Progress() Progress {
	rv := objc.CallMethod[Progress](o_, "progress")
	return rv
}

func (o_ OperationQueue) IsSuspended() bool {
	rv := objc.CallMethod[bool](o_, "isSuspended")
	return rv
}

func (o_ OperationQueue) SetSuspended(value bool) {
	objc.CallMethod[objc.Void](o_, "setSuspended:", value)
}

func (o_ OperationQueue) Name() string {
	rv := objc.CallMethod[string](o_, "name")
	return rv
}

func (o_ OperationQueue) SetName(value string) {
	objc.CallMethod[objc.Void](o_, "setName:", value)
}

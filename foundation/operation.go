// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var OperationClass = _OperationClass{objc.GetClass("NSOperation")}

type _OperationClass struct {
	objc.Class
}

type IOperation interface {
	objc.IObject
	Start()
	Main()
	Cancel()
	AddDependency(op IOperation)
	RemoveDependency(op IOperation)
	WaitUntilFinished()
	CompletionBlock() func()
	SetCompletionBlock(value func())
	IsCancelled() bool
	IsExecuting() bool
	IsFinished() bool
	IsConcurrent() bool
	IsAsynchronous() bool
	IsReady() bool
	Name() string
	SetName(value string)
	Dependencies() []Operation
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	// deprecated
	ThreadPriority() float64
	// deprecated
	SetThreadPriority(value float64)
	QueuePriority() OperationQueuePriority
	SetQueuePriority(value OperationQueuePriority)
}

type Operation struct {
	objc.Object
}

func MakeOperation(ptr unsafe.Pointer) Operation {
	return Operation{
		Object: objc.MakeObject(ptr),
	}
}

func (oc _OperationClass) Alloc() Operation {
	rv := ffi.CallMethod[Operation](oc, "alloc")
	return rv
}

func (o_ Operation) Init() Operation {
	rv := ffi.CallMethod[Operation](o_, "init")
	return rv
}

func (oc _OperationClass) New() Operation {
	rv := ffi.CallMethod[Operation](oc, "new")
	rv.Autorelease()
	return rv
}

func NewOperation() Operation {
	return OperationClass.New()
}

func (o_ Operation) Start() {
	ffi.CallMethod[ffi.Void](o_, "start")
}

func (o_ Operation) Main() {
	ffi.CallMethod[ffi.Void](o_, "main")
}

func (o_ Operation) Cancel() {
	ffi.CallMethod[ffi.Void](o_, "cancel")
}

func (o_ Operation) AddDependency(op IOperation) {
	ffi.CallMethod[ffi.Void](o_, "addDependency:", op)
}

func (o_ Operation) RemoveDependency(op IOperation) {
	ffi.CallMethod[ffi.Void](o_, "removeDependency:", op)
}

func (o_ Operation) WaitUntilFinished() {
	ffi.CallMethod[ffi.Void](o_, "waitUntilFinished")
}

func (o_ Operation) CompletionBlock() func() {
	rv := ffi.CallMethod[func()](o_, "completionBlock")
	return rv
}

func (o_ Operation) SetCompletionBlock(value func()) {
	ffi.CallMethod[ffi.Void](o_, "setCompletionBlock:", value)
}

func (o_ Operation) IsCancelled() bool {
	rv := ffi.CallMethod[bool](o_, "isCancelled")
	return rv
}

func (o_ Operation) IsExecuting() bool {
	rv := ffi.CallMethod[bool](o_, "isExecuting")
	return rv
}

func (o_ Operation) IsFinished() bool {
	rv := ffi.CallMethod[bool](o_, "isFinished")
	return rv
}

func (o_ Operation) IsConcurrent() bool {
	rv := ffi.CallMethod[bool](o_, "isConcurrent")
	return rv
}

func (o_ Operation) IsAsynchronous() bool {
	rv := ffi.CallMethod[bool](o_, "isAsynchronous")
	return rv
}

func (o_ Operation) IsReady() bool {
	rv := ffi.CallMethod[bool](o_, "isReady")
	return rv
}

func (o_ Operation) Name() string {
	rv := ffi.CallMethod[string](o_, "name")
	return rv
}

func (o_ Operation) SetName(value string) {
	ffi.CallMethod[ffi.Void](o_, "setName:", value)
}

func (o_ Operation) Dependencies() []Operation {
	rv := ffi.CallMethod[[]Operation](o_, "dependencies")
	return rv
}

func (o_ Operation) QualityOfService() QualityOfService {
	rv := ffi.CallMethod[QualityOfService](o_, "qualityOfService")
	return rv
}

func (o_ Operation) SetQualityOfService(value QualityOfService) {
	ffi.CallMethod[ffi.Void](o_, "setQualityOfService:", value)
}

// deprecated
func (o_ Operation) ThreadPriority() float64 {
	rv := ffi.CallMethod[float64](o_, "threadPriority")
	return rv
}

// deprecated
func (o_ Operation) SetThreadPriority(value float64) {
	ffi.CallMethod[ffi.Void](o_, "setThreadPriority:", value)
}

func (o_ Operation) QueuePriority() OperationQueuePriority {
	rv := ffi.CallMethod[OperationQueuePriority](o_, "queuePriority")
	return rv
}

func (o_ Operation) SetQueuePriority(value OperationQueuePriority) {
	ffi.CallMethod[ffi.Void](o_, "setQueuePriority:", value)
}

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
	rv := ffi.CallMethod[OperationQueue](oc, "alloc")
	return rv
}

func (o_ OperationQueue) Init() OperationQueue {
	rv := ffi.CallMethod[OperationQueue](o_, "init")
	return rv
}

func (oc _OperationQueueClass) New() OperationQueue {
	rv := ffi.CallMethod[OperationQueue](oc, "new")
	rv.Autorelease()
	return rv
}

func NewOperationQueue() OperationQueue {
	return OperationQueueClass.New()
}

func (o_ OperationQueue) AddOperation(op IOperation) {
	ffi.CallMethod[ffi.Void](o_, "addOperation:", op)
}

func (o_ OperationQueue) AddOperations_WaitUntilFinished(ops []IOperation, wait bool) {
	ffi.CallMethod[ffi.Void](o_, "addOperations:waitUntilFinished:", ops, wait)
}

func (o_ OperationQueue) AddOperationWithBlock(block func()) {
	ffi.CallMethod[ffi.Void](o_, "addOperationWithBlock:", block)
}

func (o_ OperationQueue) AddBarrierBlock(barrier func()) {
	ffi.CallMethod[ffi.Void](o_, "addBarrierBlock:", barrier)
}

func (o_ OperationQueue) CancelAllOperations() {
	ffi.CallMethod[ffi.Void](o_, "cancelAllOperations")
}

func (o_ OperationQueue) WaitUntilAllOperationsAreFinished() {
	ffi.CallMethod[ffi.Void](o_, "waitUntilAllOperationsAreFinished")
}

func (oc _OperationQueueClass) MainQueue() OperationQueue {
	rv := ffi.CallMethod[OperationQueue](oc, "mainQueue")
	return rv
}

func (oc _OperationQueueClass) CurrentQueue() OperationQueue {
	rv := ffi.CallMethod[OperationQueue](oc, "currentQueue")
	return rv
}

// deprecated
func (o_ OperationQueue) Operations() []Operation {
	rv := ffi.CallMethod[[]Operation](o_, "operations")
	return rv
}

// deprecated
func (o_ OperationQueue) OperationCount() uint {
	rv := ffi.CallMethod[uint](o_, "operationCount")
	return rv
}

func (o_ OperationQueue) QualityOfService() QualityOfService {
	rv := ffi.CallMethod[QualityOfService](o_, "qualityOfService")
	return rv
}

func (o_ OperationQueue) SetQualityOfService(value QualityOfService) {
	ffi.CallMethod[ffi.Void](o_, "setQualityOfService:", value)
}

func (o_ OperationQueue) MaxConcurrentOperationCount() int {
	rv := ffi.CallMethod[int](o_, "maxConcurrentOperationCount")
	return rv
}

func (o_ OperationQueue) SetMaxConcurrentOperationCount(value int) {
	ffi.CallMethod[ffi.Void](o_, "setMaxConcurrentOperationCount:", value)
}

func (o_ OperationQueue) Progress() Progress {
	rv := ffi.CallMethod[Progress](o_, "progress")
	return rv
}

func (o_ OperationQueue) IsSuspended() bool {
	rv := ffi.CallMethod[bool](o_, "isSuspended")
	return rv
}

func (o_ OperationQueue) SetSuspended(value bool) {
	ffi.CallMethod[ffi.Void](o_, "setSuspended:", value)
}

func (o_ OperationQueue) Name() string {
	rv := ffi.CallMethod[string](o_, "name")
	return rv
}

func (o_ OperationQueue) SetName(value string) {
	ffi.CallMethod[ffi.Void](o_, "setName:", value)
}

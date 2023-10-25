// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[Operation](oc, objc.SEL("alloc"))
	return rv
}

func (oc _OperationClass) New() Operation {
	rv := objc.CallMethod[Operation](oc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewOperation() Operation {
	return OperationClass.New()
}

func (o_ Operation) Init() Operation {
	rv := objc.CallMethod[Operation](o_, objc.SEL("init"))
	return rv
}

func (o_ Operation) Start() {
	objc.CallMethod[objc.Void](o_, objc.SEL("start"))
}

func (o_ Operation) Main() {
	objc.CallMethod[objc.Void](o_, objc.SEL("main"))
}

func (o_ Operation) Cancel() {
	objc.CallMethod[objc.Void](o_, objc.SEL("cancel"))
}

func (o_ Operation) AddDependency(op IOperation) {
	objc.CallMethod[objc.Void](o_, objc.SEL("addDependency:"), objc.ExtractPtr(op))
}

func (o_ Operation) RemoveDependency(op IOperation) {
	objc.CallMethod[objc.Void](o_, objc.SEL("removeDependency:"), objc.ExtractPtr(op))
}

func (o_ Operation) WaitUntilFinished() {
	objc.CallMethod[objc.Void](o_, objc.SEL("waitUntilFinished"))
}

func (o_ Operation) CompletionBlock() func() {
	rv := objc.CallMethod[func()](o_, objc.SEL("completionBlock"))
	return rv
}

func (o_ Operation) SetCompletionBlock(value func()) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setCompletionBlock:"), value)
}

func (o_ Operation) IsCancelled() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("isCancelled"))
	return rv
}

func (o_ Operation) IsExecuting() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("isExecuting"))
	return rv
}

func (o_ Operation) IsFinished() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("isFinished"))
	return rv
}

func (o_ Operation) IsConcurrent() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("isConcurrent"))
	return rv
}

func (o_ Operation) IsAsynchronous() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("isAsynchronous"))
	return rv
}

func (o_ Operation) IsReady() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("isReady"))
	return rv
}

func (o_ Operation) Name() string {
	rv := objc.CallMethod[string](o_, objc.SEL("name"))
	return rv
}

func (o_ Operation) SetName(value string) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setName:"), value)
}

func (o_ Operation) Dependencies() []Operation {
	rv := objc.CallMethod[[]Operation](o_, objc.SEL("dependencies"))
	return rv
}

func (o_ Operation) QualityOfService() QualityOfService {
	rv := objc.CallMethod[QualityOfService](o_, objc.SEL("qualityOfService"))
	return rv
}

func (o_ Operation) SetQualityOfService(value QualityOfService) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setQualityOfService:"), value)
}

// deprecated
func (o_ Operation) ThreadPriority() float64 {
	rv := objc.CallMethod[float64](o_, objc.SEL("threadPriority"))
	return rv
}

// deprecated
func (o_ Operation) SetThreadPriority(value float64) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setThreadPriority:"), value)
}

func (o_ Operation) QueuePriority() OperationQueuePriority {
	rv := objc.CallMethod[OperationQueuePriority](o_, objc.SEL("queuePriority"))
	return rv
}

func (o_ Operation) SetQueuePriority(value OperationQueuePriority) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setQueuePriority:"), value)
}

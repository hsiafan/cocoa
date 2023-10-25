// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var ProgressClass = _ProgressClass{objc.GetClass("NSProgress")}

type _ProgressClass struct {
	objc.Class
}

type IProgress interface {
	objc.IObject
	BecomeCurrentWithPendingUnitCount(unitCount int64)
	AddChild_WithPendingUnitCount(child IProgress, inUnitCount int64)
	PerformAsCurrentWithPendingUnitCount_UsingBlock(unitCount int64, work func())
	ResignCurrent()
	Cancel()
	Pause()
	Resume()
	SetUserInfoObject_ForKey(objectOrNil objc.IObject, key ProgressUserInfoKey)
	Publish()
	Unpublish()
	TotalUnitCount() int64
	SetTotalUnitCount(value int64)
	CompletedUnitCount() int64
	SetCompletedUnitCount(value int64)
	LocalizedDescription() string
	SetLocalizedDescription(value string)
	LocalizedAdditionalDescription() string
	SetLocalizedAdditionalDescription(value string)
	IsCancellable() bool
	SetCancellable(value bool)
	IsCancelled() bool
	CancellationHandler() func()
	SetCancellationHandler(value func())
	IsPausable() bool
	SetPausable(value bool)
	IsPaused() bool
	PausingHandler() func()
	SetPausingHandler(value func())
	IsIndeterminate() bool
	FractionCompleted() float64
	IsFinished() bool
	ResumingHandler() func()
	SetResumingHandler(value func())
	Kind() ProgressKind
	SetKind(value ProgressKind)
	EstimatedTimeRemaining() Number
	SetEstimatedTimeRemaining(value INumber)
	Throughput() Number
	SetThroughput(value INumber)
	UserInfo() map[ProgressUserInfoKey]objc.Object
	FileOperationKind() ProgressFileOperationKind
	SetFileOperationKind(value ProgressFileOperationKind)
	FileURL() URL
	SetFileURL(value IURL)
	FileTotalCount() Number
	SetFileTotalCount(value INumber)
	FileCompletedCount() Number
	SetFileCompletedCount(value INumber)
	IsOld() bool
}

type Progress struct {
	objc.Object
}

func MakeProgress(ptr unsafe.Pointer) Progress {
	return Progress{
		Object: objc.MakeObject(ptr),
	}
}

func (p_ Progress) InitWithParent_UserInfo(parentProgressOrNil IProgress, userInfoOrNil map[ProgressUserInfoKey]objc.IObject) Progress {
	rv := objc.CallMethod[Progress](p_, objc.SEL("initWithParent:userInfo:"), objc.ExtractPtr(parentProgressOrNil), userInfoOrNil)
	return rv
}

func (pc _ProgressClass) Alloc() Progress {
	rv := objc.CallMethod[Progress](pc, objc.SEL("alloc"))
	return rv
}

func (pc _ProgressClass) New() Progress {
	rv := objc.CallMethod[Progress](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewProgress() Progress {
	return ProgressClass.New()
}

func (p_ Progress) Init() Progress {
	rv := objc.CallMethod[Progress](p_, objc.SEL("init"))
	return rv
}

func (pc _ProgressClass) DiscreteProgressWithTotalUnitCount(unitCount int64) Progress {
	rv := objc.CallMethod[Progress](pc, objc.SEL("discreteProgressWithTotalUnitCount:"), unitCount)
	return rv
}

func (pc _ProgressClass) ProgressWithTotalUnitCount(unitCount int64) Progress {
	rv := objc.CallMethod[Progress](pc, objc.SEL("progressWithTotalUnitCount:"), unitCount)
	return rv
}

func (pc _ProgressClass) ProgressWithTotalUnitCount_Parent_PendingUnitCount(unitCount int64, parent IProgress, portionOfParentTotalUnitCount int64) Progress {
	rv := objc.CallMethod[Progress](pc, objc.SEL("progressWithTotalUnitCount:parent:pendingUnitCount:"), unitCount, objc.ExtractPtr(parent), portionOfParentTotalUnitCount)
	return rv
}

func (pc _ProgressClass) CurrentProgress() Progress {
	rv := objc.CallMethod[Progress](pc, objc.SEL("currentProgress"))
	return rv
}

func (p_ Progress) BecomeCurrentWithPendingUnitCount(unitCount int64) {
	objc.CallMethod[objc.Void](p_, objc.SEL("becomeCurrentWithPendingUnitCount:"), unitCount)
}

func (p_ Progress) AddChild_WithPendingUnitCount(child IProgress, inUnitCount int64) {
	objc.CallMethod[objc.Void](p_, objc.SEL("addChild:withPendingUnitCount:"), objc.ExtractPtr(child), inUnitCount)
}

func (p_ Progress) PerformAsCurrentWithPendingUnitCount_UsingBlock(unitCount int64, work func()) {
	objc.CallMethod[objc.Void](p_, objc.SEL("performAsCurrentWithPendingUnitCount:usingBlock:"), unitCount, work)
}

func (p_ Progress) ResignCurrent() {
	objc.CallMethod[objc.Void](p_, objc.SEL("resignCurrent"))
}

func (p_ Progress) Cancel() {
	objc.CallMethod[objc.Void](p_, objc.SEL("cancel"))
}

func (p_ Progress) Pause() {
	objc.CallMethod[objc.Void](p_, objc.SEL("pause"))
}

func (p_ Progress) Resume() {
	objc.CallMethod[objc.Void](p_, objc.SEL("resume"))
}

func (p_ Progress) SetUserInfoObject_ForKey(objectOrNil objc.IObject, key ProgressUserInfoKey) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setUserInfoObject:forKey:"), objc.ExtractPtr(objectOrNil), key)
}

func (p_ Progress) Publish() {
	objc.CallMethod[objc.Void](p_, objc.SEL("publish"))
}

func (p_ Progress) Unpublish() {
	objc.CallMethod[objc.Void](p_, objc.SEL("unpublish"))
}

func (pc _ProgressClass) RemoveSubscriber(subscriber objc.IObject) {
	objc.CallMethod[objc.Void](pc, objc.SEL("removeSubscriber:"), objc.ExtractPtr(subscriber))
}

func (p_ Progress) TotalUnitCount() int64 {
	rv := objc.CallMethod[int64](p_, objc.SEL("totalUnitCount"))
	return rv
}

func (p_ Progress) SetTotalUnitCount(value int64) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setTotalUnitCount:"), value)
}

func (p_ Progress) CompletedUnitCount() int64 {
	rv := objc.CallMethod[int64](p_, objc.SEL("completedUnitCount"))
	return rv
}

func (p_ Progress) SetCompletedUnitCount(value int64) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setCompletedUnitCount:"), value)
}

func (p_ Progress) LocalizedDescription() string {
	rv := objc.CallMethod[string](p_, objc.SEL("localizedDescription"))
	return rv
}

func (p_ Progress) SetLocalizedDescription(value string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setLocalizedDescription:"), value)
}

func (p_ Progress) LocalizedAdditionalDescription() string {
	rv := objc.CallMethod[string](p_, objc.SEL("localizedAdditionalDescription"))
	return rv
}

func (p_ Progress) SetLocalizedAdditionalDescription(value string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setLocalizedAdditionalDescription:"), value)
}

func (p_ Progress) IsCancellable() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isCancellable"))
	return rv
}

func (p_ Progress) SetCancellable(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setCancellable:"), value)
}

func (p_ Progress) IsCancelled() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isCancelled"))
	return rv
}

func (p_ Progress) CancellationHandler() func() {
	rv := objc.CallMethod[func()](p_, objc.SEL("cancellationHandler"))
	return rv
}

func (p_ Progress) SetCancellationHandler(value func()) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setCancellationHandler:"), value)
}

func (p_ Progress) IsPausable() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isPausable"))
	return rv
}

func (p_ Progress) SetPausable(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPausable:"), value)
}

func (p_ Progress) IsPaused() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isPaused"))
	return rv
}

func (p_ Progress) PausingHandler() func() {
	rv := objc.CallMethod[func()](p_, objc.SEL("pausingHandler"))
	return rv
}

func (p_ Progress) SetPausingHandler(value func()) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPausingHandler:"), value)
}

func (p_ Progress) IsIndeterminate() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isIndeterminate"))
	return rv
}

func (p_ Progress) FractionCompleted() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("fractionCompleted"))
	return rv
}

func (p_ Progress) IsFinished() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isFinished"))
	return rv
}

func (p_ Progress) ResumingHandler() func() {
	rv := objc.CallMethod[func()](p_, objc.SEL("resumingHandler"))
	return rv
}

func (p_ Progress) SetResumingHandler(value func()) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setResumingHandler:"), value)
}

func (p_ Progress) Kind() ProgressKind {
	rv := objc.CallMethod[ProgressKind](p_, objc.SEL("kind"))
	return rv
}

func (p_ Progress) SetKind(value ProgressKind) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setKind:"), value)
}

func (p_ Progress) EstimatedTimeRemaining() Number {
	rv := objc.CallMethod[Number](p_, objc.SEL("estimatedTimeRemaining"))
	return rv
}

func (p_ Progress) SetEstimatedTimeRemaining(value INumber) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setEstimatedTimeRemaining:"), objc.ExtractPtr(value))
}

func (p_ Progress) Throughput() Number {
	rv := objc.CallMethod[Number](p_, objc.SEL("throughput"))
	return rv
}

func (p_ Progress) SetThroughput(value INumber) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setThroughput:"), objc.ExtractPtr(value))
}

func (p_ Progress) UserInfo() map[ProgressUserInfoKey]objc.Object {
	rv := objc.CallMethod[map[ProgressUserInfoKey]objc.Object](p_, objc.SEL("userInfo"))
	return rv
}

func (p_ Progress) FileOperationKind() ProgressFileOperationKind {
	rv := objc.CallMethod[ProgressFileOperationKind](p_, objc.SEL("fileOperationKind"))
	return rv
}

func (p_ Progress) SetFileOperationKind(value ProgressFileOperationKind) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setFileOperationKind:"), value)
}

func (p_ Progress) FileURL() URL {
	rv := objc.CallMethod[URL](p_, objc.SEL("fileURL"))
	return rv
}

func (p_ Progress) SetFileURL(value IURL) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setFileURL:"), objc.ExtractPtr(value))
}

func (p_ Progress) FileTotalCount() Number {
	rv := objc.CallMethod[Number](p_, objc.SEL("fileTotalCount"))
	return rv
}

func (p_ Progress) SetFileTotalCount(value INumber) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setFileTotalCount:"), objc.ExtractPtr(value))
}

func (p_ Progress) FileCompletedCount() Number {
	rv := objc.CallMethod[Number](p_, objc.SEL("fileCompletedCount"))
	return rv
}

func (p_ Progress) SetFileCompletedCount(value INumber) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setFileCompletedCount:"), objc.ExtractPtr(value))
}

func (p_ Progress) IsOld() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isOld"))
	return rv
}

// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[Progress](p_, "initWithParent:userInfo:", parentProgressOrNil, userInfoOrNil)
	rv.Autorelease()
	return rv
}

func (pc _ProgressClass) Alloc() Progress {
	rv := ffi.CallMethod[Progress](pc, "alloc")
	return rv
}

func (p_ Progress) Init() Progress {
	rv := ffi.CallMethod[Progress](p_, "init")
	rv.Autorelease()
	return rv
}

func (pc _ProgressClass) New() Progress {
	rv := ffi.CallMethod[Progress](pc, "new")
	rv.Autorelease()
	return rv
}

func NewProgress() Progress {
	return ProgressClass.New()
}

func (pc _ProgressClass) DiscreteProgressWithTotalUnitCount(unitCount int64) Progress {
	rv := ffi.CallMethod[Progress](pc, "discreteProgressWithTotalUnitCount:", unitCount)
	return rv
}

func (pc _ProgressClass) ProgressWithTotalUnitCount(unitCount int64) Progress {
	rv := ffi.CallMethod[Progress](pc, "progressWithTotalUnitCount:", unitCount)
	return rv
}

func (pc _ProgressClass) ProgressWithTotalUnitCount_Parent_PendingUnitCount(unitCount int64, parent IProgress, portionOfParentTotalUnitCount int64) Progress {
	rv := ffi.CallMethod[Progress](pc, "progressWithTotalUnitCount:parent:pendingUnitCount:", unitCount, parent, portionOfParentTotalUnitCount)
	return rv
}

func (pc _ProgressClass) CurrentProgress() Progress {
	rv := ffi.CallMethod[Progress](pc, "currentProgress")
	return rv
}

func (p_ Progress) BecomeCurrentWithPendingUnitCount(unitCount int64) {
	ffi.CallMethod[ffi.Void](p_, "becomeCurrentWithPendingUnitCount:", unitCount)
}

func (p_ Progress) AddChild_WithPendingUnitCount(child IProgress, inUnitCount int64) {
	ffi.CallMethod[ffi.Void](p_, "addChild:withPendingUnitCount:", child, inUnitCount)
}

func (p_ Progress) PerformAsCurrentWithPendingUnitCount_UsingBlock(unitCount int64, work func()) {
	ffi.CallMethod[ffi.Void](p_, "performAsCurrentWithPendingUnitCount:usingBlock:", unitCount, work)
}

func (p_ Progress) ResignCurrent() {
	ffi.CallMethod[ffi.Void](p_, "resignCurrent")
}

func (p_ Progress) Cancel() {
	ffi.CallMethod[ffi.Void](p_, "cancel")
}

func (p_ Progress) Pause() {
	ffi.CallMethod[ffi.Void](p_, "pause")
}

func (p_ Progress) Resume() {
	ffi.CallMethod[ffi.Void](p_, "resume")
}

func (p_ Progress) SetUserInfoObject_ForKey(objectOrNil objc.IObject, key ProgressUserInfoKey) {
	ffi.CallMethod[ffi.Void](p_, "setUserInfoObject:forKey:", objectOrNil, key)
}

func (p_ Progress) Publish() {
	ffi.CallMethod[ffi.Void](p_, "publish")
}

func (p_ Progress) Unpublish() {
	ffi.CallMethod[ffi.Void](p_, "unpublish")
}

func (pc _ProgressClass) RemoveSubscriber(subscriber objc.IObject) {
	ffi.CallMethod[ffi.Void](pc, "removeSubscriber:", subscriber)
}

func (p_ Progress) TotalUnitCount() int64 {
	rv := ffi.CallMethod[int64](p_, "totalUnitCount")
	return rv
}

func (p_ Progress) SetTotalUnitCount(value int64) {
	ffi.CallMethod[ffi.Void](p_, "setTotalUnitCount:", value)
}

func (p_ Progress) CompletedUnitCount() int64 {
	rv := ffi.CallMethod[int64](p_, "completedUnitCount")
	return rv
}

func (p_ Progress) SetCompletedUnitCount(value int64) {
	ffi.CallMethod[ffi.Void](p_, "setCompletedUnitCount:", value)
}

func (p_ Progress) LocalizedDescription() string {
	rv := ffi.CallMethod[string](p_, "localizedDescription")
	return rv
}

func (p_ Progress) SetLocalizedDescription(value string) {
	ffi.CallMethod[ffi.Void](p_, "setLocalizedDescription:", value)
}

func (p_ Progress) LocalizedAdditionalDescription() string {
	rv := ffi.CallMethod[string](p_, "localizedAdditionalDescription")
	return rv
}

func (p_ Progress) SetLocalizedAdditionalDescription(value string) {
	ffi.CallMethod[ffi.Void](p_, "setLocalizedAdditionalDescription:", value)
}

func (p_ Progress) IsCancellable() bool {
	rv := ffi.CallMethod[bool](p_, "isCancellable")
	return rv
}

func (p_ Progress) SetCancellable(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setCancellable:", value)
}

func (p_ Progress) IsCancelled() bool {
	rv := ffi.CallMethod[bool](p_, "isCancelled")
	return rv
}

func (p_ Progress) CancellationHandler() func() {
	rv := ffi.CallMethod[func()](p_, "cancellationHandler")
	return rv
}

func (p_ Progress) SetCancellationHandler(value func()) {
	ffi.CallMethod[ffi.Void](p_, "setCancellationHandler:", value)
}

func (p_ Progress) IsPausable() bool {
	rv := ffi.CallMethod[bool](p_, "isPausable")
	return rv
}

func (p_ Progress) SetPausable(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setPausable:", value)
}

func (p_ Progress) IsPaused() bool {
	rv := ffi.CallMethod[bool](p_, "isPaused")
	return rv
}

func (p_ Progress) PausingHandler() func() {
	rv := ffi.CallMethod[func()](p_, "pausingHandler")
	return rv
}

func (p_ Progress) SetPausingHandler(value func()) {
	ffi.CallMethod[ffi.Void](p_, "setPausingHandler:", value)
}

func (p_ Progress) IsIndeterminate() bool {
	rv := ffi.CallMethod[bool](p_, "isIndeterminate")
	return rv
}

func (p_ Progress) FractionCompleted() float64 {
	rv := ffi.CallMethod[float64](p_, "fractionCompleted")
	return rv
}

func (p_ Progress) IsFinished() bool {
	rv := ffi.CallMethod[bool](p_, "isFinished")
	return rv
}

func (p_ Progress) ResumingHandler() func() {
	rv := ffi.CallMethod[func()](p_, "resumingHandler")
	return rv
}

func (p_ Progress) SetResumingHandler(value func()) {
	ffi.CallMethod[ffi.Void](p_, "setResumingHandler:", value)
}

func (p_ Progress) Kind() ProgressKind {
	rv := ffi.CallMethod[ProgressKind](p_, "kind")
	return rv
}

func (p_ Progress) SetKind(value ProgressKind) {
	ffi.CallMethod[ffi.Void](p_, "setKind:", value)
}

func (p_ Progress) EstimatedTimeRemaining() Number {
	rv := ffi.CallMethod[Number](p_, "estimatedTimeRemaining")
	return rv
}

func (p_ Progress) SetEstimatedTimeRemaining(value INumber) {
	ffi.CallMethod[ffi.Void](p_, "setEstimatedTimeRemaining:", value)
}

func (p_ Progress) Throughput() Number {
	rv := ffi.CallMethod[Number](p_, "throughput")
	return rv
}

func (p_ Progress) SetThroughput(value INumber) {
	ffi.CallMethod[ffi.Void](p_, "setThroughput:", value)
}

func (p_ Progress) UserInfo() map[ProgressUserInfoKey]objc.Object {
	rv := ffi.CallMethod[map[ProgressUserInfoKey]objc.Object](p_, "userInfo")
	return rv
}

func (p_ Progress) FileOperationKind() ProgressFileOperationKind {
	rv := ffi.CallMethod[ProgressFileOperationKind](p_, "fileOperationKind")
	return rv
}

func (p_ Progress) SetFileOperationKind(value ProgressFileOperationKind) {
	ffi.CallMethod[ffi.Void](p_, "setFileOperationKind:", value)
}

func (p_ Progress) FileURL() URL {
	rv := ffi.CallMethod[URL](p_, "fileURL")
	return rv
}

func (p_ Progress) SetFileURL(value IURL) {
	ffi.CallMethod[ffi.Void](p_, "setFileURL:", value)
}

func (p_ Progress) FileTotalCount() Number {
	rv := ffi.CallMethod[Number](p_, "fileTotalCount")
	return rv
}

func (p_ Progress) SetFileTotalCount(value INumber) {
	ffi.CallMethod[ffi.Void](p_, "setFileTotalCount:", value)
}

func (p_ Progress) FileCompletedCount() Number {
	rv := ffi.CallMethod[Number](p_, "fileCompletedCount")
	return rv
}

func (p_ Progress) SetFileCompletedCount(value INumber) {
	ffi.CallMethod[ffi.Void](p_, "setFileCompletedCount:", value)
}

func (p_ Progress) IsOld() bool {
	rv := ffi.CallMethod[bool](p_, "isOld")
	return rv
}

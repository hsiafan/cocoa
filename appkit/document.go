// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var DocumentClass = _DocumentClass{objc.GetClass("NSDocument")}

type _DocumentClass struct {
	objc.Class
}

type IDocument interface {
	objc.IObject
	ReadFromURL_OfType_Error(url foundation.IURL, typeName string, outError *foundation.Error) bool
	ReadFromFileWrapper_OfType_Error(fileWrapper foundation.IFileWrapper, typeName string, outError *foundation.Error) bool
	ReadFromData_OfType_Error(data []byte, typeName string, outError *foundation.Error) bool
	CanAsynchronouslyWriteToURL_OfType_ForSaveOperation(url foundation.IURL, typeName string, saveOperation SaveOperationType) bool
	UnblockUserInteraction()
	WriteToURL_OfType_Error(url foundation.IURL, typeName string, outError *foundation.Error) bool
	WriteSafelyToURL_OfType_ForSaveOperation_Error(url foundation.IURL, typeName string, saveOperation SaveOperationType, outError *foundation.Error) bool
	FileWrapperOfType_Error(typeName string, outError *foundation.Error) foundation.FileWrapper
	DataOfType_Error(typeName string, outError *foundation.Error) []byte
	WriteToURL_OfType_ForSaveOperation_OriginalContentsURL_Error(url foundation.IURL, typeName string, saveOperation SaveOperationType, absoluteOriginalContentsURL foundation.IURL, outError *foundation.Error) bool
	SaveToURL_OfType_ForSaveOperation_Delegate_DidSaveSelector_ContextInfo(url foundation.IURL, typeName string, saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer)
	SaveToURL_OfType_ForSaveOperation_CompletionHandler(url foundation.IURL, typeName string, saveOperation SaveOperationType, completionHandler func(errorOrNil foundation.Error))
	FileAttributesToWriteToURL_OfType_ForSaveOperation_OriginalContentsURL_Error(url foundation.IURL, typeName string, saveOperation SaveOperationType, absoluteOriginalContentsURL foundation.IURL, outError *foundation.Error) map[string]objc.Object
	WritableTypesForSaveOperation(saveOperation SaveOperationType) []string
	FileNameExtensionForType_SaveOperation(typeName string, saveOperation SaveOperationType) string
	MakeWindowControllers()
	AddWindowController(windowController IWindowController)
	RemoveWindowController(windowController IWindowController)
	WindowControllerDidLoadNib(windowController IWindowController)
	WindowControllerWillLoadNib(windowController IWindowController)
	ShouldCloseWindowController_Delegate_ShouldCloseSelector_ContextInfo(windowController IWindowController, delegate objc.IObject, shouldCloseSelector objc.Selector, contextInfo unsafe.Pointer)
	ShowWindows()
	SetWindow(window IWindow)
	SetDisplayName(displayNameOrNil string)
	DefaultDraftName() string
	CheckAutosavingSafetyAndReturnError(outError *foundation.Error) bool
	ScheduleAutosaving()
	AutosaveDocumentWithDelegate_DidAutosaveSelector_ContextInfo(delegate objc.IObject, didAutosaveSelector objc.Selector, contextInfo unsafe.Pointer)
	AutosaveWithImplicitCancellability_CompletionHandler(autosavingIsImplicitlyCancellable bool, completionHandler func(errorOrNil foundation.Error))
	BrowseDocumentVersions(sender objc.IObject)
	StopBrowsingVersionsWithCompletionHandler(completionHandler func())
	MoveDocumentToUbiquityContainer(sender objc.IObject)
	UpdateChangeCountWithToken_ForSaveOperation(changeCountToken objc.IObject, saveOperation SaveOperationType)
	UpdateChangeCount(change DocumentChangeType)
	ChangeCountTokenForSaveOperation(saveOperation SaveOperationType) objc.Object
	InvalidateRestorableState()
	RunModalSavePanelForSaveOperation_Delegate_DidSaveSelector_ContextInfo(saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer)
	PrepareSavePanel(savePanel ISavePanel) bool
	UpdateUserActivityState(activity foundation.IUserActivity)
	ValidateUserInterfaceItem(item ValidatedUserInterfaceItem) bool
	PerformSynchronousFileAccessUsingBlock(block func())
	ContinueActivityUsingBlock(block func())
	ContinueAsynchronousWorkOnMainThreadUsingBlock(block func())
	PrintDocument(sender objc.IObject)
	RunPageLayout(sender objc.IObject)
	RevertDocumentToSaved(sender objc.IObject)
	SaveDocument(sender objc.IObject)
	SaveDocumentAs(sender objc.IObject)
	SaveDocumentTo(sender objc.IObject)
	SaveDocumentWithDelegate_DidSaveSelector_ContextInfo(delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer)
	CanCloseDocumentWithDelegate_ShouldCloseSelector_ContextInfo(delegate objc.IObject, shouldCloseSelector objc.Selector, contextInfo unsafe.Pointer)
	Close()
	RevertToContentsOfURL_OfType_Error(url foundation.IURL, typeName string, outError *foundation.Error) bool
	DuplicateAndReturnError(outError *foundation.Error) Document
	DuplicateDocument(sender objc.IObject)
	DuplicateDocumentWithDelegate_DidDuplicateSelector_ContextInfo(delegate objc.IObject, didDuplicateSelector objc.Selector, contextInfo unsafe.Pointer)
	RenameDocument(sender objc.IObject)
	MoveDocument(sender objc.IObject)
	MoveDocumentWithCompletionHandler(completionHandler func(didMove bool))
	MoveToURL_CompletionHandler(url foundation.IURL, completionHandler func(param1 foundation.Error))
	LockDocument(sender objc.IObject)
	UnlockDocument(sender objc.IObject)
	LockDocumentWithCompletionHandler(completionHandler func(didLock bool))
	LockWithCompletionHandler(completionHandler func(param1 foundation.Error))
	UnlockDocumentWithCompletionHandler(completionHandler func(didUnlock bool))
	UnlockWithCompletionHandler(completionHandler func(param1 foundation.Error))
	PreparePageLayout(pageLayout IPageLayout) bool
	RunModalPageLayoutWithPrintInfo_Delegate_DidRunSelector_ContextInfo(printInfo IPrintInfo, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer)
	RunModalPrintOperation_Delegate_DidRunSelector_ContextInfo(printOperation IPrintOperation, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer)
	ShouldChangePrintInfo(newPrintInfo IPrintInfo) bool
	PrintDocumentWithSettings_ShowPrintPanel_Delegate_DidPrintSelector_ContextInfo(printSettings map[PrintInfoAttributeKey]objc.IObject, showPrintPanel bool, delegate objc.IObject, didPrintSelector objc.Selector, contextInfo unsafe.Pointer)
	PrintOperationWithSettings_Error(printSettings map[PrintInfoAttributeKey]objc.IObject, outError *foundation.Error) PrintOperation
	SaveDocumentToPDF(sender objc.IObject)
	PrepareSharingServicePicker(sharingServicePicker ISharingServicePicker)
	ShareDocumentWithSharingService_CompletionHandler(sharingService ISharingService, completionHandler func(success bool))
	HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object
	HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object
	HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object
	PresentError_ModalForWindow_Delegate_DidPresentSelector_ContextInfo(error foundation.IError, window IWindow, delegate objc.IObject, didPresentSelector objc.Selector, contextInfo unsafe.Pointer)
	PresentError(error foundation.IError) bool
	WillPresentError(error foundation.IError) foundation.Error
	WillNotPresentError(error foundation.IError)
	// deprecated
	DataRepresentationOfType(_type string) []byte
	// deprecated
	FileName() string
	// deprecated
	FileWrapperRepresentationOfType(_type string) foundation.FileWrapper
	// deprecated
	InitWithContentsOfFile_OfType(absolutePath string, typeName string) objc.Object
	// deprecated
	InitWithContentsOfURL_OfType(url foundation.IURL, typeName string) objc.Object
	// deprecated
	LoadDataRepresentation_OfType(data []byte, _type string) bool
	// deprecated
	LoadFileWrapperRepresentation_OfType(wrapper foundation.IFileWrapper, _type string) bool
	// deprecated
	PrintShowingPrintPanel(flag bool)
	// deprecated
	ReadFromFile_OfType(fileName string, _type string) bool
	// deprecated
	ReadFromURL_OfType(url foundation.IURL, _type string) bool
	// deprecated
	RevertToSavedFromFile_OfType(fileName string, _type string) bool
	// deprecated
	RevertToSavedFromURL_OfType(url foundation.IURL, _type string) bool
	// deprecated
	RunModalPageLayoutWithPrintInfo(printInfo IPrintInfo) int
	// deprecated
	SaveToFile_SaveOperation_Delegate_DidSaveSelector_ContextInfo(fileName string, saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer)
	// deprecated
	SaveToURL_OfType_ForSaveOperation_Error(url foundation.IURL, typeName string, saveOperation SaveOperationType, outError *foundation.Error) bool
	// deprecated
	SetFileName(fileName string)
	// deprecated
	WriteToFile_OfType(fileName string, _type string) bool
	// deprecated
	WriteToFile_OfType_OriginalFile_SaveOperation(fullDocumentPath string, documentTypeName string, fullOriginalDocumentPath string, saveOperationType SaveOperationType) bool
	// deprecated
	WriteToURL_OfType(url foundation.IURL, _type string) bool
	// deprecated
	WriteWithBackupToFile_OfType_SaveOperation(fullDocumentPath string, documentTypeName string, saveOperationType SaveOperationType) bool
	FileURL() foundation.URL
	SetFileURL(value foundation.IURL)
	IsEntireFileLoaded() bool
	FileModificationDate() foundation.Date
	SetFileModificationDate(value foundation.IDate)
	KeepBackupFile() bool
	IsDraft() bool
	SetDraft(value bool)
	FileType() string
	SetFileType(value string)
	IsDocumentEdited() bool
	IsInViewingMode() bool
	WindowControllers() []WindowController
	WindowNibName() NibName
	WindowForSheet() Window
	DisplayName() string
	AutosavedContentsFileURL() foundation.URL
	SetAutosavedContentsFileURL(value foundation.IURL)
	AutosavingFileType() string
	AutosavingIsImplicitlyCancellable() bool
	HasUnautosavedChanges() bool
	BackupFileURL() foundation.URL
	IsBrowsingVersions() bool
	UndoManager() foundation.UndoManager
	SetUndoManager(value foundation.IUndoManager)
	HasUndoManager() bool
	SetHasUndoManager(value bool)
	ShouldRunSavePanelWithAccessoryView() bool
	FileTypeFromLastRunSavePanel() string
	FileNameExtensionWasHiddenInLastRunSavePanel() bool
	UserActivity() foundation.UserActivity
	SetUserActivity(value foundation.IUserActivity)
	IsLocked() bool
	PrintInfo() PrintInfo
	SetPrintInfo(value IPrintInfo)
	PDFPrintOperation() PrintOperation
	AllowsDocumentSharing() bool
	ObjectSpecifier() foundation.ScriptObjectSpecifier
	LastComponentOfFileName() string
	SetLastComponentOfFileName(value string)
}

type Document struct {
	objc.Object
}

func MakeDocument(ptr unsafe.Pointer) Document {
	return Document{
		Object: objc.MakeObject(ptr),
	}
}

func (d_ Document) Init() Document {
	rv := ffi.CallMethod[Document](d_, "init")
	rv.Autorelease()
	return rv
}

func (d_ Document) InitWithContentsOfURL_OfType_Error(url foundation.IURL, typeName string, outError *foundation.Error) Document {
	rv := ffi.CallMethod[Document](d_, "initWithContentsOfURL:ofType:error:", url, typeName, unsafe.Pointer(outError))
	rv.Autorelease()
	return rv
}

func (d_ Document) InitForURL_WithContentsOfURL_OfType_Error(urlOrNil foundation.IURL, contentsURL foundation.IURL, typeName string, outError *foundation.Error) Document {
	rv := ffi.CallMethod[Document](d_, "initForURL:withContentsOfURL:ofType:error:", urlOrNil, contentsURL, typeName, unsafe.Pointer(outError))
	rv.Autorelease()
	return rv
}

func (d_ Document) InitWithType_Error(typeName string, outError *foundation.Error) Document {
	rv := ffi.CallMethod[Document](d_, "initWithType:error:", typeName, unsafe.Pointer(outError))
	rv.Autorelease()
	return rv
}

func (dc _DocumentClass) Alloc() Document {
	rv := ffi.CallMethod[Document](dc, "alloc")
	return rv
}

func (dc _DocumentClass) New() Document {
	rv := ffi.CallMethod[Document](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDocument() Document {
	return DocumentClass.New()
}

func (dc _DocumentClass) CanConcurrentlyReadDocumentsOfType(typeName string) bool {
	rv := ffi.CallMethod[bool](dc, "canConcurrentlyReadDocumentsOfType:", typeName)
	return rv
}

func (d_ Document) ReadFromURL_OfType_Error(url foundation.IURL, typeName string, outError *foundation.Error) bool {
	rv := ffi.CallMethod[bool](d_, "readFromURL:ofType:error:", url, typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) ReadFromFileWrapper_OfType_Error(fileWrapper foundation.IFileWrapper, typeName string, outError *foundation.Error) bool {
	rv := ffi.CallMethod[bool](d_, "readFromFileWrapper:ofType:error:", fileWrapper, typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) ReadFromData_OfType_Error(data []byte, typeName string, outError *foundation.Error) bool {
	rv := ffi.CallMethod[bool](d_, "readFromData:ofType:error:", data, typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) CanAsynchronouslyWriteToURL_OfType_ForSaveOperation(url foundation.IURL, typeName string, saveOperation SaveOperationType) bool {
	rv := ffi.CallMethod[bool](d_, "canAsynchronouslyWriteToURL:ofType:forSaveOperation:", url, typeName, saveOperation)
	return rv
}

func (d_ Document) UnblockUserInteraction() {
	ffi.CallMethod[ffi.Void](d_, "unblockUserInteraction")
}

func (d_ Document) WriteToURL_OfType_Error(url foundation.IURL, typeName string, outError *foundation.Error) bool {
	rv := ffi.CallMethod[bool](d_, "writeToURL:ofType:error:", url, typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) WriteSafelyToURL_OfType_ForSaveOperation_Error(url foundation.IURL, typeName string, saveOperation SaveOperationType, outError *foundation.Error) bool {
	rv := ffi.CallMethod[bool](d_, "writeSafelyToURL:ofType:forSaveOperation:error:", url, typeName, saveOperation, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) FileWrapperOfType_Error(typeName string, outError *foundation.Error) foundation.FileWrapper {
	rv := ffi.CallMethod[foundation.FileWrapper](d_, "fileWrapperOfType:error:", typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) DataOfType_Error(typeName string, outError *foundation.Error) []byte {
	rv := ffi.CallMethod[[]byte](d_, "dataOfType:error:", typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) WriteToURL_OfType_ForSaveOperation_OriginalContentsURL_Error(url foundation.IURL, typeName string, saveOperation SaveOperationType, absoluteOriginalContentsURL foundation.IURL, outError *foundation.Error) bool {
	rv := ffi.CallMethod[bool](d_, "writeToURL:ofType:forSaveOperation:originalContentsURL:error:", url, typeName, saveOperation, absoluteOriginalContentsURL, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) SaveToURL_OfType_ForSaveOperation_Delegate_DidSaveSelector_ContextInfo(url foundation.IURL, typeName string, saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](d_, "saveToURL:ofType:forSaveOperation:delegate:didSaveSelector:contextInfo:", url, typeName, saveOperation, delegate, didSaveSelector, contextInfo)
}

func (d_ Document) SaveToURL_OfType_ForSaveOperation_CompletionHandler(url foundation.IURL, typeName string, saveOperation SaveOperationType, completionHandler func(errorOrNil foundation.Error)) {
	ffi.CallMethod[ffi.Void](d_, "saveToURL:ofType:forSaveOperation:completionHandler:", url, typeName, saveOperation, completionHandler)
}

func (d_ Document) FileAttributesToWriteToURL_OfType_ForSaveOperation_OriginalContentsURL_Error(url foundation.IURL, typeName string, saveOperation SaveOperationType, absoluteOriginalContentsURL foundation.IURL, outError *foundation.Error) map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](d_, "fileAttributesToWriteToURL:ofType:forSaveOperation:originalContentsURL:error:", url, typeName, saveOperation, absoluteOriginalContentsURL, unsafe.Pointer(outError))
	return rv
}

func (dc _DocumentClass) IsNativeType(_type string) bool {
	rv := ffi.CallMethod[bool](dc, "isNativeType:", _type)
	return rv
}

func (d_ Document) WritableTypesForSaveOperation(saveOperation SaveOperationType) []string {
	rv := ffi.CallMethod[[]string](d_, "writableTypesForSaveOperation:", saveOperation)
	return rv
}

func (d_ Document) FileNameExtensionForType_SaveOperation(typeName string, saveOperation SaveOperationType) string {
	rv := ffi.CallMethod[string](d_, "fileNameExtensionForType:saveOperation:", typeName, saveOperation)
	return rv
}

func (d_ Document) MakeWindowControllers() {
	ffi.CallMethod[ffi.Void](d_, "makeWindowControllers")
}

func (d_ Document) AddWindowController(windowController IWindowController) {
	ffi.CallMethod[ffi.Void](d_, "addWindowController:", windowController)
}

func (d_ Document) RemoveWindowController(windowController IWindowController) {
	ffi.CallMethod[ffi.Void](d_, "removeWindowController:", windowController)
}

func (d_ Document) WindowControllerDidLoadNib(windowController IWindowController) {
	ffi.CallMethod[ffi.Void](d_, "windowControllerDidLoadNib:", windowController)
}

func (d_ Document) WindowControllerWillLoadNib(windowController IWindowController) {
	ffi.CallMethod[ffi.Void](d_, "windowControllerWillLoadNib:", windowController)
}

func (d_ Document) ShouldCloseWindowController_Delegate_ShouldCloseSelector_ContextInfo(windowController IWindowController, delegate objc.IObject, shouldCloseSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](d_, "shouldCloseWindowController:delegate:shouldCloseSelector:contextInfo:", windowController, delegate, shouldCloseSelector, contextInfo)
}

func (d_ Document) ShowWindows() {
	ffi.CallMethod[ffi.Void](d_, "showWindows")
}

func (d_ Document) SetWindow(window IWindow) {
	ffi.CallMethod[ffi.Void](d_, "setWindow:", window)
}

func (d_ Document) SetDisplayName(displayNameOrNil string) {
	ffi.CallMethod[ffi.Void](d_, "setDisplayName:", displayNameOrNil)
}

func (d_ Document) DefaultDraftName() string {
	rv := ffi.CallMethod[string](d_, "defaultDraftName")
	return rv
}

func (d_ Document) CheckAutosavingSafetyAndReturnError(outError *foundation.Error) bool {
	rv := ffi.CallMethod[bool](d_, "checkAutosavingSafetyAndReturnError:", unsafe.Pointer(outError))
	return rv
}

func (d_ Document) ScheduleAutosaving() {
	ffi.CallMethod[ffi.Void](d_, "scheduleAutosaving")
}

func (d_ Document) AutosaveDocumentWithDelegate_DidAutosaveSelector_ContextInfo(delegate objc.IObject, didAutosaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](d_, "autosaveDocumentWithDelegate:didAutosaveSelector:contextInfo:", delegate, didAutosaveSelector, contextInfo)
}

func (d_ Document) AutosaveWithImplicitCancellability_CompletionHandler(autosavingIsImplicitlyCancellable bool, completionHandler func(errorOrNil foundation.Error)) {
	ffi.CallMethod[ffi.Void](d_, "autosaveWithImplicitCancellability:completionHandler:", autosavingIsImplicitlyCancellable, completionHandler)
}

func (d_ Document) BrowseDocumentVersions(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "browseDocumentVersions:", sender)
}

func (d_ Document) StopBrowsingVersionsWithCompletionHandler(completionHandler func()) {
	ffi.CallMethod[ffi.Void](d_, "stopBrowsingVersionsWithCompletionHandler:", completionHandler)
}

func (d_ Document) MoveDocumentToUbiquityContainer(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "moveDocumentToUbiquityContainer:", sender)
}

func (d_ Document) UpdateChangeCountWithToken_ForSaveOperation(changeCountToken objc.IObject, saveOperation SaveOperationType) {
	ffi.CallMethod[ffi.Void](d_, "updateChangeCountWithToken:forSaveOperation:", changeCountToken, saveOperation)
}

func (d_ Document) UpdateChangeCount(change DocumentChangeType) {
	ffi.CallMethod[ffi.Void](d_, "updateChangeCount:", change)
}

func (d_ Document) ChangeCountTokenForSaveOperation(saveOperation SaveOperationType) objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "changeCountTokenForSaveOperation:", saveOperation)
	return rv
}

func (d_ Document) InvalidateRestorableState() {
	ffi.CallMethod[ffi.Void](d_, "invalidateRestorableState")
}

func (d_ Document) RunModalSavePanelForSaveOperation_Delegate_DidSaveSelector_ContextInfo(saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](d_, "runModalSavePanelForSaveOperation:delegate:didSaveSelector:contextInfo:", saveOperation, delegate, didSaveSelector, contextInfo)
}

func (d_ Document) PrepareSavePanel(savePanel ISavePanel) bool {
	rv := ffi.CallMethod[bool](d_, "prepareSavePanel:", savePanel)
	return rv
}

func (d_ Document) UpdateUserActivityState(activity foundation.IUserActivity) {
	ffi.CallMethod[ffi.Void](d_, "updateUserActivityState:", activity)
}

func (d_ Document) ValidateUserInterfaceItem(item ValidatedUserInterfaceItem) bool {
	po := ffi.CreateProtocol(item)
	defer po.Release()
	rv := ffi.CallMethod[bool](d_, "validateUserInterfaceItem:", po)
	return rv
}

func (d_ Document) PerformSynchronousFileAccessUsingBlock(block func()) {
	ffi.CallMethod[ffi.Void](d_, "performSynchronousFileAccessUsingBlock:", block)
}

func (d_ Document) ContinueActivityUsingBlock(block func()) {
	ffi.CallMethod[ffi.Void](d_, "continueActivityUsingBlock:", block)
}

func (d_ Document) ContinueAsynchronousWorkOnMainThreadUsingBlock(block func()) {
	ffi.CallMethod[ffi.Void](d_, "continueAsynchronousWorkOnMainThreadUsingBlock:", block)
}

func (d_ Document) PrintDocument(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "printDocument:", sender)
}

func (d_ Document) RunPageLayout(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "runPageLayout:", sender)
}

func (d_ Document) RevertDocumentToSaved(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "revertDocumentToSaved:", sender)
}

func (d_ Document) SaveDocument(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "saveDocument:", sender)
}

func (d_ Document) SaveDocumentAs(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "saveDocumentAs:", sender)
}

func (d_ Document) SaveDocumentTo(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "saveDocumentTo:", sender)
}

func (d_ Document) SaveDocumentWithDelegate_DidSaveSelector_ContextInfo(delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](d_, "saveDocumentWithDelegate:didSaveSelector:contextInfo:", delegate, didSaveSelector, contextInfo)
}

func (d_ Document) CanCloseDocumentWithDelegate_ShouldCloseSelector_ContextInfo(delegate objc.IObject, shouldCloseSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](d_, "canCloseDocumentWithDelegate:shouldCloseSelector:contextInfo:", delegate, shouldCloseSelector, contextInfo)
}

func (d_ Document) Close() {
	ffi.CallMethod[ffi.Void](d_, "close")
}

func (d_ Document) RevertToContentsOfURL_OfType_Error(url foundation.IURL, typeName string, outError *foundation.Error) bool {
	rv := ffi.CallMethod[bool](d_, "revertToContentsOfURL:ofType:error:", url, typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) DuplicateAndReturnError(outError *foundation.Error) Document {
	rv := ffi.CallMethod[Document](d_, "duplicateAndReturnError:", unsafe.Pointer(outError))
	return rv
}

func (d_ Document) DuplicateDocument(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "duplicateDocument:", sender)
}

func (d_ Document) DuplicateDocumentWithDelegate_DidDuplicateSelector_ContextInfo(delegate objc.IObject, didDuplicateSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](d_, "duplicateDocumentWithDelegate:didDuplicateSelector:contextInfo:", delegate, didDuplicateSelector, contextInfo)
}

func (d_ Document) RenameDocument(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "renameDocument:", sender)
}

func (d_ Document) MoveDocument(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "moveDocument:", sender)
}

func (d_ Document) MoveDocumentWithCompletionHandler(completionHandler func(didMove bool)) {
	ffi.CallMethod[ffi.Void](d_, "moveDocumentWithCompletionHandler:", completionHandler)
}

func (d_ Document) MoveToURL_CompletionHandler(url foundation.IURL, completionHandler func(param1 foundation.Error)) {
	ffi.CallMethod[ffi.Void](d_, "moveToURL:completionHandler:", url, completionHandler)
}

func (d_ Document) LockDocument(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "lockDocument:", sender)
}

func (d_ Document) UnlockDocument(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "unlockDocument:", sender)
}

func (d_ Document) LockDocumentWithCompletionHandler(completionHandler func(didLock bool)) {
	ffi.CallMethod[ffi.Void](d_, "lockDocumentWithCompletionHandler:", completionHandler)
}

func (d_ Document) LockWithCompletionHandler(completionHandler func(param1 foundation.Error)) {
	ffi.CallMethod[ffi.Void](d_, "lockWithCompletionHandler:", completionHandler)
}

func (d_ Document) UnlockDocumentWithCompletionHandler(completionHandler func(didUnlock bool)) {
	ffi.CallMethod[ffi.Void](d_, "unlockDocumentWithCompletionHandler:", completionHandler)
}

func (d_ Document) UnlockWithCompletionHandler(completionHandler func(param1 foundation.Error)) {
	ffi.CallMethod[ffi.Void](d_, "unlockWithCompletionHandler:", completionHandler)
}

func (d_ Document) PreparePageLayout(pageLayout IPageLayout) bool {
	rv := ffi.CallMethod[bool](d_, "preparePageLayout:", pageLayout)
	return rv
}

func (d_ Document) RunModalPageLayoutWithPrintInfo_Delegate_DidRunSelector_ContextInfo(printInfo IPrintInfo, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](d_, "runModalPageLayoutWithPrintInfo:delegate:didRunSelector:contextInfo:", printInfo, delegate, didRunSelector, contextInfo)
}

func (d_ Document) RunModalPrintOperation_Delegate_DidRunSelector_ContextInfo(printOperation IPrintOperation, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](d_, "runModalPrintOperation:delegate:didRunSelector:contextInfo:", printOperation, delegate, didRunSelector, contextInfo)
}

func (d_ Document) ShouldChangePrintInfo(newPrintInfo IPrintInfo) bool {
	rv := ffi.CallMethod[bool](d_, "shouldChangePrintInfo:", newPrintInfo)
	return rv
}

func (d_ Document) PrintDocumentWithSettings_ShowPrintPanel_Delegate_DidPrintSelector_ContextInfo(printSettings map[PrintInfoAttributeKey]objc.IObject, showPrintPanel bool, delegate objc.IObject, didPrintSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](d_, "printDocumentWithSettings:showPrintPanel:delegate:didPrintSelector:contextInfo:", printSettings, showPrintPanel, delegate, didPrintSelector, contextInfo)
}

func (d_ Document) PrintOperationWithSettings_Error(printSettings map[PrintInfoAttributeKey]objc.IObject, outError *foundation.Error) PrintOperation {
	rv := ffi.CallMethod[PrintOperation](d_, "printOperationWithSettings:error:", printSettings, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) SaveDocumentToPDF(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "saveDocumentToPDF:", sender)
}

func (d_ Document) PrepareSharingServicePicker(sharingServicePicker ISharingServicePicker) {
	ffi.CallMethod[ffi.Void](d_, "prepareSharingServicePicker:", sharingServicePicker)
}

func (d_ Document) ShareDocumentWithSharingService_CompletionHandler(sharingService ISharingService, completionHandler func(success bool)) {
	ffi.CallMethod[ffi.Void](d_, "shareDocumentWithSharingService:completionHandler:", sharingService, completionHandler)
}

func (d_ Document) HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "handleCloseScriptCommand:", command)
	return rv
}

func (d_ Document) HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "handlePrintScriptCommand:", command)
	return rv
}

func (d_ Document) HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "handleSaveScriptCommand:", command)
	return rv
}

func (d_ Document) PresentError_ModalForWindow_Delegate_DidPresentSelector_ContextInfo(error foundation.IError, window IWindow, delegate objc.IObject, didPresentSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](d_, "presentError:modalForWindow:delegate:didPresentSelector:contextInfo:", error, window, delegate, didPresentSelector, contextInfo)
}

func (d_ Document) PresentError(error foundation.IError) bool {
	rv := ffi.CallMethod[bool](d_, "presentError:", error)
	return rv
}

func (d_ Document) WillPresentError(error foundation.IError) foundation.Error {
	rv := ffi.CallMethod[foundation.Error](d_, "willPresentError:", error)
	return rv
}

func (d_ Document) WillNotPresentError(error foundation.IError) {
	ffi.CallMethod[ffi.Void](d_, "willNotPresentError:", error)
}

// deprecated
func (d_ Document) DataRepresentationOfType(_type string) []byte {
	rv := ffi.CallMethod[[]byte](d_, "dataRepresentationOfType:", _type)
	return rv
}

// deprecated
func (d_ Document) FileName() string {
	rv := ffi.CallMethod[string](d_, "fileName")
	return rv
}

// deprecated
func (d_ Document) FileWrapperRepresentationOfType(_type string) foundation.FileWrapper {
	rv := ffi.CallMethod[foundation.FileWrapper](d_, "fileWrapperRepresentationOfType:", _type)
	return rv
}

// deprecated
func (d_ Document) InitWithContentsOfFile_OfType(absolutePath string, typeName string) objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "initWithContentsOfFile:ofType:", absolutePath, typeName)
	rv.Autorelease()
	return rv
}

// deprecated
func (d_ Document) InitWithContentsOfURL_OfType(url foundation.IURL, typeName string) objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "initWithContentsOfURL:ofType:", url, typeName)
	rv.Autorelease()
	return rv
}

// deprecated
func (d_ Document) LoadDataRepresentation_OfType(data []byte, _type string) bool {
	rv := ffi.CallMethod[bool](d_, "loadDataRepresentation:ofType:", data, _type)
	return rv
}

// deprecated
func (d_ Document) LoadFileWrapperRepresentation_OfType(wrapper foundation.IFileWrapper, _type string) bool {
	rv := ffi.CallMethod[bool](d_, "loadFileWrapperRepresentation:ofType:", wrapper, _type)
	return rv
}

// deprecated
func (d_ Document) PrintShowingPrintPanel(flag bool) {
	ffi.CallMethod[ffi.Void](d_, "printShowingPrintPanel:", flag)
}

// deprecated
func (d_ Document) ReadFromFile_OfType(fileName string, _type string) bool {
	rv := ffi.CallMethod[bool](d_, "readFromFile:ofType:", fileName, _type)
	return rv
}

// deprecated
func (d_ Document) ReadFromURL_OfType(url foundation.IURL, _type string) bool {
	rv := ffi.CallMethod[bool](d_, "readFromURL:ofType:", url, _type)
	return rv
}

// deprecated
func (d_ Document) RevertToSavedFromFile_OfType(fileName string, _type string) bool {
	rv := ffi.CallMethod[bool](d_, "revertToSavedFromFile:ofType:", fileName, _type)
	return rv
}

// deprecated
func (d_ Document) RevertToSavedFromURL_OfType(url foundation.IURL, _type string) bool {
	rv := ffi.CallMethod[bool](d_, "revertToSavedFromURL:ofType:", url, _type)
	return rv
}

// deprecated
func (d_ Document) RunModalPageLayoutWithPrintInfo(printInfo IPrintInfo) int {
	rv := ffi.CallMethod[int](d_, "runModalPageLayoutWithPrintInfo:", printInfo)
	return rv
}

// deprecated
func (d_ Document) SaveToFile_SaveOperation_Delegate_DidSaveSelector_ContextInfo(fileName string, saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](d_, "saveToFile:saveOperation:delegate:didSaveSelector:contextInfo:", fileName, saveOperation, delegate, didSaveSelector, contextInfo)
}

// deprecated
func (d_ Document) SaveToURL_OfType_ForSaveOperation_Error(url foundation.IURL, typeName string, saveOperation SaveOperationType, outError *foundation.Error) bool {
	rv := ffi.CallMethod[bool](d_, "saveToURL:ofType:forSaveOperation:error:", url, typeName, saveOperation, unsafe.Pointer(outError))
	return rv
}

// deprecated
func (d_ Document) SetFileName(fileName string) {
	ffi.CallMethod[ffi.Void](d_, "setFileName:", fileName)
}

// deprecated
func (d_ Document) WriteToFile_OfType(fileName string, _type string) bool {
	rv := ffi.CallMethod[bool](d_, "writeToFile:ofType:", fileName, _type)
	return rv
}

// deprecated
func (d_ Document) WriteToFile_OfType_OriginalFile_SaveOperation(fullDocumentPath string, documentTypeName string, fullOriginalDocumentPath string, saveOperationType SaveOperationType) bool {
	rv := ffi.CallMethod[bool](d_, "writeToFile:ofType:originalFile:saveOperation:", fullDocumentPath, documentTypeName, fullOriginalDocumentPath, saveOperationType)
	return rv
}

// deprecated
func (d_ Document) WriteToURL_OfType(url foundation.IURL, _type string) bool {
	rv := ffi.CallMethod[bool](d_, "writeToURL:ofType:", url, _type)
	return rv
}

// deprecated
func (d_ Document) WriteWithBackupToFile_OfType_SaveOperation(fullDocumentPath string, documentTypeName string, saveOperationType SaveOperationType) bool {
	rv := ffi.CallMethod[bool](d_, "writeWithBackupToFile:ofType:saveOperation:", fullDocumentPath, documentTypeName, saveOperationType)
	return rv
}

func (d_ Document) FileURL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](d_, "fileURL")
	return rv
}

func (d_ Document) SetFileURL(value foundation.IURL) {
	ffi.CallMethod[ffi.Void](d_, "setFileURL:", value)
}

func (d_ Document) IsEntireFileLoaded() bool {
	rv := ffi.CallMethod[bool](d_, "isEntireFileLoaded")
	return rv
}

func (d_ Document) FileModificationDate() foundation.Date {
	rv := ffi.CallMethod[foundation.Date](d_, "fileModificationDate")
	return rv
}

func (d_ Document) SetFileModificationDate(value foundation.IDate) {
	ffi.CallMethod[ffi.Void](d_, "setFileModificationDate:", value)
}

func (d_ Document) KeepBackupFile() bool {
	rv := ffi.CallMethod[bool](d_, "keepBackupFile")
	return rv
}

func (d_ Document) IsDraft() bool {
	rv := ffi.CallMethod[bool](d_, "isDraft")
	return rv
}

func (d_ Document) SetDraft(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setDraft:", value)
}

func (d_ Document) FileType() string {
	rv := ffi.CallMethod[string](d_, "fileType")
	return rv
}

func (d_ Document) SetFileType(value string) {
	ffi.CallMethod[ffi.Void](d_, "setFileType:", value)
}

func (d_ Document) IsDocumentEdited() bool {
	rv := ffi.CallMethod[bool](d_, "isDocumentEdited")
	return rv
}

func (d_ Document) IsInViewingMode() bool {
	rv := ffi.CallMethod[bool](d_, "isInViewingMode")
	return rv
}

func (dc _DocumentClass) ReadableTypes() []string {
	rv := ffi.CallMethod[[]string](dc, "readableTypes")
	return rv
}

func (dc _DocumentClass) WritableTypes() []string {
	rv := ffi.CallMethod[[]string](dc, "writableTypes")
	return rv
}

func (d_ Document) WindowControllers() []WindowController {
	rv := ffi.CallMethod[[]WindowController](d_, "windowControllers")
	return rv
}

func (d_ Document) WindowNibName() NibName {
	rv := ffi.CallMethod[NibName](d_, "windowNibName")
	return rv
}

func (d_ Document) WindowForSheet() Window {
	rv := ffi.CallMethod[Window](d_, "windowForSheet")
	return rv
}

func (d_ Document) DisplayName() string {
	rv := ffi.CallMethod[string](d_, "displayName")
	return rv
}

func (dc _DocumentClass) AutosavesInPlace() bool {
	rv := ffi.CallMethod[bool](dc, "autosavesInPlace")
	return rv
}

func (dc _DocumentClass) AutosavesDrafts() bool {
	rv := ffi.CallMethod[bool](dc, "autosavesDrafts")
	return rv
}

func (dc _DocumentClass) PreservesVersions() bool {
	rv := ffi.CallMethod[bool](dc, "preservesVersions")
	return rv
}

func (d_ Document) AutosavedContentsFileURL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](d_, "autosavedContentsFileURL")
	return rv
}

func (d_ Document) SetAutosavedContentsFileURL(value foundation.IURL) {
	ffi.CallMethod[ffi.Void](d_, "setAutosavedContentsFileURL:", value)
}

func (d_ Document) AutosavingFileType() string {
	rv := ffi.CallMethod[string](d_, "autosavingFileType")
	return rv
}

func (d_ Document) AutosavingIsImplicitlyCancellable() bool {
	rv := ffi.CallMethod[bool](d_, "autosavingIsImplicitlyCancellable")
	return rv
}

func (d_ Document) HasUnautosavedChanges() bool {
	rv := ffi.CallMethod[bool](d_, "hasUnautosavedChanges")
	return rv
}

func (d_ Document) BackupFileURL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](d_, "backupFileURL")
	return rv
}

func (d_ Document) IsBrowsingVersions() bool {
	rv := ffi.CallMethod[bool](d_, "isBrowsingVersions")
	return rv
}

func (dc _DocumentClass) UsesUbiquitousStorage() bool {
	rv := ffi.CallMethod[bool](dc, "usesUbiquitousStorage")
	return rv
}

func (d_ Document) UndoManager() foundation.UndoManager {
	rv := ffi.CallMethod[foundation.UndoManager](d_, "undoManager")
	return rv
}

func (d_ Document) SetUndoManager(value foundation.IUndoManager) {
	ffi.CallMethod[ffi.Void](d_, "setUndoManager:", value)
}

func (d_ Document) HasUndoManager() bool {
	rv := ffi.CallMethod[bool](d_, "hasUndoManager")
	return rv
}

func (d_ Document) SetHasUndoManager(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setHasUndoManager:", value)
}

func (dc _DocumentClass) RestorableStateKeyPaths() []string {
	rv := ffi.CallMethod[[]string](dc, "restorableStateKeyPaths")
	return rv
}

func (d_ Document) ShouldRunSavePanelWithAccessoryView() bool {
	rv := ffi.CallMethod[bool](d_, "shouldRunSavePanelWithAccessoryView")
	return rv
}

func (d_ Document) FileTypeFromLastRunSavePanel() string {
	rv := ffi.CallMethod[string](d_, "fileTypeFromLastRunSavePanel")
	return rv
}

func (d_ Document) FileNameExtensionWasHiddenInLastRunSavePanel() bool {
	rv := ffi.CallMethod[bool](d_, "fileNameExtensionWasHiddenInLastRunSavePanel")
	return rv
}

func (d_ Document) UserActivity() foundation.UserActivity {
	rv := ffi.CallMethod[foundation.UserActivity](d_, "userActivity")
	return rv
}

func (d_ Document) SetUserActivity(value foundation.IUserActivity) {
	ffi.CallMethod[ffi.Void](d_, "setUserActivity:", value)
}

func (d_ Document) IsLocked() bool {
	rv := ffi.CallMethod[bool](d_, "isLocked")
	return rv
}

func (d_ Document) PrintInfo() PrintInfo {
	rv := ffi.CallMethod[PrintInfo](d_, "printInfo")
	return rv
}

func (d_ Document) SetPrintInfo(value IPrintInfo) {
	ffi.CallMethod[ffi.Void](d_, "setPrintInfo:", value)
}

func (d_ Document) PDFPrintOperation() PrintOperation {
	rv := ffi.CallMethod[PrintOperation](d_, "PDFPrintOperation")
	return rv
}

func (d_ Document) AllowsDocumentSharing() bool {
	rv := ffi.CallMethod[bool](d_, "allowsDocumentSharing")
	return rv
}

func (d_ Document) ObjectSpecifier() foundation.ScriptObjectSpecifier {
	rv := ffi.CallMethod[foundation.ScriptObjectSpecifier](d_, "objectSpecifier")
	return rv
}

func (d_ Document) LastComponentOfFileName() string {
	rv := ffi.CallMethod[string](d_, "lastComponentOfFileName")
	return rv
}

func (d_ Document) SetLastComponentOfFileName(value string) {
	ffi.CallMethod[ffi.Void](d_, "setLastComponentOfFileName:", value)
}

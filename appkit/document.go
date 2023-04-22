// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	EncodeRestorableStateWithCoder_BackgroundQueue(coder foundation.ICoder, queue foundation.IOperationQueue)
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
	EncodeRestorableStateWithCoder(coder foundation.ICoder)
	RestoreStateWithCoder(coder foundation.ICoder)
	InvalidateRestorableState()
	RestoreDocumentWindowWithIdentifier_State_CompletionHandler(identifier UserInterfaceItemIdentifier, state foundation.ICoder, completionHandler func(param1 Window, param2 foundation.Error))
	RunModalSavePanelForSaveOperation_Delegate_DidSaveSelector_ContextInfo(saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer)
	PrepareSavePanel(savePanel ISavePanel) bool
	UpdateUserActivityState(activity foundation.IUserActivity)
	ValidateUserInterfaceItem(item ValidatedUserInterfaceItem) bool
	ValidateUserInterfaceItem0(item objc.IObject) bool
	PerformSynchronousFileAccessUsingBlock(block func())
	PerformAsynchronousFileAccessUsingBlock(block func(param1 func()))
	PerformActivityWithSynchronousWaiting_UsingBlock(waitSynchronously bool, block func(param1 func()))
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
	DataRepresentationOfType(type_ string) []byte
	// deprecated
	FileName() string
	// deprecated
	FileWrapperRepresentationOfType(type_ string) foundation.FileWrapper
	// deprecated
	InitWithContentsOfFile_OfType(absolutePath string, typeName string) objc.Object
	// deprecated
	InitWithContentsOfURL_OfType(url foundation.IURL, typeName string) objc.Object
	// deprecated
	LoadDataRepresentation_OfType(data []byte, type_ string) bool
	// deprecated
	LoadFileWrapperRepresentation_OfType(wrapper foundation.IFileWrapper, type_ string) bool
	// deprecated
	PrintShowingPrintPanel(flag bool)
	// deprecated
	ReadFromFile_OfType(fileName string, type_ string) bool
	// deprecated
	ReadFromURL_OfType(url foundation.IURL, type_ string) bool
	// deprecated
	RevertToSavedFromFile_OfType(fileName string, type_ string) bool
	// deprecated
	RevertToSavedFromURL_OfType(url foundation.IURL, type_ string) bool
	// deprecated
	RunModalPageLayoutWithPrintInfo(printInfo IPrintInfo) int
	// deprecated
	SaveToFile_SaveOperation_Delegate_DidSaveSelector_ContextInfo(fileName string, saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer)
	// deprecated
	SaveToURL_OfType_ForSaveOperation_Error(url foundation.IURL, typeName string, saveOperation SaveOperationType, outError *foundation.Error) bool
	// deprecated
	SetFileName(fileName string)
	// deprecated
	WriteToFile_OfType(fileName string, type_ string) bool
	// deprecated
	WriteToFile_OfType_OriginalFile_SaveOperation(fullDocumentPath string, documentTypeName string, fullOriginalDocumentPath string, saveOperationType SaveOperationType) bool
	// deprecated
	WriteToURL_OfType(url foundation.IURL, type_ string) bool
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
	rv := objc.CallMethod[Document](d_, objc.GetSelector("init"))
	return rv
}

func (d_ Document) InitWithContentsOfURL_OfType_Error(url foundation.IURL, typeName string, outError *foundation.Error) Document {
	rv := objc.CallMethod[Document](d_, objc.GetSelector("initWithContentsOfURL:ofType:error:"), url, typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) InitForURL_WithContentsOfURL_OfType_Error(urlOrNil foundation.IURL, contentsURL foundation.IURL, typeName string, outError *foundation.Error) Document {
	rv := objc.CallMethod[Document](d_, objc.GetSelector("initForURL:withContentsOfURL:ofType:error:"), urlOrNil, contentsURL, typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) InitWithType_Error(typeName string, outError *foundation.Error) Document {
	rv := objc.CallMethod[Document](d_, objc.GetSelector("initWithType:error:"), typeName, unsafe.Pointer(outError))
	return rv
}

func (dc _DocumentClass) Alloc() Document {
	rv := objc.CallMethod[Document](dc, objc.GetSelector("alloc"))
	return rv
}

func (dc _DocumentClass) New() Document {
	rv := objc.CallMethod[Document](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDocument() Document {
	return DocumentClass.New()
}

func (dc _DocumentClass) CanConcurrentlyReadDocumentsOfType(typeName string) bool {
	rv := objc.CallMethod[bool](dc, objc.GetSelector("canConcurrentlyReadDocumentsOfType:"), typeName)
	return rv
}

func (d_ Document) ReadFromURL_OfType_Error(url foundation.IURL, typeName string, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("readFromURL:ofType:error:"), url, typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) ReadFromFileWrapper_OfType_Error(fileWrapper foundation.IFileWrapper, typeName string, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("readFromFileWrapper:ofType:error:"), fileWrapper, typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) ReadFromData_OfType_Error(data []byte, typeName string, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("readFromData:ofType:error:"), data, typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) CanAsynchronouslyWriteToURL_OfType_ForSaveOperation(url foundation.IURL, typeName string, saveOperation SaveOperationType) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("canAsynchronouslyWriteToURL:ofType:forSaveOperation:"), url, typeName, saveOperation)
	return rv
}

func (d_ Document) UnblockUserInteraction() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("unblockUserInteraction"))
}

func (d_ Document) WriteToURL_OfType_Error(url foundation.IURL, typeName string, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("writeToURL:ofType:error:"), url, typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) WriteSafelyToURL_OfType_ForSaveOperation_Error(url foundation.IURL, typeName string, saveOperation SaveOperationType, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("writeSafelyToURL:ofType:forSaveOperation:error:"), url, typeName, saveOperation, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) FileWrapperOfType_Error(typeName string, outError *foundation.Error) foundation.FileWrapper {
	rv := objc.CallMethod[foundation.FileWrapper](d_, objc.GetSelector("fileWrapperOfType:error:"), typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) DataOfType_Error(typeName string, outError *foundation.Error) []byte {
	rv := objc.CallMethod[[]byte](d_, objc.GetSelector("dataOfType:error:"), typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) WriteToURL_OfType_ForSaveOperation_OriginalContentsURL_Error(url foundation.IURL, typeName string, saveOperation SaveOperationType, absoluteOriginalContentsURL foundation.IURL, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("writeToURL:ofType:forSaveOperation:originalContentsURL:error:"), url, typeName, saveOperation, absoluteOriginalContentsURL, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) SaveToURL_OfType_ForSaveOperation_Delegate_DidSaveSelector_ContextInfo(url foundation.IURL, typeName string, saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveToURL:ofType:forSaveOperation:delegate:didSaveSelector:contextInfo:"), url, typeName, saveOperation, delegate, didSaveSelector, contextInfo)
}

func (d_ Document) SaveToURL_OfType_ForSaveOperation_CompletionHandler(url foundation.IURL, typeName string, saveOperation SaveOperationType, completionHandler func(errorOrNil foundation.Error)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveToURL:ofType:forSaveOperation:completionHandler:"), url, typeName, saveOperation, completionHandler)
}

func (d_ Document) FileAttributesToWriteToURL_OfType_ForSaveOperation_OriginalContentsURL_Error(url foundation.IURL, typeName string, saveOperation SaveOperationType, absoluteOriginalContentsURL foundation.IURL, outError *foundation.Error) map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](d_, objc.GetSelector("fileAttributesToWriteToURL:ofType:forSaveOperation:originalContentsURL:error:"), url, typeName, saveOperation, absoluteOriginalContentsURL, unsafe.Pointer(outError))
	return rv
}

func (dc _DocumentClass) IsNativeType(type_ string) bool {
	rv := objc.CallMethod[bool](dc, objc.GetSelector("isNativeType:"), type_)
	return rv
}

func (d_ Document) WritableTypesForSaveOperation(saveOperation SaveOperationType) []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("writableTypesForSaveOperation:"), saveOperation)
	return rv
}

func (d_ Document) FileNameExtensionForType_SaveOperation(typeName string, saveOperation SaveOperationType) string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("fileNameExtensionForType:saveOperation:"), typeName, saveOperation)
	return rv
}

func (d_ Document) MakeWindowControllers() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("makeWindowControllers"))
}

func (d_ Document) AddWindowController(windowController IWindowController) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("addWindowController:"), windowController)
}

func (d_ Document) RemoveWindowController(windowController IWindowController) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("removeWindowController:"), windowController)
}

func (d_ Document) WindowControllerDidLoadNib(windowController IWindowController) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("windowControllerDidLoadNib:"), windowController)
}

func (d_ Document) WindowControllerWillLoadNib(windowController IWindowController) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("windowControllerWillLoadNib:"), windowController)
}

func (d_ Document) ShouldCloseWindowController_Delegate_ShouldCloseSelector_ContextInfo(windowController IWindowController, delegate objc.IObject, shouldCloseSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("shouldCloseWindowController:delegate:shouldCloseSelector:contextInfo:"), windowController, delegate, shouldCloseSelector, contextInfo)
}

func (d_ Document) ShowWindows() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("showWindows"))
}

func (d_ Document) SetWindow(window IWindow) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setWindow:"), window)
}

func (d_ Document) SetDisplayName(displayNameOrNil string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDisplayName:"), displayNameOrNil)
}

func (d_ Document) DefaultDraftName() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("defaultDraftName"))
	return rv
}

func (d_ Document) EncodeRestorableStateWithCoder_BackgroundQueue(coder foundation.ICoder, queue foundation.IOperationQueue) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("encodeRestorableStateWithCoder:backgroundQueue:"), coder, queue)
}

func (d_ Document) CheckAutosavingSafetyAndReturnError(outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("checkAutosavingSafetyAndReturnError:"), unsafe.Pointer(outError))
	return rv
}

func (d_ Document) ScheduleAutosaving() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("scheduleAutosaving"))
}

func (d_ Document) AutosaveDocumentWithDelegate_DidAutosaveSelector_ContextInfo(delegate objc.IObject, didAutosaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("autosaveDocumentWithDelegate:didAutosaveSelector:contextInfo:"), delegate, didAutosaveSelector, contextInfo)
}

func (d_ Document) AutosaveWithImplicitCancellability_CompletionHandler(autosavingIsImplicitlyCancellable bool, completionHandler func(errorOrNil foundation.Error)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("autosaveWithImplicitCancellability:completionHandler:"), autosavingIsImplicitlyCancellable, completionHandler)
}

func (d_ Document) BrowseDocumentVersions(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("browseDocumentVersions:"), sender)
}

func (d_ Document) StopBrowsingVersionsWithCompletionHandler(completionHandler func()) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("stopBrowsingVersionsWithCompletionHandler:"), completionHandler)
}

func (d_ Document) MoveDocumentToUbiquityContainer(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("moveDocumentToUbiquityContainer:"), sender)
}

func (d_ Document) UpdateChangeCountWithToken_ForSaveOperation(changeCountToken objc.IObject, saveOperation SaveOperationType) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("updateChangeCountWithToken:forSaveOperation:"), changeCountToken, saveOperation)
}

func (d_ Document) UpdateChangeCount(change DocumentChangeType) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("updateChangeCount:"), change)
}

func (d_ Document) ChangeCountTokenForSaveOperation(saveOperation SaveOperationType) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("changeCountTokenForSaveOperation:"), saveOperation)
	return rv
}

func (dc _DocumentClass) AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	rv := objc.CallMethod[[]objc.Class](dc, objc.GetSelector("allowedClassesForRestorableStateKeyPath:"), keyPath)
	return rv
}

func (d_ Document) EncodeRestorableStateWithCoder(coder foundation.ICoder) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("encodeRestorableStateWithCoder:"), coder)
}

func (d_ Document) RestoreStateWithCoder(coder foundation.ICoder) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("restoreStateWithCoder:"), coder)
}

func (d_ Document) InvalidateRestorableState() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("invalidateRestorableState"))
}

func (d_ Document) RestoreDocumentWindowWithIdentifier_State_CompletionHandler(identifier UserInterfaceItemIdentifier, state foundation.ICoder, completionHandler func(param1 Window, param2 foundation.Error)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("restoreDocumentWindowWithIdentifier:state:completionHandler:"), identifier, state, completionHandler)
}

func (d_ Document) RunModalSavePanelForSaveOperation_Delegate_DidSaveSelector_ContextInfo(saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("runModalSavePanelForSaveOperation:delegate:didSaveSelector:contextInfo:"), saveOperation, delegate, didSaveSelector, contextInfo)
}

func (d_ Document) PrepareSavePanel(savePanel ISavePanel) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("prepareSavePanel:"), savePanel)
	return rv
}

func (d_ Document) UpdateUserActivityState(activity foundation.IUserActivity) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("updateUserActivityState:"), activity)
}

func (d_ Document) ValidateUserInterfaceItem(item ValidatedUserInterfaceItem) bool {
	po := objc.CreateProtocol("NSValidatedUserInterfaceItem", item)
	defer po.Release()
	rv := objc.CallMethod[bool](d_, objc.GetSelector("validateUserInterfaceItem:"), po)
	return rv
}

func (d_ Document) ValidateUserInterfaceItem0(item objc.IObject) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("validateUserInterfaceItem:"), item)
	return rv
}

func (d_ Document) PerformSynchronousFileAccessUsingBlock(block func()) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("performSynchronousFileAccessUsingBlock:"), block)
}

func (d_ Document) PerformAsynchronousFileAccessUsingBlock(block func(param1 func())) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("performAsynchronousFileAccessUsingBlock:"), block)
}

func (d_ Document) PerformActivityWithSynchronousWaiting_UsingBlock(waitSynchronously bool, block func(param1 func())) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("performActivityWithSynchronousWaiting:usingBlock:"), waitSynchronously, block)
}

func (d_ Document) ContinueActivityUsingBlock(block func()) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("continueActivityUsingBlock:"), block)
}

func (d_ Document) ContinueAsynchronousWorkOnMainThreadUsingBlock(block func()) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("continueAsynchronousWorkOnMainThreadUsingBlock:"), block)
}

func (d_ Document) PrintDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("printDocument:"), sender)
}

func (d_ Document) RunPageLayout(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("runPageLayout:"), sender)
}

func (d_ Document) RevertDocumentToSaved(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("revertDocumentToSaved:"), sender)
}

func (d_ Document) SaveDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveDocument:"), sender)
}

func (d_ Document) SaveDocumentAs(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveDocumentAs:"), sender)
}

func (d_ Document) SaveDocumentTo(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveDocumentTo:"), sender)
}

func (d_ Document) SaveDocumentWithDelegate_DidSaveSelector_ContextInfo(delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveDocumentWithDelegate:didSaveSelector:contextInfo:"), delegate, didSaveSelector, contextInfo)
}

func (d_ Document) CanCloseDocumentWithDelegate_ShouldCloseSelector_ContextInfo(delegate objc.IObject, shouldCloseSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("canCloseDocumentWithDelegate:shouldCloseSelector:contextInfo:"), delegate, shouldCloseSelector, contextInfo)
}

func (d_ Document) Close() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("close"))
}

func (d_ Document) RevertToContentsOfURL_OfType_Error(url foundation.IURL, typeName string, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("revertToContentsOfURL:ofType:error:"), url, typeName, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) DuplicateAndReturnError(outError *foundation.Error) Document {
	rv := objc.CallMethod[Document](d_, objc.GetSelector("duplicateAndReturnError:"), unsafe.Pointer(outError))
	return rv
}

func (d_ Document) DuplicateDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("duplicateDocument:"), sender)
}

func (d_ Document) DuplicateDocumentWithDelegate_DidDuplicateSelector_ContextInfo(delegate objc.IObject, didDuplicateSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("duplicateDocumentWithDelegate:didDuplicateSelector:contextInfo:"), delegate, didDuplicateSelector, contextInfo)
}

func (d_ Document) RenameDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("renameDocument:"), sender)
}

func (d_ Document) MoveDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("moveDocument:"), sender)
}

func (d_ Document) MoveDocumentWithCompletionHandler(completionHandler func(didMove bool)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("moveDocumentWithCompletionHandler:"), completionHandler)
}

func (d_ Document) MoveToURL_CompletionHandler(url foundation.IURL, completionHandler func(param1 foundation.Error)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("moveToURL:completionHandler:"), url, completionHandler)
}

func (d_ Document) LockDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("lockDocument:"), sender)
}

func (d_ Document) UnlockDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("unlockDocument:"), sender)
}

func (d_ Document) LockDocumentWithCompletionHandler(completionHandler func(didLock bool)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("lockDocumentWithCompletionHandler:"), completionHandler)
}

func (d_ Document) LockWithCompletionHandler(completionHandler func(param1 foundation.Error)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("lockWithCompletionHandler:"), completionHandler)
}

func (d_ Document) UnlockDocumentWithCompletionHandler(completionHandler func(didUnlock bool)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("unlockDocumentWithCompletionHandler:"), completionHandler)
}

func (d_ Document) UnlockWithCompletionHandler(completionHandler func(param1 foundation.Error)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("unlockWithCompletionHandler:"), completionHandler)
}

func (d_ Document) PreparePageLayout(pageLayout IPageLayout) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("preparePageLayout:"), pageLayout)
	return rv
}

func (d_ Document) RunModalPageLayoutWithPrintInfo_Delegate_DidRunSelector_ContextInfo(printInfo IPrintInfo, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("runModalPageLayoutWithPrintInfo:delegate:didRunSelector:contextInfo:"), printInfo, delegate, didRunSelector, contextInfo)
}

func (d_ Document) RunModalPrintOperation_Delegate_DidRunSelector_ContextInfo(printOperation IPrintOperation, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("runModalPrintOperation:delegate:didRunSelector:contextInfo:"), printOperation, delegate, didRunSelector, contextInfo)
}

func (d_ Document) ShouldChangePrintInfo(newPrintInfo IPrintInfo) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("shouldChangePrintInfo:"), newPrintInfo)
	return rv
}

func (d_ Document) PrintDocumentWithSettings_ShowPrintPanel_Delegate_DidPrintSelector_ContextInfo(printSettings map[PrintInfoAttributeKey]objc.IObject, showPrintPanel bool, delegate objc.IObject, didPrintSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("printDocumentWithSettings:showPrintPanel:delegate:didPrintSelector:contextInfo:"), printSettings, showPrintPanel, delegate, didPrintSelector, contextInfo)
}

func (d_ Document) PrintOperationWithSettings_Error(printSettings map[PrintInfoAttributeKey]objc.IObject, outError *foundation.Error) PrintOperation {
	rv := objc.CallMethod[PrintOperation](d_, objc.GetSelector("printOperationWithSettings:error:"), printSettings, unsafe.Pointer(outError))
	return rv
}

func (d_ Document) SaveDocumentToPDF(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveDocumentToPDF:"), sender)
}

func (d_ Document) PrepareSharingServicePicker(sharingServicePicker ISharingServicePicker) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("prepareSharingServicePicker:"), sharingServicePicker)
}

func (d_ Document) ShareDocumentWithSharingService_CompletionHandler(sharingService ISharingService, completionHandler func(success bool)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("shareDocumentWithSharingService:completionHandler:"), sharingService, completionHandler)
}

func (d_ Document) HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("handleCloseScriptCommand:"), command)
	return rv
}

func (d_ Document) HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("handlePrintScriptCommand:"), command)
	return rv
}

func (d_ Document) HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("handleSaveScriptCommand:"), command)
	return rv
}

func (d_ Document) PresentError_ModalForWindow_Delegate_DidPresentSelector_ContextInfo(error foundation.IError, window IWindow, delegate objc.IObject, didPresentSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), error, window, delegate, didPresentSelector, contextInfo)
}

func (d_ Document) PresentError(error foundation.IError) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("presentError:"), error)
	return rv
}

func (d_ Document) WillPresentError(error foundation.IError) foundation.Error {
	rv := objc.CallMethod[foundation.Error](d_, objc.GetSelector("willPresentError:"), error)
	return rv
}

func (d_ Document) WillNotPresentError(error foundation.IError) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("willNotPresentError:"), error)
}

// deprecated
func (d_ Document) DataRepresentationOfType(type_ string) []byte {
	rv := objc.CallMethod[[]byte](d_, objc.GetSelector("dataRepresentationOfType:"), type_)
	return rv
}

// deprecated
func (d_ Document) FileName() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("fileName"))
	return rv
}

// deprecated
func (d_ Document) FileWrapperRepresentationOfType(type_ string) foundation.FileWrapper {
	rv := objc.CallMethod[foundation.FileWrapper](d_, objc.GetSelector("fileWrapperRepresentationOfType:"), type_)
	return rv
}

// deprecated
func (d_ Document) InitWithContentsOfFile_OfType(absolutePath string, typeName string) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("initWithContentsOfFile:ofType:"), absolutePath, typeName)
	return rv
}

// deprecated
func (d_ Document) InitWithContentsOfURL_OfType(url foundation.IURL, typeName string) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("initWithContentsOfURL:ofType:"), url, typeName)
	return rv
}

// deprecated
func (d_ Document) LoadDataRepresentation_OfType(data []byte, type_ string) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("loadDataRepresentation:ofType:"), data, type_)
	return rv
}

// deprecated
func (d_ Document) LoadFileWrapperRepresentation_OfType(wrapper foundation.IFileWrapper, type_ string) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("loadFileWrapperRepresentation:ofType:"), wrapper, type_)
	return rv
}

// deprecated
func (d_ Document) PrintShowingPrintPanel(flag bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("printShowingPrintPanel:"), flag)
}

// deprecated
func (d_ Document) ReadFromFile_OfType(fileName string, type_ string) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("readFromFile:ofType:"), fileName, type_)
	return rv
}

// deprecated
func (d_ Document) ReadFromURL_OfType(url foundation.IURL, type_ string) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("readFromURL:ofType:"), url, type_)
	return rv
}

// deprecated
func (d_ Document) RevertToSavedFromFile_OfType(fileName string, type_ string) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("revertToSavedFromFile:ofType:"), fileName, type_)
	return rv
}

// deprecated
func (d_ Document) RevertToSavedFromURL_OfType(url foundation.IURL, type_ string) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("revertToSavedFromURL:ofType:"), url, type_)
	return rv
}

// deprecated
func (d_ Document) RunModalPageLayoutWithPrintInfo(printInfo IPrintInfo) int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("runModalPageLayoutWithPrintInfo:"), printInfo)
	return rv
}

// deprecated
func (d_ Document) SaveToFile_SaveOperation_Delegate_DidSaveSelector_ContextInfo(fileName string, saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("saveToFile:saveOperation:delegate:didSaveSelector:contextInfo:"), fileName, saveOperation, delegate, didSaveSelector, contextInfo)
}

// deprecated
func (d_ Document) SaveToURL_OfType_ForSaveOperation_Error(url foundation.IURL, typeName string, saveOperation SaveOperationType, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("saveToURL:ofType:forSaveOperation:error:"), url, typeName, saveOperation, unsafe.Pointer(outError))
	return rv
}

// deprecated
func (d_ Document) SetFileName(fileName string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setFileName:"), fileName)
}

// deprecated
func (d_ Document) WriteToFile_OfType(fileName string, type_ string) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("writeToFile:ofType:"), fileName, type_)
	return rv
}

// deprecated
func (d_ Document) WriteToFile_OfType_OriginalFile_SaveOperation(fullDocumentPath string, documentTypeName string, fullOriginalDocumentPath string, saveOperationType SaveOperationType) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("writeToFile:ofType:originalFile:saveOperation:"), fullDocumentPath, documentTypeName, fullOriginalDocumentPath, saveOperationType)
	return rv
}

// deprecated
func (d_ Document) WriteToURL_OfType(url foundation.IURL, type_ string) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("writeToURL:ofType:"), url, type_)
	return rv
}

// deprecated
func (d_ Document) WriteWithBackupToFile_OfType_SaveOperation(fullDocumentPath string, documentTypeName string, saveOperationType SaveOperationType) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("writeWithBackupToFile:ofType:saveOperation:"), fullDocumentPath, documentTypeName, saveOperationType)
	return rv
}

func (d_ Document) FileURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](d_, objc.GetSelector("fileURL"))
	return rv
}

func (d_ Document) SetFileURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setFileURL:"), value)
}

func (d_ Document) IsEntireFileLoaded() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isEntireFileLoaded"))
	return rv
}

func (d_ Document) FileModificationDate() foundation.Date {
	rv := objc.CallMethod[foundation.Date](d_, objc.GetSelector("fileModificationDate"))
	return rv
}

func (d_ Document) SetFileModificationDate(value foundation.IDate) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setFileModificationDate:"), value)
}

func (d_ Document) KeepBackupFile() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("keepBackupFile"))
	return rv
}

func (d_ Document) IsDraft() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isDraft"))
	return rv
}

func (d_ Document) SetDraft(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDraft:"), value)
}

func (d_ Document) FileType() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("fileType"))
	return rv
}

func (d_ Document) SetFileType(value string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setFileType:"), value)
}

func (d_ Document) IsDocumentEdited() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isDocumentEdited"))
	return rv
}

func (d_ Document) IsInViewingMode() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isInViewingMode"))
	return rv
}

func (dc _DocumentClass) ReadableTypes() []string {
	rv := objc.CallMethod[[]string](dc, objc.GetSelector("readableTypes"))
	return rv
}

func (dc _DocumentClass) WritableTypes() []string {
	rv := objc.CallMethod[[]string](dc, objc.GetSelector("writableTypes"))
	return rv
}

func (d_ Document) WindowControllers() []WindowController {
	rv := objc.CallMethod[[]WindowController](d_, objc.GetSelector("windowControllers"))
	return rv
}

func (d_ Document) WindowNibName() NibName {
	rv := objc.CallMethod[NibName](d_, objc.GetSelector("windowNibName"))
	return rv
}

func (d_ Document) WindowForSheet() Window {
	rv := objc.CallMethod[Window](d_, objc.GetSelector("windowForSheet"))
	return rv
}

func (d_ Document) DisplayName() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("displayName"))
	return rv
}

func (dc _DocumentClass) AutosavesInPlace() bool {
	rv := objc.CallMethod[bool](dc, objc.GetSelector("autosavesInPlace"))
	return rv
}

func (dc _DocumentClass) AutosavesDrafts() bool {
	rv := objc.CallMethod[bool](dc, objc.GetSelector("autosavesDrafts"))
	return rv
}

func (dc _DocumentClass) PreservesVersions() bool {
	rv := objc.CallMethod[bool](dc, objc.GetSelector("preservesVersions"))
	return rv
}

func (d_ Document) AutosavedContentsFileURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](d_, objc.GetSelector("autosavedContentsFileURL"))
	return rv
}

func (d_ Document) SetAutosavedContentsFileURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setAutosavedContentsFileURL:"), value)
}

func (d_ Document) AutosavingFileType() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("autosavingFileType"))
	return rv
}

func (d_ Document) AutosavingIsImplicitlyCancellable() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("autosavingIsImplicitlyCancellable"))
	return rv
}

func (d_ Document) HasUnautosavedChanges() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("hasUnautosavedChanges"))
	return rv
}

func (d_ Document) BackupFileURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](d_, objc.GetSelector("backupFileURL"))
	return rv
}

func (d_ Document) IsBrowsingVersions() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isBrowsingVersions"))
	return rv
}

func (dc _DocumentClass) UsesUbiquitousStorage() bool {
	rv := objc.CallMethod[bool](dc, objc.GetSelector("usesUbiquitousStorage"))
	return rv
}

func (d_ Document) UndoManager() foundation.UndoManager {
	rv := objc.CallMethod[foundation.UndoManager](d_, objc.GetSelector("undoManager"))
	return rv
}

func (d_ Document) SetUndoManager(value foundation.IUndoManager) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setUndoManager:"), value)
}

func (d_ Document) HasUndoManager() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("hasUndoManager"))
	return rv
}

func (d_ Document) SetHasUndoManager(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setHasUndoManager:"), value)
}

func (dc _DocumentClass) RestorableStateKeyPaths() []string {
	rv := objc.CallMethod[[]string](dc, objc.GetSelector("restorableStateKeyPaths"))
	return rv
}

func (d_ Document) ShouldRunSavePanelWithAccessoryView() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("shouldRunSavePanelWithAccessoryView"))
	return rv
}

func (d_ Document) FileTypeFromLastRunSavePanel() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("fileTypeFromLastRunSavePanel"))
	return rv
}

func (d_ Document) FileNameExtensionWasHiddenInLastRunSavePanel() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("fileNameExtensionWasHiddenInLastRunSavePanel"))
	return rv
}

func (d_ Document) UserActivity() foundation.UserActivity {
	rv := objc.CallMethod[foundation.UserActivity](d_, objc.GetSelector("userActivity"))
	return rv
}

func (d_ Document) SetUserActivity(value foundation.IUserActivity) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setUserActivity:"), value)
}

func (d_ Document) IsLocked() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isLocked"))
	return rv
}

func (d_ Document) PrintInfo() PrintInfo {
	rv := objc.CallMethod[PrintInfo](d_, objc.GetSelector("printInfo"))
	return rv
}

func (d_ Document) SetPrintInfo(value IPrintInfo) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setPrintInfo:"), value)
}

func (d_ Document) PDFPrintOperation() PrintOperation {
	rv := objc.CallMethod[PrintOperation](d_, objc.GetSelector("PDFPrintOperation"))
	return rv
}

func (d_ Document) AllowsDocumentSharing() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("allowsDocumentSharing"))
	return rv
}

func (d_ Document) ObjectSpecifier() foundation.ScriptObjectSpecifier {
	rv := objc.CallMethod[foundation.ScriptObjectSpecifier](d_, objc.GetSelector("objectSpecifier"))
	return rv
}

func (d_ Document) LastComponentOfFileName() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("lastComponentOfFileName"))
	return rv
}

func (d_ Document) SetLastComponentOfFileName(value string) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setLastComponentOfFileName:"), value)
}

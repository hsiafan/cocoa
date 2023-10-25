// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var FileWrapperClass = _FileWrapperClass{objc.GetClass("NSFileWrapper")}

type _FileWrapperClass struct {
	objc.Class
}

type IFileWrapper interface {
	objc.IObject
	// deprecated
	InitWithPath(path string) objc.Object
	// deprecated
	InitSymbolicLinkWithDestination(path string) objc.Object
	AddFileWrapper(child IFileWrapper) string
	RemoveFileWrapper(child IFileWrapper)
	// deprecated
	AddFileWithPath(path string) string
	AddRegularFileWithContents_PreferredFilename(data []byte, fileName string) string
	// deprecated
	AddSymbolicLinkWithDestination_PreferredFilename(path string, filename string) string
	KeyForFileWrapper(child IFileWrapper) string
	// deprecated
	SymbolicLinkDestination() string
	// deprecated
	NeedsToBeUpdatedFromPath(path string) bool
	MatchesContentsOfURL(url IURL) bool
	// deprecated
	UpdateFromPath(path string) bool
	ReadFromURL_Options_Error(url IURL, options FileWrapperReadingOptions, outError *Error) bool
	// deprecated
	WriteToFile_Atomically_UpdateFilenames(path string, atomicFlag bool, updateFilenamesFlag bool) bool
	WriteToURL_Options_OriginalContentsURL_Error(url IURL, options FileWrapperWritingOptions, originalContentsURL IURL, outError *Error) bool
	IsRegularFile() bool
	IsDirectory() bool
	IsSymbolicLink() bool
	FileWrappers() map[string]FileWrapper
	SymbolicLinkDestinationURL() URL
	SerializedRepresentation() []byte
	Filename() string
	SetFilename(value string)
	PreferredFilename() string
	SetPreferredFilename(value string)
	FileAttributes() map[string]objc.Object
	SetFileAttributes(value map[string]objc.IObject)
	RegularFileContents() []byte
}

type FileWrapper struct {
	objc.Object
}

func MakeFileWrapper(ptr unsafe.Pointer) FileWrapper {
	return FileWrapper{
		Object: objc.MakeObject(ptr),
	}
}

func (f_ FileWrapper) InitWithURL_Options_Error(url IURL, options FileWrapperReadingOptions, outError *Error) FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.SEL("initWithURL:options:error:"), objc.ExtractPtr(url), options, unsafe.Pointer(outError))
	return rv
}

func (f_ FileWrapper) InitDirectoryWithFileWrappers(childrenByPreferredName map[string]IFileWrapper) FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.SEL("initDirectoryWithFileWrappers:"), childrenByPreferredName)
	return rv
}

func (f_ FileWrapper) InitRegularFileWithContents(contents []byte) FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.SEL("initRegularFileWithContents:"), contents)
	return rv
}

func (f_ FileWrapper) InitSymbolicLinkWithDestinationURL(url IURL) FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.SEL("initSymbolicLinkWithDestinationURL:"), objc.ExtractPtr(url))
	return rv
}

func (f_ FileWrapper) InitWithSerializedRepresentation(serializeRepresentation []byte) FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.SEL("initWithSerializedRepresentation:"), serializeRepresentation)
	return rv
}

func (fc _FileWrapperClass) Alloc() FileWrapper {
	rv := objc.CallMethod[FileWrapper](fc, objc.SEL("alloc"))
	return rv
}

func (fc _FileWrapperClass) New() FileWrapper {
	rv := objc.CallMethod[FileWrapper](fc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewFileWrapper() FileWrapper {
	return FileWrapperClass.New()
}

func (f_ FileWrapper) Init() FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.SEL("init"))
	return rv
}

// deprecated
func (f_ FileWrapper) InitWithPath(path string) objc.Object {
	rv := objc.CallMethod[objc.Object](f_, objc.SEL("initWithPath:"), path)
	return rv
}

// deprecated
func (f_ FileWrapper) InitSymbolicLinkWithDestination(path string) objc.Object {
	rv := objc.CallMethod[objc.Object](f_, objc.SEL("initSymbolicLinkWithDestination:"), path)
	return rv
}

func (f_ FileWrapper) AddFileWrapper(child IFileWrapper) string {
	rv := objc.CallMethod[string](f_, objc.SEL("addFileWrapper:"), objc.ExtractPtr(child))
	return rv
}

func (f_ FileWrapper) RemoveFileWrapper(child IFileWrapper) {
	objc.CallMethod[objc.Void](f_, objc.SEL("removeFileWrapper:"), objc.ExtractPtr(child))
}

// deprecated
func (f_ FileWrapper) AddFileWithPath(path string) string {
	rv := objc.CallMethod[string](f_, objc.SEL("addFileWithPath:"), path)
	return rv
}

func (f_ FileWrapper) AddRegularFileWithContents_PreferredFilename(data []byte, fileName string) string {
	rv := objc.CallMethod[string](f_, objc.SEL("addRegularFileWithContents:preferredFilename:"), data, fileName)
	return rv
}

// deprecated
func (f_ FileWrapper) AddSymbolicLinkWithDestination_PreferredFilename(path string, filename string) string {
	rv := objc.CallMethod[string](f_, objc.SEL("addSymbolicLinkWithDestination:preferredFilename:"), path, filename)
	return rv
}

func (f_ FileWrapper) KeyForFileWrapper(child IFileWrapper) string {
	rv := objc.CallMethod[string](f_, objc.SEL("keyForFileWrapper:"), objc.ExtractPtr(child))
	return rv
}

// deprecated
func (f_ FileWrapper) SymbolicLinkDestination() string {
	rv := objc.CallMethod[string](f_, objc.SEL("symbolicLinkDestination"))
	return rv
}

// deprecated
func (f_ FileWrapper) NeedsToBeUpdatedFromPath(path string) bool {
	rv := objc.CallMethod[bool](f_, objc.SEL("needsToBeUpdatedFromPath:"), path)
	return rv
}

func (f_ FileWrapper) MatchesContentsOfURL(url IURL) bool {
	rv := objc.CallMethod[bool](f_, objc.SEL("matchesContentsOfURL:"), objc.ExtractPtr(url))
	return rv
}

// deprecated
func (f_ FileWrapper) UpdateFromPath(path string) bool {
	rv := objc.CallMethod[bool](f_, objc.SEL("updateFromPath:"), path)
	return rv
}

func (f_ FileWrapper) ReadFromURL_Options_Error(url IURL, options FileWrapperReadingOptions, outError *Error) bool {
	rv := objc.CallMethod[bool](f_, objc.SEL("readFromURL:options:error:"), objc.ExtractPtr(url), options, unsafe.Pointer(outError))
	return rv
}

// deprecated
func (f_ FileWrapper) WriteToFile_Atomically_UpdateFilenames(path string, atomicFlag bool, updateFilenamesFlag bool) bool {
	rv := objc.CallMethod[bool](f_, objc.SEL("writeToFile:atomically:updateFilenames:"), path, atomicFlag, updateFilenamesFlag)
	return rv
}

func (f_ FileWrapper) WriteToURL_Options_OriginalContentsURL_Error(url IURL, options FileWrapperWritingOptions, originalContentsURL IURL, outError *Error) bool {
	rv := objc.CallMethod[bool](f_, objc.SEL("writeToURL:options:originalContentsURL:error:"), objc.ExtractPtr(url), options, objc.ExtractPtr(originalContentsURL), unsafe.Pointer(outError))
	return rv
}

func (f_ FileWrapper) IsRegularFile() bool {
	rv := objc.CallMethod[bool](f_, objc.SEL("isRegularFile"))
	return rv
}

func (f_ FileWrapper) IsDirectory() bool {
	rv := objc.CallMethod[bool](f_, objc.SEL("isDirectory"))
	return rv
}

func (f_ FileWrapper) IsSymbolicLink() bool {
	rv := objc.CallMethod[bool](f_, objc.SEL("isSymbolicLink"))
	return rv
}

func (f_ FileWrapper) FileWrappers() map[string]FileWrapper {
	rv := objc.CallMethod[map[string]FileWrapper](f_, objc.SEL("fileWrappers"))
	return rv
}

func (f_ FileWrapper) SymbolicLinkDestinationURL() URL {
	rv := objc.CallMethod[URL](f_, objc.SEL("symbolicLinkDestinationURL"))
	return rv
}

func (f_ FileWrapper) SerializedRepresentation() []byte {
	rv := objc.CallMethod[[]byte](f_, objc.SEL("serializedRepresentation"))
	return rv
}

func (f_ FileWrapper) Filename() string {
	rv := objc.CallMethod[string](f_, objc.SEL("filename"))
	return rv
}

func (f_ FileWrapper) SetFilename(value string) {
	objc.CallMethod[objc.Void](f_, objc.SEL("setFilename:"), value)
}

func (f_ FileWrapper) PreferredFilename() string {
	rv := objc.CallMethod[string](f_, objc.SEL("preferredFilename"))
	return rv
}

func (f_ FileWrapper) SetPreferredFilename(value string) {
	objc.CallMethod[objc.Void](f_, objc.SEL("setPreferredFilename:"), value)
}

func (f_ FileWrapper) FileAttributes() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](f_, objc.SEL("fileAttributes"))
	return rv
}

func (f_ FileWrapper) SetFileAttributes(value map[string]objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.SEL("setFileAttributes:"), value)
}

func (f_ FileWrapper) RegularFileContents() []byte {
	rv := objc.CallMethod[[]byte](f_, objc.SEL("regularFileContents"))
	return rv
}

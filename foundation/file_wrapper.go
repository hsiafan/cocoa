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
	rv := objc.CallMethod[FileWrapper](f_, objc.GetSelector("initWithURL:options:error:"), url, options, unsafe.Pointer(outError))
	return rv
}

func (f_ FileWrapper) InitDirectoryWithFileWrappers(childrenByPreferredName map[string]IFileWrapper) FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.GetSelector("initDirectoryWithFileWrappers:"), childrenByPreferredName)
	return rv
}

func (f_ FileWrapper) InitRegularFileWithContents(contents []byte) FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.GetSelector("initRegularFileWithContents:"), contents)
	return rv
}

func (f_ FileWrapper) InitSymbolicLinkWithDestinationURL(url IURL) FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.GetSelector("initSymbolicLinkWithDestinationURL:"), url)
	return rv
}

func (f_ FileWrapper) InitWithSerializedRepresentation(serializeRepresentation []byte) FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.GetSelector("initWithSerializedRepresentation:"), serializeRepresentation)
	return rv
}

func (fc _FileWrapperClass) Alloc() FileWrapper {
	rv := objc.CallMethod[FileWrapper](fc, objc.GetSelector("alloc"))
	return rv
}

func (fc _FileWrapperClass) New() FileWrapper {
	rv := objc.CallMethod[FileWrapper](fc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewFileWrapper() FileWrapper {
	return FileWrapperClass.New()
}

func (f_ FileWrapper) Init() FileWrapper {
	rv := objc.CallMethod[FileWrapper](f_, objc.GetSelector("init"))
	return rv
}

// deprecated
func (f_ FileWrapper) InitWithPath(path string) objc.Object {
	rv := objc.CallMethod[objc.Object](f_, objc.GetSelector("initWithPath:"), path)
	return rv
}

// deprecated
func (f_ FileWrapper) InitSymbolicLinkWithDestination(path string) objc.Object {
	rv := objc.CallMethod[objc.Object](f_, objc.GetSelector("initSymbolicLinkWithDestination:"), path)
	return rv
}

func (f_ FileWrapper) AddFileWrapper(child IFileWrapper) string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("addFileWrapper:"), child)
	return rv
}

func (f_ FileWrapper) RemoveFileWrapper(child IFileWrapper) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("removeFileWrapper:"), child)
}

// deprecated
func (f_ FileWrapper) AddFileWithPath(path string) string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("addFileWithPath:"), path)
	return rv
}

func (f_ FileWrapper) AddRegularFileWithContents_PreferredFilename(data []byte, fileName string) string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("addRegularFileWithContents:preferredFilename:"), data, fileName)
	return rv
}

// deprecated
func (f_ FileWrapper) AddSymbolicLinkWithDestination_PreferredFilename(path string, filename string) string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("addSymbolicLinkWithDestination:preferredFilename:"), path, filename)
	return rv
}

func (f_ FileWrapper) KeyForFileWrapper(child IFileWrapper) string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("keyForFileWrapper:"), child)
	return rv
}

// deprecated
func (f_ FileWrapper) SymbolicLinkDestination() string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("symbolicLinkDestination"))
	return rv
}

// deprecated
func (f_ FileWrapper) NeedsToBeUpdatedFromPath(path string) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("needsToBeUpdatedFromPath:"), path)
	return rv
}

func (f_ FileWrapper) MatchesContentsOfURL(url IURL) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("matchesContentsOfURL:"), url)
	return rv
}

// deprecated
func (f_ FileWrapper) UpdateFromPath(path string) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("updateFromPath:"), path)
	return rv
}

func (f_ FileWrapper) ReadFromURL_Options_Error(url IURL, options FileWrapperReadingOptions, outError *Error) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("readFromURL:options:error:"), url, options, unsafe.Pointer(outError))
	return rv
}

// deprecated
func (f_ FileWrapper) WriteToFile_Atomically_UpdateFilenames(path string, atomicFlag bool, updateFilenamesFlag bool) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("writeToFile:atomically:updateFilenames:"), path, atomicFlag, updateFilenamesFlag)
	return rv
}

func (f_ FileWrapper) WriteToURL_Options_OriginalContentsURL_Error(url IURL, options FileWrapperWritingOptions, originalContentsURL IURL, outError *Error) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("writeToURL:options:originalContentsURL:error:"), url, options, originalContentsURL, unsafe.Pointer(outError))
	return rv
}

func (f_ FileWrapper) IsRegularFile() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isRegularFile"))
	return rv
}

func (f_ FileWrapper) IsDirectory() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isDirectory"))
	return rv
}

func (f_ FileWrapper) IsSymbolicLink() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isSymbolicLink"))
	return rv
}

func (f_ FileWrapper) FileWrappers() map[string]FileWrapper {
	rv := objc.CallMethod[map[string]FileWrapper](f_, objc.GetSelector("fileWrappers"))
	return rv
}

func (f_ FileWrapper) SymbolicLinkDestinationURL() URL {
	rv := objc.CallMethod[URL](f_, objc.GetSelector("symbolicLinkDestinationURL"))
	return rv
}

func (f_ FileWrapper) SerializedRepresentation() []byte {
	rv := objc.CallMethod[[]byte](f_, objc.GetSelector("serializedRepresentation"))
	return rv
}

func (f_ FileWrapper) Filename() string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("filename"))
	return rv
}

func (f_ FileWrapper) SetFilename(value string) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setFilename:"), value)
}

func (f_ FileWrapper) PreferredFilename() string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("preferredFilename"))
	return rv
}

func (f_ FileWrapper) SetPreferredFilename(value string) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setPreferredFilename:"), value)
}

func (f_ FileWrapper) FileAttributes() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](f_, objc.GetSelector("fileAttributes"))
	return rv
}

func (f_ FileWrapper) SetFileAttributes(value map[string]objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setFileAttributes:"), value)
}

func (f_ FileWrapper) RegularFileContents() []byte {
	rv := objc.CallMethod[[]byte](f_, objc.GetSelector("regularFileContents"))
	return rv
}

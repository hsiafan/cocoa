// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var PrinterClass = _PrinterClass{objc.GetClass("NSPrinter")}

type _PrinterClass struct {
	objc.Class
}

type IPrinter interface {
	objc.IObject
	PageSizeForPaper(paperName PrinterPaperName) foundation.Size
	// deprecated
	IsKey_InTable(key string, table string) bool
	// deprecated
	StringForKey_InTable(key string, table string) string
	// deprecated
	StringListForKey_InTable(key string, table string) []objc.Object
	// deprecated
	BooleanForKey_InTable(key string, table string) bool
	// deprecated
	FloatForKey_InTable(key string, table string) float32
	// deprecated
	IntForKey_InTable(key string, table string) int32
	// deprecated
	RectForKey_InTable(key string, table string) foundation.Rect
	// deprecated
	SizeForKey_InTable(key string, table string) foundation.Size
	// deprecated
	StatusForTable(tableName string) PrinterTableStatus
	// deprecated
	AcceptsBinary() bool
	// deprecated
	Domain() string
	// deprecated
	Host() string
	// deprecated
	ImageRectForPaper(paperName string) foundation.Rect
	// deprecated
	IsColor() bool
	// deprecated
	IsFontAvailable(faceName string) bool
	// deprecated
	IsOutputStackInReverseOrder() bool
	// deprecated
	Note() string
	Name() string
	Type() PrinterTypeName
	LanguageLevel() int
	DeviceDescription() map[DeviceDescriptionKey]objc.Object
}

type Printer struct {
	objc.Object
}

func MakePrinter(ptr unsafe.Pointer) Printer {
	return Printer{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PrinterClass) Alloc() Printer {
	rv := objc.CallMethod[Printer](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PrinterClass) New() Printer {
	rv := objc.CallMethod[Printer](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPrinter() Printer {
	return PrinterClass.New()
}

func (p_ Printer) Init() Printer {
	rv := objc.CallMethod[Printer](p_, objc.SEL("init"))
	return rv
}

func (pc _PrinterClass) PrinterWithName(name string) Printer {
	rv := objc.CallMethod[Printer](pc, objc.SEL("printerWithName:"), name)
	return rv
}

func (pc _PrinterClass) PrinterWithType(type_ PrinterTypeName) Printer {
	rv := objc.CallMethod[Printer](pc, objc.SEL("printerWithType:"), type_)
	return rv
}

func (p_ Printer) PageSizeForPaper(paperName PrinterPaperName) foundation.Size {
	rv := objc.CallMethod[foundation.Size](p_, objc.SEL("pageSizeForPaper:"), paperName)
	return rv
}

// deprecated
func (p_ Printer) IsKey_InTable(key string, table string) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isKey:inTable:"), key, table)
	return rv
}

// deprecated
func (p_ Printer) StringForKey_InTable(key string, table string) string {
	rv := objc.CallMethod[string](p_, objc.SEL("stringForKey:inTable:"), key, table)
	return rv
}

// deprecated
func (p_ Printer) StringListForKey_InTable(key string, table string) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](p_, objc.SEL("stringListForKey:inTable:"), key, table)
	return rv
}

// deprecated
func (p_ Printer) BooleanForKey_InTable(key string, table string) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("booleanForKey:inTable:"), key, table)
	return rv
}

// deprecated
func (p_ Printer) FloatForKey_InTable(key string, table string) float32 {
	rv := objc.CallMethod[float32](p_, objc.SEL("floatForKey:inTable:"), key, table)
	return rv
}

// deprecated
func (p_ Printer) IntForKey_InTable(key string, table string) int32 {
	rv := objc.CallMethod[int32](p_, objc.SEL("intForKey:inTable:"), key, table)
	return rv
}

// deprecated
func (p_ Printer) RectForKey_InTable(key string, table string) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](p_, objc.SEL("rectForKey:inTable:"), key, table)
	return rv
}

// deprecated
func (p_ Printer) SizeForKey_InTable(key string, table string) foundation.Size {
	rv := objc.CallMethod[foundation.Size](p_, objc.SEL("sizeForKey:inTable:"), key, table)
	return rv
}

// deprecated
func (p_ Printer) StatusForTable(tableName string) PrinterTableStatus {
	rv := objc.CallMethod[PrinterTableStatus](p_, objc.SEL("statusForTable:"), tableName)
	return rv
}

// deprecated
func (pc _PrinterClass) PrinterWithName_Domain_IncludeUnavailable(name string, domain string, flag bool) Printer {
	rv := objc.CallMethod[Printer](pc, objc.SEL("printerWithName:domain:includeUnavailable:"), name, domain, flag)
	return rv
}

// deprecated
func (p_ Printer) AcceptsBinary() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("acceptsBinary"))
	return rv
}

// deprecated
func (p_ Printer) Domain() string {
	rv := objc.CallMethod[string](p_, objc.SEL("domain"))
	return rv
}

// deprecated
func (p_ Printer) Host() string {
	rv := objc.CallMethod[string](p_, objc.SEL("host"))
	return rv
}

// deprecated
func (p_ Printer) ImageRectForPaper(paperName string) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](p_, objc.SEL("imageRectForPaper:"), paperName)
	return rv
}

// deprecated
func (p_ Printer) IsColor() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isColor"))
	return rv
}

// deprecated
func (p_ Printer) IsFontAvailable(faceName string) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isFontAvailable:"), faceName)
	return rv
}

// deprecated
func (p_ Printer) IsOutputStackInReverseOrder() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isOutputStackInReverseOrder"))
	return rv
}

// deprecated
func (p_ Printer) Note() string {
	rv := objc.CallMethod[string](p_, objc.SEL("note"))
	return rv
}

func (pc _PrinterClass) PrinterNames() []string {
	rv := objc.CallMethod[[]string](pc, objc.SEL("printerNames"))
	return rv
}

func (pc _PrinterClass) PrinterTypes() []PrinterTypeName {
	rv := objc.CallMethod[[]PrinterTypeName](pc, objc.SEL("printerTypes"))
	return rv
}

func (p_ Printer) Name() string {
	rv := objc.CallMethod[string](p_, objc.SEL("name"))
	return rv
}

func (p_ Printer) Type() PrinterTypeName {
	rv := objc.CallMethod[PrinterTypeName](p_, objc.SEL("type"))
	return rv
}

func (p_ Printer) LanguageLevel() int {
	rv := objc.CallMethod[int](p_, objc.SEL("languageLevel"))
	return rv
}

func (p_ Printer) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	rv := objc.CallMethod[map[DeviceDescriptionKey]objc.Object](p_, objc.SEL("deviceDescription"))
	return rv
}

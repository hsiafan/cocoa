// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var CharacterSetClass = _CharacterSetClass{objc.GetClass("NSCharacterSet")}

type _CharacterSetClass struct {
	objc.Class
}

type ICharacterSet interface {
	objc.IObject
	CharacterIsMember(aCharacter unichar) bool
	IsSupersetOfSet(theOtherSet ICharacterSet) bool
	LongCharacterIsMember(theLongChar uint32) bool
	BitmapRepresentation() []byte
	InvertedSet() CharacterSet
}

type CharacterSet struct {
	objc.Object
}

func MakeCharacterSet(ptr unsafe.Pointer) CharacterSet {
	return CharacterSet{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CharacterSetClass) Alloc() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CharacterSetClass) New() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCharacterSet() CharacterSet {
	return CharacterSetClass.New()
}

func (c_ CharacterSet) Init() CharacterSet {
	rv := objc.CallMethod[CharacterSet](c_, objc.SEL("init"))
	return rv
}

func (cc _CharacterSetClass) CharacterSetWithCharactersInString(aString string) CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("characterSetWithCharactersInString:"), aString)
	return rv
}

func (cc _CharacterSetClass) CharacterSetWithRange(aRange Range) CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("characterSetWithRange:"), aRange)
	return rv
}

func (cc _CharacterSetClass) CharacterSetWithBitmapRepresentation(data []byte) CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("characterSetWithBitmapRepresentation:"), data)
	return rv
}

func (cc _CharacterSetClass) CharacterSetWithContentsOfFile(fName string) CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("characterSetWithContentsOfFile:"), fName)
	return rv
}

func (c_ CharacterSet) CharacterIsMember(aCharacter unichar) bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("characterIsMember:"), aCharacter)
	return rv
}

func (c_ CharacterSet) IsSupersetOfSet(theOtherSet ICharacterSet) bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isSupersetOfSet:"), objc.ExtractPtr(theOtherSet))
	return rv
}

func (c_ CharacterSet) LongCharacterIsMember(theLongChar uint32) bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("longCharacterIsMember:"), theLongChar)
	return rv
}

func (cc _CharacterSetClass) AlphanumericCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("alphanumericCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) CapitalizedLetterCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("capitalizedLetterCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) ControlCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("controlCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) DecimalDigitCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("decimalDigitCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) DecomposableCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("decomposableCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) IllegalCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("illegalCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) LetterCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("letterCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) LowercaseLetterCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("lowercaseLetterCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) NewlineCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("newlineCharacterSet"))
	rv.Autorelease()
	return rv
}

func (cc _CharacterSetClass) NonBaseCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("nonBaseCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) PunctuationCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("punctuationCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) SymbolCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("symbolCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) UppercaseLetterCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("uppercaseLetterCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) WhitespaceAndNewlineCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("whitespaceAndNewlineCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) WhitespaceCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("whitespaceCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) URLFragmentAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("URLFragmentAllowedCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) URLHostAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("URLHostAllowedCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) URLPasswordAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("URLPasswordAllowedCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) URLPathAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("URLPathAllowedCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) URLQueryAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("URLQueryAllowedCharacterSet"))
	return rv
}

func (cc _CharacterSetClass) URLUserAllowedCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](cc, objc.SEL("URLUserAllowedCharacterSet"))
	return rv
}

func (c_ CharacterSet) BitmapRepresentation() []byte {
	rv := objc.CallMethod[[]byte](c_, objc.SEL("bitmapRepresentation"))
	return rv
}

func (c_ CharacterSet) InvertedSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](c_, objc.SEL("invertedSet"))
	return rv
}

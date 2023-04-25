// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var SoundClass = _SoundClass{objc.GetClass("NSSound")}

type _SoundClass struct {
	objc.Class
}

type ISound interface {
	objc.IObject
	SetName(string_ SoundName) bool
	Pause() bool
	Play() bool
	Resume() bool
	Stop() bool
	WriteToPasteboard(pasteboard IPasteboard)
	// deprecated
	ChannelMapping() []objc.Object
	// deprecated
	SetChannelMapping(channelMapping []objc.IObject)
	Delegate() SoundDelegateWrapper
	SetDelegate(value SoundDelegate)
	SetDelegate0(value objc.IObject)
	Name() SoundName
	Volume() float32
	SetVolume(value float32)
	CurrentTime() foundation.TimeInterval
	SetCurrentTime(value foundation.TimeInterval)
	Loops() bool
	SetLoops(value bool)
	PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier
	SetPlaybackDeviceIdentifier(value SoundPlaybackDeviceIdentifier)
	Duration() foundation.TimeInterval
	IsPlaying() bool
}

type Sound struct {
	objc.Object
}

func MakeSound(ptr unsafe.Pointer) Sound {
	return Sound{
		Object: objc.MakeObject(ptr),
	}
}

func (s_ Sound) InitWithContentsOfFile_ByReference(path string, byRef bool) Sound {
	rv := objc.CallMethod[Sound](s_, objc.GetSelector("initWithContentsOfFile:byReference:"), path, byRef)
	return rv
}

func (s_ Sound) InitWithContentsOfURL_ByReference(url foundation.IURL, byRef bool) Sound {
	rv := objc.CallMethod[Sound](s_, objc.GetSelector("initWithContentsOfURL:byReference:"), url, byRef)
	return rv
}

func (s_ Sound) InitWithData(data []byte) Sound {
	rv := objc.CallMethod[Sound](s_, objc.GetSelector("initWithData:"), data)
	return rv
}

func (s_ Sound) InitWithPasteboard(pasteboard IPasteboard) Sound {
	rv := objc.CallMethod[Sound](s_, objc.GetSelector("initWithPasteboard:"), pasteboard)
	return rv
}

func (sc _SoundClass) Alloc() Sound {
	rv := objc.CallMethod[Sound](sc, objc.GetSelector("alloc"))
	return rv
}

func (sc _SoundClass) New() Sound {
	rv := objc.CallMethod[Sound](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSound() Sound {
	return SoundClass.New()
}

func (s_ Sound) Init() Sound {
	rv := objc.CallMethod[Sound](s_, objc.GetSelector("init"))
	return rv
}

func (sc _SoundClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](sc, objc.GetSelector("canInitWithPasteboard:"), pasteboard)
	return rv
}

func (s_ Sound) SetName(string_ SoundName) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("setName:"), string_)
	return rv
}

func (sc _SoundClass) SoundNamed(name SoundName) Sound {
	rv := objc.CallMethod[Sound](sc, objc.GetSelector("soundNamed:"), name)
	return rv
}

func (s_ Sound) Pause() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("pause"))
	return rv
}

func (s_ Sound) Play() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("play"))
	return rv
}

func (s_ Sound) Resume() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("resume"))
	return rv
}

func (s_ Sound) Stop() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("stop"))
	return rv
}

func (s_ Sound) WriteToPasteboard(pasteboard IPasteboard) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("writeToPasteboard:"), pasteboard)
}

// deprecated
func (s_ Sound) ChannelMapping() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](s_, objc.GetSelector("channelMapping"))
	return rv
}

// deprecated
func (s_ Sound) SetChannelMapping(channelMapping []objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setChannelMapping:"), channelMapping)
}

// deprecated
func (sc _SoundClass) SoundUnfilteredFileTypes() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](sc, objc.GetSelector("soundUnfilteredFileTypes"))
	return rv
}

// deprecated
func (sc _SoundClass) SoundUnfilteredPasteboardTypes() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](sc, objc.GetSelector("soundUnfilteredPasteboardTypes"))
	return rv
}

func (s_ Sound) Delegate() SoundDelegateWrapper {
	rv := objc.CallMethod[SoundDelegateWrapper](s_, objc.GetSelector("delegate"))
	return rv
}

func (s_ Sound) SetDelegate(value SoundDelegate) {
	po := objc.CreateProtocol("NSSoundDelegate", value)
	objc.SetAssociatedObject(s_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), po)
}

func (s_ Sound) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), value)
}

func (s_ Sound) Name() SoundName {
	rv := objc.CallMethod[SoundName](s_, objc.GetSelector("name"))
	return rv
}

func (s_ Sound) Volume() float32 {
	rv := objc.CallMethod[float32](s_, objc.GetSelector("volume"))
	return rv
}

func (s_ Sound) SetVolume(value float32) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setVolume:"), value)
}

func (s_ Sound) CurrentTime() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](s_, objc.GetSelector("currentTime"))
	return rv
}

func (s_ Sound) SetCurrentTime(value foundation.TimeInterval) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setCurrentTime:"), value)
}

func (s_ Sound) Loops() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("loops"))
	return rv
}

func (s_ Sound) SetLoops(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setLoops:"), value)
}

func (s_ Sound) PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier {
	rv := objc.CallMethod[SoundPlaybackDeviceIdentifier](s_, objc.GetSelector("playbackDeviceIdentifier"))
	return rv
}

func (s_ Sound) SetPlaybackDeviceIdentifier(value SoundPlaybackDeviceIdentifier) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setPlaybackDeviceIdentifier:"), value)
}

func (sc _SoundClass) SoundUnfilteredTypes() []string {
	rv := objc.CallMethod[[]string](sc, objc.GetSelector("soundUnfilteredTypes"))
	return rv
}

func (s_ Sound) Duration() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](s_, objc.GetSelector("duration"))
	return rv
}

func (s_ Sound) IsPlaying() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isPlaying"))
	return rv
}

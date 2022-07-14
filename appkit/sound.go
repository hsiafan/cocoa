// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	SetName(_string SoundName) bool
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
	rv := ffi.CallMethod[Sound](s_, "initWithContentsOfFile:byReference:", path, byRef)
	rv.Autorelease()
	return rv
}

func (s_ Sound) InitWithContentsOfURL_ByReference(url foundation.IURL, byRef bool) Sound {
	rv := ffi.CallMethod[Sound](s_, "initWithContentsOfURL:byReference:", url, byRef)
	rv.Autorelease()
	return rv
}

func (s_ Sound) InitWithData(data []byte) Sound {
	rv := ffi.CallMethod[Sound](s_, "initWithData:", data)
	rv.Autorelease()
	return rv
}

func (s_ Sound) InitWithPasteboard(pasteboard IPasteboard) Sound {
	rv := ffi.CallMethod[Sound](s_, "initWithPasteboard:", pasteboard)
	rv.Autorelease()
	return rv
}

func (sc _SoundClass) Alloc() Sound {
	rv := ffi.CallMethod[Sound](sc, "alloc")
	return rv
}

func (s_ Sound) Init() Sound {
	rv := ffi.CallMethod[Sound](s_, "init")
	rv.Autorelease()
	return rv
}

func (sc _SoundClass) New() Sound {
	rv := ffi.CallMethod[Sound](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSound() Sound {
	return SoundClass.New()
}

func (sc _SoundClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](sc, "canInitWithPasteboard:", pasteboard)
	return rv
}

func (s_ Sound) SetName(_string SoundName) bool {
	rv := ffi.CallMethod[bool](s_, "setName:", _string)
	return rv
}

func (sc _SoundClass) SoundNamed(name SoundName) Sound {
	rv := ffi.CallMethod[Sound](sc, "soundNamed:", name)
	return rv
}

func (s_ Sound) Pause() bool {
	rv := ffi.CallMethod[bool](s_, "pause")
	return rv
}

func (s_ Sound) Play() bool {
	rv := ffi.CallMethod[bool](s_, "play")
	return rv
}

func (s_ Sound) Resume() bool {
	rv := ffi.CallMethod[bool](s_, "resume")
	return rv
}

func (s_ Sound) Stop() bool {
	rv := ffi.CallMethod[bool](s_, "stop")
	return rv
}

func (s_ Sound) WriteToPasteboard(pasteboard IPasteboard) {
	ffi.CallMethod[ffi.Void](s_, "writeToPasteboard:", pasteboard)
}

// deprecated
func (s_ Sound) ChannelMapping() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](s_, "channelMapping")
	return rv
}

// deprecated
func (s_ Sound) SetChannelMapping(channelMapping []objc.IObject) {
	ffi.CallMethod[ffi.Void](s_, "setChannelMapping:", channelMapping)
}

// deprecated
func (sc _SoundClass) SoundUnfilteredFileTypes() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](sc, "soundUnfilteredFileTypes")
	return rv
}

// deprecated
func (sc _SoundClass) SoundUnfilteredPasteboardTypes() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](sc, "soundUnfilteredPasteboardTypes")
	return rv
}

func (s_ Sound) Delegate() SoundDelegateWrapper {
	rv := ffi.CallMethod[SoundDelegateWrapper](s_, "delegate")
	return rv
}

func (s_ Sound) SetDelegate(value SoundDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(s_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](s_, "setDelegate:", po)
}

func (s_ Sound) Name() SoundName {
	rv := ffi.CallMethod[SoundName](s_, "name")
	return rv
}

func (s_ Sound) Volume() float32 {
	rv := ffi.CallMethod[float32](s_, "volume")
	return rv
}

func (s_ Sound) SetVolume(value float32) {
	ffi.CallMethod[ffi.Void](s_, "setVolume:", value)
}

func (s_ Sound) CurrentTime() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](s_, "currentTime")
	return rv
}

func (s_ Sound) SetCurrentTime(value foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](s_, "setCurrentTime:", value)
}

func (s_ Sound) Loops() bool {
	rv := ffi.CallMethod[bool](s_, "loops")
	return rv
}

func (s_ Sound) SetLoops(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setLoops:", value)
}

func (s_ Sound) PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier {
	rv := ffi.CallMethod[SoundPlaybackDeviceIdentifier](s_, "playbackDeviceIdentifier")
	return rv
}

func (s_ Sound) SetPlaybackDeviceIdentifier(value SoundPlaybackDeviceIdentifier) {
	ffi.CallMethod[ffi.Void](s_, "setPlaybackDeviceIdentifier:", value)
}

func (sc _SoundClass) SoundUnfilteredTypes() []string {
	rv := ffi.CallMethod[[]string](sc, "soundUnfilteredTypes")
	return rv
}

func (s_ Sound) Duration() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](s_, "duration")
	return rv
}

func (s_ Sound) IsPlaying() bool {
	rv := ffi.CallMethod[bool](s_, "isPlaying")
	return rv
}

type SoundDelegate interface {
	ImplementsSound_DidFinishPlaying() bool
	// optional
	Sound_DidFinishPlaying(sound Sound, flag bool)
}

type SoundDelegateImpl struct {
	_Sound_DidFinishPlaying func(sound Sound, flag bool)
}

func (di *SoundDelegateImpl) ImplementsSound_DidFinishPlaying() bool {
	return di._Sound_DidFinishPlaying != nil
}

func (di *SoundDelegateImpl) SetSound_DidFinishPlaying(f func(sound Sound, flag bool)) {
	di._Sound_DidFinishPlaying = f
}

func (di *SoundDelegateImpl) Sound_DidFinishPlaying(sound Sound, flag bool) {
	di._Sound_DidFinishPlaying(sound, flag)
}

type SoundDelegateWrapper struct {
	objc.Object
}

func (s_ *SoundDelegateWrapper) ImplementsSound_DidFinishPlaying() bool {
	return s_.RespondsToSelector(objc.GetSelector("sound:didFinishPlaying:"))
}

func (s_ SoundDelegateWrapper) Sound_DidFinishPlaying(sound ISound, flag bool) {
	ffi.CallMethod[ffi.Void](s_, "sound:didFinishPlaying:", sound, flag)
}

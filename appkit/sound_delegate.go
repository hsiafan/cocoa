// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

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
	objc.CallMethod[objc.Void](s_, objc.GetSelector("sound:didFinishPlaying:"), sound, flag)
}

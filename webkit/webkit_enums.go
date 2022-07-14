// auto-generated code, do not modify
package webkit

import "github.com/hsiafan/cocoa/foundation"

const WebViewDidBeginEditingNotification foundation.NotificationName = "WebViewDidBeginEditingNotification"
const WebViewDidChangeNotification foundation.NotificationName = "WebViewDidChangeNotification"
const WebViewDidChangeSelectionNotification foundation.NotificationName = "WebViewDidChangeSelectionNotification"
const WebViewDidChangeTypingStyleNotification foundation.NotificationName = "WebViewDidChangeTypingStyleNotification"
const WebViewDidEndEditingNotification foundation.NotificationName = "WebViewDidEndEditingNotification"

type AudiovisualMediaTypes uint

const AudiovisualMediaTypeNone AudiovisualMediaTypes = 0
const AudiovisualMediaTypeAudio AudiovisualMediaTypes = 1
const AudiovisualMediaTypeVideo AudiovisualMediaTypes = 2
const AudiovisualMediaTypeAll AudiovisualMediaTypes = 18446744073709551615

type ContentMode int

const ContentModeRecommended ContentMode = 0
const ContentModeDesktop ContentMode = 2
const ContentModeMobile ContentMode = 1

type FullscreenState int
type MediaCaptureState int
type MediaCaptureType int
type MediaPlaybackState int

const MediaPlaybackStateNone MediaPlaybackState = 0
const MediaPlaybackStatePaused MediaPlaybackState = 1
const MediaPlaybackStatePlaying MediaPlaybackState = 3
const MediaPlaybackStateSuspended MediaPlaybackState = 2

type NavigationActionPolicy int

const NavigationActionPolicyCancel NavigationActionPolicy = 0
const NavigationActionPolicyAllow NavigationActionPolicy = 1
const NavigationActionPolicyDownload NavigationActionPolicy = 2

type NavigationResponsePolicy int

const NavigationResponsePolicyCancel NavigationResponsePolicy = 0
const NavigationResponsePolicyAllow NavigationResponsePolicy = 1
const NavigationResponsePolicyDownload NavigationResponsePolicy = 2

type NavigationType int

const NavigationTypeLinkActivated NavigationType = 0
const NavigationTypeFormSubmitted NavigationType = 1
const NavigationTypeBackForward NavigationType = 2
const NavigationTypeReload NavigationType = 3
const NavigationTypeFormResubmitted NavigationType = 4
const NavigationTypeOther NavigationType = -1

type PermissionDecision int
type UserInterfaceDirectionPolicy int

const UserInterfaceDirectionPolicyContent UserInterfaceDirectionPolicy = 0
const UserInterfaceDirectionPolicySystem UserInterfaceDirectionPolicy = 1

type UserScriptInjectionTime int

const UserScriptInjectionTimeAtDocumentStart UserScriptInjectionTime = 0
const UserScriptInjectionTimeAtDocumentEnd UserScriptInjectionTime = 1

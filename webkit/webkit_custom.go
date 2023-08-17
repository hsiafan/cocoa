package webkit

import (
	"errors"
	"io"
	"io/fs"
	"strconv"
	"strings"

	"github.com/hsiafan/cocoa/dispatch"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

// LoadURL is convinent method for loading url into webview
func LoadURL(v IWebView, url string) {
	req := foundation.URLRequestClass.RequestWithURL(foundation.URLClass.URLWithString(url))
	v.LoadRequest(req)
}

// AddScriptMessageHandler is convinent method for adding ScriptMessageHandler to WebView ContentController
func AddScriptMessageHandler(v IWebView, name string, handler func(message objc.Object)) {
	h := &scriptMessageHandler{
		handler: handler,
	}
	v.Configuration().UserContentController().AddScriptMessageHandler_Name(WrapScriptMessageHandler(h), name)
}

// AddScriptMessageHandlerWithReply is convinent method for adding ScriptMessageHandlerWithReply to WebView ContentController
func AddScriptMessageHandlerWithReply(v IWebView, name string, handler func(message objc.Object) (objc.IObject, error)) {
	h := &scriptMessageHandlerWithReply{
		handler: handler,
	}
	v.Configuration().UserContentController().AddScriptMessageHandlerWithReply_ContentWorld_Name(
		WrapScriptMessageHandlerWithReply(h), ContentWorldClass.PageWorld(), name)
}

var _ ScriptMessageHandler = (*scriptMessageHandler)(nil)

type scriptMessageHandler struct {
	handler func(message objc.Object)
}

// UserContentController_DidReceiveScriptMessage implements ScriptMessageHandler
func (h *scriptMessageHandler) UserContentController_DidReceiveScriptMessage(
	userContentController UserContentController, message ScriptMessage) {
	message.Retain()
	go func() {
		defer func() {
			message.Release()
		}()
		h.handler(message.Body())
	}()

}

var _ ScriptMessageHandlerWithReply = (*scriptMessageHandlerWithReply)(nil)

type scriptMessageHandlerWithReply struct {
	handler func(message objc.Object) (objc.IObject, error)
}

// UserContentController_DidReceiveScriptMessage_ReplyHandler implements ScriptMessageHandlerWithReply
func (h *scriptMessageHandlerWithReply) UserContentController_DidReceiveScriptMessage_ReplyHandler(
	userContentController UserContentController, message ScriptMessage,
	replyHandler func(reply objc.IObject, errorMessage foundation.String)) {
	message.Retain()
	go func() {
		defer message.Release()
		reply, err := h.handler(message.Body())
		var errMsg = foundation.String{} // nil
		if err != nil {
			errMsg = foundation.NewString(err.Error())
		}
		if !reply.IsNil() {
			reply.Retain()
		}

		dispatch.GetMainQueue().DispatchAsync(func() {
			defer func() {
				if !reply.IsNil() {
					reply.Release()
				}
			}()
			replyHandler(reply, errMsg)
		})
	}()
}

var _ URLSchemeHandler = (*FileSystemURLSchemeHandler)(nil)

// FileSystemURLSchemeHandler is a WebView URLSchemeHandler supportting the go FileSystem.
type FileSystemURLSchemeHandler struct {
	FS fs.FS
}

// WebView_StartURLSchemeTask implements URLSchemeHandler
func (h *FileSystemURLSchemeHandler) WebView_StartURLSchemeTask(webView WebView, urlSchemeTask URLSchemeTaskWrapper) {
	url := urlSchemeTask.Request().URL()
	path := url.Path()
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}
	f, err := h.FS.Open(path)
	if err != nil {
		h.handleError(urlSchemeTask, err)
		return
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		h.handleError(urlSchemeTask, err)
		return
	}

	mime := h.getMime(path)
	foundation.NewURLResponse()
	response := foundation.HTTPURLResponseClass.Alloc().InitWithURL_MIMEType_ExpectedContentLength_TextEncodingName(
		url, mime, len(data), "utf-8",
	)

	urlSchemeTask.DidReceiveResponse(response)
	urlSchemeTask.DidReceiveData(data)
	urlSchemeTask.DidFinish()
}

func (h *FileSystemURLSchemeHandler) handleError(urlSchemeTask URLSchemeTaskWrapper, err error) {
	url := urlSchemeTask.Request().URL()
	var status = 500
	if errors.Is(err, fs.ErrNotExist) {
		status = 404
	}
	data := []byte(err.Error())
	response := foundation.HTTPURLResponseClass.Alloc().InitWithURL_StatusCode_HTTPVersion_HeaderFields(
		url, status, "", map[string]string{
			"Content-type":   "text/plain; charset=utf-8",
			"Content-length": strconv.Itoa(len(data)),
		})
	urlSchemeTask.DidReceiveResponse(response)
	urlSchemeTask.DidReceiveData(data)
	urlSchemeTask.DidFinish()
}

// WebView_StopURLSchemeTask implements URLSchemeHandler
func (h *FileSystemURLSchemeHandler) WebView_StopURLSchemeTask(webView WebView, urlSchemeTask URLSchemeTaskWrapper) {
}

func (h *FileSystemURLSchemeHandler) getMime(path string) string {
	var suffix string
	idx := strings.LastIndexByte(path, '.')
	if idx >= 0 {
		suffix = path[idx+1:]
	}
	switch strings.ToLower(suffix) {
	case "html", "htm":
		return "text/html"
	case "xhtml":
		return "application/xhtml+xml"
	case "css":
		return "text/css"
	case "txt":
		return "text/plain"
	case "jpeg", "jpg":
		return "image/jpeg"
	case "png":
		return "image/png"
	case "webp":
		return "image/webp"
	case "gif":
		return "image/gif"
	case "ico":
		return "image/vnd.microsoft.icon"
	case "bmp":
		return "image/bmp"
	case "js":
		return "text/javascript"
	case "json":
		return "application/json"
	case "xml":
		return "application/xml"
	case "pdf":
		return "application/pdf"
	case "ttf":
		return "font/ttf"
	case "otf":
		return "font/otf"
	case "woff":
		return "font/woff"
	case "woff2":
		return "font/woff2"
	case "mp3", "mpeg":
		return "audio/mpeg"
	case "oga":
		return "audio/ogg"
	case "weba":
		return "audio/webm"
	case "wav":
		return "audio/wav"
	case "mid", "midi":
		return "audio/midi"
	case "ogv":
		return "video/ogg"
	case "mp4":
		return "video/mp4"
	case "avi":
		return "video/x-msvideo"
	case "webm":
		return "video/webm"
	default:
		return "application/octet-stream"
	}
}

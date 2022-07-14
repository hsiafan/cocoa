package coface

import "github.com/hsiafan/cocoa/appkit"

// NewLabel create a text field, which looks like a Label
func NewLabel(title string) appkit.TextField {
	tf := appkit.TextFieldClass.LabelWithString(title)
	tf.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return tf
}

// NewTextField return a plain TextField
func NewTextField() appkit.TextField {
	tf := appkit.TextFieldClass.TextFieldWithString("")
	tf.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return tf
}

// NewSecureTextField return a plain SecureTextField
func NewSecureTextField() appkit.SecureTextField {
	tf := appkit.SecureTextFieldClass.TextFieldWithString("")
	tf.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return tf
}

// ScrollableTextView is appkit.ScrollView that contains a TextView
type ScrollableTextView struct {
	appkit.ScrollView
}

// ContentTextView return the inner TextView
func (t *ScrollableTextView) ContentTextView() appkit.TextView {
	return appkit.MakeTextView(t.DocumentView().Ptr())
}

// NewScrollableTextView create and return new scrollable text view.
func NewScrollableTextView() *ScrollableTextView {
	stv := appkit.TextViewClass.ScrollableTextView()
	stv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tv := appkit.MakeTextView(stv.DocumentView().Ptr())
	tv.SetAutomaticQuoteSubstitutionEnabled(false)
	tv.SetAutomaticDashSubstitutionEnabled(false)
	tv.SetAllowsUndo(true)
	return &ScrollableTextView{
		ScrollView: stv,
	}
}

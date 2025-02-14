package hb

import (
	"encoding/json"
	"strconv"
)

// SwalOptions represents configuration options for SweetAlert2 dialogs
// See https://sweetalert2.github.io/ for full documentation
type SwalOptions struct {
	Background         string `json:"background,omitempty"`
	Backdrop           string `json:"backdrop,omitempty"`
	CancelButtonColor  string `json:"cancelButtonColor,omitempty"`
	CancelButtonText   string `json:"cancelButtonText,omitempty"`
	Color              string `json:"color,omitempty"`
	ConfirmButtonText  string `json:"confirmButtonText,omitempty"`
	ConfirmButtonColor string `json:"confirmButtonColor,omitempty"`
	ConfirmCallback    string `json:"-"`
	DenyButtonText     string `json:"denyButtonText,omitempty"`
	CustomClass        string `json:"customClass,omitempty"`
	Grow               string `json:"grow,omitempty"`
	HeightAuto         bool   `json:"heightAuto,omitempty"`
	HTML               string `json:"html,omitempty"`
	Icon               string `json:"icon,omitempty"`
	IconColor          string `json:"iconColor,omitempty"`
	IconHtml           string `json:"iconHtml,omitempty"`
	ImageURL           string `json:"imageUrl,omitempty"`
	ImageWidth         string `json:"imageWidth,omitempty"`
	ImageHeight        string `json:"imageHeight,omitempty"`
	ImageAlt           string `json:"imageAlt,omitempty"`
	Padding            string `json:"padding,omitempty"`
	Position           string `json:"position,omitempty"`
	ShowCancelButton   bool   `json:"showCancelButton"`
	ShowConfirmButton  bool   `json:"showConfirmButton"`
	ShowDenyButton     bool   `json:"showDenyButton"`
	Text               string `json:"text,omitempty"`
	TimerProgressBar   bool   `json:"timerProgressBar,omitempty"`
	Title              string `json:"title,omitempty"`
	Timer              int    `json:"timer,omitempty"`
	Toast              bool   `json:"toast,omitempty"`
	Width              string `json:"width,omitempty"`

	// The following are not standard Sweetalert2 options, do not export to JSON
	RedirectURL     string `json:"-"`
	RedirectSeconds int    `json:"-"`
}

// swalToJS converts SwalOptions to JavaScript code for SweetAlert2
// Returns empty string if options cannot be marshaled to JSON
func swalToJS(options SwalOptions) string {
	// Validate required fields
	if options.Title == "" && options.Text == "" && options.HTML == "" {
		return ""
	}
	// By default, show the cancel button if text is specified
	if options.CancelButtonText != "" && !options.ShowCancelButton {
		options.ShowCancelButton = true
	}

	// By default, show the confirm button if text is specified
	if options.ConfirmButtonText != "" && !options.ShowConfirmButton {
		options.ShowConfirmButton = true
	}

	// By default, show the deny button if text is specified
	if options.DenyButtonText != "" && !options.ShowDenyButton {
		options.ShowDenyButton = true
	}

	optionsBytes, err := json.Marshal(options)

	var optionsJSON string

	if err != nil {
		optionsJSON = ""
	} else {
		optionsJSON = string(optionsBytes)
	}

	swal := `Swal.fire(` + optionsJSON + `)`

	if options.ConfirmCallback != "" {
		swal += `.then((result) => {
			if (result.isConfirmed) {
				` + options.ConfirmCallback + `
			}
		});`
	} else {
		swal += `;`
	}

	if options.RedirectURL != "" {
		swal += `setTimeout(() => {
			window.location.href = "` + options.RedirectURL + `";
		}, ` + strconv.Itoa(options.RedirectSeconds) + `000);`
	}

	return swal
}

// Swal generates a script with a SweetAlert2 dialog
// Note! you must include the library yourself (i.e. CDN)
func Swal(options SwalOptions) TagInterface {
	return NewSwal(options)
}

// SwalSuccess creates a SweetAlert2 success dialog
func SwalSuccess(options SwalOptions) TagInterface {
	return NewSwalSuccess(options)
}

// SwalError creates a SweetAlert2 error dialog
func SwalError(options SwalOptions) TagInterface {
	return NewSwalError(options)
}

// SwalWarning creates a SweetAlert2 warning dialog
func SwalWarning(options SwalOptions) TagInterface {
	return NewSwalWarning(options)
}

// SwalInfo creates a SweetAlert2 info dialog
func SwalInfo(options SwalOptions) TagInterface {
	return NewSwalInfo(options)
}

// NewSwal generates a script with a Sweetalert2 dialog
// Note! you must include the library yourself (i.e. CDN)
// Shortcut method exists: Swal()
func NewSwal(options SwalOptions) TagInterface {
	return NewScript(swalToJS(options))
}

// NewSwalSuccess creates a SweetAlert2 success dialog with the provided options
func NewSwalSuccess(options SwalOptions) TagInterface {
	options.Icon = "success"
	return NewScript(swalToJS(options))
}

// NewSwalError creates a SweetAlert2 error dialog with the provided options
func NewSwalError(options SwalOptions) TagInterface {
	options.Icon = "error"
	return NewScript(swalToJS(options))
}

// NewSwalWarning creates a SweetAlert2 warning dialog with the provided options
func NewSwalWarning(options SwalOptions) TagInterface {
	options.Icon = "warning"
	return NewScript(swalToJS(options))
}

// NewSwalInfo creates a SweetAlert2 info dialog with the provided options
func NewSwalInfo(options SwalOptions) TagInterface {
	options.Icon = "info"
	return NewScript(swalToJS(options))
}

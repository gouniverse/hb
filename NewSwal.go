package hb

import (
	"encoding/json"
)

type SwalOptions struct {
	Background         string `json:"background,omitempty"`
	Backdrop           string `json:"backdrop,omitempty"`
	CancelButtonColor  string `json:"cancelButtonColor,omitempty"`
	CancelButtonText   string `json:"cancelButtonText,omitempty"`
	Color              string `json:"color,omitempty"`
	ConfirmButtonText  string `json:"confirmButtonText,omitempty"`
	ConfirmButtonColor string `json:"confirmButtonColor,omitempty"`
	ConfirmCallback    string `json:"-"`
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
	ShowCancelButton   bool   `json:"showCancelButton,omitempty"`
	ShowConfirmButton  bool   `json:"showConfirmButton,omitempty"`
	Text               string `json:"text,omitempty"`
	TimerProgressBar   bool   `json:"timerProgressBar,omitempty"`
	Title              string `json:"title,omitempty"`
	Timer              int    `json:"timer,omitempty"`
	Toast              string `json:"toast,omitempty"`
	Width              string `json:"width,omitempty"`
}

// NewSwal generates a script with a Sweetalert2 dialog
// Note! you must include the library yourself (i.e. CDN)
// Deprecated: replaced by the new method Swal()
func NewSwal(options SwalOptions) *Tag {
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
	}

	return NewScript(swal)
}

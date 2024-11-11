package hb

import (
	"encoding/json"
)

type SwalOptions struct {
	Icon               string `json:"icon,omitempty"`
	Text               string `json:"text,omitempty"`
	Title              string `json:"title,omitempty"`
	Position           string `json:"position,omitempty"`
	Timer              int    `json:"timer,omitempty"`
	TimerProgressBar   bool    `json:"timerProgressBar,omitempty"`
	ShowCancelButton   bool   `json:"showCancelButton,omitempty"`
	ShowConfirmButton  bool   `json:"showConfirmButton,omitempty"`
	CancelButtonColor  string `json:"cancelButtonColor,omitempty"`
	CancelButtonText   string `json:"cancelButtonText,omitempty"`
	ConfirmButtonText  string `json:"confirmButtonText,omitempty"`
	ConfirmButtonColor string `json:"confirmButtonColor,omitempty"`
	ConfirmCallback    string `json:"-"`
	ImageURL	   string `json:"imageUrl,omitempty"`
	ImageWidth	   string `json:"imageWidth,omitempty"`
	ImageHeight	   string `json:"imageHeight,omitempty"`
	ImageAlt	   string `json:"imageAlt,omitempty"`
	Width		   string `json:"width,omitempty"`
	Padding		   string `json:"padding,omitempty"`
	Color		   string `json:"color,omitempty"`
	Background	   string `json:"background,omitempty"`
	Backdrop	   string `json:"backdrop,omitempty"`
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

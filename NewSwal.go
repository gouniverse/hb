package hb

import (
	"github.com/gouniverse/utils"
)

type SwalOptions struct {
	Icon               string `json:"icon,omitempty"`
	Text               string `json:"text,omitempty"`
	Title              string `json:"title,omitempty"`
	ShowCancelButton   bool   `json:"showCancelButton,omitempty"`
	CancelButtonColor  string `json:"cancelButtonColor,omitempty"`
	CancelButtonText   string `json:"cancelButtonText,omitempty"`
	ConfirmButtonText  string `json:"confirmButtonText,omitempty"`
	ConfirmButtonColor string `json:"confirmButtonColor,omitempty"`
	ConfirmCallback    string `json:"-"`
}

// NewSwal generates a script with a Sweetalert2 dialog
// Note! you must include the library yourself (i.e. CDN)
func NewSwal(options SwalOptions) *Tag {
	optionsJSON, _ := utils.ToJSON(options)
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

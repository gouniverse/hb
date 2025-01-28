package hb

import (
	"encoding/json"
	"strconv"
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

func swalToJS(options SwalOptions) string {
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

package hb

// ENCTYPE_FORM_URLENCODED ("application/x-www-form-urlencoded") default encoding of forms.
// All characters are converted. Spaces are converted to "+" symbols,
// and special characters are converted to ASCII HEX values)
const ENCTYPE_FORM_URLENCODED = "application/x-www-form-urlencoded"

// ENCTYPE_MULTIPART_FORM_DATA ("multipart/form-data")
// Necessary if the user will upload a file through the form
const ENCTYPE_FORM_MULTIPART = "multipart/form-data"

// ENCTYPE_FORM_TEXT ("text/plain") plain text
// An ambiguous format, human-readable content not reliably interpretable by computer
const ENCTYPE_FORM_TEXT = "text/plain"

const TYPE_BUTTON = "button"
const TYPE_CHECKBOX = "checkbox"
const TYPE_COLOR = "color"
const TYPE_DATE = "date"
const TYPE_DATETIME = "datetime-local"
const TYPE_EMAIL = "email"
const TYPE_FILE = "file"
const TYPE_HIDDEN = "hidden"
const TYPE_IMAGE = "image"
const TYPE_MONTH = "month"
const TYPE_NUMBER = "number"
const TYPE_PASSWORD = "password"
const TYPE_RADIO = "radio"
const TYPE_RANGE = "range"
const TYPE_RESET = "reset"
const TYPE_SEARCH = "search"
const TYPE_SUBMIT = "submit"
const TYPE_TEL = "tel"
const TYPE_TEXT = "text"
const TYPE_TIME = "time"
const TYPE_URL = "url"
const TYPE_WEEK = "week"

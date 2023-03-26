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

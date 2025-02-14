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

const DEFAULT_FAVICON = "data:image/x-icon;base64,AAABAAEAEBAAAAEAIABoBAAAFgAAACgAAAAQAAAAIAAAAAEAIAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACoqKgNTV1Vg3+De/jutb/4/uHb/QLJ0+FRUVFQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFRUVCw/uHb/i8qn/+zr6/7p6Of+2dzZ/j63df4/uHb/P7h2/1RUVCUAAAAAAAAAAAAAAAAAAAAAAAAAAFRUVC8/uHb/Qrh4//Tz8/7x8fD+7u7t/uvr6v7o5+f+P7h2/z+4dv8/uHb/VFRUJQAAAAAAAAAAAAAAAFJcVwJGu33/Tr+E//r6+v/4+Pj/9vX1/vPz8v/w8O//ZcCP/z+4dv8/uHb/P7h2/z+4dv8AAAAAAAAAAAAAAABKtn3uVMKJ/13Gkv/9/f3//Pv7/vr5+f749/f/W8WQ/lHBh/5Gu33+P7h2/z+4dv8/uHb/VFRUVAAAAAAAAAAAVcKK/2HIlf9ox5r/n93B/ovTs/910qn+cdCk/mjMnP5dxpL+UcCH/kS6ev4/uHb/PrV0/m6tivYAAAAAAAAAAF7Hk/9szaD/ydXP/3/Wsf95x6f+g9e0/n7WsP510qj+aMyc/lvFkP5MvoL+8/Ly/+/v7v7s6+r+AAAAAAAAAABmypr/zNrT/4DWsv+K2br/kNy//4/cvv6H2bj+ftWw/nHQpP5hyJb+UsGH/vj39/719PT+8fHw/gAAAAAAAAAAac2e//X09P52vqH+fr+m/5vgyf+Z4Mf+j9y+/oPXtP510qn+Zcqa/lXDiv7r9e/++fn5/vb29f4AAAAAAAAAAGrNnv75+Pj/9/b2//T08/6d38n+m+DJ/5Dcv/+D17X+dtOp/mbLmv9Vw4v+Rbp7/v39/f9NT05zAAAAAAAAAABcgG8E/Pz8/vr6+v/4+Pj+9vb1/uLm5P6Ayqz+f9ax/3LRpv9jyZf/U8GI/+no5/4+t3X/KSsqAQAAAAAAAAAAAAAAAG7MoPH9/f3/+/v7/4G7pP749/f+esmo/3jTq/9rzZ//XcaS/0ihc//w7+/+U1NTLwAAAAAAAAAAAAAAAAAAAAAAAAAAbsyg8f7+/v/g5+P+gLGc/23DnP739/b+brOR/1TCif/F1sz+UlJSNQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB4jIIEaMKX/v7+/v78/Pz/+/r6/1XCiv+93szuaWpqAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA//8AAPw/AADwDwAA4AcAAMADAACAAwAAgAEAAIABAACAAQAAgAEAAIADAADAAwAAwAcAAOAPAAD4HwAA//8AAA=="

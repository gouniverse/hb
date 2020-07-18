package html

// Addslashes addslashes()
func addslashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// Tag represents an HTML tag
type Tag struct {
	name        string
	text       string
	attributes map[string]string
	children   []*Tag
}

// ToHTML returns HTML from Node
func (t Tag) ToHTML() string {
	attrString := ""
	
	for key, value := range t.attributes {
		attrString += `"` + key + `"="` + addslashes(value) + `"`
	}
  
	return `<` + t.name + ` ` + attrString + `>` + t.text + `</` + t.name + `>`
}

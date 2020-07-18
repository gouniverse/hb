package html

// Tag represents an HTML tag
type Tag struct {
	name        string
	text       string
	attributes map[string]string
	children   []*Node
}

// ToHTML returns HTML from Node
func (t Tag) ToHTML() string {
	attrString := ""
	
  for key, value := range t.attributes {
		attrString += `"` + key + `"="` + value + `"`
	}
  
	return `<` + t.name + ` ` + attrString + `>` + t.text + `</` + t.name + `>`
}

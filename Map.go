package hb

// Map a slice of anything to a slice of tags.
func Map[T any](items []T, callback func(item T, index int) TagInterface) []TagInterface {
	tags := make([]TagInterface, len(items))
	for i, item := range items {
		tags[i] = callback(item, i)
	}
	return tags
}

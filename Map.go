package hb

// Map a slice of anything to a slice of tags.
func Map[T any](items []T, callback func(item T, index int) TagInterface) []TagInterface {
	var tags []TagInterface
	for index, item := range items {
		tags = append(tags, callback(item, index))
	}
	return tags
}

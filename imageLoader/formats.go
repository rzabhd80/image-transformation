package imageLoader

var imageFormats = []string{"jpg", "jpeg", "png", "gif", "webp"}

func validateFormat(format string) bool {
	for _, frmt := range imageFormats {
		if frmt == format {
			return true
		}
	}
	return false
}

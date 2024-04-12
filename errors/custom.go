package errors

type UnsupportedFormat struct {
	message string
}

func (e UnsupportedFormat) Error() string {
	return "Unsupported image format"
}

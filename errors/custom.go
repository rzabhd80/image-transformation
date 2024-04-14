package errors

type UnsupportedFormat struct {
	message string
}

type InstanceAlreadyExists struct {
	message string
}

func (e UnsupportedFormat) Error() string {
	return "Unsupported image format"
}

func (e InstanceAlreadyExists) Error() string {
	return "Instance already exists"
}

type CommandNotFound struct {
	message string
}

func (e CommandNotFound) Error() string {
	return "Command not found"
}

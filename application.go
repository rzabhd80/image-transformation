package image_transformation

import (
	"flag"
	"image_transformation/commands"
	"image_transformation/encoder"
	"image_transformation/errors"
	"image_transformation/imageLoader"
	"os"
	"strings"
)

var hasInstance bool = false

type Application struct {
	commands commands.Commands
	image    imageLoader.Image
}

func NewApplication(path string, format string) (*Application, error) {
	if hasInstance {
		return nil, errors.InstanceAlreadyExists{}
	}
	loader, err := imageLoader.NewImage(format, path)
	if err != nil {
		return nil, err
	}
	command := commands.NewCommands()
	hasInstance = true
	return &Application{commands: *command, image: *loader}, nil
}

func registerCommands(command commands.Commands) {
	commandList := commands.Transformations
	for name, funcName := range commandList {
		command.AddCommand(name, funcName.Apply)
	}
}

func runOperation(application Application, operation string) (*imageLoader.Image, error) {
	result := application.commands.GetCommand(operation)
	if result == nil {
		return nil, errors.CommandNotFound{}
	}
	data := result(application.image.Data)
	application.image.Data = data
	return &application.image, nil
}

func storeResult(image imageLoader.Image) bool {
	var path string = "./out"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755) // 0755 is the default permission
		if err != nil {
			return false
		}
	}
	parts := strings.Split(image.Path, "/")
	var fileName string = strings.TrimSpace(parts[len(parts)-1])
	err := encoder.EncodeAndWriteData(image, path+fileName)
	if err != nil {
		return false
	}
	return true
}

func InitApplication() (bool, error) {
	var path string
	var format string
	var operation string
	flag.StringVar(&path, "path", "Anonymous", "the file path")
	flag.StringVar(&format, "format", "Anonymous", "the file format required")
	flag.StringVar(&operation, "operation", "Anonymous", "the operation required")
	flag.Parse()
	application, err := NewApplication(path, format)
	if err != nil {
		return false, err
	}
	registerCommands(application.commands)
	result, err := runOperation(*application, operation)
	if err != nil {
		return false, err
	}
	storeResult(*result)
	return true, nil
}

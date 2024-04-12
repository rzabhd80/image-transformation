package image_transformation

import (
	"flag"
	"image_transformation/commands"
	"image_transformation/errors"
	"image_transformation/imageLoader"
)

var hasInstance bool = false

type Application struct {
	commands    commands.Commands
	imageLoader imageLoader.ImageLoader
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
	return &Application{commands: *command, imageLoader: loader}, nil
}

func registerCommands(command commands.Commands) {
	commandList := commands.Transformations
	for name, funcName := range commandList {
		command.AddCommand(name, funcName.Apply)
	}
}

func InitApplication() (bool, error) {
	//TODO implement the application instance
	var path string
	var format string
	flag.StringVar(&path, "path", "Anonymous", "the file path")
	flag.StringVar(&format, "operation", "Anonymous", "the operation required")
	flag.Parse()
	return hasInstance, nil
}

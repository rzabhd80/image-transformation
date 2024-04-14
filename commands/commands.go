package commands

import "image"

type Commands struct {
	commands map[string]func(img image.Image) image.Image
}

func NewCommands() *Commands {
	return &Commands{commands: map[string]func(img image.Image) image.Image{}}

}

func (cmds *Commands) AddCommand(name string, f func(img image.Image) image.Image) {
	cmds.commands[name] = f
}

func (cmds *Commands) Run(command string) func(img image.Image) image.Image {
	return cmds.commands[command]
}
func (cmds *Commands) GetCommand(operation string) func(img image.Image) image.Image {
	op, err := cmds.commands[operation]
	if err {
		return nil
	} else {
		return op
	}
}

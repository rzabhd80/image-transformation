package io

type Commands struct {
	commands map[string]func()
}

func NewCommands() *Commands {
	return &Commands{commands: map[string]func(){}}

}

func (cmds *Commands) AddCommand(name string, f func()) {
	cmds.commands[name] = f
}

func (cmds *Commands) Run(command string) func() {
	return cmds.commands[command]
}

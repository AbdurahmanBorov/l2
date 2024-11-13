package pattern

import "fmt"

type Command interface {
	Execute()
}

type StartCommand struct{}

func (c *StartCommand) Execute() { fmt.Println("Запуск") }

type StopCommand struct{}

func (c *StopCommand) Execute() { fmt.Println("Остановка") }

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(c Command) {
	i.command = c
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func CommandPattern() {
	invoker := &Invoker{}

	start := &StartCommand{}
	stop := &StopCommand{}

	invoker.SetCommand(start)
	invoker.ExecuteCommand()

	invoker.SetCommand(stop)
	invoker.ExecuteCommand()
}

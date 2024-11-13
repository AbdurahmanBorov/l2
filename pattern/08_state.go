package pattern

import "fmt"

type State interface {
	DoAction(context *Context)
}

type StartState struct{}

func (s *StartState) DoAction(context *Context) {
	fmt.Println("Включение")
	context.SetState(s)
}

type StopState struct{}

func (s *StopState) DoAction(context *Context) {
	fmt.Println("Выключение")
	context.SetState(s)
}

type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.DoAction(c)
}

func (c *Context) GetState() State {
	return c.state
}

func StatePattern() {
	context := &Context{}

	startState := &StartState{}
	startState.DoAction(context)

	stopState := &StopState{}
	stopState.DoAction(context)
}

package games

type Event struct {
	gid   int
	wid   int
	value interface{}
}

type Game interface {
	Init()
	Tick()
	Update(event Event)
	GetInputsState() []interface{}
	GetOutputsState() []interface{}
	GetObjectives() []string
}

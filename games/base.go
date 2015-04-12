package games

type Event struct {
	Gid   int
	Wid   int
	Value interface{}
}

type Game interface {
	Init()
	Tick()
	GetGid() int
	Update(event Event)
	GetInputsState() []interface{}
	GetOutputsState() []interface{}
	GetObjectives() []string
	CheckObjectives() bool
}

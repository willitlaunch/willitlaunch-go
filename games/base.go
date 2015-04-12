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
	SetObjectives([]string)
	CheckObjectives() bool
}

type GameBase struct {
	Gid int
}

func (g *GameBase) GetGid() int {
	return g.Gid
}

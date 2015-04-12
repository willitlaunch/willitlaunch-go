package games

import (
	"github.com/willitlaunch/willitlaunch-go/widgets"
	"math/rand"
)

const (
	lmin = 0
	lmax = 100
	fmin = -1.0
	fmax = 1.0
)

type CryogenicLevelsGame struct {
	GameBase
	CryogenicLevel int
	CLDial         widgets.Dial
	CLBool         widgets.Bool
	CLSlider       widgets.Slider

	cryogenicFlow float32
	broken        bool
}

//Cabin Pressure will be more or less identical

func (g *CryogenicLevelsGame) Init() {
	g.CLDial = widgets.Dial{WidgetBase: widgets.WidgetBase{Gid: g.Gid, Wid: 0, Label: "Cryogenic liquid level"}, Value: 50, Min: lmin, Max: lmax}
	g.CLSlider = widgets.Slider{WidgetBase: widgets.WidgetBase{Gid: g.Gid, Wid: 1, Label: "Cryogenic liquid flow"}, Value: 0.5, Min: fmin, Max: fmax}
	g.CLBool = widgets.Bool{WidgetBase: widgets.WidgetBase{Gid: g.Gid, Wid: 1, Label: "Fuel cells active"}, Value: true}
	g.CLDial.Init()
	g.CLSlider.Init()
	g.CLBool.Init()
	g.CryogenicLevel = 50
	g.cryogenicFlow = 0.5
	g.broken = false
}

func (g *CryogenicLevelsGame) Tick() {
	g.UserInteractionUpdate()

	g.CLDial.Value = float32(g.CryogenicLevel)
}

func flowEffect(g *CryogenicLevelsGame) float32 {
	return float32(g.CryogenicLevel) * (1.0 + g.cryogenicFlow*rand.Float32()*0.3 + 0.05*(rand.Float32()-0.5))
}

func (g *CryogenicLevelsGame) UserInteractionUpdate() {
	level := int(flowEffect(g))

	// if out of bound -> broken
	if g.broken {
		return
	}

	if level > lmax || level < lmin {
		g.broken = true
		level = 0
		g.CLBool.Value = false
	}

	g.CryogenicLevel = level
}

func (g *CryogenicLevelsGame) Update(event Event) {
	//TODO: check value is between fmin/fmax
	g.cryogenicFlow = event.Value.(float32)
}

func (g *CryogenicLevelsGame) GetInputsState() []interface{} {
	return []interface{}{&g.CLSlider}
}

func (g *CryogenicLevelsGame) GetOutputsState() []interface{} {
	return []interface{}{&g.CLDial, &g.CLBool}
}

func (g *CryogenicLevelsGame) GetObjectives() []string {
	//healthy:  35 < mid range < 65
	return []string{"Fuel cells must be active", "Cryogenic Liquid must stay in a mid range"}
}

func (g *CryogenicLevelsGame) CheckObjectives() bool {
	return !g.broken && g.CryogenicLevel > 35 && g.CryogenicLevel < 65
}

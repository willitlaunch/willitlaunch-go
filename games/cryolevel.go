package games

import (
	"github.com/willitlaunch/willitlaunch-go/widgets"
	"math/rand"
)

const (
	lmin      = 0
	lmax      = 100
	fmin      = -10
	fmax      = 10
	maxbroken = 100
)

type CryogenicLevelsGame struct {
	GameBase
	CryogenicLevel int
	CLDial         widgets.Dial
	CLBool         widgets.Bool
	CLSlider       widgets.Slider

	cryogenicFlow float64
	broken        bool
	count         int
}

//Cabin Pressure will be more or less identical

func (g *CryogenicLevelsGame) Init() {
	g.CLDial = widgets.Dial{WidgetBase: widgets.WidgetBase{Gid: g.Gid, Wid: 0, Label: "Cryogenic liquid level"}, Value: 50, Min: lmin, Max: lmax}
	g.CLSlider = widgets.Slider{WidgetBase: widgets.WidgetBase{Gid: g.Gid, Wid: 1, Label: "Cryogenic liquid flow"}, Value: 2, Min: fmin, Max: fmax}
	g.CLBool = widgets.Bool{WidgetBase: widgets.WidgetBase{Gid: g.Gid, Wid: 1, Label: "Fuel cells active"}, Value: true}
	g.CLDial.Init()
	g.CLSlider.Init()
	g.CLBool.Init()
	g.CryogenicLevel = 50
	g.cryogenicFlow = 5
	g.broken = false
	g.count = 0
}

func (g *CryogenicLevelsGame) Tick() {
	g.UserInteractionUpdate()

	g.CLDial.Value = float32(g.CryogenicLevel)
}

func flowEffect(g *CryogenicLevelsGame) float64 {
	return float64(g.CryogenicLevel) * (1.0 + g.cryogenicFlow*rand.Float64()*0.03/20 + 0.005*(rand.Float64()-0.5))
}

func (g *CryogenicLevelsGame) UserInteractionUpdate() {
	level := int(flowEffect(g))

	// if out of bound -> broken, but not for too long
	if g.count > maxbroken {
		level = rand.Intn(lmax-lmin) + lmin
		g.count = 0
		g.broken = false
		g.CLBool.Value = false
	}
	if g.broken {
		level = rand.Intn(lmax-lmin) + lmin
		g.count += 1
	}

	if level >= lmax || level <= lmin {
		g.broken = true
	}

	if level > 100 {
		level = 100
	} else if level <= 0 {
		level = 0
		g.CLBool.Value = false
	}

	g.CryogenicLevel = level
}

func (g *CryogenicLevelsGame) Update(event Event) {
	//TODO: check value is between fmin/fmax
	g.cryogenicFlow = event.Value.(float64)
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

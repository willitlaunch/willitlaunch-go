package games

import (
	"github.com/willitlaunch/willitlaunch-go/widgets"
	"math/rand"
  //"fmt"
)

type CabinPressureGame struct {
	GameBase
	CabinPressure       int
	CLDial              widgets.Dial
	CLPressuriserSlider widgets.Slider
	CLResetButton       widgets.Button

	pressuriser float64
	broken      bool
}

const (
	pmin         = 0
	pmax         = 100 //Find a sensible value here...
	pminoutbound = 66
	smin         = -10
	smax         = 10
)

//Cabin Pressure will be more or less identical

func (g *CabinPressureGame) Init() {
	g.CLDial = widgets.Dial{WidgetBase: widgets.WidgetBase{Gid: g.Gid, Wid: 0, Label: "Cabin pressure level"}, Value: 50, Min: pmin, Max: pmax}
	g.CLPressuriserSlider = widgets.Slider{WidgetBase: widgets.WidgetBase{Gid: g.Gid, Wid: 1, Label: "Pressuriser"}, Value: 2, Min: smin, Max: smax}
	g.CLResetButton = widgets.Button{WidgetBase: widgets.WidgetBase{Gid: g.Gid, Wid: 2, Label: "Reboot pressuirser system"}, Value: false}
	g.CLDial.Init()
	g.CLPressuriserSlider.Init()
	g.CLResetButton.Init()
	g.CabinPressure = 50
	g.pressuriser = 2
	g.broken = false
}

func (g *CabinPressureGame) Tick() {
	g.UserInteractionUpdate()

	g.CLDial.Value = float32(g.CabinPressure)
}

func pressuriserEffect(g *CabinPressureGame) float64 {
  val := float64(g.CabinPressure) + g.pressuriser/2.0 + 3.0 * (rand.Float64()-0.5)
  //fmt.Println("{Pressure: ", g.CabinPressure, ", pressuriser: ", g.pressuriser, "pressuriser effect: ", val,"}")
	return val
}

func (g *CabinPressureGame) UserInteractionUpdate() {

	var pressure int

	if g.CLResetButton.Value {
		g.CLResetButton.Value = false
		g.broken = []bool{true, false, false}[rand.Intn(3)]
	}

	// if out of bound -> broken
	if g.broken {
		// if broken pressure oscillates wildly and is out of control
		pressure = rand.Intn(pmax-pminoutbound) + []int{pminoutbound, 0}[rand.Intn(2)]
	} else {
		pressure = int(pressuriserEffect(g))
	}

	if pressure > pmax || pressure < pmin {
		g.broken = true
		pressure = pmin
	}

	g.CabinPressure = pressure
}

func (g *CabinPressureGame) Update(event Event) {
	switch event.Wid {
	case 1:
		g.pressuriser = event.Value.(float64)
	case 2:
		g.CLResetButton.Value = event.Value.(bool)
	}
}

func (g *CabinPressureGame) GetInputsState() []interface{} {
	return []interface{}{&g.CLPressuriserSlider, &g.CLResetButton}
}

func (g *CabinPressureGame) GetOutputsState() []interface{} {
	return []interface{}{&g.CLDial}
}

func (g *CabinPressureGame) GetObjectives() []string {
	//healthy:  35 < mid range < 65
	return []string{"Cryogenic Liquid must not overflow", "Cryogenic Liquid must not completely vanish", "Cryogenic Liquid must stay in a mid range"}
}

func (g *CabinPressureGame) CheckObjectives() bool {
	return !g.broken && g.CabinPressure > 35 && g.CabinPressure < 65
}

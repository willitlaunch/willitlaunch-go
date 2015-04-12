package games

import (
	"github.com/willitlaunch/willitlaunch-go/widgets"
	"math/rand"
	//"fmt"
)

type WindGame struct {
	GameBase
	WindStrength  float32
	LaunchAngle   float32
	InputAngle    float32
	AngleWidget   widgets.Dial
	WindWidget    widgets.Bar
	ControlSlider widgets.Slider

	objectives []string
}

func (w *WindGame) Init() {
	w.WindStrength = 0.0
	w.LaunchAngle = 90.0
	w.InputAngle = 5.0
	w.AngleWidget = widgets.Dial{WidgetBase: widgets.WidgetBase{Gid: w.Gid, Wid: 0, Label: "Rocket Angle"}, Value: 90, Min: 0, Max: 180}
	w.WindWidget = widgets.Bar{WidgetBase: widgets.WidgetBase{Gid: w.Gid, Wid: 1, Label: "Wind Strength"}, Value: 0, Min: -10, Max: 10}
	w.ControlSlider = widgets.Slider{WidgetBase: widgets.WidgetBase{Gid: w.Gid, Wid: 2, Label: "Wind Compensation"}, Value: 0, Min: -30, Max: 30}
	w.AngleWidget.Init()
	w.WindWidget.Init()
	w.ControlSlider.Init()
	w.objectives = []string{"Keep the launch angle up!"}
}

func (w *WindGame) Tick() {
	w.WindStrength += float32((rand.Intn(3) - 1) * rand.Intn(2))
	if w.WindStrength < -10 {
		w.WindStrength = -10
	} else if w.WindStrength > 10 {
		w.WindStrength = 10
	}

	//delta := w.InputAngle
	w.LaunchAngle += 0.5 * w.InputAngle
	w.LaunchAngle += 0.2 * w.WindStrength

	if w.LaunchAngle < 0 {
		w.LaunchAngle = 0
	} else if w.LaunchAngle > 180 {
		w.LaunchAngle = 180
	}

	//fmt.Printf("(input, wind, LaunchAngle) -> (%v, %v, %v)\n", w.InputAngle, w.WindStrength, w.LaunchAngle)

	w.WindWidget.Value = float32(w.WindStrength)
	w.AngleWidget.Value = float32(w.LaunchAngle)
}

func (w *WindGame) Update(event Event) {
	w.InputAngle = float32(event.Value.(float64))
}

func (w *WindGame) GetInputsState() []interface{} {
	return []interface{}{&w.ControlSlider}
}

func (w *WindGame) GetOutputsState() []interface{} {
	return []interface{}{&w.WindWidget, &w.AngleWidget}
}

func (w *WindGame) GetObjectives() []string {
	return w.objectives
}

func (g *WindGame) SetObjectives(objs []string) {
	g.objectives = objs
}

func (w *WindGame) CheckObjectives() bool {
	return w.LaunchAngle > 65 && w.LaunchAngle < 115
}

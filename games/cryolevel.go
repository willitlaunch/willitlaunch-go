package games

import (
  "github.com/willitlaunch/willitlaunch-go/widgets"
  "math/rand"
)

type CryogenicLevelsGame struct {
  GameImpl
  CryogenicLevel         int
  CLWidget      widgets.Dial
  CLSlider      widgets.Slider

  cryogenicFlow float32
  broken bool
}

//Cabin Pressure will be more or less identical

func (g *CryogenicLevelsGame) Init() {
  g.CLWidget = widgets.Dial{WidgetBase: widgets.WidgetBase{Gid: g.Gid, Wid: 0, Label: "Cryogenic liquid level"}, Value: 50, Min: 0, Max: 100}
  g.CLSlider = widgets.Slider{WidgetBase: widgets.WidgetBase{Gid: g.Gid, Wid: 1, Label: "Cryogenic liquid flow"}, Value: 0.5, Min: -1.0, Max: 1.0}
  g.CLWidget.Init()
  g.CLSlider.Init()
  g.CryogenicLevel = 50
  g.cryogenicFlow = 0.5
  g.broken = false
}

func (g *CryogenicLevelsGame) Tick() {
  g.UserInteractionUpdate()

  g.CLWidget.Value = float32(g.CryogenicLevel)
}

func flowEffect(g *CryogenicLevelsGame) float32 {
  return float32(g.CryogenicLevel)*(1.0 + g.cryogenicFlow * rand.Float32() * 0.3 + 0.05*(rand.Float32() - 0.5))
}

func (g *CryogenicLevelsGame) UserInteractionUpdate() {
  level := int(flowEffect(g))

  // if out of bound -> broken
  if g.broken {
    return
  }

  if level > 100 || level < 0 {
    g.broken = true
    level = 0
  }

  g.CryogenicLevel = level
}

func (g *CryogenicLevelsGame) Update(event Event) {
  g.cryogenicFlow = event.Value.(float32)
}

func (g *CryogenicLevelsGame) GetInputsState() []interface{} {
  return []interface{}{&g.CLSlider}
}

func (g *CryogenicLevelsGame) GetOutputsState() []interface{} {
  return []interface{}{&g.CLWidget}
}

func (g *CryogenicLevelsGame) GetObjectives() []string {
  //healthy:  35 < mid range < 65
  return []string{"Cryogenic Liquid must not overflow", "Cryogenic Liquid must not completely vanish", "Cryogenic Liquid must stay in a mid range"}
}

func (g *CryogenicLevelsGame) CheckObjectives() bool {
  return !g.broken && g.CryogenicLevel > 35 && g.CryogenicLevel < 65
}

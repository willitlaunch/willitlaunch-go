package controllers

import (
	"github.com/willitlaunch/willitlaunch-go/games"
)

type FIDOController struct {
	FlightControllerBase
}

func (c *FIDOController) Init() {
	c.FlightControllerBase.Init()
	c.Name = "FIDO"
	XWGame := games.WindGame{GameBase: games.GameBase{Gid: 0}}
	XWGame.Init()
  XWGame.AngleWidget.Label = "X Rocket Angle"
  XWGame.WindWidget.Label = "X Wind Strength"
  XWGame.ControlSlider.Label = "X Wind Compensation"
	c.Games = append(c.Games, &XWGame)

  YWGame := games.WindGame{GameBase: games.GameBase{Gid: 1}}
  YWGame.Init()
  YWGame.AngleWidget.Label = "Y Rocket Angle"
  YWGame.WindWidget.Label = "Y Wind Strength"
  YWGame.ControlSlider.Label = "Y Wind Compensation"
  c.Games = append(c.Games, &YWGame)
}

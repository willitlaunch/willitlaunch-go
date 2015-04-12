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
	WGame := games.WindGame{GameBase: games.GameBase{Gid: 0}}
	WGame.Init()
	c.Games = append(c.Games, &WGame)
}

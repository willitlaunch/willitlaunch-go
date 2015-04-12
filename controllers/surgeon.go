package controllers

import (
	"github.com/willitlaunch/willitlaunch-go/games"
)

type SurgeonController struct {
	FlightControllerBase
}

func (c *SurgeonController) Init() {
	c.FlightControllerBase.Init()
	c.Name = "Surgeon"
	HRGame := games.HeartRateGame{GameBase: games.GameBase{Gid: 0}}
	HRGame.Init()
	c.Games = append(c.Games, &HRGame)
}

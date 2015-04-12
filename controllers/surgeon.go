package controllers

import (
	"github.com/willitlaunch/willitlaunch-go/games"
)

type SurgeonController struct {
	FlightControllerImpl
}

func (c *SurgeonController) Init() {
	c.Name = "Surgeon"
	HRGame := games.HeartRateGame{Gid: 0}
	HRGame.Init()
	c.Games = append(c.Games, &HRGame)
}

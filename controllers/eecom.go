package controllers

import (
	"github.com/willitlaunch/willitlaunch-go/games"
)

type EECOMController struct {
	FlightControllerBase
}

func (c *EECOMController) Init() {
	c.FlightControllerBase.Init()
	c.Name = "EECOM"

	CLGame := games.CryogenicLevelsGame{GameBase: games.GameBase{Gid: 0}}
	CLGame.Init()
	c.Games = append(c.Games, &CLGame)

	CPGame := games.CabinPressureGame{GameBase: games.GameBase{Gid: 1}}
	CPGame.Init()
	c.Games = append(c.Games, &CPGame)

	// BPGame := games.BatteryPowerGame{Gid: 2}
	// BPGame.Init()
	// c.Games = append(c.Games, &BPGame)
}

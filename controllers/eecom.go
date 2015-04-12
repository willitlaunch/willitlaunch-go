package controllers

import (
  "github.com/willitlaunch/willitlaunch-go/games"
)

type EECOMController struct {
  FlightControllerImpl
}

func (c *EECOMController) Init() {
  c.Name = "Electrical, Environmental and Consumables Manager (EECOM)"

  CLGame := games.CryogenicLevelsGame{GameImpl: games.GameImpl{Gid: 0}}
  CLGame.Init()
  c.Games = append(c.Games, &CLGame)

  // CPGame := games.CabinPressureGame{Gid: 1}
  // CPGame.Init()
  // c.Games = append(c.Games, &CPGame)

  // BPGame := games.BatteryPowerGame{Gid: 2}
  // BPGame.Init()
  // c.Games = append(c.Games, &BPGame)
}

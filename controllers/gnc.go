package controllers

import (
  "github.com/willitlaunch/willitlaunch-go/games"
)

type GNCController struct {
	FlightControllerBase
}

func (c *GNCController) Init() {
  c.FlightControllerBase.Init()
  c.Name = "GNC"

  CPGame1 := games.CabinPressureGame{GameBase: games.GameBase{Gid: 0}}
  CPGame1.Init()
  CPGame1.CLDial.Label = "Fuel Pressurisation Level"
  CPGame1.SetObjectives([]string{"Fuel pressurisation level must stay stable and in a mid range"})
  CPGame1.CLResetButton.Label = "Depressurise"

  CryoGame1 := games.CryogenicLevelsGame{GameBase: games.GameBase{Gid: 1}}
  CryoGame1.Init()
  CryoGame1.CLDial.Label = "Fuel Temperature"
  CryoGame1.CLBool.Label = "Fuel Temperature Warning"
  CryoGame1.SetObjectives([]string{"Keep fuel temperature stable and in the appropriate range"})

  CryoGame2 := games.CryogenicLevelsGame{GameBase: games.GameBase{Gid: 2}}
  CryoGame2.Init()
  CryoGame2.CLDial.Label = "Fuel Density"
  CryoGame2.CLBool.Label = "Fuel Density Warning"
  CryoGame2.SetObjectives([]string{"Keep fuel density stable and in the appropriate range"})
  CryoGame2.SetObjectives([]string{})

  c.Games = append(c.Games, &CPGame1)
  c.Games = append(c.Games, &CryoGame1)
  c.Games = append(c.Games, &CryoGame2)
}

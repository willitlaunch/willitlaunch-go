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
	HRGame1 := games.HeartRateGame{GameBase: games.GameBase{Gid: 0}}
	HRGame1.Init()
	HRGame2 := games.HeartRateGame{GameBase: games.GameBase{Gid: 1}}
  HRGame2.SetObjectives([]string{})
	HRGame2.Init()
	CryoGame1 := games.CryogenicLevelsGame{GameBase: games.GameBase{Gid: 2}}
	CryoGame1.Init()
	CryoGame1.CLDial.Label = "Fluid Level"
	CryoGame1.CLBool.Label = "No Headache Detected"
  CryoGame1.SetObjectives([]string{"Keep fluid level stable and in mid range", "Keep headaches under control"})
	CryoGame2 := games.CryogenicLevelsGame{GameBase: games.GameBase{Gid: 3}}
	CryoGame2.Init()
	CryoGame2.CLDial.Label = "BAC"
	CryoGame2.CLBool.Label = "Vital Signs Present"
  CryoGame2.SetObjectives([]string{"Thou shall not drink!", "Keep the pilot alive"})
	c.Games = append(c.Games, &HRGame1)
	c.Games = append(c.Games, &CryoGame1)
	c.Games = append(c.Games, &HRGame2)
	c.Games = append(c.Games, &CryoGame2)
}

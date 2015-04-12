package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/willitlaunch/willitlaunch-go/games"
)

type SurgeonController struct {
	Name  string
	Games []games.Game
}

func (c *SurgeonController) Init() {
	c.Name = "Surgeon"
	HRGame := games.HeartRateGame{Gid: 0}
	HRGame.Init()
	c.Games = append(c.Games, &HRGame)
}

func (c *SurgeonController) Tick() {
	for _, game := range c.Games {
		game.Tick()
	}
}

func (c *SurgeonController) Update(event Event) {
	for _, game := range c.Games {
		if game.GetGid() == event.Gid {
			evt := games.Event{Gid: event.Gid, Wid: event.Wid, Value: event.Value}
			game.Update(evt)
			return
		}
	}
}

func (c *SurgeonController) GetInitJSON() []byte {
	var inputStates []interface{}
	var outputStates []interface{}
	var objectives []interface{}
	for _, game := range c.Games {
		for _, state := range game.GetOutputsState() {
			outputStates = append(outputStates, state)
		}
		for _, state := range game.GetInputsState() {
			inputStates = append(inputStates, state)
		}
		for _, objective := range game.GetObjectives() {
			objectives = append(objectives, objective)
		}
	}

	output := map[string]interface{}{
		"inputWidgets":  inputStates,
		"outputWidgets": outputStates,
		"objectives":    objectives,
		"name":          c.Name,
	}
	out, err := json.Marshal(output)
	if err != nil {
		out = []byte{'E', 'r', 'r', 'o', 'r'}
		fmt.Println("error:", err)
		return out
	}
	return out
}

func (c *SurgeonController) GetTickJSON() []byte {
	var outputStates []interface{}
	for _, game := range c.Games {
		for _, state := range game.GetOutputsState() {
			outputStates = append(outputStates, state)
		}
	}
	output := map[string]interface{}{
		"outputWidgets": outputStates,
	}

	out, err := json.Marshal(output)
	if err != nil {
		out := []byte{'E', 'r', 'r', 'o', 'r'}
		fmt.Println("error:", err)
		return out
	}

	return out
}

func (c *SurgeonController) CheckObjectives() bool {
	// Emptyset gives true!
	won := true
	for _, game := range c.Games {
		won = game.CheckObjectives() && won
	}
	return won
}

package games

import (
	"github.com/willitlaunch/willitlaunch-go/widgets"
	"math/rand"
)

type HeartRateGame struct {
	HeartRate     int
	Gid           int
	HRWidget      widgets.Number
	TranqButton   widgets.Button
	SteroidButton widgets.Button
	AEDButton     widgets.Button
}

func (hr *HeartRateGame) Init() {
	hr.HRWidget = widgets.Number{Gid: hr.Gid, Wid: 0, Label: "Heart Rate", Value: 100, Min: 0, Max: 200}
	hr.TranqButton = widgets.Button{Gid: hr.Gid, Wid: 1, Label: "Tranquilizers", Value: false}
	hr.SteroidButton = widgets.Button{Gid: hr.Gid, Wid: 2, Label: "Steroids", Value: false}
	hr.AEDButton = widgets.Button{Gid: hr.Gid, Wid: 3, Label: "AED", Value: false}
	hr.HRWidget.Init()
	hr.TranqButton.Init()
	hr.SteroidButton.Init()
	hr.AEDButton.Init()
}

func (hr *HeartRateGame) Tick() {
	hr.UserInteractionUpdate()
	if rand.Intn(500) == 0 {
		// Heart attack! Better AED.
		hr.HeartRate = 0
	} else {
		hr.HeartRate += (rand.Intn(3) - 1) * rand.Intn(3)
	}

	if hr.HeartRate > 300 || hr.HeartRate < 0 {
		hr.HeartRate = 0
	}
}

func (hr *HeartRateGame) GetGid() int {
	return hr.Gid
}

func medicineEffect() int {
	return rand.Intn(20) + 5
}

func (hr *HeartRateGame) UserInteractionUpdate() {
	if hr.TranqButton.Value {
		hr.TranqButton.Value = false
		hr.HeartRate -= medicineEffect()
	} else if hr.SteroidButton.Value {
		hr.SteroidButton.Value = false
		hr.HeartRate += medicineEffect()
	} else if hr.AEDButton.Value {
		hr.AEDButton.Value = false
		if hr.HeartRate != 0 {
			// Fibrillation
			hr.HeartRate = rand.Intn(450) + 50
		} else {
			hr.HeartRate += rand.Intn(20) + 60
		}
	}
}

func (hr *HeartRateGame) Update(event Event) {
	switch event.Wid {
	case 1:
		hr.TranqButton.Value = event.Value.(bool)
	case 2:
		hr.SteroidButton.Value = event.Value.(bool)
	case 3:
		hr.AEDButton.Value = event.Value.(bool)
	}
}

func (hr *HeartRateGame) GetInputsState() []interface{} {
	return []interface{}{&hr.TranqButton, &hr.SteroidButton, &hr.AEDButton}
}

func (hr *HeartRateGame) GetOutputsState() []interface{} {
	return []interface{}{&hr.HRWidget}
}

func (hr *HeartRateGame) GetObjectives() []string {
	//healthy:  60 < HR < 120
	return []string{"Astronaut must have a healthy HR", "Astronaut heart must be beating"}
}

func (hr *HeartRateGame) CheckObjectives() bool {
	return hr.HeartRate > 60 && hr.HeartRate < 120
}

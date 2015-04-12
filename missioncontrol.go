package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"time"
)

type timeMsg struct {
	timeLeft time.Duration
}

type statusMsg struct {
	status string
}

type pollMsg struct {
	status         string
	controllerName string
}

func missionTimer() {
	World.GameLength = 60e3 * time.Millisecond
	World.StartTime = time.Now()
	for {
		World.TimeLeft = World.GameLength - time.Now().Sub(World.StartTime)
		if World.TimeLeft <= 0 {
			World.TimeLeft = 0
		}
		time_msg := timeMsg{World.TimeLeft}
		msg, _ := json.Marshal(time_msg)

		for _, player := range World.players {
			player.ws.WriteMessage(websocket.TextMessage, msg)
		}

		if World.TimeLeft == 0 {
			missionFailed()
			return
		}
	}
}

func missionControl() {
	for {
		allReady := true
		for _, player := range World.players {
			allReady = allReady && player.controller.GetIsReady()
		}
		if allReady {
			missionPoll()
		}
	}
}

func missionFailed() {
	sendAllPlayers("FAILED")
}

func missionSuccess() {
	sendAllPlayers("SUCCESS")
}

func missionPoll() {
	sendAllPlayers("POLL")
	for _, player := range World.players {
		activeName := player.controller.GetName()
		msg, _ := json.Marshal(pollMsg{"POLLCONT", activeName})
		for _, player := range World.players {
			player.ws.WriteMessage(websocket.TextMessage, msg)
		}
		result := <-World.GoNoGo
		if !result {
			sendAllPlayers("NOPOLL")
			return
		}
	}

	allOk := true
	for _, player := range World.players {
		allOk = allOk && player.controller.CheckObjectives()
	}

	if allOk {
		missionSuccess()
	} else {
		missionFailed()
	}
}

func sendAllPlayers(s string) {
	msg, _ := json.Marshal(statusMsg{s})
	for _, player := range World.players {
		player.ws.WriteMessage(websocket.TextMessage, msg)
	}
}

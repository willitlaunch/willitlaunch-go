package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type timeMsg struct {
	TimeLeft int
}

type statusMsg struct {
	Status string
}

type pollMsg struct {
	Status         string
	ControllerName string
}

func missionTimer() {
	World.GameLength = 60e3 * time.Millisecond
	for len(World.players) == 0 {
		time.Sleep(100 * time.Millisecond)
	}
	World.StartTime = time.Now()
	for {
		World.TimeLeft = World.GameLength - time.Now().Sub(World.StartTime)
		if World.TimeLeft <= 0 {
			World.TimeLeft = 0
		}
		msg, _ := json.Marshal(timeMsg{int(World.TimeLeft) / 1e9})

		for _, player := range World.players {
			player.ws.WriteMessage(websocket.TextMessage, msg)
		}

		if World.TimeLeft == 0 {
			missionFailed()
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func missionControl() {
	for {
		time.Sleep(500 * time.Millisecond)
		if len(World.players) == 0 {
			continue
		}
		allReady := true
		for _, player := range World.players {
			allReady = allReady && player.controller.GetIsReady()
		}
		if allReady && len(World.players) > 0 {
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
	msg, err := json.Marshal(statusMsg{s})
	if err != nil {
		fmt.Println("err encoding json:", err)
	}
	for _, player := range World.players {
		player.ws.WriteMessage(websocket.TextMessage, msg)
	}
}

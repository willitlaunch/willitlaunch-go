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
		if World.Finished {
			return
		}

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
		if World.Finished {
			return
		}
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
			return
		}
	}
}

func missionFailed() {
	sendAllPlayers("FAILED")
	World.Finished = true
	fmt.Println("Mission Failed")
}

func missionSuccess() {
	sendAllPlayers("SUCCESS")
	World.Finished = true
	fmt.Println("Mission Successful")
}

func missionPoll() {
	World.GoNoGo = make(chan bool)
	sendAllPlayers("POLL")
	fmt.Println("OK all mission controllers, Going round the room")
	time.Sleep(4 * time.Second)
	for _, player := range World.players {
		activeName := player.controller.GetName()
		fmt.Println(activeName, "?")
		msg, _ := json.Marshal(pollMsg{"POLLCONT", activeName})
		for _, player := range World.players {
			player.ws.WriteMessage(websocket.TextMessage, msg)
		}
		result := <-World.GoNoGo
		fmt.Println(result, "!")
		if !result {
			sendAllPlayers("NOPOLL")
			fmt.Println("NO GO")
			return
		}
	}

	allOk := true
	fmt.Println("GO! WE ARE GO!")
	for _, player := range World.players {
		allOk = allOk && player.controller.CheckObjectives()
	}

	if allOk {
		missionSuccess()
	} else {
		missionSuccess()
		//missionFailed()
	}

	for _, player := range World.players {
		player.ws.Close()
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

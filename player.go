package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/willitlaunch/willitlaunch-go/controllers"
	"time"
)

const (
	wsWriteWait  = 10 * time.Second
	wsPongWait   = 60 * time.Second
	wsPingPeriod = (wsPongWait * 9) / 10
	tickPeriod   = time.Second
)

type Player struct {
	id         string
	ws         *websocket.Conn
	controller controllers.FlightController
}

func (p *Player) init() {
	p.controller = controllers.GetRandomController()
	init_msg := p.controller.GetInitJSON()
	err := p.ws.WriteMessage(websocket.TextMessage, init_msg)
	if err != nil {
		go p.listen()
		go p.run()
	} else {
		p.ws.Close()
	}
}

func (p *Player) listen() {
	defer p.ws.Close()
	p.ws.SetReadLimit(512)
	p.ws.SetReadDeadline(time.Now().Add(wsPongWait))
	p.ws.SetPongHandler(func(string) error {
		p.ws.SetReadDeadline(time.Now().Add(wsPongWait))
		return nil
	})
	for {
		msgType, msg, err := p.ws.ReadMessage()
		if err != nil {
			break
		}
		if msgType == websocket.TextMessage {
			p.update(msg)
		}
	}
}

func (p *Player) update(msg []byte) {
}

func (p *Player) run() {
	pingTicker := time.NewTicker(wsPingPeriod)
	tickTicker := time.NewTicker(tickPeriod)

	defer func() {
		pingTicker.Stop()
		tickTicker.Stop()
		p.ws.Close()
	}()

	for {
		select {
		case <-pingTicker.C:
			p.ws.SetWriteDeadline(time.Now().Add(wsWriteWait))
			err := p.ws.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				return
			}
		case <-tickTicker.C:
			p.ws.SetWriteDeadline(time.Now().Add(wsWriteWait))
			p.tick()
			msg := p.getTickJSON()
			err := p.ws.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				return
			}
		}
	}
}

func (p *Player) tick() {
	fmt.Printf("Player %s tick\n", p.id)
	p.controller.Tick()
}

func (p *Player) getTickJSON() []byte {
	return p.controller.GetTickJSON()
}

package main

import (
  "github.com/gorilla/websocket"
  "net/http"
  "time"
  "log"
)

const (
  // Time allowed to write the file to the client.
  writeWait = 10 * time.Second

  // Time allowed to read the next pong message from the client.
  pongWait = 60 * time.Second

  // Send pings to client with this period. Must be less than pongWait.
  pingPeriod = (pongWait * 9) / 10

  // Tick period
  tickPeriod = time.Second / 10
)

var (
  upgrader  = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
  }
)

func reader(ws *websocket.Conn) {
  defer ws.Close()
  ws.SetReadLimit(512)
  ws.SetReadDeadline(time.Now().Add(pongWait))
  ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
  for {
    _, _, err := ws.ReadMessage()
    if err != nil {
      break
    }
  }
}

func writer(ws *websocket.Conn) {
  pingTicker := time.NewTicker(pingPeriod)
  tickTicker := time.NewTicker(tickPeriod)
  defer func() {
    pingTicker.Stop()
    tickTicker.Stop()
    ws.Close()
  }()
  for {
    select {
    case <-pingTicker.C:
      ws.SetWriteDeadline(time.Now().Add(writeWait))
      if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
        return
      }
    case <-tickTicker.C:
      ws.SetWriteDeadline(time.Now().Add(writeWait))
      if err := ws.WriteMessage(websocket.TextMessage, []byte{}); err != nil {
        return
      }
    }
  }
}

func serveWs(w http.ResponseWriter, r *http.Request) {
  ws, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    if _, ok := err.(websocket.HandshakeError); !ok {
      log.Println(err)
    }
    return
  }

  go writer(ws)
  reader(ws)
}


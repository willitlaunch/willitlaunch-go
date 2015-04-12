package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

// Store all the players we know about
var World struct {
	players    map[string]Player
	StartTime  time.Time
	TimeLeft   time.Duration
	GameLength time.Duration
	GoNoGo     chan bool
}

type Event struct {
	gid   int
	wid   int
	value interface{}
}

var addr = flag.String("addr", ":8080", "http host:port")

// Start and run the websockets server on the commandline supplied port
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	World.players = make(map[string]Player, 0)
	go missionTimer()
	go missionControl()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	fmt.Printf("Starting web server on %s...\n", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}

// Serves debug JS
func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var v = struct{ Host string }{r.Host}
	homeTmpl.Execute(w, &v)
}

var homeTmpl, _ = template.New("").Parse(homeHTML)

const homeHTML = `<!DOCTYPE html>
<style>
div pre {
	padding: 5px;
	margin: 5px;
	background-color: #eee;
}
</style>
<div id=data></div>
<button id=tranq onclick="tranq()">Tranqs</button>
<button id=steroid onclick="steroid()">Steroids</button>
<button id=aed onclick="aed()">AED</button>
<script>
	d = document.getElementById("data");
	function log (msg) {
		var n = document.createElement("pre");
		n.textContent = msg;
		d.appendChild(n);
	};
	log("Connecting...");
	c = new WebSocket("ws://{{.Host}}/ws");
	c.onclose = function(e) {
		log("Connection closed.");
	};
	c.onmessage = function(e) {
		log(e.data);
	};
	window.c = c;
	function tranq(e) {
		window.c.send('{"Gid":0,"Wid":1,"Value":true}');
	}
	function steroid(e) {
		window.c.send('{"Gid":0,"Wid":2,"Value":true}');
	}
	function aed(e) {
		window.c.send('{"Gid":0,"Wid":3,"Value":true}');
	}
</script>
`

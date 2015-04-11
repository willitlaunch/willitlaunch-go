package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Store all the players we know about
var World struct {
	players map[string]Player
}

type Event struct {
	gid   int
	wid   int
	value interface{}
}

var addr = flag.String("addr", ":8080", "http host:port")

// Start and run the websockets server on the commandline supplied port
func main() {
	World.players = make(map[string]Player, 0)
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
</script>
`

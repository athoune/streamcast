package main

import (
	"net/http"

	"github.com/athoune/streamcast/icecast"
	"github.com/athoune/streamcast/web"
)

func main() {
	streamer := web.NewStreamer()
	mux := http.NewServeMux()
	mux.Handle("/stream", streamer)
	mux.Handle("/", http.FileServer(http.Dir("./static/")))

	i := icecast.New(streamer.Pubsub)
	go i.Listen("0.0.0.0:5000")
	http.ListenAndServe(":5001", mux)
}

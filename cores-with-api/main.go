package main

import (
	"net/http"

	"github.com/go-puzzles/cores"
	httppuzzle "github.com/go-puzzles/http-puzzle"
	"github.com/go-puzzles/pflags"
	"github.com/go-puzzles/plog"
	"github.com/gorilla/mux"
)

var (
	port = pflags.IntRequired("port", "service run port")
)

func main() {
	pflags.Parse()
	router := mux.NewRouter()

	router.Path("/hello").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	core := cores.NewPuzzleCore(
		httppuzzle.WithCoreHttpPuzzle("/api", router),
	)

	plog.PanicError(cores.Start(core, port()))
}

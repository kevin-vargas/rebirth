package main

import (
	"bytes"
	_ "embed"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kevin-vargas/rebirth/follower/webserver/config"
)

//go:embed index.html
var indexTemplate []byte

var (
	URI_TOKEN = []byte(`${{uri}}`)
)

func runServer(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index := bytes.ReplaceAll(indexTemplate, URI_TOKEN, []byte(r.URL.String()))
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write(index)
	})

	if err := http.ListenAndServe(addr, mux); err != nil {
		return err
	}
	return nil
}

func main() {
	godotenv.Load()
	cfg := config.Make()

	if err := runServer(cfg.Address); err != nil {
		panic(err)
	}
}

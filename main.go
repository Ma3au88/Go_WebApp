package main

import (
	"flag"
	"github.com/Ma3au88/Go_WebApp/deamon"
	"log"
	"net/http"
)

var assetsPath string

func processFlags() *deamon.Config {
	cfg := &deamon.Config{}

	flag.StringVar(&cfg.ListenSpec, "listen", "localhost:3000", "HTTP listen spec")
	flag.StringVar(&cfg.Db.ConnectString, "db-connect", "host=/var/run/postgresql dbname=gowebapp sslmode=disable", "DB Connect String")
	flag.StringVar(&assetsPath, "assets-path", "assets", "Path to assets dir")

	flag.Parse()
	return cfg
}

func setupHttpAssets(cfg *deamon.Config) {
	log.Printf("Assets served from %q.", assetsPath)
	cfg.UI.Assets = http.Dir(assetsPath)
}

func main() {
	cfg := processFlags()

	setupHttpAssets(cfg)

	if err := deamon.Run(cfg); err != nil {
		log.Printf("Error in main(): %v", err)
	}
}

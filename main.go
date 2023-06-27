/* Пакет main должен быть минимальным, насколько это возможно. Единственный код, который тут находится, —
это анализ аргументов команды. Если у приложения был бы конфигурационный файл, я поместил бы парсинг
и проверку этого файла в еще один пакет, который, скорее всего, назвал бы config. После этого main
должен передать управление пакету daemon */

package main

import (
	"flag"
	"github.com/Ma3au88/Go_WebApp/deamon"
	"log"
	"net/http"
)

var assetsPath string

func processFlags() *deamon.Config {
	// Основная цель daemon.Config — представить конфигурацию в структурированном и понятном формате
	cfg := &deamon.Config{}

	flag.StringVar(&cfg.ListenSpec, "listen", "localhost:3000", "HTTP listen spec")
	flag.StringVar(&cfg.Db.ConnectString, "db-connect", "host=/var/run/postgresql dbname=gowebapp sslmode=disable", "DB Connect String")
	flag.StringVar(&assetsPath, "assets-path", "assets", "Path to assets dir")

	flag.Parse()
	return cfg
}

func setupHttpAssets(cfg *deamon.Config) {
	log.Printf("Assets served from %q.", assetsPath)
	// http.Dir(assetsPath) — это подготовка к тому, как мы будем обслуживать статику в пакете ui.
	// Сделано это именно так, чтобы оставить возможность для другой реализации cfg.UI.Assets
	// (который является интерфейсом http.FileSystem), например, отдавать этот контент из оперативной памяти
	cfg.UI.Assets = http.Dir(assetsPath)
}

func main() {
	cfg := processFlags()

	setupHttpAssets(cfg)
	// daemon.Run(cfg) - фактически запускает наше приложение и блокируется до момента завершения работы
	if err := deamon.Run(cfg); err != nil {
		log.Printf("Error in main(): %v", err)
	}
}

/* Пакет daemon содержит все, что связано с запуском процесса. Сюда относится, например, какой порт
будет прослушиваться, здесь будет определен пользовательский журнал, а также все,
что связано с вежливым перезапуском и т.д. Поскольку задачей пакета daemon является инициализация
подключения к базе данных, ему нужно импортировать пакет db. Он также отвечает за прослушивание TCP порта
и запуск пользовательского интерфейса для этого слушателя, поэтому ему необходимо импортировать пакет ui,
а поскольку пакету ui необходимо иметь доступ к данным, который обеспечивается пакетом model,
ему также необходимо импортировать пакет model. */

package deamon

import (
	"gowebapp/db"
	"gowebapp/model"
	"gowebapp/ui"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	ListenSpec string

	Db db.Config
	UI ui.Config
}

// Run(*Config) - инициализирует соединение с базой данных, создает экземпляр model.Model
// и запускает ui, передавая ему настройки, указатели на модель и слушателя
func Run(cfg *Config) error {
	log.Printf("Starting, HTTP on: %s\n", cfg.ListenSpec)

	db, err := db.InitDb(cfg.Db)
	if err != nil {
		log.Printf("Error initializing database: %v\n", err)
		return err
	}

	m := model.New(db)

	l, err := net.Listen("tcp", cfg.ListenSpec)
	if err != nil {
		log.Printf("Error creating listener: %v\n", err)
		return err
	}

	ui.Start(cfg.UI, m, l)

	waitForSignal()

	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Printf("Got signal: %v, exiting.", s)
}

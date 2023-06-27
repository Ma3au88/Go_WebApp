/* ui обслуживает через HTTP */

package ui

import (
	"fmt"
	"github.com/Ma3au88/Go_WebApp/model"
	"net"
	"net/http"
	"time"
)

type Config struct {
	Assets http.FileSystem
}

func Start(cfg Config, m *model.Model, listener net.Listener) {
	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16}

	http.Handle("/", indexHandler(m))

	go server.Serve(listener)
}

const indexHTML = `
<!DOCTYPE HTML>
<html>
  <head>
    <meta charset="utf-8">
    <title>Simple Go Web App</title>
  </head>
  <body>
    <div id='root'></div>
  </body>
</html>
`

// indexHandler() — это не сам обработчик, он возвращает функцию-обработчик. Это делается таким образом для того,
// чтобы мы могли передать *model.Model через замыкание, так как сигнатура функции-обработчика HTTP
// фиксирована и указатель на модель не является одним из ее параметров
func indexHandler(m *model.Model) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, indexHTML)
	})
}

package server

import (
	"fmt"
	"klocalcp/common/ip"
	"log"
	"net/http"
	"sync"
)

func Start() (srv *http.Server) {
	return startHttpServer()
}

func startHttpServer() (srv *http.Server) {
	myip, ok := ip.MyLocalIP()
	if !ok {
		log.Fatal("local ip could not be found")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HAHAHAHEHEIGŞLQEWŞF"))
	})
	srv = &http.Server{Addr: myip + ":" + fmt.Sprint(ip.PORT)}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})

	go srv.ListenAndServe()
	// returning reference so caller can call Shutdown()
	return srv
}

type Server struct{
	sync.Mutex
	




}

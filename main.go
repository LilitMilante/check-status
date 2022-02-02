package main

import (
	"check-status/api"
	"check-status/service"
	"log"
)

func main() {
	var httpPort = "8080"
	ss := service.NewStatusService()
	srv := api.NewServer(ss)
	err := srv.Start(httpPort)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

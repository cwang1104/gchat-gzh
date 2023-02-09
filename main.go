package main

import (
	"gchat-gzh/api"
	"gchat-gzh/pkg/logger"
)

const (
	port = "12369"
	host = "0.0.0.0"
)

func main() {
	srv := api.NewServer(port)
	err := srv.RunServer(host)
	if err != nil {
		logger.Log.Errorf("server run error: %v", err)
		panic(err)
	}
}

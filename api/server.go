package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	Router *gin.Engine
}

func NewServer(port string) *Server {
	server := Server{
		Port: port,
	}

	server.setRouter()
	return &server
}

func (s *Server) setRouter() {

	r := router()

	s.Router = r
}
func (s *Server) RunServer(address string) error {
	url := fmt.Sprintf("%s:%s", address, s.Port)
	return s.Router.Run(url)
}

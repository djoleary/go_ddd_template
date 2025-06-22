package server

func (s *server) routes() {
	s.webserver.GET("/", s.handleSayHello(), middlewareAddRequestId())
}

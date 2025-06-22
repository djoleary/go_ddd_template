package server

func (s *server) routes() {
	s.webserver.GET("/", s.handleGreeting(), middlewareAddRequestId())
	s.webserver.GET("/healthcheck", s.handleHealthcheck(), middlewareAddRequestId())
	s.webserver.GET("/:name", s.handleGreetingByName(), middlewareAddRequestId())
}

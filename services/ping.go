package services

type PingService struct{}

func (s *PingService) Ping() map[string]string {
	return map[string]string{
		"ping": "pong",
	}
}

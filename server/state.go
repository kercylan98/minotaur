package server

import "time"

type State struct {
	server     *server
	LaunchedAt time.Time // 服务器启动时间
	Ip         string    // 服务器 IP 地址
}

func (s *State) init(srv *server) *State {
	s.server = srv
	return s
}

func (s *State) Status() *State {
	return &State{
		LaunchedAt: s.LaunchedAt,
		Ip:         s.Ip,
	}
}

func (s *State) onLaunched(ip string, t time.Time) {
	s.LaunchedAt = t
	s.Ip = ip
	s.server.events.onLaunched()
}

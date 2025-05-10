package main

type Server struct {
	conns map[*websocket.Conn]bool
}
package main

import "todo/internal/server"

func main() {
	s := server.New()
	s.Init()
	s.Run()
}

package main

import "github.com/christopherriley/3dttt/server/game_server"

func main() {
	gs := game_server.GameServer{}
	gs.Start(8080)
}

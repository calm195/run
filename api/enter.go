package api

import "run/service"

var RootGroup = new(Group)

type Group struct {
	GameApi
}

var (
	gameService = service.RootGroup.GameService
)

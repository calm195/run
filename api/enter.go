package api

import "run/service"

var RootGroup = new(Group)

type Group struct {
	GameApi
	RecordApi
}

var (
	gameService   = service.RootGroup.GameService
	recordService = service.RootGroup.RecordService
)

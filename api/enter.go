package api

import "run/service"

var RootGroup = new(Group)

type Group struct {
	GameApi
	RecordApi
	EventApi
	StandardApi
}

var (
	gameService     = service.RootGroup.GameService
	recordService   = service.RootGroup.RecordService
	eventService    = service.RootGroup.EventService
	standardService = service.RootGroup.StandardService
)

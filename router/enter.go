package router

import "run/api"

var RootGroup = new(Group)

type Group struct {
	Game     GameRouter
	Record   RecordRouter
	Event    EventRouter
	Standard StandardRouter
}

var (
	gameApi     = api.RootGroup.GameApi
	recordApi   = api.RootGroup.RecordApi
	eventApi    = api.RootGroup.EventApi
	standardApi = api.RootGroup.StandardApi
)
